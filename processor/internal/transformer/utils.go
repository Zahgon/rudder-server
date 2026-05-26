package utils

import (
	"context"
	"errors"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	transformerclient "github.com/rudderlabs/rudder-server/internal/transformer-client"
	"github.com/rudderlabs/rudder-server/processor/types"
)

const (
	StatusCPDown                 = 809
	StatusColdStartWindowFailure = 819
	StatusMirrorFiltered         = 297
	TransformerRequestFailure    = 909
	TransformerRequestTimeout    = 919
)

var ErrColdStart = errors.New("cold start error")

func IsJobTerminated(status int) bool { _ = "STUB: not implemented"; return false }

func TransformerClientConfig(conf *config.Config, configPrefix string) *transformerclient.ClientConfig {
	_ = "STUB: not implemented"
	return nil
}

func TrackLongRunningTransformation(ctx context.Context, stage string, timeout time.Duration, log logger.Logger) {
	_ = "STUB: not implemented"
	return
}

// PythonTransformConfig holds version-based filtering config for Python transformations.
type PythonTransformConfig struct {
	Enabled    bool
	VersionIDs map[string]struct{}
}

// LoadPythonTransformConfig reads python transform version filtering from config.
func LoadPythonTransformConfig(conf *config.Config) PythonTransformConfig {
	_ = "STUB: not implemented"
	return *new(PythonTransformConfig)
}

// IsVersionAllowed returns true if version filtering is disabled or the versionID is in the allowlist.
func (c PythonTransformConfig) IsVersionAllowed(versionID string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetTransformationInfo extracts language, versionID, and transformationID from the first event's first transformation.
func GetTransformationInfo(events []types.TransformerEvent) (language, versionID, transformationID string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// GetEndpointFromURL is a helper function to extract hostname from URL
func GetEndpointFromURL(urlStr string) string {
	_ = "STUB: not implemented"
	// Parse URL and extract hostname
	return ""
}

// WithProcTransformReqTimeStat is a wrapper function to measure time taken by a request function as stats, capturing error rates as well, through a [success] label
func WithProcTransformReqTimeStat(request func() error, stat stats.Stats, labels types.TransformerMetricLabels) func() error {
	_ = "STUB: not implemented"
	return nil
}
