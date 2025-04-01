package gateway

import (
	"fmt"
	"game/gate/conf"
	"game/gate/gateway/gatewaymodels"
	"game/logs"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/kercylan98/minotaur/utils/network"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/super"
	"github.com/kercylan98/minotaur/utils/times"
	"go.uber.org/zap"
)

var cli = resty.New().RemoveProxy() // REST 客户端
var instanceId string               // 实例 ID
var gatewayUrl string               // 网关地址
var port int                        // 端口
var forged bool                     // 是否为伪造的非腾讯云实例
var allowForged bool                // 是否允许伪造的非腾讯云实例
var currVersion int                 // 当前版本序列
var usePort bool

type response struct {
	Suc  bool   `json:"suc"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}

// InitGatewayClient 初始化网关客户端
func InitGatewayClient(gateway string, programPort int, isAllowForged bool, engine *gin.Engine, statusHandler func() *gatewaymodels.ServiceStatus, exitNotify func()) {
	gatewayUrl = gateway
	allowForged = isAllowForged
	port = programPort
	var statusHandlerCast = func(ctx *gin.Context) {
		status := statusHandler()
		status.GameRouter = fmt.Sprintf("/%s%s", GetInstanceId(), conf.ServerConfig.GameRouter)
		status.Forged = forged
		status.UsePort = usePort
		status.Ok = true
		ctx.JSON(200, status)
	}
	var pingHandler = func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
	}
	var upgradeHandler = func(ctx *gin.Context) {
		var params struct {
			Version int `json:"version"`
		}

		if err := ctx.BindJSON(&params); err != nil {
			ctx.JSON(200, gin.H{
				"msg": "参数错误",
				"suc": false,
			})
			return
		}

		if params.Version == currVersion {
			return // 无需更新
		}

		ctx.JSON(200, gin.H{
			"msg": "ok",
		})

		logs.Info("收到更新请求，即将退出", zap.String("实例 ID", GetInstanceId()))
		exitNotify()
	}

	engine.Any("/"+GetInstanceId()+conf.ServerConfig.GameRouter+"/action", func(context *gin.Context) {
		switch context.Query("type") {
		case "status":
			statusHandlerCast(context)
		case "ping":
			pingHandler(context)
		case "upgrade":
			upgradeHandler(context)
		}
	})

	ip, err := network.IP()
	if err == nil {
		logs.Info("本机 IP", zap.String("IP", ip.String()))
	}
	super.RetryForever(times.Second*5, func() error {
		var err = regService()
		if err != nil {
			logs.Error("注册服务到网关", zap.String("状态", "失败"), zap.String("实例 ID", GetInstanceId()), zap.String("网关地址", gatewayUrl), zap.Error(err))
		} else {
			logs.Info("注册服务到网关", zap.String("状态", "成功"), zap.String("实例 ID", GetInstanceId()), zap.Int("实例版本", GetVersion()), zap.String("网关地址", gatewayUrl))
		}
		return err
	})
}

// SetAllowForged 设置是否允许伪造的非腾讯云实例
func SetAllowForged(allow bool, useForgedInstanceId ...string) {
	allowForged = allow
	if allowForged && len(useForgedInstanceId) > 0 {
		if id := useForgedInstanceId[0]; len(id) > 0 {
			instanceId = id
			forged = true
		}
	}
}

// SetUsePort 设置是否需要携带端口访问
func SetUsePort(use bool) {
	usePort = use
}

// GetInstanceIP 获取实例 IP
func GetInstanceIP() string {
	if conf.ServerConfig.IsAws() {
		result, err := cli.NewRequest().
			SetHeader("X-aws-ec2-metadata-token-ttl-seconds", "21600").
			Put("http://169.254.169.254/latest/api/token")
		if err != nil {
			panic(fmt.Errorf("获取实例 IP 失败: %w", err))
		}

		token := string(result.Body())

		result, err = cli.NewRequest().
			SetHeader("X-aws-ec2-metadata-token", token).
			Get("http://169.254.169.254/latest/meta-data/public-ipv4")
		if err != nil {
			panic(fmt.Errorf("获取实例 IP 失败: %w", err))
		}
		ip := string(result.Body())
		logs.Info("初始化 AWS 实例 IP", zap.String("实例 IP", ip))
		return ip
	}

	if conf.ServerConfig.IsTencent() {
		result, err := cli.NewRequest().Get("http://metadata.tencentyun.com/latest/meta-data/public-ipv4")
		if err != nil {
			panic(fmt.Errorf("获取实例 IP 失败: %w", err))
		}
		ip := string(result.Body())
		logs.Info("初始化腾讯云实例 IP", zap.String("实例 IP", ip))
		return ip
	}

	return ""
}

// GetInstanceId 获取实例 ID
func GetInstanceId() string {
	if instanceId != "" {
		return instanceId
	}

	if conf.ServerConfig.IsAws() {
		id, err := os.ReadFile("/var/lib/cloud/data/instance-id")
		if err != nil {
			panic(fmt.Errorf("获取实例 ID 失败: %w", err))
		}
		instanceId = strings.TrimSpace(string(id))
		logs.Info("初始化 AWS 实例 ID", zap.String("实例 ID", instanceId))
	}

	if conf.ServerConfig.IsTencent() {
		result, err := cli.NewRequest().Get("http://metadata.tencentyun.com/latest/meta-data/instance-id")
		if err != nil {
			if strings.Contains(err.Error(), "no such host") && allowForged {
				forged = true
				instanceId = random.HostName()
				logs.Warn("初始化非腾讯云实例 ID", zap.String("实例 ID", instanceId))
				return instanceId
			}
			panic(fmt.Errorf("获取实例 ID 失败: %w", err))
		}
		if instanceId == "" {
			instanceId = string(result.Body())
			logs.Info("初始化腾讯云实例 ID", zap.String("实例 ID", instanceId))
		}
	}

	return instanceId
}

// regService 注册服务
func regService() error {
	result, err := cli.NewRequest().
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(gatewaymodels.BasicAuthAccount, gatewaymodels.BasicAuthPassword).
		SetBody(map[string]interface{}{
			"instance_id": GetInstanceId(),
			"instance_ip": GetInstanceIP(),
			"gameid":      conf.ServerConfig.GameId,
			"port":        port,
			"forged":      forged,
			"use_port":    usePort,
			"game_router": fmt.Sprintf("/%s%s", GetInstanceId(), conf.ServerConfig.GameRouter),
		}).
		SetResult(new(response)).
		Post(gatewayUrl + "/register")
	if err != nil {
		return fmt.Errorf("向网关 [%s] 注册服务 [%s] 失败: %w", gatewayUrl, GetInstanceId(), err)
	}

	resp := result.Result().(*response)
	if !resp.Suc {
		return fmt.Errorf("向网关 [%s] 注册服务 [%s] 失败: %s", gatewayUrl, GetInstanceId(), resp.Msg)
	}

	currVersion = resp.Data
	return nil
}

// GetVersion 获取当前版本
func GetVersion() int {
	return currVersion
}

// UnRegService 注销服务，注销后新的请求将无法到达，但已有的请求会继续处理
func UnRegService() error {
	result, err := cli.NewRequest().
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(gatewaymodels.BasicAuthAccount, gatewaymodels.BasicAuthPassword).
		SetBody(map[string]interface{}{
			"instance_id": GetInstanceId(),
			"gameid":      conf.ServerConfig.GameId,
		}).
		SetResult(new(response)).
		Delete(gatewayUrl + "/unRegister")
	if err != nil {
		return fmt.Errorf("向网关 [%s] 注销服务 [%s] 失败: %w", gatewayUrl, GetInstanceId(), err)
	}

	resp := result.Result().(*response)
	if !resp.Suc {
		return fmt.Errorf("向网关 [%s] 注销服务 [%s] 失败: %s", gatewayUrl, GetInstanceId(), resp.Msg)
	}

	return nil
}
