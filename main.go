package main

import (
	"anyiblog/conf"
	"anyiblog/server"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// 登录拦截中间件
func main() {
	conf.Init()
	// debug release
	runMode := os.Getenv("GIN_MODE")
	if runMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := server.NewRouter()
	if runMode == "debug" {
		err := r.Run(":8080")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := r.RunTLS(":8080", "./cert/ssl.crt", "./cert/ssl.key")
		if err != nil {
			fmt.Println(err)
		}
	}
}
