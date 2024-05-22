package api

import (
	"log"
	"net/http"

	"todoApi/auth"
	"todoApi/database"
	"todoApi/model"

	"github.com/gin-gonic/gin"
)

type IdentityUpdateDto struct {
	Password string
}

func createIdentity(ctx *gin.Context) {
	var identity model.Identity

	if err := ctx.BindJSON(&identity); err != nil {
		log.Print(err)
		return
	}

	if !auth.ValidatePassword(identity.Password) {
		ctx.IndentedJSON(http.StatusBadRequest, "Password is too long")
		return
	}

	hash, err := auth.HashPassword(identity.Password)
	if err != nil {
		log.Println(err)
		return
	}

	identity.Password = hash
	err = database.Instance.InsertIdentity(ctx, identity)
	if err != nil {
		log.Print(err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, identity)
}

func changePassword(ctx *gin.Context) {
	log.Println("Change password")
	err := authorizeRequest(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	identityId := ctx.Param("id")
	var identityUpdateDto IdentityUpdateDto

	if err := ctx.BindJSON(&identityUpdateDto); err != nil {
		log.Println(err)
		return
	}

	hash, err := auth.HashPassword(identityUpdateDto.Password)
	if err != nil {
		log.Println(err)
		return
	}

	err = database.Instance.UpdateIdentityPassword(ctx, identityId, hash)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, nil)
}

func deleteIdentity(ctx *gin.Context) {
	err := authorizeRequest(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	identityId := ctx.Param("id")

	err = database.Instance.DeleteIdentity(ctx, identityId)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, nil)
}
