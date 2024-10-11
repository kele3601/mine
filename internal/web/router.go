package web

import (
	"github.com/gin-gonic/gin"
	"mine/internal/app"
	h "mine/internal/web/handler"
	"mine/internal/web/middleware"
	"net/http"
)

func InitRouter(app *app.App) {
	router := app.Web.Group("v1")
	// 通用接口
	public := router.Group("public")
	// 内部接口
	protected := router.Group("protected")
	protected.Use(middleware.AuthMiddleware())
	handler := h.NewHandler(app.DB)
	{
		// 注册API
		hello(public)
		handler.Register(public)
		handler.Login(public)
		handler.Account(protected)
		handler.List(protected)
	}
}

func hello(router *gin.RouterGroup) {
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})
}
