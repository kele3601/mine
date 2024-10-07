package handler

import (
	"github.com/gin-gonic/gin"
	"mine/common/r"
	"mine/common/utils"
	"mine/internal/db/model"
)

type UserService interface {
	Register(user *model.User) error
	Login(user *model.User) (string, error)
	CheckUserByClaims(ctx *gin.Context) (*utils.JwtUserClaims, error)
}

func (h *Handler) Login(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("参数错误"))
		}
		if jwtStr, err := h.us.Login(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes(err.Error()))
		} else {
			r.Return(ctx, r.OK().SetData(jwtStr))
		}
	})
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.POST("/register", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("参数错误"))
		}
		if err := h.us.Register(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("注册失败："+err.Error()))
		} else {
			r.Return(ctx, r.OK().SetMes("注册成功"))
		}
	})
}
