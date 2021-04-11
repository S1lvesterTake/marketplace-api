package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/S1lvesterTake/marketplace-api/domain"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //for import mysql
	"github.com/joho/godotenv"
)

var dB *gorm.DB

// DBInit create connection to database
func DBInit() *gorm.DB {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", dbUser, password, host, port, dbName, "Asia%2FJakarta")

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Panicf("failed to connect to database with err : %s ", err)
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)

	db.LogMode(true)

	dB = db
	db.AutoMigrate(
		&domain.Transaction{},
		&domain.TransactionDetail{},
		&domain.Cart{},
		&domain.CartDetail{},
		&domain.Product{},
		&domain.StatusCode{},
	)
	return dB
}

// GetDB getdb
func GetDB() *gorm.DB {
	return dB
}
