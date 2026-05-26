package reporting

import (
	"regexp"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

const (
	responseKey = "response"
	errorKey    = "error"
	spaceStr    = " "

	errorsKey = "errors"
)

var (
	urlRegex         = regexp.MustCompile(`\b((?:https?://|www\.)\S+)\b`)
	ipRegex          = regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	emailRegex       = regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
	notWordRegex     = regexp.MustCompile(`\W+`)
	idRegex          = regexp.MustCompile(`\b([a-zA-Z0-9-_]*\d[a-zA-Z0-9-_]*)\b`)
	spaceRegex       = regexp.MustCompile(`\s+`)
	whitespacesRegex = regexp.MustCompile("[ \t\n\r]*") // used in checking if string is a valid json to remove extra-spaces

	defaultErrorMessageKeys = []string{"error_message", "message", "description", "detail", errorKey, "title"}
	deprecationKeywordSets  = map[string][][]string{
		"version": {
			{"action required", "api"},
			{"api", "removed"},
			{"api", "retired"},
			{"deprecated"},
			{"discontinued"},
			{"end of life"},
			{"end of service"},
			{"end of support"},
			{"expiring"},
			{"expired"},
			{"maintenance mode"},
			{"no longer available"},
			{"no longer supported"},
			{"not active"},
			{"outdated"},
			{"phased out"},
			{"please upgrade"},
			{"scheduled", "deprecation"},
			{"sunset"},
			{"support ending"},
			{"unsupported"},
			{"not supported"},
			{"upgrade", "required"},
		},
		"endpoint": {
			{"deprecated"},
			{"removed"},
			{"unsupported"},
			{"unavailable"},
			{"obsolete"},
			{"outdated"},
			{"not supported"},
			{"end of life"},
			{"end of service"},
			{"end of support"},
			{"expiring"},
			{"maintenance mode"},
			{"no longer available"},
			{"no longer supported"},
		},
		"api": {
			{"deprecated"},
			{"no longer supported"},
			{"end of life"},
			{"end of service"},
			{"end of support"},
			{"maintenance mode"},
			{"no longer available"},
			{"no longer supported"},
		},
	}
)

type ExtractorHandle struct {
	log              logger.Logger
	ErrorMessageKeys []string // the keys where in we may have error message
	maxMessageLength config.ValueLoader[int]
}

func NewErrorDetailExtractor(log logger.Logger, conf *config.Config) *ExtractorHandle {
	_ = "STUB: not implemented"
	return nil
}

// adding to default message keys

// Functions used for error message extraction -- STARTS
func checkForGoMapOrList(value any) bool { _ = "STUB: not implemented"; return false }

func (ext *ExtractorHandle) getSimpleMessage(sampleResponse string) string {
	_ = "STUB: not implemented"
	return ""
}

// First, try the specific key handlers (response, error, etc.)
// This handles nested JSON responses where the error message is in a "response" field

// If no specific keys were found, try to find message keys directly in the parsed JSON
// This handles cases where the JSON has a direct message field without a response wrapper
// This enhancement improves error extraction for various JSON response formats

func (ext *ExtractorHandle) handleKey(key string, value any) string {
	_ = "STUB: not implemented"
	return ""
}

// nolint:forbidigo

// Allow handleWarehouseError to process the value, regardless of its type

func handleError(valueStr string) string { _ = "STUB: not implemented"; return "" }

func (ext *ExtractorHandle) handleResponseOrErrorKey(valueStr string) string {
	_ = "STUB: not implemented"
	return ""
}

// isHTMLString checks if a string contains HTML content
func isHTMLString(s string) bool { _ = "STUB: not implemented"; return false }

// Check for common HTML patterns

func (ext *ExtractorHandle) handleWarehouseError(value any, key string) string {
	_ = "STUB: not implemented"
	return ""
}

// nolint:forbidigo

func getHTMLErrorMessage(erResStr string) string { _ = "STUB: not implemented"; return "" }

// truncateMessage truncates error message to the configured maximum length
func (ext *ExtractorHandle) truncateMessage(message string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ext *ExtractorHandle) GetErrorMessage(sampleResponse string) string {
	_ = "STUB: not implemented"
	return ""
}

func findKeys(keys []string, jsonObj any) map[string]any { _ = "STUB: not implemented"; return nil }

// recursively search for keys in nested JSON objects

// if jsonObj is a map

// if jsonObj is a slice

// return the map of keys and values

// This function takes a list of keys and a JSON object as input, and returns the value of the first key that exists in the JSON object.
func findFirstExistingKey(keys []string, jsonObj any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func getFirstNonNilValue(keys []string, jsonObj map[string]any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func convertInterfaceArrToStrArrWithDelimitter(arrI []any, delimitter string) string {
	_ = "STUB: not implemented"
	return ""
}

func getErrorMessageFromResponse(resp any, messageKeys []string) string {
	_ = "STUB: not implemented"
	return ""
}

func getErrorFromWarehouse(resp map[string]any) string { _ = "STUB: not implemented"; return "" }

func IsJSON(s string) bool { _ = "STUB: not implemented"; return false }

// Scenarios where we might have problems if the below logic is not included
// 1. Parsing of a string which contains { or [ at the start of the string could be parsed successfully
// 2. A valid with spacing before { or [ can also be deemed as not valid string

// We are making sure we remove white-spaces when we check the string for being an array or an object (scenario-2 is covered)

// We are making sure we check for end-braces for array or object(scenario-1 is covered)

func (ext *ExtractorHandle) CleanUpErrorMessage(errMsg string) string {
	_ = "STUB: not implemented"
	return ""
}

// Trim whitespace only

func getErrorCodeFromStatTags(statTags map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func containsDeprecationKey(errorMessage, key string) bool { _ = "STUB: not implemented"; return false }

func containsAllKeywords(errorMessage string, keywordSets [][]string) bool {
	_ = "STUB: not implemented"
	return false
}

func (ext *ExtractorHandle) isVersionDeprecationError(errorMessage string) bool {
	_ = "STUB: not implemented"
	// Normalize error message
	return false
}

func (ext *ExtractorHandle) GetErrorCode(errorMessage string, statTags map[string]string, destType string) string {
	_ = "STUB: not implemented"
	return ""
}

// Skip deprecation error detection for warehouse destinations
