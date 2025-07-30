package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Noname2812/go-ecommerce-backend-api/pkg/logger"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
)

func NewMySqlC(config setting.MySQLSetting, logger *logger.LoggerZap) *sql.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.Dbname)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic("Init MySql Failed")
	}

	// Configure connection pool from config
	if config.MaxOpenConns > 0 {
		db.SetMaxOpenConns(config.MaxOpenConns)
	} else {
		db.SetMaxOpenConns(100) // Default for high concurrency
	}

	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	} else {
		db.SetMaxIdleConns(25) // Default idle connections
	}

	if config.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)
	} else {
		db.SetConnMaxLifetime(5 * time.Minute) // Default 5 minutes
	}

	logger.Logger.Info("Initializing MySQL Successfully with connection pool config",
		zap.Int("maxOpenConns", config.MaxOpenConns),
		zap.Int("maxIdleConns", config.MaxIdleConns),
		zap.Int("connMaxLifetime", config.ConnMaxLifetime))
	return db
}
