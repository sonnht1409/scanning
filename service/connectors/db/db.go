package db

import (
	"fmt"

	"github.com/sonnht1409/scanning/service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Values.DB.User,
		config.Values.DB.Password, config.Values.DB.Address, config.Values.DB.Port, config.Values.DB.Name)
	logLevel := gormLogger.Error
	if config.Values.Env != "prod" {
		logLevel = gormLogger.Info
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	err = sqlDB.Ping()
	if err != nil {
		return nil
	}
	return db
}
