// * Database intialization and configuration
package database

import (
	"finalAssing/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(cfg config.Config) (*gorm.DB, error) {
	// dsn := "host=localhost user=postgres password=admin dbname=finalAssing port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DbConfig.DbHost, cfg.DbConfig.DbUser, cfg.DbConfig.DbPassword, cfg.DbConfig.DbName, cfg.DbConfig.DbPort, cfg.DbConfig.DbSSLMode, cfg.DbConfig.DbTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
