package testhelper

import (
	"testing"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	whclient "github.com/rudderlabs/rudder-server/warehouse/client"
)

func verifyEventsInStagingFiles(t testing.TB, testConfig *TestConfig) {
	_ = "STUB: not implemented"
	return
}

func verifyEventsInTableUploads(t testing.TB, testConfig *TestConfig) {
	_ = "STUB: not implemented"
	return
}

func verifyEventsInWareHouse(t testing.TB, testConfig *TestConfig) {
	_ = "STUB: not implemented"
	return
}

func queryCount(cl *whclient.Client, statement string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func verifySourceJob(t testing.TB, tc *TestConfig) { _ = "STUB: not implemented"; return }

func VerifyConfigurationTest(t testing.TB, destination backendconfig.DestinationT) {
	_ = "STUB: not implemented"
	return
}
