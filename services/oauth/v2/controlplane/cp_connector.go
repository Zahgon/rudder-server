package controlplane

//go:generate mockgen -destination=../../../../mocks/services/oauthV2/mock_cp_connector.go -package=mock_oauthV2 github.com/rudderlabs/rudder-server/services/oauth/v2/controlplane Connector

import (
	"net/http"
	"regexp"
	"syscall"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
)

var (
	errTypMap = map[syscall.Errno]string{
		syscall.ECONNRESET:   "econnreset",
		syscall.ECONNREFUSED: "econnrefused",
		syscall.ECONNABORTED: "econnaborted",
		syscall.ECANCELED:    "ecanceled",
	}
	contentTypePattern = regexp.MustCompile(`text|application/json|application/xml`)
)

type Connector interface {
	CpApiCall(cpReq *Request) (int, string)
}

type connector struct {
	client  HttpClient
	logger  logger.Logger
	timeout time.Duration
	stats   stats.Stats
}

func NewConnector(conf *config.Config, options ...func(*connector)) Connector {
	_ = "STUB: not implemented"
	return *new(Connector)
}

func WithStats(stats stats.Stats) func(*connector) { _ = "STUB: not implemented"; return nil }

// WithClient is a functional option to set the client for the Connector
func WithClient(client HttpClient) func(*connector) { _ = "STUB: not implemented"; return nil }

// WithLogger is a functional option to set the parent logger for the Connector
func WithLogger(parentLogger logger.Logger) func(*connector) { _ = "STUB: not implemented"; return nil }

// WithCpClientTimeout is a functional option to set the timeout for the Connector
func WithCpClientTimeout(timeout time.Duration) func(*connector) {
	_ = "STUB: not implemented"
	return nil
}

// processResponse is a helper function to process the response from the control plane
func processResponse(resp *http.Response) (statusCode int, respBody string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// Detecting content type of the respData

// If content type is not of type "*text*", overriding it with empty string

// CpApiCall is a function to make a call to the control plane, handle the response and return the status code and response body
func (c *connector) CpApiCall(cpReq *Request) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// Abort on receiving an error in request formation

// Authorisation setting

// Set content-type in order to send the body in request correctly

// Abort on receiving an error

// got some valid response from cp

func GetErrorType(err error) string { _ = "STUB: not implemented"; return "" }
