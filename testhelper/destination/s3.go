package destination

import (
	"github.com/rudderlabs/rudder-go-kit/testhelper/docker/resource/minio"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

func MINIOFromResource(id string, m *minio.Resource) backendconfig.DestinationT {
	_ = "STUB: not implemented"
	return *new(backendconfig.DestinationT)
}
