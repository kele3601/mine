package web

import (
	"github.com/gin-gonic/gin"
	"mine/internal/app"
	"mine/internal/web/handler"
	"mine/internal/web/service"
	"net/http"
)

func InitRouter(app *app.App) {
	router := app.Web.Group("v1")
	// 通用接口
	group := router.Group("public")
	ur := handler.UserHandler{US: service.NewUserService(app.DB)}
	{
		hello(group)
		ur.Register(group)
	}
}

func hello(router *gin.RouterGroup) {
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})
}
