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
	List() ([]*model.User, error)
}

func (h *Handler) UserList(router *gin.RouterGroup) {
	router.POST("/user/list", func(ctx *gin.Context) {
		if users, err := h.us.List(); err != nil {
			r.Return(ctx, r.Fail().SetMes(err.Error()))
		} else {
			r.Return(ctx, r.OK().SetData(users))
		}
	})
}

func (h *Handler) Login(router *gin.RouterGroup) {
	router.POST("/user/login", func(ctx *gin.Context) {
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
	router.POST("/user/register", func(ctx *gin.Context) {
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
