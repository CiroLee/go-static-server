package main

import (
	"github.com/CiroLee/go-static-server/middlemare"
	"github.com/CiroLee/go-static-server/upload"
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
	router.MaxMultipartMemory = 10 << 20 // 10Mib
	router.StaticFS("/statics", http.Dir("./statics"))

	router.POST("/statics/api/upload", middlemare.Authorization(), upload.ImageUploadHandler)

	// listen on port
	err := router.Run(":" + env.Port)
	if err != nil {
		log.Fatal("server start ERROR:", err)
	}
}
