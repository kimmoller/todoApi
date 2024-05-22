package api

import (
	"context"
	"net/http"
	"todoApi/auth"
)

func addAuthHeaders(req *http.Request, username string) {
	token, _ := auth.Login(context.Background(), username, username)
	req.Header.Add("Identity", username)
	req.Header.Add("Authorization", "Bearer "+token)
}
