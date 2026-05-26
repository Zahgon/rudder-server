package destination_transformer

import (
	"context"

	"github.com/rudderlabs/rudder-server/processor/types"
)

func (c *Client) CompareAndLog(
	ctx context.Context,
	embeddedResponse, legacyResponse types.Response,
) {
	_ = "STUB: not implemented"
	return
}

// Cannot upload, we should just report the issue with no diff

func (c *Client) compareAndLog(
	ctx context.Context,
	embeddedResponse, legacyResponse types.Response,
) {
	_ = "STUB: not implemented"
	return
}

// upload sample diff and differing response to s3

func (c *Client) differingEvents(
	embeddedResponse, legacyResponse types.Response,
) ([]types.TransformerResponse, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// Collect the mismatched event response and break (sample only)

// Collect the mismatched event response and break (sample only)
