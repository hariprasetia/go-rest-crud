package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest-crud/services"
	"log"
	"os"
)

var router = gin.Default()

func Init() {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/create", services.Create)
		v1.PUT("/user/update", services.Update)
		v1.GET("/user/:id", services.FindByID)
		v1.GET("/user", services.FindAll)
		v1.DELETE("/user/delete/:id", services.Delete)
	}

	log.Println("API service well initialized. Let's work!")

	err := router.Run(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err.Error(), "router error")
	}

}


