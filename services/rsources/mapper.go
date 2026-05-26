package rsources

func statusFromQueryResult(jobRunId string, statMap map[JobTargetKey]Stats) JobStatus {
	_ = "STUB: not implemented"
	return *new(JobStatus)
}

// task run id -> index
// task run id -> source id -> index

func failedRecordsFromQueryResult[T any](jobRunId string, recordsMap map[JobTargetKey][]T) JobFailedRecords[T] {
	_ = "STUB: not implemented"
	return nil
}

// task run id -> index
// task run id -> source id -> index
