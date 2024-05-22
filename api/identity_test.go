package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoApi/auth"
	"todoApi/database"
	"todoApi/model"
	"todoApi/testutils"

	"github.com/stretchr/testify/assert"
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
		Username: "newTestIdentity",
		Password: "newPassword",
	}
	jsonValue, _ := json.Marshal(identity)
	req, err := http.NewRequest("POST", "/identity", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)

	dbIdentity, err := database.Instance.GetIdentity(context.Background(), "newTestIdentity")
	require.NoError(t, err)

	assert.NotEqual(t, identity.Password, dbIdentity.Password)
	assert.True(t, auth.CheckPasswordHash(identity.Password, dbIdentity.Password))
}
