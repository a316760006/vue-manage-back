package bootstrap

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"
	"vue-manage-back/app/models"
	"vue-manage-back/global"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 自定义 gorm Writer
// 自定义 Logger 就已经实现了，这里只简单替换了 logger.Writer 的实现
func InitializeDB() *gorm.DB {
	// 根据驱动配置进行初始化
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()
	}
}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
		models.Media{},
	)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}

// 初始化 mysql gorm.DB
func initMySqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Database

	// 如果yaml中没有配置数据库名称,则结束初始化连接
	if dbConfig.Database == "" {
		return nil
	}
	// 字符串拼接成连接语句
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	// 数据库设置
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	// 打开连接
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   getGormLogger(), // 使用自定义 Logger
	}); err != nil {
		global.App.Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initMySqlTables(db)
		return db
	}
}

// 切换默认 Logger 使用的 Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Database.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

// 自定义 gorm Writer
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // 慢 SQL 阈值
		LogLevel:                  logMode,                                         // 日志级别
		IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter, // 禁用彩色打印
	})
}
