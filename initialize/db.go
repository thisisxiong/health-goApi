package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"health/global"
	"log"
	"os"
	"time"
)

func Initdb() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Conf.Mysql.User,
		global.Conf.Mysql.Password,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Dbname)
	logLevel := logger.Info //日志级别
	switch global.Conf.Mysql.Level {
	case 1:
		logLevel = logger.Silent
		break
	case 2:
		logLevel = logger.Error
		break
	case 3:
		logLevel = logger.Warn
		break
	case 4:
		logLevel = logger.Info
		break
	}

	logOutPUt := os.Stdout //日志输出
	if !global.Conf.Debug && global.Conf.LogPath != "" {
		err := mkdir(global.Conf.LogPath)
		if err != nil {
			panic(err)
		}

		path := global.Conf.LogPath + "/" + time.Now().Format("20060102") + "_sql.log"
		f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logOutPUt = f
	}
	newLogger := logger.New(
		log.New(logOutPUt, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logLevel,    // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//用于处理invalid connection 错误 长时间没有操作 会导致数据库连接等待超时 这个设置时间要小于数据库配置的wait_timeout
	sqlDb, _ := db.DB()
	sqlDb.SetConnMaxLifetime(time.Minute)
	global.Db = db

	//初始化表结构
	//db.AutoMigrate(&models.Health{}, &models.User{})
}

func mkdir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return err
}
