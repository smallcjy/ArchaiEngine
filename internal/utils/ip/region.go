package ip

import (
	"encoding/json"
	"game/internal/logs"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"go.uber.org/zap"
)

var (
	Search      *xdb.Searcher
	LimitCounty = []string{""}
	LimitCity   = []string{""}
	CN          = "中国"
	DBFile      = "ip2region.xdb"
)

// ### 使用ip包前先调用,内存中搜索
// InitLocalData 初始化本地搜索对象
// dbPath 本地ip库路径
func InitLocalData(dbPath string) {

	cbuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		panic(err)
	}

	search, err := xdb.NewWithBuffer(cbuff)
	if err != nil {
		panic(err)
	}

	Search = search
}

type Region struct {
	CallBack string `json:"callback"`
}
type RegionCity struct {
	Country  string `json:"country"`
	Province string `json:"province"`
}

const (
	//
	GetUserLoction = "https://ip.useragentinfo.com/jsonp?&ip=IP"
)

// CheckIsCountry 检查是否限制
func CheckIsCountry(r *http.Request) bool {

	// 获取ip归属
	region := GetCountry(r)

	if region.Country == CN {
		city := region.Province
		//需要区分 香港,澳门 和台湾地区 没法按国家分,拿城市做区分
		// 限制大陆地区
		if !strings.Contains(city, "香港") && !strings.Contains(city, "澳门") && !strings.Contains(city, "台湾") {
			return true
		}
	}
	// 其他地区,
	for _, v := range LimitCity {
		if v == region.Province {
			return true
		}
	}
	// 其他国家
	for _, v := range LimitCounty {
		if v == region.Country {
			return true
		}
	}

	return false
}

// 获取玩家ip 归属
func GetCountry(r *http.Request) *RegionCity {
	ip := GexExenIp(r)
	region, err := ServerGetUserLoction(ip)
	logs.Info("GetCountry", zap.Any("ip", ip), zap.Any("region", region), zap.Any("err", err))
	if err == nil {
		return region
	} else {
		region, _ = LocalGetRegion(ip)
	}
	return region
}

// 查询用户ip地址
func GexExenIp(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// ServerGetUserLoction 用户ip归属
func ServerGetUserLoction(ip string) (*RegionCity, error) {

	url := GetUserLoction
	url = strings.ReplaceAll(url, "IP", ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	region := &RegionCity{}

	// 修改成json格式
	data := regexpReplaceJsonP(string(b))

	if err := json.Unmarshal([]byte(data), region); err != nil {
		return nil, err
	}

	return region, nil
}

// ? 本地的,网络拿不到备用(这个用代理没法识别准确地址
func LocalGetRegion(ip string) (*RegionCity, error) {

	var region = &RegionCity{}
	data, err := Search.SearchByStr(ip)
	if err != nil {
		logs.Warn("failed to SearchIP", zap.Error(err), zap.String("ip", ip))
		return region, err
	}

	now := strings.Split(data, "|")

	region.Country = now[0]
	region.Province = now[2]

	return region, nil

}

// 把jsonp格式转为json
func regexpReplaceJsonP(content string) string {
	content = strings.ReplaceAll(content, "(", "")
	content = strings.ReplaceAll(content, ")", "")
	content = strings.ReplaceAll(content, ";", "")
	content = strings.ReplaceAll(content, "callback", "")

	//fmt.Println("now", content)
	return content
}
