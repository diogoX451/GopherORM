package database

import (
	"time"
)

type Database interface {
	Connect() error
	SetConnection(n int32)
	SetMinConnections(n int32)
	SetCloseAutomaticConn(timeout time.Duration)
	Insert(query string, params ...interface{}) (interface{}, error)
	Query(query string, params ...interface{}) (interface{}, error)
	Close()
	TableExists(tableName string) (bool, error)
	CreateTable(tableName string, types DatabaseTypes) error
}

type IDatabaseType interface {
	TableName() string
	Id(column string) func(*DatabaseTypes)
	String(column string, length int) func(*DatabaseTypes)
	Bool(column string) func(*DatabaseTypes)
	Int(column string) func(*DatabaseTypes)
	Float(column string) func(*DatabaseTypes)
	Null(column string) func(*DatabaseTypes)
	NotNull(column string) func(*DatabaseTypes)
	PrimaryKey(column string) func(*DatabaseTypes)
	AutoIncrement(name string) func(*DatabaseTypes)
	Unique(column string) func(*DatabaseTypes)
	Timestamp() func(*DatabaseTypes)
}
