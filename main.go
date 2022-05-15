package main

import (
	"fmt"
	"goblog/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	bootstrap.SetupRoute(router)
	err := router.Run(":8888")

	if err != nil {
		fmt.Println(err.Error())
	}
}
