package testhelper

import (
	"testing"

	"cloud.google.com/go/bigquery"
)

type TestCredentials struct {
	ProjectID   string `json:"projectID"`
	Location    string `json:"location"`
	BucketName  string `json:"bucketName"`
	Credentials string `json:"credentials"`
}

const TestKey = "BIGQUERY_INTEGRATION_TEST_CREDENTIALS"

func GetBQTestCredentials() (*TestCredentials, error) { _ = "STUB: not implemented"; return nil, nil }

// RetrieveRecordsFromWarehouse retrieves records from the warehouse based on the given query.
// It returns a slice of slices, where each inner slice represents a record's values.
func RetrieveRecordsFromWarehouse(
	t testing.TB,
	db *bigquery.Client,
	query string,
) [][]string {
	_ = "STUB: not implemented"
	return nil
}
