package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	env, errEnv := GetEnv()
	if errEnv != nil {
		log.Fatal(errEnv)
	}
	gin.SetMode(env.Mode)
	router := gin.Default()

	// rotues
	router.POST("/assets-api/upload", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "test",
		})
	})

	// listen on port
	err := router.Run(":" + env.Port)
	if err != nil {
		log.Fatal("server start ERROR:", err)
	} else {
		log.Println("server is running")
	}
}
