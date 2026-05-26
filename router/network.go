//go:generate mockgen -destination=../mocks/router/mock_network.go -package mock_network github.com/rudderlabs/rudder-server/router NetHandle

package router

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/netutil"

	"github.com/rudderlabs/rudder-server/processor/integrations"
	"github.com/rudderlabs/rudder-server/router/utils"
	"github.com/rudderlabs/rudder-server/utils/sysUtils"
)

var (
	contentTypeRegex = regexp.MustCompile(`^(text/[a-z0-9.-]+)|(application/([a-z0-9.-]+\+)?(json|xml))$`)
	ErrDenyPrivateIP = errors.New("access to private IPs is blocked")
)

// netHandle is the wrapper holding private variables
type netHandle struct {
	disableEgress        bool
	httpClient           sysUtils.HTTPClientI
	logger               logger.Logger
	blockPrivateIPs      bool
	blockPrivateIPsCIDRs netutil.CIDRs
	destType             string
	instanceID           string
}

// NetHandle interface
type NetHandle interface {
	SendPost(ctx context.Context, structData integrations.PostParametersT) *utils.SendPostResponse
}

// temp solution for handling complex query params
func handleQueryParam(param any) string { _ = "STUB: not implemented"; return "" }

// SendPost takes the EventPayload of a transformed job, gets the necessary values from the payload and makes a call to destination to push the event to it
// this returns the statusCode, status and response body from the response of the destination call
func (network *netHandle) SendPost(ctx context.Context, structData integrations.PostParametersT) *utils.SendPostResponse {
	_ = "STUB: not implemented"
	return nil
}

// going forward we may want to support GraphQL and multipart requests
// the files key in the response is specifically to handle the multipart use case
// for type GraphQL may need to support more keys like expected response format etc.
// in future it's expected that we will build on top of this response type
// so, code addition should be done here instead of version bumping of response.

// support for JSON and FORM body type

// support for JSON ARRAY

// transformer ensures top level string values, still val.(string) would be restrictive

// add query params to the url
// support of array type in params is handled if the
// response from transformers are "," separated

// Detecting content type of the respBody

// If media type is not in some human-readable format (text,json,xml), override the response with an empty string
// https://www.iana.org/assignments/media-types/media-types.xhtml

// returning 200 with a message in case of unsupported processing
// so that we don't process again. can change this code to anything
// to be not picked up by router again

// Setup initializes the module
func (network *netHandle) Setup(config *config.Config, netClientTimeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// In block mode, reject the connection
