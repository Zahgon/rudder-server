package testhelper

import (
	"testing"

	"github.com/rudderlabs/rudder-server/processor/types"
)

func ValidateExpectedEvents(t testing.TB, expectedResponse, embeddedResponse, legacyResponse types.Response) {
	_ = "STUB: not implemented"
	return
}

func ValidateEvents(t *testing.T, embeddedResponse, legacyResponse types.Response) {
	_ = "STUB: not implemented"
	return
}

func cmpEvents(t testing.TB, expected, actual []types.TransformerResponse) {
	_ = "STUB: not implemented"
	return
}

func checkForMarshalledFieldsAndRemove(t testing.TB, expected, embedded, legacy []types.TransformerResponse, fields ...string) {
	_ = "STUB: not implemented"
	return
}

func cmpFailedEvents(t testing.TB, expected, actual []types.TransformerResponse) {
	_ = "STUB: not implemented"
	return
}

func deepCopyResponse(t testing.TB, res types.Response) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}
