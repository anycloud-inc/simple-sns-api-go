package db

import (
	"context"
	"log"
	"simple_sns_api/src/ent"
)

var (
	Client *ent.Client
)

func CreateConnection() {
	client, err := ent.Open("mysql", "mysql:password@tcp(db:3306)/simple_sns?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	Client = client
}
