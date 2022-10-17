package api

import (
	"fmt"
	"os"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var apiEnginge *gin.Engine

func createRouter() {
	apiEnginge = gin.Default()
	apiEnginge.Use(cors.Default())

	apiEnginge.GET("/file/:id/:encryptionKey/:name", GetFile)
	apiEnginge.POST("/template/:name", SubmitTemplate)

	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func StartServer() {
	createRouter()
	listenHost := os.Getenv("LISTEN_HOST")
	if listenHost == "" {
		listenHost = config.CurrentConfig.Server.Listen.Host
	}
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		listenPort = config.CurrentConfig.Server.Listen.Port
	}
	listen := fmt.Sprintf("%s:%s", listenHost, listenPort)
	fmt.Printf("Server is running on http://%s\n", listen)
	apiEnginge.Run(listen)
}
