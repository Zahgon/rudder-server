package api

// This is simply going to make any API call to transfomer as per the API spec here: https://www.notion.so/rudderstacks/GDPR-Transformer-API-Spec-c3303b5b70c64225815d56c7767f8d22
// to get deletion done.
// called by delete/deleteSvc with (model.Job, model.Destination).
// returns final status,error ({successful, failure}, err)
import (
	"context"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
	"github.com/rudderlabs/rudder-server/services/transformer"
)

var (
	pkgLogger             = logger.NewLogger().Child("api")
	SupportedDestinations = []string{"BRAZE", "AM", "INTERCOM", "CLEVERTAP", "AF", "MP", "GA", "ITERABLE", "ENGAGE", "CUSTIFY", "SENDGRID", "SPRIG"}
)

type APIManager struct {
	Client                       *http.Client
	DestTransformURL             string
	MaxOAuthRefreshRetryAttempts int
	TransformerFeaturesService   transformer.FeaturesService
}

func GetAuthErrorCategoryFromResponse(bodyBytes []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *APIManager) GetSupportedDestinations() []string {
	_ = "STUB: not implemented"
	// Wait for transformer features to be available
	return nil
}

// Fallback to default supported destinations

func (m *APIManager) deleteWithRetry(ctx context.Context, job model.Job, destination *backendconfig.DestinationT, currentOauthRetryAttempt int) model.JobStatus {
	_ = "STUB: not implemented"
	return *new(model.JobStatus)
}

// reflection might be an expensive operation, checking if we need to print it
// nolint:forbidigo

// Post response work to be done for OAuthV2

// We don't need to handle it, as we can receive a string response even before executing OAuth operations like Refresh Token.
// It's acceptable if the structure of bodyBytes doesn't match the oauthv2.TransportResponse struct.

// most probably it was thrown before postRoundTrip through interceptor itself
// setting original response

// Update the same error response to all as the response received would be []JobRespSchema

// prepares payload based on (job,destination) & make an API call to transformer.
// gets (status, failure_reason) which is converted to appropriate model.Error & returned to caller.
func (m *APIManager) Delete(ctx context.Context, job model.Job, destination *backendconfig.DestinationT) model.JobStatus {
	_ = "STUB: not implemented"
	return *new(model.JobStatus)
}

func getJobStatus(statusCode int, jobResp []JobRespSchema) model.JobStatus {
	_ = "STUB: not implemented"
	return *new(model.JobStatus)
}

func mapJobToPayload(job model.Job, destName string, destConfig map[string]any) []apiDeletionPayloadSchema {
	_ = "STUB: not implemented"
	return nil
}

func getOAuthErrorJob(jobResponses []JobRespSchema) (JobRespSchema, bool) {
	_ = "STUB: not implemented"
	return *new(JobRespSchema), false
}

type PostResponseParams struct {
	destination              *backendconfig.DestinationT
	isOAuthEnabled           bool
	currentOAuthRetryAttempt int
	job                      model.Job
	responseBodyBytes        []byte
	responseStatusCode       int
}

func (m *APIManager) PostResponse(ctx context.Context, params PostResponseParams) model.JobStatus {
	_ = "STUB: not implemented"
	return *new(model.JobStatus)
}

// new oauth handling

// All the handling related to OAuth has been done(inside api.Client.Do() itself)!
// retry the request

// Abort the regulation request
