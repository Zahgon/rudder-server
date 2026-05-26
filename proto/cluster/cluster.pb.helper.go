package proto

import (
	"github.com/rudderlabs/rudder-server/jobsdb"
)

func (x *JobsBatchChunk) GetJobsdbJobs() ([]*jobsdb.JobT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func JobFromJobsdbJob(j *jobsdb.JobT, _ int) *Job { _ = "STUB: not implemented"; return nil }
