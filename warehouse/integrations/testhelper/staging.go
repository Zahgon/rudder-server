package testhelper

import (
	"testing"

	"github.com/rudderlabs/rudder-go-kit/filemanager"

	warehouseclient "github.com/rudderlabs/rudder-server/warehouse/client"
)

func createStagingFile(t testing.TB, testConfig *TestConfig) { _ = "STUB: not implemented"; return }

func prepareStagingFilePathUsingStagingFile(t testing.TB, testConfig *TestConfig) string {
	_ = "STUB: not implemented"
	return ""
}

func prepareStagingFilePathUsingEventsFile(t testing.TB, testConfig *TestConfig) string {
	_ = "STUB: not implemented"
	return ""
}

func uploadStagingFile(t testing.TB, testConfig *TestConfig, stagingFile string) filemanager.UploadedFile {
	_ = "STUB: not implemented"
	return *new(filemanager.UploadedFile)
}

func prepareStagingPayload(t testing.TB, testConfig *TestConfig, stagingFile string, uploadOutput filemanager.UploadedFile) warehouseclient.StagingFile {
	_ = "STUB: not implemented"
	return *new(warehouseclient.StagingFile)
}

// merge rules and mappings events will not contain received_at, ignoring those
