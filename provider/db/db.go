package db

import (
	"context"
	"log"

	"github.com/XiroXD/fiber-ent-crud-app/ent"
	_ "github.com/mattn/go-sqlite3"
)

var client *ent.Client

func Connect() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func Disconnect() {
	client.Close()
}

func GetClient() *ent.Client {
	return client
}
