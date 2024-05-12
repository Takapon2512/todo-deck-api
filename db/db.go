package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	// 環境変数からデータベース接続情報を取得
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	// DSN (Data Source Name) の構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	// GORMを使ってデータベース接続（接続できない時があるので、5回までリトライする機能を実装）
	for i := 0; i < 5; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connected to database")
			return db
		}
		fmt.Println("Retrying to connect to database")
		time.Sleep(5 * time.Second)
	}
	log.Fatalf("Failed to connect to database: %v", err)
	return nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}