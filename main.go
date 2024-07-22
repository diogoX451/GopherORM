package main

import (
	"flag"

	"github.com/diogoX451/gopherORM/internal/database/migrations"
	"github.com/diogoX451/gopherORM/internal/providers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	tableName := flag.String("migrations", "default", "Run migrations")
	migrate := flag.Bool("migrate", false, "Run migrations")
	flag.Parse()

	if *tableName != "default" {
		cm := migrations.NewCommand(*tableName, "default")
		cm.Run()
	}

	dbProvider := providers.NewDatabaseProviders()

	if *migrate {
		mg := migrations.NewMigrations(dbProvider.Connect())
		mg.Run()
	}
}
