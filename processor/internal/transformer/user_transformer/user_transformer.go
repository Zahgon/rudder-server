package user_transformer

import (
	"context"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	transformerclient "github.com/rudderlabs/rudder-server/internal/transformer-client"
	transformerutils "github.com/rudderlabs/rudder-server/processor/internal/transformer"
	"github.com/rudderlabs/rudder-server/processor/types"
)

type Opt func(*Client)

func WithClient(client transformerclient.Client) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func ForMirroring() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func New(conf *config.Config, log logger.Logger, stat stats.Stats, opts ...Opt) *Client {
	_ = "STUB: not implemented"
	return nil
}

type Client struct {
	config struct {
		userTransformationURL         string
		pythonTransformationURL       string
		pythonTransformConfig         transformerutils.PythonTransformConfig
		forMirroring                  bool
		maxRetry                      config.ValueLoader[int]
		cpDownEndlessRetries          config.ValueLoader[bool]
		failOnUserTransformTimeout    config.ValueLoader[bool]
		failOnError                   config.ValueLoader[bool]
		maxRetryBackoffInterval       config.ValueLoader[time.Duration]
		timeoutDuration               time.Duration
		collectInstanceLevelStats     bool
		batchSize                     config.ValueLoader[int]
		perWorkspacePyTEnabled        config.ValueLoader[bool]
		perWorkspacePyTURLTemplate    string
		perWorkspacePyTEndlessRetries config.ValueLoader[bool]
	}
	conf   *config.Config
	log    logger.Logger
	stat   stats.Stats
	client transformerclient.Client
}

func (u *Client) Transform(ctx context.Context, clientEvents []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// If any batch was mirror-filtered, the whole response is mirror-filtered.
// All batches share the same transformation, so this is all-or-nothing.

// Transform is one to many mapping so returned
// response for each is an array. We flatten it out

func (u *Client) sendBatch(
	ctx context.Context,
	url string,
	labels types.TransformerMetricLabels,
	clientEvents []types.TransformerEvent,
) (
	[]types.TransformerResponse,
	bool, // is mirror filtered
) {
	_ = "STUB: not implemented"
	return nil, false
}

// Call remote transformation

// flip sourceID and originalSourceID if it's a replay source for the purpose of any user transformation
// flip back afterward

// endless retry if transformer-control plane connection is down

// endless backoff loop, only nil error or panics inside

// no max time -> ends only when no error

// control plane back up

// This is returned by our JS engine so should be parseable
// Panic the processor to avoid replays

func (u *Client) doPost(ctx context.Context, rawJSON []byte, url string, labels types.TransformerMetricLabels) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Header to let transformer know that the client understands event filter code

// Record metrics with labels

// This metric is to track cold start errors for PyT, which are expected to be higher than usual due to the nature of PyT scaling.

// We'll count response events after unmarshaling in the request method

// Per-workspace PyT: a persistent transport failure on the workspace-
// scoped URL is most likely a cold-start window (pod not yet scaled by
// HPA, EndpointSlice lag, kube-proxy 5xx). Surface a dedicated status
// code so sendBatch retries the whole call until the pod is up instead of treating it as a failed transformation.

// perform version compatibility check only on success

// isPerWorkspacePyTPath returns true when the request is targeting the
// per-workspace PyT URL. Single source of truth shared by URL resolution and
// the cold-start error / counter path so the two can't drift.
func (u *Client) isPerWorkspacePyTPath(language, workspaceID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (u *Client) shouldThrowPythonColdStartErr(labels types.TransformerMetricLabels, err error, resp *http.Response) bool {
	_ = "STUB: not implemented"
	return false
}

// isColdStartError returns true for transient errors that mean the target
// PyT deployment isn't ready yet — zero endpoints, pod not yet ready, or
// kube-proxy returning a no-endpoints status.
func isColdStartError(err error, resp *http.Response) bool {
	_ = "STUB: not implemented"

	// ECONNREFUSED: Service has no endpoints (Deployment at 0 replicas).
	// EHOSTUNREACH ("no route to host"): stale iptables / EndpointSlice
	// after a pod replacement or scale-down — same transient signal.
	return false
}

func isPythonTransformation(language string) bool { _ = "STUB: not implemented"; return false }

func (u *Client) userTransformURL(language, versionID, workspaceID string) string {
	_ = "STUB: not implemented"
	return ""
}

// Per-workspace PyT: a global version allowlist doesn't apply — each
// workspace runs its own pod with its own version.

// Panic so the bug surfaces immediately as this should not happen

// Legacy shared-PyT path: the version allowlist is a rollout gate for the shared service.
