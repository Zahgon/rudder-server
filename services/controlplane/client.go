package controlplane

import (
	"context"
	"net/http"
	"time"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
)

var (
	defaultTimeout    = 30 * time.Second
	defaultMaxRetries = 3
)

type OptFn func(c *commonClient)

func WithHTTPClient(httpClient *http.Client) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithTimeout(timeout time.Duration) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithMaxRetries(retries int) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

func WithRegion(region string) OptFn { _ = "STUB: not implemented"; return *new(OptFn) }

type commonClient struct {
	client  *http.Client
	retries int
	ua      string
	url     string
	region  string
}

type Client struct {
	*commonClient
	identity identity.Identifier
}

type AdminClient struct {
	*commonClient
	authorizer identity.Authorizer
}

type payloadSchema struct {
	Components []componentSchema `json:"components"`
}

type componentSchema struct {
	Name     string   `json:"name"`
	Features []string `json:"features"`
}

type SSHKeyPair struct {
	PublicKey  string
	PrivateKey string
}

func hostname() string { _ = "STUB: not implemented"; return "" }

func newCommonClient(baseURL string, fns ...OptFn) *commonClient {
	_ = "STUB: not implemented"
	return nil
}

func NewClient(baseURL string, identity identity.Identifier, fns ...OptFn) *Client {
	_ = "STUB: not implemented"
	return nil
}

func NewAdminClient(baseURL string, authorizer identity.Authorizer, fns ...OptFn) *AdminClient {
	_ = "STUB: not implemented"
	return nil
}

type PerComponent = map[string][]string

func (c *commonClient) retry(ctx context.Context, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *AdminClient) GetDestinationSSHKeyPair(ctx context.Context, destID string) (kp SSHKeyPair, err error) {
	_ = "STUB: not implemented"
	return *new(SSHKeyPair), nil
}

func (c *Client) SendFeatures(ctx context.Context, component string, features []string) error {
	_ = "STUB: not implemented"
	return nil
}

// we don't expect a body, unless there is an error

func (c *Client) DestinationHistory(ctx context.Context, revisionID string) (backendconfig.DestinationT, error) {
	_ = "STUB: not implemented"
	return *new(backendconfig.DestinationT), nil
}
