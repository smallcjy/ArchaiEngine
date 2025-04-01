package logs

import (
	"go.uber.org/zap"
)

func Info(method string, otherFields ...zap.Field) {
	Logger.Info(method, otherFields...)
}

func Debug(method string, otherFields ...zap.Field) {
	Logger.Debug(method, otherFields...)
}

func Warn(method string, otherFields ...zap.Field) {
	Logger.Warn(method, otherFields...)
}
func Error(method string, otherFields ...zap.Field) {
	Logger.Error(method, otherFields...)
}
