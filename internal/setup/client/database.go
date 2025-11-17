package client

import (
	"log"
	"time"

	"github.com/gunawanpras/oc-app-backend/config"

	_ "github.com/godror/godror"

	"github.com/jmoiron/sqlx"
)

func InitOracle(conf *config.Config) *sqlx.DB {
	oracle, err := sqlx.Connect(conf.Oracle.Primary.Dialect, conf.Oracle.Primary.ConnString)
	if err != nil {
		log.Panic("failed to open oracle client for service:", err)
	}

	if err := oracle.Ping(); err != nil {
		log.Panic("failed to ping oracle: %w", err)
	}

	oracle.SetMaxOpenConns(conf.Oracle.Primary.MaxOpenConn)
	oracle.SetMaxIdleConns(conf.Oracle.Primary.MaxIdleConn)
	oracle.SetConnMaxLifetime(time.Second * time.Duration(conf.Oracle.Primary.MaxConnLifeTimeInSecond))

	log.Println("connected to oracle successfully!")

	return oracle
}
