package jobsdb

import (
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/stats"
)

type sqlDbOrTx interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

const preDropTableComment = "rudder:pre_drop:v1"

/*
Function to return an ordered list of datasets and datasetRanges
Most callers use the in-memory list of dataset and datasetRanges
*/
func getDSList(jd assertInterface, dbHandle sqlDbOrTx, tablePrefix string) ([]dataSetT, error) {
	_ = "STUB: not implemented"
	return nil,

		// Read the table names from PG
		nil
}

// Tables are of form jobs_ and job_status_. Iterate
// through them and sort them to produce and
// ordered list of datasets

// Create the structure

/*
sortDnumList Function to sort table suffixes. We should not have any use case
for having > 2 len suffixes (e.g. 1_1_1 - see comment below)
but this sort handles the general case
*/
func sortDnumList(dnumList []string) { _ = "STUB: not implemented"; return }

// getAllTableNames gets all table names from Postgres, excluding tables marked as pre-drop.
func getAllTableNames(dbHandle sqlDbOrTx) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func preDropDatasetTables(tableName string) (jobTable, statusTable string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// checkValidJobState Function to check validity of states
func checkValidJobState(jd assertInterface, stateFilters []string) {
	_ = "STUB: not implemented"
	return
}

// constructQueryOR construct a query were paramKey is any of the values in paramValues
func constructQueryOR(paramKey string, paramList []string, additionalPredicates ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// constructParameterJSONQuery construct and return query
func constructParameterJSONQuery(alias string, parameterFilters []ParameterFilterT) string {
	_ = "STUB: not implemented"
	// eg. query with optional destination_id (batch_rt_jobs_1.parameters @> '{"source_id":"<source_id>","destination_id":"<destination_id>"}'  OR (batch_rt_jobs_1.parameters @> '{"source_id":"<source_id>"}' AND batch_rt_jobs_1.parameters -> 'destination_id' IS NULL))
	return ""
}

// statTags is a struct to hold tags for stats
type statTags struct {
	CustomValFilters []string
	ParameterFilters []ParameterFilterT
	StateFilters     []string
	WorkspaceID      string
	MoreToken        bool
}

func (jd *Handle) getTimerStat(stat string, tags *statTags) stats.Measurement {
	_ = "STUB: not implemented"
	return *new(stats.Measurement)
}

func (tags *statTags) getStatsTags(tablePrefix string) stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}
