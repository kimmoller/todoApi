package database

import (
	"context"
	"fmt"
	"todoApi/model"

	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) FetchTodos(ctx context.Context, identityId string) ([]model.Todo, error) {
	query := "SELECT * FROM todo WHERE identityId=@identityId"
	args := pgx.NamedArgs{
		"identityId": identityId,
	}

	rows, err := pg.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user %s todos, %w", identityId, err)
	}

	todos, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Todo])
	if err != nil {
		return nil, fmt.Errorf("error while collecting rows for user %s, %w", identityId, err)
	}

	return todos, nil
}

func (pg *Postgres) InsertTodo(ctx context.Context, todo model.Todo) error {
	query := "INSERT INTO todo (task, status, identityId) VALUES (@task, 'ONGOING', @identityId)"
	args := pgx.NamedArgs{
		"task":       todo.Task,
		"identityId": todo.IdentityId,
	}

	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error while creating todo %w", err)
	}
	return nil
}

func (pg *Postgres) UpdateTodo(ctx context.Context, id string, status string) error {
	query := "UPDATE todo SET status=@status WHERE id=@id"
	args := pgx.NamedArgs{
		"status": status,
		"id":     id,
	}

	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error while updating todo %s, %w", id, err)
	}
	return nil
}

func (pg *Postgres) DeleteTodo(ctx context.Context, id string) error {
	query := "DELETE FROM todo WHERE id=@id"
	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error while deleting todo %s, %w", id, err)
	}

	return nil
}
