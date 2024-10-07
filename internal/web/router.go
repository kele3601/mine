package web

import (
	"github.com/gin-gonic/gin"
	"mine/internal/app"
	"mine/internal/web/handler"
	"net/http"
)

func InitRouter(app *app.App) {
	router := app.Web.Group("v1")
	// 通用接口
	group := router.Group("public")
	ur := handler.NewUserHandler(app.DB)
	{
		hello(group)
		ur.Register(group)
		ur.Login(group)
	}
}

func hello(router *gin.RouterGroup) {
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})
}
