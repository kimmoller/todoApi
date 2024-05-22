package database

import (
	"context"
	"fmt"
	"todoApi/model"

	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) InsertIdentity(ctx context.Context, identity model.Identity) error {
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

func (pg *Postgres) GetIdentity(ctx context.Context, username string) (*model.Identity, error) {
	query := "SELECT * FROM identity WHERE username=@username"
	args := pgx.NamedArgs{
		"username": username,
	}

	row := pg.db.QueryRow(ctx, query, args)
	var identity model.Identity
	err := row.Scan(&identity.ID, &identity.Username, &identity.Password)
	if err != nil {
		return nil, fmt.Errorf("error while fetching identity %s, %w", username, err)
	}
	return &identity, nil
}
