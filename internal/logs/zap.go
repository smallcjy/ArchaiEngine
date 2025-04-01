package logs

import (
	"fmt"
	"game/internal/conf"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLog() {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(fmt.Sprintf("%s/%s.log", conf.GetServerConf().LogPath, conf.GetServerConf().AppName), conf.GetServerConf().LogTime)
	warnWriter := getWriter(fmt.Sprintf("%s/%s_error.log", conf.GetServerConf().LogPath, conf.GetServerConf().AppName), conf.GetServerConf().LogTime)

	var cores zapcore.Core
	// 最后创建具体的Logger
	if conf.GetServerConf().IsDebug {
		cores = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), debugLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), debugLevel))

	} else {
		cores = zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel)
	}
	warnCor := zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel)

	core := zapcore.NewTee(cores, warnCor)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
}

func getWriter(filename string, times int32) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// //保存7天内的日志，每天分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(times)),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
