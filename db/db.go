package db

import (
	"fmt"
	"github.com/dwisulfahnur/todo-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var db *gorm.DB
var err error


func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Init() {
	user := getEnv("DB_USER", "dev")
	password := getEnv("DB_PASSWORD", "dev")
	dbname := getEnv("DB_NAME", "todoapiGo")
	//host := getEnv("DB_HOST", "127.0.0.1")
	//port := getEnv("DB_PORT", "3306")

	dbinfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		dbname,
	)
	fmt.Println(dbinfo)

	db, err = gorm.Open("mysql", dbinfo)
	if err != nil {
		log.Println("Couldn't connect to database")
		panic(err)
	}

	//Database check Process
	if !db.HasTable(&models.Task{}) {
		err := db.CreateTable(&models.Task{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate()
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}