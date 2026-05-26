package marketobulkupload

import (
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

type MarketoAPIServiceInterface interface {
	ImportLeads(csvFilePath, deduplicationField string) (string, *APIError)
	PollImportStatus(importId string) (*MarketoResponse, *APIError)
	GetLeadStatus(url string) ([]map[string]string, *APIError)
}

type MarketoAPIService struct {
	logger       logger.Logger
	statsFactory stats.Stats
	httpClient   *http.Client
	munchkinId   string
	authService  MarketoAuthServiceInterface
	maxRetries   int
}

type APIError struct {
	StatusCode int
	Category   string
	Message    string
}

func (m *MarketoAPIService) checkForCSVLikeResponse(resp *http.Response) bool {
	_ = "STUB: not implemented"
	// check for csv like response by checking the headers
	return false
}

func (m *MarketoAPIService) attemptImport(uploadURL, csvFilePath, deduplicationField string, uploadTimeStat stats.Measurement) (string, *APIError) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *MarketoAPIService) ImportLeads(csvFilePath, deduplicationField string) (string, *APIError) {
	_ = "STUB: not implemented"
	return "", nil
}

// Initial attempt

// If it's not a token refresh error, don't retry

func (m *MarketoAPIService) PollImportStatus(importId string) (*MarketoResponse, *APIError) {
	_ = "STUB: not implemented"
	// poll for the import status
	return nil, nil
}

// Make the API request

func (m *MarketoAPIService) GetLeadStatus(url string) ([]map[string]string, *APIError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the response is not a csv like response, then it should be a json response

// if the response is a csv like response
// parse the csv response

// read each row one by one

// The first row is the header
