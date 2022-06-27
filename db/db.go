package db

import (
	"context"
	"todogql/ent"
	"todogql/ent/migrate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *ent.Client
)

func Init() {
	db, err := ent.Open(dialect.SQLite, "file:db.sqlite?_fk=1")

	if err != nil {
		panic(err)
	}

	db.Schema.Create(context.Background(),
		schema.WithAtlas(true),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)

	Client = db
}
