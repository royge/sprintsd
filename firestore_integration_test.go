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
	fs := sprintsd.NewFirestore(client, "sprintsd-testautomation")
	if fs == nil {
		t.Fatal("unexpected nil firestore instance")
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
