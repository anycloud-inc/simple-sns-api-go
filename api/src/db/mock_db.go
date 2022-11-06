package db

import (
	"context"
	"simple_sns_api/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTestConnection(t *testing.T) context.Context {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	Client = client
	return context.Background()
}
