package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoApi/database"
	"todoApi/testutils"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())
	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	router := GetApi()

	userAuth := UserAuth{
		Username: "loginTest",
		Password: "loginTest",
	}
	jsonValue, _ := json.Marshal(userAuth)
	req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	headers := w.Header()

	identityHeader := headers.Get("Identity")
	assert.Equal(t, "loginTest", identityHeader)

	token := headers.Get("Authorization")
	assert.NotEmpty(t, token)
}
