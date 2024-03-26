package middleware

import (
	"desafioFinalGo/pkg/web"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")

		if token == "" {
			web.FailureResponse(ctx, 401, errors.New("Token no encontrado"))
			ctx.Abort()
			return
		}

		if token != os.Getenv("TOKEN") {
			web.FailureResponse(ctx, 401, errors.New("Token inválido"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
