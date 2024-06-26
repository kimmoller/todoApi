package api

import "github.com/gin-gonic/gin"

func GetApi() *gin.Engine {

	router := gin.Default()

	router.POST("/auth", login)

	router.POST("/identity", createIdentity)
	router.PATCH("/identity/:id", changePassword)
	router.DELETE("/identity/:id", deleteIdentity)

	router.GET("/todo/identity/:identityId", getTodos)
	router.POST("/todo", createTodo)
	router.PATCH("/todo/:id", updateTodo)
	router.DELETE("/todo/:id", deleteTodo)

	return router
}
