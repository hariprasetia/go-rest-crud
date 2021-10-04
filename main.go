package main

/**
 * Created by GoLand.
 * Author: hariprasetia
 * Date: 2021-10-02
 * Time: 10:07:33
 */

import (
	"github.com/joho/godotenv"
	"go-rest-crud/datasources/mysql"
	"go-rest-crud/models"
	"go-rest-crud/routers"
	"log"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Panic(err.Error(), ": env load error")
	}

	mysql.Init()
	models.Migration()

	routers.Init()
}
