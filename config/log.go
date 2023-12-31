package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func InitLog() {
	Log = logrus.New()

	// 设置 logrus 日志级别
	Log.SetLevel(logrus.InfoLevel)

	// 创建日志分割器 Lumberjack
	infoLog := &lumberjack.Logger{
		Filename:   "D:/gin_test/logs/info.log", // Info 级别日志文件
		MaxSize:    100,                         // 单个日志文件最大大小，单位 MB
		MaxBackups: 3,                           // 最大保留的日志文件数
		MaxAge:     30,                          // 保留日志的最大天数
		Compress:   true,                        // 是否压缩日志
	}

	errorLog := &lumberjack.Logger{
		Filename:   "D:/gin_test/logs/error.log", // Error 级别日志文件
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}

	// 设置日志输出为文件
	Log.SetOutput(os.Stdout)
	Log.AddHook(&LogrusWriter{writer: infoLog, level: logrus.InfoLevel})
	Log.AddHook(&LogrusWriter{writer: errorLog, level: logrus.ErrorLevel})
}

// 自定义 logrus 的钩子（Hook），将日志写入文件
type LogrusWriter struct {
	writer *lumberjack.Logger
	level  logrus.Level
}

// Fire 方法将日志写入文件
func (hook *LogrusWriter) Fire(entry *logrus.Entry) error {
	// 根据日志级别写入不同的日志文件
	if entry.Level <= hook.level {
		logEntry, _ := entry.String()
		_, err := hook.writer.Write([]byte(logEntry))
		return err
	}
	return nil
}

// Levels 方法返回支持的日志级别
func (hook *LogrusWriter) Levels() []logrus.Level {
	return logrus.AllLevels[:hook.level+1]
}
