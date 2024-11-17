package main

import (
	"awesomeProject/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRoutes(r)
	r.Run(":8080") // 运行在指定端口
}
