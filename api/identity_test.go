package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoApi/database"
	"todoApi/model"
	"todoApi/testutils"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

var (
	dbUrl    = "postgres://user:password@localhost:5432/postgres"
	filePath = "../database/testMigrations"
)

func TestIdentityCreate(t *testing.T) {
	container, err := testutils.GetTestContainer()
	require.NoError(t, err)

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	_, err = database.NewPG(context.Background(), dbUrl)

	require.NoError(t, err)

	router := GetApi()

	identity := model.Identity{
		Username: "test",
		Password: "test",
	}
	jsonValue, _ := json.Marshal(identity)
	req, err := http.NewRequest("POST", "/identity", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
}
