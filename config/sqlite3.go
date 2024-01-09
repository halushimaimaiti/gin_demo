package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlLiteDB *gorm.DB

// InitDB 初始化数据库连接
func InitSqlLite3() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 将日志输出到标准输出
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 开启彩色打印
		},
	)
	// file, err := os.OpenFile("gorm_sql.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	logrus.SetOutput(file)
	// } else {
	// 	logrus.Info("Failed to log to file, using default stderr")
	// }
	SqlLiteDB, err = gorm.Open(sqlite.Open("D:\\doctor\\gin_demo\\db.sqlite3"), &gorm.Config{
		// Logger: logger.New(
		// 	logrus.New(), // 使用logrus实例
		// 	logger.Config{
		// 		SlowThreshold:             200,         // 慢查询阈值
		// 		LogLevel:                  logger.Info, // 日志级别
		// 		IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
		// 		Colorful:                  true,        // 彩色打印
		// 	},
		// ),
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
}
