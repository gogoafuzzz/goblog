package main

import (
	"flag"
	"fmt"
	"goblog/bootstrap"
	btsConfig "goblog/config"
	"goblog/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加載 .env 檔案，如 --env=testing 加載的是 .env.testing 檔案")
	flag.Parse()
	config.InitConfig(env)
	router := gin.New()

	bootstrap.SetupDB()
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))

	if err != nil {
		fmt.Println(err.Error())
	}
}
