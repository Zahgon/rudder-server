package lyticsBulkUpload

import (
	"github.com/tidwall/gjson"
)

func (u *LyticsBulkUploader) PopulateCsvFile(actionFile *ActionFileInfo, streamTraitsMapping []StreamTraitMapping, line string, data Data) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a map for quick lookups of LyticsProperty based on RudderProperty

// Unmarshal Properties into a map of json.RawMessage

// Initialize an empty CSV row

// Populate the CSV row based on streamTraitsMapping

// Convert the json.RawMessage value to a string

// Convert the json.RawMessage value to a string

// Handle if the value is directly a []byte

// If the value is already a string, use it directly

// Handle other types (e.g., numbers, booleans) by converting to a string

// Append an empty string if the RudderProperty is not found in fields

// Write the CSV header only once

// Write the CSV row

func createCSVWriter(fileName string) (*ActionFileInfo, error) {
	_ = "STUB: not implemented"
	// Open or create the file where the CSV will be written
	return nil, nil
}

// Create a new CSV writer using the file

// Return the ActionFileInfo struct with the CSV writer, file, and file path

func (u *LyticsBulkUploader) createCSVFile(existingFilePath string, streamTraitsMapping []StreamTraitMapping) (*ActionFileInfo, error) {
	_ = "STUB: not implemented"
	// Create a temporary directory using misc.GetTmpDir
	return nil, nil
}

// Define a local directory name within the temp directory

// Combine the temporary directory with the local directory name and generate a unique file path

// Initialize the CSV writer with the generated file path

// Ensure the file is closed when done

// Store the CSV file path in the ActionFileInfo struct

// Create a scanner to read the existing file line by line

// Adjust the buffer size if necessary

// Collect the failed job ID

// Calculate the payload size and observe it

// Populate the CSV file and collect success/failure job IDs

// After processing, calculate the final file size

func convertGjsonToStreamTraitMapping(result gjson.Result) []StreamTraitMapping {
	_ = "STUB: not implemented"
	return nil
}

// Iterate through the array in the result

// Continue iteration
