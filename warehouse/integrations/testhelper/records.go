package testhelper

// SampleTestRecordsTemplate returns a set of records for testing default loading scenarios.
// It uses testdata/load.template as the source of data.
func SampleTestRecordsTemplate(recordSetIndex int) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// SampleTestRecords returns a set of records for testing default loading scenarios.
// It uses testdata/load.* as the source of data.
func SampleTestRecords() [][]string { _ = "STUB: not implemented"; return nil }

// AppendTestRecords returns a set of records for testing append scenarios.
// It uses testdata/load.* twice as the source of data.
func AppendTestRecords() [][]string { _ = "STUB: not implemented"; return nil }

// DiscardTestRecords returns a set of records for testing rudder discards.
// It uses testdata/discards.* as the source of data.
func DiscardTestRecords() [][]string { _ = "STUB: not implemented"; return nil }

// DedupTestRecords returns a set of records for testing deduplication scenarios.
// It uses testdata/dedup.* as the source of data.
func DedupTestRecords() [][]string { _ = "STUB: not implemented"; return nil }

// DedupTwiceTestRecords returns a set of records for testing deduplication scenarios.
// It uses testdata/dedup.* twice as the source of data.
func DedupTwiceTestRecords() [][]string { _ = "STUB: not implemented"; return nil }

// MismatchSchemaTestRecords returns a set of records for testing schema mismatch scenarios.
// It uses testdata/mismatch-schema.* as the source of data.
func MismatchSchemaTestRecords() [][]string { _ = "STUB: not implemented"; return nil }

// UploadJobIdentifiesRecords returns a set of records for testing upload job identifies scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobIdentifiesRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobIdentifiesAppendRecords returns a set of records for testing upload job identifies scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobIdentifiesAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobIdentifiesMergeRecords returns a set of records for testing upload job identifies scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobIdentifiesMergeRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobUsersRecords returns a set of records for testing upload job users scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobUsersRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobUsersRecordsForDatalake returns a set of records for testing upload job users scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
// sent_at, timestamp, original_timestamp will not be present in the records.
func UploadJobUsersRecordsForDatalake(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobUsersAppendRecords returns a set of records for testing upload job users scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobUsersAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobUsersAppendRecordsUsingUsersLoadFiles returns a set of records for testing upload job users scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobUsersAppendRecordsUsingUsersLoadFiles(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobUsersMergeRecord returns a set of records for testing upload job users scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobUsersMergeRecord(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobUsersRecordsUsingUsersLoadFilesForClickhouse returns a set of records for testing upload job users scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
// For AggregatingMergeTree ClickHouse replaces all rows with the same primary key (or more accurately, with the same sorting key) with a single row (within a one data part) that stores a combination of states of aggregate functions.
// So received_at will be record for the first record only.
func UploadJobUsersRecordsUsingUsersLoadFilesForClickhouse(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobTracksRecords returns a set of records for testing upload job tracks scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobTracksRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobTracksAppendRecords returns a set of records for testing upload job tracks scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobTracksAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobTracksMergeRecords returns a set of records for testing upload job tracks scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobTracksMergeRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobProductTrackRecords returns a set of records for testing upload job product track scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobProductTrackRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobProductTrackAppendRecords returns a set of records for testing upload job product track scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobProductTrackAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobProductTrackMergeRecords returns a set of records for testing upload job product track scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobProductTrackMergeRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobPagesRecords returns a set of records for testing upload job pages scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobPagesRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobPagesAppendRecords returns a set of records for testing upload job pages scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobPagesAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobPagesMergeRecords returns a set of records for testing upload job pages scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobPagesMergeRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobScreensRecords returns a set of records for testing upload job screens scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobScreensRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobScreensAppendRecords returns a set of records for testing upload job screens scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobScreensAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobScreensMergeRecords returns a set of records for testing upload job screens scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobScreensMergeRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobAliasesRecords returns a set of records for testing upload job aliases scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobAliasesRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobAliasesAppendRecords returns a set of records for testing upload job aliases scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobAliasesAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobAliasesMergeRecords returns a set of records for testing upload job aliases scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobAliasesMergeRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobGroupsRecords returns a set of records for testing upload job groups scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func UploadJobGroupsRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobGroupsAppendRecords returns a set of records for testing upload job groups scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobGroupsAppendRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// UploadJobGroupsMergeRecords returns a set of records for testing upload job groups scenarios.
// It uses twice upload-job.events-1.json as the source of data.
func UploadJobGroupsMergeRecords(userIDFormat, sourceID, destID, destType string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// SourceJobTracksRecords returns a set of records for testing source job tracks scenarios.
// It uses source-job.events-1.json, source-job.events-2.json as the source of data.
func SourceJobTracksRecords(userIDFormat, sourceID, destID, destType, jobRunID, taskRunID string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// SourceJobGoogleSheetRecords returns a set of records for testing source job google sheet scenarios.
// It uses upload-job.events-1.json, upload-job.events-2.json as the source of data.
func SourceJobGoogleSheetRecords(userIDFormat, sourceID, destID, destType, jobRunID, taskRunID string) [][]string {
	_ = "STUB: not implemented"
	return nil
}
