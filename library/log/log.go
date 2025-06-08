package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LOG_LEVEL_DEBUG = "DEBUG"
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_WARN  = "WARN"
	LOG_LEVEL_ERROR = "ERROR"
	LOG_LEVEL_FATAL = "FATAL"
	LOG_LEVEL_PANIC = "PANIC"
)

const (
	DEFAULT_LOG_MAXSIZE    = 100
	DEFAULT_LOG_MAXBACKUPS = 5
	DEFAULT_LOG_MAXAGE     = 30
)

var zapLog *zap.Logger

func NewLogConfig() *LogConfig {
	var logconfig LogConfig
	logconfig.UseLogRotation = false
	return &logconfig
}

func Valid(lc *LogConfig) error {
	if lc.UseLogRotation {
		if lc.LogProps.FileName == "" {
			return fmt.Errorf("log.file.name (%v)is required", lc.LogProps.FileName)
		} else {
			if lc.LogProps.MaxSize == 0 {
				lc.LogProps.MaxSize = DEFAULT_LOG_MAXSIZE
			}
			if lc.LogProps.MaxBackups == 0 {
				lc.LogProps.MaxBackups = DEFAULT_LOG_MAXBACKUPS
			}
			if lc.LogProps.MaxAge == 0 {
				lc.LogProps.MaxAge = DEFAULT_LOG_MAXAGE
			}

		}
	}
	return nil
}

func Init(lc *LogConfig) error {
	if err := Valid(lc); err != nil {
		fmt.Fprintf(os.Stderr, "log config invalid, fallback to stdout: %v\n", err)
		lc.UseLogRotation = false
	}
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(setLogLevel(lc.LogLevel))
	config.EncoderConfig.EncodeTime = jSTTimeEncoder
	if lc.UseLogRotation {
		enc := zapcore.NewJSONEncoder(config.EncoderConfig)
		sinki := zapcore.AddSync(&lumberjack.Logger{
			Filename:   lc.LogProps.FileName,
			MaxSize:    lc.LogProps.MaxSize,
			MaxBackups: lc.LogProps.MaxBackups,
			MaxAge:     lc.LogProps.MaxAge,
		})
		syncer := zap.CombineWriteSyncers(os.Stdout, sinki)
		logger := zap.New(zapcore.NewCore(enc, syncer, config.Level), zap.AddStacktrace(zap.NewAtomicLevelAt(zap.WarnLevel)))
		defer logger.Sync()
		zapLog = logger
	} else {
		logger, _ := config.Build()

		defer logger.Sync()
		zapLog = logger
	}
	return nil
}

func jSTTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	time := t.In(loc)
	enc.AppendString(time.Format("2006-01-02 15:04:05.000"))
}
func setLogLevel(level string) zapcore.Level {
	switch level {
	case LOG_LEVEL_DEBUG:
		return zapcore.DebugLevel
	case LOG_LEVEL_INFO:
		return zapcore.InfoLevel
	case LOG_LEVEL_WARN:
		return zapcore.WarnLevel
	case LOG_LEVEL_ERROR:
		return zapcore.ErrorLevel
	case LOG_LEVEL_FATAL:
		return zapcore.FatalLevel
	case LOG_LEVEL_PANIC:
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

func Debug(msg string, fields ...zap.Field) {
	zapLog.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLog.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLog.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLog.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zapLog.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zapLog.Panic(msg, fields...)
}
