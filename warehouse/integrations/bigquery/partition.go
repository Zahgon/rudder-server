package bigquery

import (
	"errors"

	"cloud.google.com/go/bigquery"
)

var (
	errPartitionColumnNotSupported = errors.New("partition column not supported")
	errPartitionTypeNotSupported   = errors.New("partition type not supported")
)

var supportedPartitionColumnMap = map[string]struct{}{
	"_PARTITIONTIME":     {},
	"loaded_at":          {},
	"received_at":        {},
	"sent_at":            {},
	"timestamp":          {},
	"original_timestamp": {},
}

var supportedPartitionTypeMap = map[string]bigquery.TimePartitioningType{
	"hour": bigquery.HourPartitioningType,
	"day":  bigquery.DayPartitioningType,
}

// avoidPartitionDecorator returns true if custom partition is enabled via destination or global config
// if we know beforehand that the data is in a single partition, specifying the partition decorator can improve write performance.
// However, in case of time-unit column and integer-range partitioned tables, the partition ID specified in the decorator must match the data being written. Otherwise, an error occurs.
// Therefore, if we are not sure about the partition decorator, we should not specify it i.e. in case when it is enabled via destination config.
func (bq *BigQuery) avoidPartitionDecorator() bool { _ = "STUB: not implemented"; return false }

func (bq *BigQuery) customPartitionEnabledViaGlobalConfig() bool {
	_ = "STUB: not implemented"
	return false
}

func (bq *BigQuery) isTimeUnitPartitionColumn() bool { _ = "STUB: not implemented"; return false }

func (bq *BigQuery) partitionColumn() string { _ = "STUB: not implemented"; return "" }

func (bq *BigQuery) partitionType() string { _ = "STUB: not implemented"; return "" }

func (bq *BigQuery) checkValidPartitionColumn(partitionColumn string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bq *BigQuery) bigqueryPartitionType(partitionType string) (bigquery.TimePartitioningType, error) {
	_ = "STUB: not implemented"
	return *new(bigquery.TimePartitioningType), nil
}

func (bq *BigQuery) partitionDate() (string, error) { _ = "STUB: not implemented"; return "", nil }

func partitionedTable(tableName, partitionDate string) string { _ = "STUB: not implemented"; return "" }
