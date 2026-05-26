package eloqua

import (
	"os"

	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

const bufferSize = 5000 * 1024

func getEventDetails(file *os.File) (*EventDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getKeys(m map[string]any) []string { _ = "STUB: not implemented"; return nil }

func createCSVFile(fields []string, file *os.File, uploadJobInfo *JobInfo, jobIdRowMap map[int64]int64) (string, int64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func createBodyForImportDefinition(eventDetails *EventDetails, eloquaFields *Fields) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateErrorString(item RejectedItem) string { _ = "STUB: not implemented"; return "" }

func parseRejectedData(data *HttpRequestData, importingList []*jobsdb.JobT, eloqua *EloquaBulkUploader) (*common.EventStatMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/**
We have a limit of 1000 rejected events per api call. And the offset starts from 0. So we are fetching 1000 rejected events every time making the api call and updating the offset accordingly.
*/

func parseFailedData(syncId string, importingList []*jobsdb.JobT) *common.EventStatMeta {
	_ = "STUB: not implemented"
	return nil
}

func getUniqueKeys(eloquaFields *Fields) []string { _ = "STUB: not implemented"; return nil }
