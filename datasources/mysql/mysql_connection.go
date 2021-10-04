package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var Connect *gorm.DB

func Init() {
	var err error
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"))

	Connect, err = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
	}})

	if err != nil {
		log.Fatal("failed to initialize connection to mysql")
	} else {
		log.Println("initialize connection to mysql")
	}

}

