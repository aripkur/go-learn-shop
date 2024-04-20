package database

import (
	"fmt"
	"time"

	"github.com/aripkur/go-learn-shop/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectMysql(cfg config.DBConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		db = nil
		return
	}

	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime((time.Duration(cfg.ConnectionPool.MaxLifetimeConnection) * time.Second))

	return
}
