package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SQL *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbc := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	SQL, err = gorm.Open(mysql.Open(dbc), &gorm.Config{})
	if err != nil {
		panic("Cant connect to database")
	}
	//SQL = SQL.Set("gorm:table_options", "COLLATE=utf8_bin")

	fmt.Println("Database connection established")
}
