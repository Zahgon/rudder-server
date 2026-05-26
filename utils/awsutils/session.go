package awsutils

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/awsutil"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

func NewSimpleSessionConfigForDestination(destination *backendconfig.DestinationT, serviceName string) (*awsutil.SessionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/**
In order prevent confused deputy problem, we are using
workspace token as external ID.
Ref: https://docs.aws.amazon.com/IAM/latest/UserGuide/confused-deputy.html
*/

func NewSessionConfigForDestination(destination *backendconfig.DestinationT, timeout time.Duration, serviceName string) (*awsutil.SessionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
