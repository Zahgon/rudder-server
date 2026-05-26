//go:generate mockgen --build_flags=--mod=mod -destination=./../mocks/mockwebhook.go -package mockwebhook github.com/rudderlabs/rudder-server/gateway/webhook Gateway

package webhook

import (
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/stats"

	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
	"github.com/rudderlabs/rudder-server/gateway/webhook/model"
)

type Gateway interface {
	RequestMetricsTracker
	WebhookRequestProcessor
}

type WebhookRequestProcessor interface {
	// ProcessTransformedWebhookRequest processes the transformed webhook request and save it to the gw jobsDB
	ProcessTransformedWebhookRequest(writer *http.ResponseWriter, req *http.Request, reqType string, requestPayload []byte, arctx *gwtypes.AuthRequestContext) string
	// SaveWebhookFailures saves the webhook failures to the procErr jobsDB
	SaveWebhookFailures([]*model.FailedWebhookPayload) error
}

// StatReporterCreator is a function type that creates StatReporter instances
type StatReporterCreator func(authContext *gwtypes.AuthRequestContext, requestType string) gwtypes.StatReporter

// RequestMetricsTracker is used to track webhook request metrics on a request basis for OSS customers
type RequestMetricsTracker interface {
	TrackRequestMetrics(errorMessage string)
}

type TransformerFeaturesService interface {
	SourceTransformerVersion() string
}

func newWebhookStats(stat stats.Stats) *webhookStatsT { _ = "STUB: not implemented"; return nil }

func Setup(gwHandle Gateway, transformerFeaturesService TransformerFeaturesService, stat stats.Stats, conf *config.Config, statReporterCreator StatReporterCreator, opts ...batchTransformerOption) *HandleT {
	_ = "STUB: not implemented"
	return nil
}

// Number of incoming webhooks that are batched before calling source transformer

// Timeout after which batch is formed anyway with whatever webhooks are available

// Multiple source transformers are used to generate rudder events from webhooks

// Parse all query params from sources mentioned in this list

// Maximum request size to gateway

// enable webhook v2 handler

// lowercasing the strings in sourceListForParsingParams
