package auth

import (
	"context"
	"testing"
	"time"
	"todoApi/database"
	"todoApi/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	dbUrl    = "postgres://user:password@localhost:5432/postgres"
	filePath = "../database/testMigrations"
)

func TestValidateToken(t *testing.T) {
	token, err := generateJWT("username", time.Now().Add(5*time.Minute))
	require.NoError(t, err)
	valid, err := ValidateToken(token)
	require.NoError(t, err)
	assert.True(t, valid.Token.Valid)
}

func TestExpiredToken(t *testing.T) {
	token, err := generateJWT("username", time.Now().Add(1*time.Second))
	require.NoError(t, err)
	time.Sleep(2 * time.Second)
	_, err = ValidateToken(token)
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	container, _ := testutils.GetTestContainer()

	defer container.Terminate(context.Background())
	database.Migratedb(dbUrl, filePath)
	database.NewPG(context.Background(), dbUrl)

	_, err := Login(context.Background(), "loginTest", "loginTest")

	assert.Nil(t, err)
}
