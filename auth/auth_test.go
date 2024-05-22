package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
