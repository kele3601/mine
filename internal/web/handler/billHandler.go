package handler

import (
	"github.com/gin-gonic/gin"
	"mine/common/r"
	"mine/internal/db/model"
)

type BillService interface {
	Account(bill *model.Bill) error
}

func (h *Handler) Account(router *gin.RouterGroup) {
	router.POST("/account", func(ctx *gin.Context) {
		claims, err := h.us.CheckUserByClaims(ctx)
		if nil != err {
			r.Return(ctx, r.Fail().SetMes(err.Error()))
			ctx.Abort()
			return
		}
		var bill model.Bill
		if err := ctx.ShouldBindJSON(&bill); nil != err {
			r.Return(ctx, r.Fail().SetMes("参数错误"))
			ctx.Abort()
			return
		}
		bill.User = claims.UserID
		if err := h.bs.Account(&bill); nil != err {
			r.Return(ctx, r.Fail().SetMes(err.Error()))
		} else {
			r.Return(ctx, r.OK())
		}
	})
}
