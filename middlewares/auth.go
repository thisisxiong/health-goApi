package middlewares

import (
	"github.com/gin-gonic/gin"
	"health/api"
	"health/models"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("x-token")
		claims, err := models.ParseToken(token)
		if err != nil {
			api.Fail(ctx, "token 验证失败")
			ctx.Abort()
			return
		}
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			api.Fail(ctx, "token 已失效")
			ctx.Abort()
			return
		}

		ctx.Set("uid", uint32(claims.Id))
		ctx.Set("phone", claims.Phone)
		ctx.Next()
	}
}
