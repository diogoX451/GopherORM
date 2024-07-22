package database

import "os"

func InitializeDatabaseFactory() Database {
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		return NewPostgres()
	default:
		return nil
	}
}
