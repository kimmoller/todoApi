package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoApi/database"
	"todoApi/model"
	"todoApi/testutils"

	"github.com/stretchr/testify/assert"
)

func TestListTodosForUser(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	req, _ := http.NewRequest("GET", "/todo/identity/3", nil)
	addAuthHeaders(req, "todoFetch")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var todos = []model.Todo{}
	body, _ := io.ReadAll(w.Body)

	json.Unmarshal(body, &todos)
	assert.Equal(t, 3, len(todos))
}

func TestTodoCreate(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	todo := model.Todo{
		Task:       "new todo",
		IdentityId: 4,
	}

	jsonValue, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonValue))
	addAuthHeaders(req, "todoCreate")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	todos, _ := database.Instance.FetchTodos(context.Background(), "4")
	assert.Equal(t, 1, len(todos))
}

func TestUpdateTodo(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	jsonValue, _ := json.Marshal("COMPLETED")
	req, _ := http.NewRequest("PATCH", "/todo/4", bytes.NewBuffer(jsonValue))
	addAuthHeaders(req, "todoUpdate")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	todos, _ := database.Instance.FetchTodos(context.Background(), "5")
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, "COMPLETED", todos[0].Status)
}

func TestTodoDelete(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	req, _ := http.NewRequest("DELETE", "/todo/5", nil)
	addAuthHeaders(req, "todoDelete")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	todos, _ := database.Instance.FetchTodos(context.Background(), "6")
	assert.Equal(t, 0, len(todos))
}
