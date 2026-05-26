//go:generate mockgen --build_flags=--mod=mod -destination=../../../mocks/services/streammanager/common/mock_streammanager.go -package mock_streammanager github.com/rudderlabs/rudder-server/services/streammanager/common StreamProducer

package common

import (
	"encoding/json"
	"io"
	"time"

	"github.com/aws/smithy-go"
)

type Producer interface {
	io.Closer
	Produce(jsonData json.RawMessage, _ any) (int, string, string)
}

type StreamProducer interface {
	io.Closer
	Produce(jsonData json.RawMessage, destConfig any) (int, string, string)
}

type Opts struct {
	Timeout time.Duration
}

func mapErrorMessageToStatusCode(errorMessage string, defaultStatusCode int) int {
	_ = "STUB: not implemented"
	return 0
}

// aws returns  "ThrottlingException"
// for throttling requests server will retry

// Retryable

func getStatusCodeFromFault(fault smithy.ErrorFault) int { _ = "STUB: not implemented"; return 0 }

func ParseAWSError(err error) (statusCode int, respStatus, responseMessage string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}
