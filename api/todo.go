package api

import (
	"log"
	"net/http"
	"todoApi/database"
	"todoApi/model"

	"github.com/gin-gonic/gin"
)

func getTodos(ctx *gin.Context) {
	err := authorizeRequest(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	identityId := ctx.Param("identityId")

	todos, err := database.Instance.FetchTodos(ctx, identityId)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, todos)
}

func createTodo(ctx *gin.Context) {
	err := authorizeRequest(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	var newTodo model.Todo

	if err := ctx.BindJSON(&newTodo); err != nil {
		log.Println(err)
		return
	}

	err = database.Instance.InsertTodo(ctx, newTodo)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodo(ctx *gin.Context) {
	err := authorizeRequest(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	id := ctx.Param("id")
	var status string

	if err := ctx.BindJSON(&status); err != nil {
		log.Println(err)
		return
	}

	err = database.Instance.UpdateTodo(ctx, id, status)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, updateTodo)
}

func deleteTodo(ctx *gin.Context) {
	err := authorizeRequest(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	todoId := ctx.Param("id")

	err = database.Instance.DeleteTodo(ctx, todoId)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, todoId)
}
