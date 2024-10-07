package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mine/common/r"
	"mine/internal/db/model"
	"mine/internal/web/service"
)

type UserService interface {
	Register(user *model.User) error
	Login(user *model.User) (string, error)
}

type UserHandler struct {
	us UserService
}

func (uh *UserHandler) Login(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("参数错误"))
		}
		if jwtStr, err := uh.us.Login(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes(err.Error()))
		} else {
			r.Return(ctx, r.OK().SetData(jwtStr))
		}
	})
}

func (uh *UserHandler) Register(router *gin.RouterGroup) {
	router.POST("/register", func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("参数错误"))
		}
		if err := uh.us.Register(&user); err != nil {
			r.Return(ctx, r.Fail().SetMes("注册失败："+err.Error()))
		} else {
			r.Return(ctx, r.OK().SetMes("注册成功"))
		}
	})
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		us: service.NewUserService(db),
	}
}
