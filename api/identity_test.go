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
)

var (
	dbUrl    = "postgres://user:password@localhost:5432/postgres"
	filePath = "../database/testMigrations"
)

func TestIdentityCreate(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	identity := model.Identity{
		Username: "newTestIdentity",
		Password: "newPassword",
	}
	jsonValue, _ := json.Marshal(identity)
	req, _ := http.NewRequest("POST", "/identity", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	dbIdentity, _ := database.Instance.GetIdentity(context.Background(), "newTestIdentity")

	assert.NotEqual(t, identity.Password, dbIdentity.Password)
	assert.True(t, auth.CheckPasswordHash(identity.Password, dbIdentity.Password))
}

func TestUpdatePassword(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	updateDto := IdentityUpdateDto{
		Password: "newPassword",
	}
	jsonValue, _ := json.Marshal(updateDto)

	req, _ := http.NewRequest("PATCH", "/identity/1", bytes.NewBuffer(jsonValue))
	addAuthHeaders(req, "identityToUpdate")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	_, err := auth.Login(context.Background(), "identityToUpdate", "newPassword")
	assert.Nil(t, err)
}

func TestDeleteIdentity(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())

	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	req, _ := http.NewRequest("DELETE", "/identity/2", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	_, err := database.Instance.GetIdentity(context.Background(), "identityToDelete")
	assert.NotNil(t, err)
}
