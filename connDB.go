package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	connDB *gorm.DB
	err    error
)

type DB struct{}

func NewDatabase() *DB {

	DB_NAME := os.Getenv(`DB_NAME`)
	DB_USER := os.Getenv(`DB_USER`)
	DB_PASS := os.Getenv(`DB_PASS`)
	DB_HOST := os.Getenv(`DB_HOST`)

	dsn := fmt.Sprintf(`%[1]s:%[2]s@tcp(%[3]s:3306)/%[4]s?charset=utf8mb4&parseTime=True&loc=Asia%%2FHong_Kong`, DB_USER, DB_PASS, DB_HOST, DB_NAME)
	// fmt.Println(dsn)
	connDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &DB{}
}

func (db *DB) MigrateTables() {

	// connDB.AutoMigrate(&ProductVouchers{})

}

func (db *DB) ConnDB() *gorm.DB {

	if os.Getenv(`DEBUG`) == `Y` {
		return connDB.Debug()
	}

	sqlDB, _ := connDB.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute)

	return connDB
}
