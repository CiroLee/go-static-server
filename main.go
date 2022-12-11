package main

import (
	"github.com/CiroLee/go-static-server/images"
	"github.com/CiroLee/go-static-server/middlemare"
	"github.com/CiroLee/go-static-server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	env, errEnv := utils.GetEnv()
	if errEnv != nil {
		log.Fatal(errEnv)
	}
	gin.SetMode(env.Mode)
	router := gin.Default()
	router.MaxMultipartMemory = 1000 << 20 // 100Mib
	router.StaticFS("/statics", http.Dir("./statics"))

	imagesGroup := router.Group("/statics/api/images")
	{
		imagesGroup.POST("/upload", middlemare.Authorization(), images.ImageUploadHandler)
		imagesGroup.POST("/list", images.ImageListHandler)
		imagesGroup.POST("/delete", middlemare.Authorization(), images.ImageDeleteHandler)
	}

	// listen on port
	err := router.Run(":" + env.Port)
	if err != nil {
		log.Fatal("server start ERROR:", err)
	}
}
