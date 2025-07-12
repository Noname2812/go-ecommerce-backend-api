package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Noname2812/go-ecommerce-backend-api/pkg/logger"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
)

func NewMySqlC(config setting.MySQLSetting, logger *logger.LoggerZap) *sql.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.Dbname)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic("Init MySql Failed")
	}
	logger.Logger.Info("Initializing MySQL Successfully sql")
	return db
}
