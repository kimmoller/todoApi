package api

import (
	"todoApi/auth"

	"github.com/gin-gonic/gin"
)

func authorizeRequest(ctx *gin.Context) error {
	identityHeader := ctx.GetHeader("Identity")
	authHeader := ctx.GetHeader("Authorization")
	return auth.ValidateRequest(ctx, identityHeader, authHeader)
}
