package api

import (
	"log"
	"net/http"

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

	err := database.Instance.InsertIdentity(ctx, identity)
	if err != nil {
		log.Print(err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, identity)
}
