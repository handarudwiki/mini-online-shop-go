package database

import (
	"fmt"
	"handarudwiki/mini-online-shop-go/internal/config"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {
	// Connect to postgresdsn
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		db = nil
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection))
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection))
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIddleConnection))
	return
}
