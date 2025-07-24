package logger

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// LogLevel 日志级别
type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
)

// Config 日志配置
type Config struct {
	Level      LogLevel `json:"level" yaml:"level"`
	Format     string   `json:"format" yaml:"format"` // json, text
	Output     string   `json:"output" yaml:"output"` // stdout, stderr, file
	FilePath   string   `json:"file_path" yaml:"file_path"`
	MaxSize    int      `json:"max_size" yaml:"max_size"`       // MB
	MaxBackups int      `json:"max_backups" yaml:"max_backups"` // 保留的备份文件数量
	MaxAge     int      `json:"max_age" yaml:"max_age"`         // 保留天数
	Compress   bool     `json:"compress" yaml:"compress"`       // 是否压缩
}

// Logger 日志接口
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger
}

// logrusLogger logrus实现的日志器
type logrusLogger struct {
	logger *logrus.Logger
}

// NewLogger 创建新的日志器
func NewLogger(config *Config) (Logger, error) {
	logger := logrus.New()
	
	// 设置日志级别
	level, err := logrus.ParseLevel(string(config.Level))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	
	// 设置日志格式
	switch config.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	}
	
	// 设置输出
	switch config.Output {
	case "stderr":
		logger.SetOutput(os.Stderr)
	case "file":
		if config.FilePath != "" {
			file, err := os.OpenFile(config.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				return nil, err
			}
			logger.SetOutput(file)
		}
	default:
		logger.SetOutput(os.Stdout)
	}
	
	return &logrusLogger{logger: logger}, nil
}

// 实现Logger接口的方法
func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogger{logger: l.logger.WithField(key, value).Logger}
}

func (l *logrusLogger) WithFields(fields map[string]interface{}) Logger {
	return &logrusLogger{logger: l.logger.WithFields(logrus.Fields(fields)).Logger}
}

func (l *logrusLogger) WithError(err error) Logger {
	return &logrusLogger{logger: l.logger.WithError(err).Logger}
}

func (l *logrusLogger) WithContext(ctx context.Context) Logger {
	return &logrusLogger{logger: l.logger.WithContext(ctx).Logger}
}

// DefaultLogger 默认日志器
var DefaultLogger Logger

func init() {
	config := &Config{
		Level:  InfoLevel,
		Format: "text",
		Output: "stdout",
	}
	
	var err error
	DefaultLogger, err = NewLogger(config)
	if err != nil {
		panic(err)
	}
}

// SetDefaultLogger 设置默认日志器
func SetDefaultLogger(logger Logger) {
	DefaultLogger = logger
}