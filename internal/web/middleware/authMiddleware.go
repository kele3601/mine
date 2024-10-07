package middleware

import (
	"github.com/gin-gonic/gin"
	"mine/common/r"
	"mine/common/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if "" == token {
			r.Return(ctx, r.Fail().SetMes("请先登录"))
			ctx.Abort()
			return
		}
		if claims, err := utils.ParseJwtToken(token); nil != err {
			r.Return(ctx, r.Fail().SetMes(err.Error()))
			ctx.Abort()
			return
		} else {
			ctx.Set("claims", claims)
			ctx.Next()
		}
	}
}
