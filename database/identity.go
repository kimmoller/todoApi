package database

import (
	"fmt"
	"todoApi/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) InsertIdentity(ctx *gin.Context, identity model.Identity) error {
	query := "INSERT INTO identity(username, password) VALUES(@username, @password)"
	args := pgx.NamedArgs{
		"username": identity.Username,
		"password": identity.Password,
	}

	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error while inserting identity %s, %w", identity.Username, err)
	}
	return nil
}
