package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlDB struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionMySql() (*MySqlDB, error) {
	// Example: user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySql.User,
		cfg.MySql.Password,
		cfg.MySql.Host,
		cfg.MySql.Port,
		cfg.MySql.DBName,
	)

	db, err := gorm.Open(mysql.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Printf("[ConnectionMySqlDB-1] failed to connect to postgres database %s: %v", cfg.MySql.Host, err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("[ConnectionMySqlDB-2] failed to get database instance: %v", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MySql.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MySql.DBMaxIdleConns)

	return &MySqlDB{DB: db}, nil
}
