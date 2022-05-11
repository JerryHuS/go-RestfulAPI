/**
 * @Author: alessonhu
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2022/5/5 11:42
 */

package logger

import (
	"apidemo/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	Log *zap.Logger
)

// InitLogger ...
func InitLogger() {
	initConsoleLogger()
}

func initConsoleLogger() {
	var err error
	if config.ServerConfig.Mode == config.ModeDebug {
		Log, err = zap.NewDevelopment()
	} else {
		Log = zap.New(zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.Lock(os.Stdout),
				zap.LevelEnablerFunc(func(level zapcore.Level) bool {
					return level == zapcore.InfoLevel
				}),
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.Lock(os.Stderr),
				zap.LevelEnablerFunc(func(level zapcore.Level) bool {
					return level == zapcore.ErrorLevel || level == zapcore.FatalLevel
				}),
			),
		), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	if err != nil {
		panic(err)
	}
}

// Sync ...
func Sync() {
	_ = Log.Sync()
}
