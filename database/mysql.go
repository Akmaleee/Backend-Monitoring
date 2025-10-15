package database

import (
	"it-backend/internal/helper"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseMySQL struct {
	DBInfra *gorm.DB
}

func NewDatabaseMySQL(config *helper.Config) DatabaseMySQL {
	dbInfra := newConnectionMySQL(config, config.INFRASTRUCTURE_MYSQL_DSN)

	return DatabaseMySQL{
		DBInfra: dbInfra,
	}
}

func newConnectionMySQL(config *helper.Config, dsn string) *gorm.DB {
	dialect := mysql.Open(dsn)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	}

	if config.APP_ENVIRONMENT != "production" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		log.Printf("Failed to connect to database: %s", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
