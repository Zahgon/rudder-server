package warehouse

import (
	"context"
	"io"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/processor/types"
)

func (t *Transformer) CompareResponsesAndUpload(ctx context.Context, events []types.TransformerEvent, legacyResponse types.Response) {
	_ = "STUB: not implemented"
	return
}

func (t *Transformer) compareResponsesAndUpload(ctx context.Context, events []types.TransformerEvent, legacyResponse types.Response) {
	_ = "STUB: not implemented"
	return
}

func getSamplingUploader(conf *config.Config, log logger.Logger) (*filemanager.S3Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) sampleDiff(events []types.TransformerEvent, legacyResponse, embeddedResponse types.Response) string {
	_ = "STUB: not implemented"
	return ""
}

// Don't diff in case there is no response from transformer

// If the event counts differ, return all events in the transformation

// JS converts new Date('0001-01-01 00:00').toISOString() to 2001-01-01T00:00:00.000Z
// https://www.programiz.com/online-compiler/4SqZcIH5k6Yli

// If messageID's are not present, we add it in rudder-transformer
// https://github.com/rudderlabs/rudder-transformer/blob/develop/src/warehouse/index.js#L675-L677

// Ignore Unicode diffs caused by Go vs JavaScript serialization differences

func write(w io.WriteCloser, data []string) error { _ = "STUB: not implemented"; return nil }
