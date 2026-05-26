//go:generate mockgen -destination=../../../mocks/services/streammanager/googlecloudfunction/mock_googlecloudfunction.go -package mock_googlecloudfunction github.com/rudderlabs/rudder-server/services/streammanager/googlecloudfunction GoogleCloudFunctionClient

package googlecloudfunction

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type Config struct {
	Credentials           string        `json:"credentials"`
	RequireAuthentication bool          `json:"requireAuthentication"`
	FunctionUrl           string        `json:"googleCloudFunctionUrl"`
	Token                 *oauth2.Token `json:"token"`
	TokenCreatedAt        time.Time     `json:"tokenCreatedAt"`
	TokenTimeout          time.Duration `json:"tokenTimeout"`
}

func (config *Config) shouldGenerateToken() bool { _ = "STUB: not implemented"; return false }

func (config *Config) generateToken(ctx context.Context, client GoogleCloudFunctionClient) error {
	_ = "STUB: not implemented"
	return nil
}

var pkgLogger logger.Logger

func Init() { _ = "STUB: not implemented"; return }

func init() {
	Init()
}

type GoogleCloudFunctionProducer struct {
	client     GoogleCloudFunctionClient
	config     *Config
	httpClient *http.Client
}

type GoogleCloudFunctionClient interface {
	GetToken(ctx context.Context, functionUrl string, opts ...option.ClientOption) (*oauth2.Token, error)
}

type GoogleCloudFunctionClientImpl struct{}

func (c *GoogleCloudFunctionClientImpl) GetToken(ctx context.Context, functionUrl string, opts ...option.ClientOption) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the ID token, to make an authenticated call to the target audience.

func getFunctionConfig(fnConfig Config) *Config { _ = "STUB: not implemented"; return nil }

// NewProducer creates a producer based on destination config
func NewProducer(destination *backendconfig.DestinationT, _ common.Opts) (*GoogleCloudFunctionProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (producer *GoogleCloudFunctionProducer) Produce(jsonData json.RawMessage, _ any) (statusCode int, respStatus, responseMessage string) {
	_ = "STUB: not implemented"
	// Create a POST request
	return 0, "", ""
}

// Set the appropriate headers

// Make the request using the client

func (producer *GoogleCloudFunctionProducer) Close() error { _ = "STUB: not implemented"; return nil }
