package marketobulkupload

import "net/http"

// mockTransport implements http.RoundTripper interface
type mockTransport struct {
	response *http.Response
	err      error
}

func (m *mockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil,

		// createMockClient creates a new http.Client with mocked transport
		nil
}

func createMockClient(response *http.Response, err error) *http.Client {
	_ = "STUB: not implemented"
	return nil
}
