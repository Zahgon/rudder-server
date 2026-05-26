package webhook

import (
	"context"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/retryablehttp"
	"github.com/rudderlabs/rudder-go-kit/stats"

	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

type webhookT struct {
	request     *http.Request
	writer      http.ResponseWriter
	done        chan<- transformerResponse
	sourceID    string
	sourceType  string
	authContext *gwtypes.AuthRequestContext
}

type batchWebhookT struct {
	batchRequest []*webhookT
	sourceType   string
}

//go:generate mockgen -destination=../../mocks/gateway/webhook.go -package=mocks_gateway github.com/rudderlabs/rudder-server/gateway/webhook WebhookRequestHandler
type WebhookRequestHandler interface {
	// RequestHandler handles the incoming webhook request
	RequestHandler(w http.ResponseWriter, r *http.Request)
	// Register registers a new webhook source type and starts a goroutine to process requests for that source type
	Register(name string)
	// Shutdown shuts down the webhook handler, closing all channels and waiting for goroutines to finish
	Shutdown() error
}

type HandleT struct {
	logger        logger.Logger
	requestQMu    sync.RWMutex
	requestQ      map[string]chan *webhookT
	batchRequestQ chan *batchWebhookT
	gwHandle      Gateway
	stats         stats.Stats
	ackCount      atomic.Uint64
	recvCount     atomic.Uint64

	batchRequestsWg  sync.WaitGroup
	backgroundWait   func() error
	backgroundCancel context.CancelFunc

	config struct {
		maxReqSize                 config.ValueLoader[int]
		webhookBatchTimeout        config.ValueLoader[time.Duration]
		maxWebhookBatchSize        config.ValueLoader[int]
		sourceListForParsingParams []string
		forwardGetRequestForSrcMap map[string]struct{}
		webhookV2HandlerEnabled    bool
	}
	statReporterCreator StatReporterCreator
	httpClient          retryablehttp.HttpClient
}

type webhookSourceStatT struct {
	id              string
	numEvents       stats.Measurement
	numOutputEvents stats.Measurement
	sourceTransform stats.Measurement
}

type webhookStatsT struct {
	sentStat           stats.Measurement
	receivedStat       stats.Measurement
	failedStat         stats.Measurement
	transformTimerStat stats.Measurement
	sourceStats        map[string]*webhookSourceStatT
}

type batchWebhookTransformerT struct {
	webhook                *HandleT
	stats                  *webhookStatsT
	statsFactory           stats.Stats
	sourceTransformAdapter func(ctx context.Context) (sourceTransformAdapter, error)
}

type batchTransformerOption func(bt *batchWebhookTransformerT)

func (webhook *HandleT) failRequest(w http.ResponseWriter, r *http.Request, reason string, code int) {
	_ = "STUB: not implemented"
	return
}

func (wb *HandleT) IsGetAndNotAllow(reqMethod, sourceDefName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (webhook *HandleT) RequestHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Wait for batcher process to be done

func getWebhookFailureReason(errMsg string, statusCode int) string {
	_ = "STUB: not implemented"
	return ""
}

func (webhook *HandleT) batchRequests(sourceDef string, requestQ chan *webhookT) {
	_ = "STUB: not implemented"
	return
}

// If there are requests in the buffer, send them to the batcher

// Append to request buffer

// TODO : return back immediately for blank request body. its waiting till timeout
func (bt *batchWebhookTransformerT) batchTransformLoop() { _ = "STUB: not implemented"; return }

// If unable to fetch features from transformer, send GatewayTimeout to all requests
// TODO: Remove timeout from here after timeout handler is added in gateway

// stats

// stats

func (bt *batchWebhookTransformerT) getWebhookFailureReason(errMessage, reason string) string {
	_ = "STUB: not implemented"
	return ""
}

func (webhook *HandleT) enqueueInGateway(req *webhookT, payload []byte) string {
	_ = "STUB: not implemented"
	// replace body with transformed event (it comes in a batch format)
	return ""
}

func (webhook *HandleT) Register(name string) { _ = "STUB: not implemented"; return }

func (webhook *HandleT) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (webhook *HandleT) countWebhookErrors(sourceType string, arctx *gwtypes.AuthRequestContext, reason string, statusCode, count int) {
	_ = "STUB: not implemented"
	return
}

func (webhook *HandleT) recordWebhookErrors(sourceType, reason string, reqs []*webhookT, statusCode int) {
	_ = "STUB: not implemented"
	return
}

// TODO: Check if correct
func (bt *batchWebhookTransformerT) newWebhookStat(sourceType string) *webhookSourceStatT {
	_ = "STUB: not implemented"
	return nil
}

func (webhook *HandleT) printStats(ctx context.Context) { _ = "STUB: not implemented"; return }
