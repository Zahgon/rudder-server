package salesforcebulkupload

import (
	"io"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

func newAPIService(
	logger logger.Logger,
	destination *backendconfig.DestinationT,
	client *http.Client,
) APIServiceInterface {
	_ = "STUB: not implemented"
	return *new(APIServiceInterface)
}

func (s *apiService) CreateJob(
	objectName, operation, externalIDField string,
) (string, *APIError) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *apiService) UploadData(jobID, csvFilePath string) *APIError {
	_ = "STUB: not implemented"
	return nil
}

func (s *apiService) CloseJob(jobID string) *APIError { _ = "STUB: not implemented"; return nil }

func (s *apiService) GetJobStatus(jobID string) (*JobResponse, *APIError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *apiService) GetFailedRecords(jobID string) ([]map[string]string, *APIError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *apiService) GetSuccessfulRecords(jobID string) ([]map[string]string, *APIError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *apiService) DeleteJob(jobID string) *APIError { _ = "STUB: not implemented"; return nil }

func (s *apiService) getCSVRecords(endpoint string) ([]map[string]string, *APIError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *apiService) makeRequest(
	method, endpoint string,
	body io.Reader,
	contentType string,
) ([]byte, *APIError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func categorizeError(statusCode int) string { _ = "STUB: not implemented"; return "" }
