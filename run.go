package sprintsd

import (
	"context"
	"errors"
	"strings"
	"fmt"
	"log"

	"cloud.google.com/go/compute/metadata"
	run "cloud.google.com/go/run/apiv2"
	runpb "cloud.google.com/go/run/apiv2/runpb"
)

var ErrNotAvailable = errors.New("not available")

// GetRunServiceURL retrieves the URL of a cloud run service.
func GetRunServiceURL(ctx context.Context, name string) (string, error) {
	c, err := run.NewServicesClient(ctx)
	if err != nil {
		return "", fmt.Errorf(
			"unable to to create new services client: %w: %v",
			ErrNotAvailable,
			err,
		)
	}
	defer c.Close()

	projectID, err := metadata.NumericProjectID()
	if err != nil {
		return "", fmt.Errorf(
			"unable to get project id: %w: %v",
			ErrNotAvailable,
			err,
		)
	}

	region, err := metadata.Get("instance/region")
	if err != nil {
		return "", fmt.Errorf(
			"unable to get instance region: %w: %v",
			ErrNotAvailable,
			err,
		)
	}

	location := extractLocation(region)

	svcName := fmt.Sprintf(
		"projects/%s/locations/%s/services/%s",
		projectID,
		location,
		name,
	)
	log.Println("service name:", svcName)
	req := &runpb.GetServiceRequest{
		Name: svcName,
	}
	resp, err := c.GetService(ctx, req)
	if err != nil {
		return "", fmt.Errorf(
			"unable to get service information: %w: %v",
			ErrNotAvailable,
			err,
		)
	}

	return resp.GetUri(), nil
}

func extractLocation(region string) string {
	parts := strings.Split(region, "/regions/")
	if len(parts) >= 2 {
		return parts[1]
	}

	return ""
}
