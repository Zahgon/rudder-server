package warehouse

import (
	"context"
	"regexp"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	wtypes "github.com/rudderlabs/rudder-server/processor/internal/transformer/destination_transformer/embedded/warehouse/internal/types"
	"github.com/rudderlabs/rudder-server/processor/types"
)

const (
	violationErrors     = "violationErrors"
	redshiftStringLimit = 512
)

var unicodePattern = regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)

type Opts func(t *Transformer)

func WithNow(now func() time.Time) Opts { _ = "STUB: not implemented"; return *new(Opts) }

func WithUUIDGenerator(uuidGenerator func() string) Opts {
	_ = "STUB: not implemented"
	return *new(Opts)
}

func WithSorter(sorter func([]string) []string) Opts { _ = "STUB: not implemented"; return *new(Opts) }

func New(conf *config.Config, logger logger.Logger, statsFactory stats.Stats, opts ...Opts) *Transformer {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transformer) Transform(_ context.Context, clientEvents []types.TransformerEvent) (res types.Response) {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// TODO: Currently, it's getting ignored during JSON marshalling Remove this once we start using it.

func (t *Transformer) processWarehouseMessage(cache *cache, event *types.TransformerEvent) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) checkValidContext(event *wtypes.TransformerEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transformer) eventContext(tec *transformEventContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (t *Transformer) handleEvent(event *wtypes.TransformerEvent, cache *cache) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func transformerResponseFromErr(metadata *types.Metadata, statTags map[string]string, err error) types.TransformerResponse {
	_ = "STUB: not implemented"
	return *new(types.TransformerResponse)
}

func (t *Transformer) getColumns(
	destType string,
	data map[string]any, metadata map[string]string,
) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// uuid_ts and loaded_at datatypes are passed from here to create appropriate columns.
// Corresponding values are inserted when loading into the warehouse
