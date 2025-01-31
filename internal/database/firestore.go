package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitFirestore() {
	ctx := context.Background()
	projectID := os.Getenv("FIRESTORE_PROJECT_ID")

	opt := option.WithCredentialsFile("config/todo-config.json")
	var err error
	FirestoreClient, err = firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
}

func CloseFirestore() {
	if FirestoreClient != nil {
		err := FirestoreClient.Close()
		if err != nil {
			log.Printf("Error closing Firestore client: %v", err)
		}
	}
}
