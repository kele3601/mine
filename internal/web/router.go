package web

import (
	"github.com/gin-gonic/gin"
	"mine/internal/app"
	"net/http"
)

func InitRouter(app *app.App) {
	router := app.Web.Group("v1")
	// 通用接口
	group := router.Group("public")
	{
		hello(group)
	}
}

func hello(router *gin.RouterGroup) {
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})
}
