package main

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	memberController := controller.NewMemberController()
	r.POST("/member/register", memberController.Register)
	r.GET("/member/list", memberController.List)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
