package marketobulkupload

import (
	"net/http"
	"time"
)

const (
	defaultTimeout             = 30 * time.Second
	defaultIdleConnTimeout     = 90 * time.Second
	defaultMaxIdleConnsPerHost = 50
	defaultMaxConnsPerHost     = 100
)

// getDefaultHTTPClient returns an http.Client with standard configuration
func getDefaultHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }

// Disable compression to prevent BREACH attacks
