package handler

import (
	"github.com/gin-gonic/gin"
	"mine/common/r"
	"mine/internal/db/model"
)

type UserService interface {
	Register(user *model.User) error
}

type UserHandler struct {
	US UserService
}

func (uh *UserHandler) Register(router *gin.RouterGroup) {
	router.POST("/register", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("参数错误"))
		}
		if err := uh.US.Register(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("注册失败："+err.Error()))
		} else {
			r.Return(ctx, r.OK().SetMes("注册成功"))
		}
	})
}
