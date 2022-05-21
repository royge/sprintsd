//go:build integration
// +build integration

package sprintsd_test

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/royge/sprintsd"
)

func Test_NewFirestore_Integration(t *testing.T) {
	client := createFirestoreClient(t)
	defer client.Close()

	fs := sprintsd.NewFirestore(client, "sprintsd-testautomation")
	if fs == nil {
		t.Fatal("unexpected nil firestore instance")
	}
}

func Test_Firestore_Save_Integration(t *testing.T) {
	client := createFirestoreClient(t)
	defer client.Close()

	ctx := context.Background()

	fs := sprintsd.NewFirestore(client, "sprintsd-testautomation")
	if err := fs.Save(ctx, &sprintsd.Registration{
		ID: "test-1234",
		Name: "test",
		Address: "http://localhost",
		Port: 80,
	}); err != nil {
		t.Errorf("unable to save new registration: %v", err)
	}
}

func createFirestoreClient(t *testing.T) *firestore.Client {
	t.Helper()

	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	if projectID == "" {
		t.Fatalf("no google project id defined")
	}

	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		t.Fatalf("error creating firestore client: %v", err)
	}

	return client
}
