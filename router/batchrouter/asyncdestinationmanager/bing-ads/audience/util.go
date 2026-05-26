package audience

// Upload related utils

// returns the clientID struct
func newClientID(jobID int64, hashedEmail string) ClientID {
	_ = "STUB: not implemented"
	return *new(ClientID)
}

/*
returns the csv file and zip file path, along with the csv writer that
contains the template of the uploadable file.
*/
func createActionFile(audienceId, actionType string) (*ActionFileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertCsvToZip(actionFile *ActionFileInfo) error { _ = "STUB: not implemented"; return nil }

// Close the ZIP writer

// Remove the csv file after creating the zip file

// populateZipFile only if it is within the file size limit 100mb and row number limit 4000000
// Otherwise event is appended to the failedJobs and will be retried.
func (b *BingAdsBulkUploader) populateZipFile(actionFile *ActionFileInfo, audienceId, line string, data Data) error {
	_ = "STUB: not implemented"
	return nil
}

/*
Depending on add, remove and update action we are creating 3 different zip files using this function
It is also returning the list of succeed and failed events lists.
The following map indicates the index->actionType mapping
0-> Add
1-> Remove
2-> Update
*/
func (b *BingAdsBulkUploader) createZipFile(filePath, audienceId string) ([]*ActionFileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Poll Related Utils

/*
From the ResultFileUrl, it downloads the zip file and extracts the contents of the zip file
and finally Provides file paths containing error information as an array string
*/
func (b *BingAdsBulkUploader) downloadAndGetUploadStatusFile(ResultFileUrl string) ([]string, error) {
	_ = "STUB: not implemented"
	// the final status file needs to be downloaded
	return nil, nil
}

// Create output directory if it doesn't exist

// Download the zip file

// Create a temporary file to save the downloaded zip file

// Save the downloaded zip file to the temporary file

// Extract the contents of the zip file to the output directory

// unzips the file downloaded from bingads, which contains error informations
// of a particular event.
func unzip(zipFile, targetDir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Open each file in the zip archive

// Create the corresponding file in the target directory

// Create directories if the file is a directory

// Create the file and copy the contents

// Append the file path to the list

/*
ReadPollResults reads the CSV file and returns the records
In the below format (only adding relevant keys)

	[][]string{
		{"Client Id", "Error", "Type"},
		{"1<<>>client1", "error1", "Customer List Error"},
		{"1<<>>client2", "error1", "Customer List Item Error"},
		{"1<<>>client2", "error2", "Customer List Item Error"},
	}
*/
func (b *BingAdsBulkUploader) readPollResults(filePath string) ([][]string, error) {
	_ = "STUB: not implemented"
	// Open the CSV file
	return nil, nil
}

// defer file.Close() and remove

// remove the file after the response has been written

// Create a new CSV reader

// Read all records from the CSV file

// converting the string clientID to ClientID struct

func newClientIDFromString(clientID string) (*ClientID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
records is the output of ReadPollResults function which is in the below format

	[][]string{
		{"Client Id", "Error", "Type"},
		{"1<<>>client1", "error1", "Customer List Error"},
		{"1<<>>client2", "error1", "Customer List Item Error"},
		{"1<<>>client2", "error2", "Customer List Item Error"},
	}

This function processes the CSV records and returns the JobIDs and the corresponding error messages
In the below format:

	map[string]map[string]struct{}{
		"1": {
			"error1": {},
		},
		"2": {
			"error1": {},
			"error2": {},
		},
	}

** we are using map[int64]map[string]struct{} for storing the error messages
** because we want to avoid duplicate error messages
*/
func processPollStatusData(records [][]string) (map[int64]map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Declare variables for storing data

// Iterate over the remaining rows and filter based on the 'Type' field containing the substring 'Error'
// The error messages are present on the rows where the corresponding Type column values are "Customer List Error", "Customer List Item Error" etc

// expecting the client ID is present as jobId<<>>clientId

// making the structure as jobId: [error1, error2]

// GetUploadStats Related utils

// get the list of unique error messages for a particular jobId.
func getAbortedReasons(clientIDErrors map[int64]map[string]struct{}) map[int64]string {
	_ = "STUB: not implemented"
	return nil
}

// filtering out failed jobIds from the total array of jobIds
// in order to get jobIds of the successful jobs
func getSuccessJobIDs(failedEventList, initialEventList []int64) []int64 {
	_ = "STUB: not implemented"
	return nil
}
