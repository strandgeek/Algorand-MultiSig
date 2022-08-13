package loggerutil

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	var cfg zap.Config
	level, err := zap.ParseAtomicLevel(viper.GetString("logger.level"))
	if err != nil {
		panic(err)
	}
	cfg.Level = level
	cfg.Encoding = viper.GetString("logger.encoding")
	cfg.OutputPaths = viper.GetStringSlice("logger.output_paths")
	cfg.ErrorOutputPaths = viper.GetStringSlice("logger.error_output_paths")
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	var err2 error
	logger, err2 := cfg.Build()
	if err2 != nil {
		panic(err2)
	}
	logger.Info("Zap logger construction succeeded")
	zap.ReplaceGlobals(logger)
	return logger, err2
}

//SyslogTimeEncoder - time stamp format for zap logger
func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 01, 2006  15:04:05"))
}
