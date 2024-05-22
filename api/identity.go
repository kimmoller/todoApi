package api

import (
	"log"
	"net/http"

	"todoApi/auth"
	"todoApi/database"
	"todoApi/model"

	"github.com/gin-gonic/gin"
)

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
