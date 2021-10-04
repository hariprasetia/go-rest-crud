package models

import (
	"go-rest-crud/datasources/mysql"
	"log"
)

func Migration() {
	if !mysql.Connect.Migrator().HasTable(&User{}) {
		err := mysql.Connect.Migrator().CreateTable(&User{})
		if err != nil {
			log.Println(err.Error())
		}
	}
}
