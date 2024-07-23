package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func ConnectDataBase() {

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	// "mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(DbDriver, DBURL)

	if err != nil {
		log.Fatalln("Data Base connection error", err)
	} else {
		fmt.Println("Data Base connection success")
		fmt.Println("Connection -> ", DbDriver)
	}

	DB.AutoMigrate(&User{})

}
