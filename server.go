package main

import (
	"context"

	"todoApi/api"
	"todoApi/database"
)

func main() {
	dbUrl := "postgres://user:password@localhost:5432/postgres"
	database.Migratedb(dbUrl, "database/migrations")
	database.NewPG(context.Background(), dbUrl)

	router := api.GetApi()
	router.Run(":8080")
}
