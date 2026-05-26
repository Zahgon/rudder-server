package common

var (
	asyncDestinations = []string{"MARKETO_BULK_UPLOAD", "BINGADS_AUDIENCE", "ELOQUA", "YANDEX_METRICA_OFFLINE_EVENTS", "BINGADS_OFFLINE_CONVERSIONS", "KLAVIYO_BULK_UPLOAD", "LYTICS_BULK_UPLOAD", "SNOWPIPE_STREAMING", "SALESFORCE_BULK_UPLOAD"}
	sftpDestinations  = []string{"SFTP"}
)

func IsSFTPDestination(destination string) bool { _ = "STUB: not implemented"; return false }

func IsAsyncRegularDestination(destination string) bool { _ = "STUB: not implemented"; return false }

func IsAsyncDestination(destination string) bool { _ = "STUB: not implemented"; return false }

// FormatCSVValue stringifies a JSON-derived value for a CSV cell.
// Top-level nil renders as an empty cell so destinations that treat empty
// cells as null (e.g. Salesforce Bulk) get the expected semantics. Floats
// are rendered without scientific notation, and arrays/maps are emitted
// as JSON so nested numbers stay plain and nested nulls become `null`.
// Returns an error when JSON marshalling of a composite value fails.
func FormatCSVValue(value any) (string, error) { _ = "STUB: not implemented"; return "", nil }
