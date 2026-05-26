package client

import (
	"net/rpc"
)

// Setup Unix Domain socket client to communicate with
// the rudder server running on the same machine as this cli
func GetUDSClient() *rpc.Client { _ = "STUB: not implemented"; return nil }
