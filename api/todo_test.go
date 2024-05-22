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

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestListTodosForUser(t *testing.T) {
	container, err := testutils.GetTestContainer()
	require.NoError(t, err)

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	_, err = database.NewPG(context.Background(), dbUrl)

	require.NoError(t, err)

	router := GetApi()

	req, err := http.NewRequest("GET", "/todo/identity/3", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var todos = []model.Todo{}
	body, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	err = json.Unmarshal(body, &todos)
	require.NoError(t, err)
	assert.Equal(t, 3, len(todos))
}

func TestTodoCreate(t *testing.T) {
	container, err := testutils.GetTestContainer()
	require.NoError(t, err)

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	_, err = database.NewPG(context.Background(), dbUrl)

	require.NoError(t, err)

	router := GetApi()

	todo := model.Todo{
		Task:       "new todo",
		IdentityId: 4,
	}

	jsonValue, _ := json.Marshal(todo)
	req, err := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonValue))
	require.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	todos, _ := database.Instance.FetchTodos(context.Background(), "4")
	assert.Equal(t, 1, len(todos))
}

func TestUpdateTodo(t *testing.T) {
	container, err := testutils.GetTestContainer()
	require.NoError(t, err)

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	_, err = database.NewPG(context.Background(), dbUrl)

	require.NoError(t, err)

	router := GetApi()

	jsonValue, _ := json.Marshal("COMPLETED")
	req, err := http.NewRequest("PATCH", "/todo/4", bytes.NewBuffer(jsonValue))
	require.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	todos, _ := database.Instance.FetchTodos(context.Background(), "5")
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, "COMPLETED", todos[0].Status)
}

func TestTodoDelete(t *testing.T) {
	container, err := testutils.GetTestContainer()
	require.NoError(t, err)

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	_, err = database.NewPG(context.Background(), dbUrl)

	require.NoError(t, err)

	router := GetApi()

	req, err := http.NewRequest("DELETE", "/todo/5", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	todos, _ := database.Instance.FetchTodos(context.Background(), "6")
	assert.Equal(t, 0, len(todos))
}
