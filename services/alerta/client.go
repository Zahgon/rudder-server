package alerta

import (
	"context"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

type OptFn func(*Client)

type Tags map[string]string

type Priority string

type Severity string

type Environment string

type Alert struct {
	Resource    string      `json:"resource"`    // warehouse-upload-aborted
	Event       string      `json:"event"`       // <tags_list>
	Environment Environment `json:"environment"` // [PRODUCTION,DEVELOPMENT,PROXYMODE,CARBONCOPY]
	Severity    Severity    `json:"severity"`    // warning,critical,normal // Get the full list from https://docs.alerta.io/api/alert.html#severity-table
	Text        string      `json:"text"`        // <event> is critical
	Timeout     int         `json:"timeout"`     // 86400
	TagList     []string    `json:"tags"`        // {priority=P1,destID=27CHciD6leAhurSyFAeN4dp14qZ,destType=RS,namespace=hosted,cluster=rudder}
}

type SendAlertOpts struct {
	Tags        Tags
	Text        string
	Severity    Severity
	Priority    Priority
	Environment Environment
}

const (
	SeverityCritical Severity = "critical"
	SeverityWarning  Severity = "warning"
	SeverityNormal   Severity = "normal"
	SeverityOk       Severity = "ok"
)

const (
	PriorityP1 Priority = "P1"
	PriorityP2 Priority = "P2"
	PriorityP3 Priority = "P3"
)

const (
	PRODUCTION  Environment = "PRODUCTION"
	DEVELOPMENT Environment = "DEVELOPMENT"
	PROXYMODE   Environment = "PROXYMODE"
	CARBONCOPY  Environment = "CARBONCOPY"
)

var (
	defaultTimeout     = 30 * time.Second
	defaultMaxRetries  = 3
	defaultPriority    = PriorityP1
	defaultSeverity    = SeverityCritical
	defaultEnvironment = DEVELOPMENT
)

type AlertSender interface {
	SendAlert(ctx context.Context, resource string, opts SendAlertOpts) error
}

type Client struct {
	client         *http.Client
	retries        int
	url            string
	config         *config.Config
	alertTimeout   int
	kuberNamespace string
}

func WithHTTPClient(httpClient *http.Client) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithTimeout(timeout time.Duration) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithMaxRetries(retries int) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithConfig(config *config.Config) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithAlertTimeout(timeout int) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithKubeNamespace(namespace string) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("alerta")
}

func NewClient(baseURL string, fns ...OptFn) AlertSender {
	_ = "STUB: not implemented"
	return *new(AlertSender)
}

func (c *Client) retry(ctx context.Context, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) defaultTags(opts *SendAlertOpts) Tags {
	_ = "STUB: not implemented"
	return *new(Tags)
}

func (c *Client) isEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *Client) setDefaultsOpts(resource string, opts *SendAlertOpts) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) SendAlert(ctx context.Context, resource string, opts SendAlertOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// default tags
