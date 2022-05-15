package bootstrap

import (
	"errors"
	"fmt"
	"goblog/app/models/user"
	"goblog/pkg/config"
	"goblog/pkg/database"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=True&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)

		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.max_idle_connections"))
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.max_open_connections"))
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.max_connection_lifetime")) * time.Second)

	database.DB.AutoMigrate(&user.User{})
}
