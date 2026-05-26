package salesforcebulkupload

import (
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

func extractObjectInfo(jobs []common.AsyncJob) (*ObjectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractFromVDM(externalIDRaw any) (*ObjectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createCSVFile(
	destinationID string,
	input []common.AsyncJob,
) (string, []string, map[string][]int64, error) {
	_ = "STUB: not implemented"
	return "", nil, nil, nil
}

func calculateHashCode(row []string) string { _ = "STUB: not implemented"; return "" }

func calculateHashFromRecord(record map[string]string, csvHeaders []string) string {
	_ = "STUB: not implemented"
	return ""
}
