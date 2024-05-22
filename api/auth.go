package api

import (
	"net/http"
	"todoApi/auth"

	"github.com/gin-gonic/gin"
)

type UserAuth struct {
	Username string
	Password string
}

func login(ctx *gin.Context) {
	var userAuth UserAuth

	if err := ctx.BindJSON(&userAuth); err != nil {
		return
	}

	token, err := auth.Login(ctx, userAuth.Username, userAuth.Password)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, "Invalid username or password")
		return
	}
	ctx.Writer.Header().Set("Identity", userAuth.Username)
	ctx.Writer.Header().Set("Authorization", token)
	ctx.IndentedJSON(http.StatusOK, nil)
}

func authorizeRequest(ctx *gin.Context) error {
	identityHeader := ctx.GetHeader("Identity")
	authHeader := ctx.GetHeader("Authorization")
	return auth.ValidateRequest(ctx, identityHeader, authHeader)
}
