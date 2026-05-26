package misc

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/rudderlabs/rudder-go-kit/jsonrs"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

var (
	reservedFolderPaths []*RFP

	regexGwHa               = regexp.MustCompile(`^.*-gw-ha-\d+-\w+-\w+$`)
	regexGwNonHaOrProcessor = regexp.MustCompile(`^.*-\d+$`)
)

const (
	// RFC3339Milli with milli sec precision
	RFC3339Milli          = "2006-01-02T15:04:05.000Z07:00"
	NOTIMEZONEFORMATPARSE = "2006-01-02T15:04:05"
)

const (
	RudderAsyncDestinationLogs    = "rudder-async-destination-logs"
	RudderArchives                = "rudder-archives"
	RudderWarehouseStagingUploads = "rudder-warehouse-staging-uploads"
	RudderRawDataDestinationLogs  = "rudder-raw-data-destination-logs"
	RudderWarehouseLoadUploadsTmp = "rudder-warehouse-load-uploads-tmp"
	RudderIdentityMergeRulesTmp   = "rudder-identity-merge-rules-tmp"
	RudderIdentityMappingsTmp     = "rudder-identity-mappings-tmp"
	RudderRedshiftManifests       = "rudder-redshift-manifests"
	RudderWarehouseJsonUploadsTmp = "rudder-warehouse-json-uploads-tmp"
	RudderTestPayload             = "rudder-test-payload"
	RudderWarehouseGRPCDownloads  = "rudder-warehouse-grpc-downloads"
	RudderReportingErrorIndex     = "rudder-reporting-error-index"
)

// ErrorStoreT : DS to store the app errors
type ErrorStoreT struct {
	Errors []RudderError
}

// RudderError : to store rudder error
type RudderError struct {
	StartTime         int64
	CrashTime         int64
	ReadableStartTime string
	ReadableCrashTime string
	Message           string
	StackTrace        string
	Code              int
}

type RFP struct {
	path         string
	levelsToKeep int
}

var pkgLogger logger.Logger

func init() {
	uuid.EnableRandPool()
}

func Init() { _ = "STUB: not implemented"; return }

func BatchDestinations() []string { _ = "STUB: not implemented"; return nil }

func GetHash(s string) int { _ = "STUB: not implemented"; return 0 }

// GetMD5Hash returns EncodeToString(md5 hash of the input string)
func GetMD5Hash(input string) string { _ = "STUB: not implemented"; return "" }

// skipcq: GO-S1023

// RemoveFilePaths removes filePaths as well as cleans up the empty folder structure.
func RemoveFilePaths(filePaths ...string) { _ = "STUB: not implemented"; return }

// GetReservedFolderPaths returns all temporary folder paths.
func GetReservedFolderPaths() []*RFP { _ = "STUB: not implemented"; return nil }

func checkMatch(currDir string) bool { _ = "STUB: not implemented"; return false }

func (r *RFP) matches(currDir string) (match bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// RemoveContents removes all the contents of the directory
func RemoveContents(dir string) error { _ = "STUB: not implemented"; return nil }

// RemoveEmptyFolderStructureForFilePath recursively cleans up everything till it reaches the stage where the folders are not empty or parent.
func RemoveEmptyFolderStructureForFilePath(fp string) { _ = "STUB: not implemented"; return }

// Checking if the currDir is present in the temporary folders or not
// If present we should stop at that point.

var logOnce sync.Once

// GetTmpDir gets tmp dir at path configured via RUDDER_TMPDIR env var
func GetTmpDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// second chance: fallback to /tmp if this folder exists

// Copy copies the exported fields from src to dest
// Used for copying the default transport
func Copy(dst, src any) { _ = "STUB: not implemented"; return }

// First src and dst must be pointers, so that dst can be assignable.

// Then src must be assignable to dst and both must be structs (but this is
// already guaranteed).

// Finally, copy all exported fields.  Since the types are the same, we
// have no problems and we only have to ignore unexported fields.

// Unexported field.

// Returns chronological timestamp of the event using the formula
// timestamp = receivedAt - (sentAt - originalTimestamp)
func GetChronologicalTimeStamp(receivedAt, sentAt, originalTimestamp time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func TruncateStr(str string, limit int) string { _ = "STUB: not implemented"; return "" }

// TailTruncateStr returns the last `count` digits of a string
func TailTruncateStr(str string, count int) string { _ = "STUB: not implemented"; return "" }

func ReplaceMultiRegex(str string, expList map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ConvertStringInterfaceToIntArray(interfaceT any) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeHTTPRequestWithTimeout(url string, payload io.Reader, timeout time.Duration) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func ConvertInterfaceToStringArray(input []any) []string { _ = "STUB: not implemented"; return nil }

func HTTPCallWithRetryWithTimeout(url string, payload []byte, timeout time.Duration) ([]byte, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func IntArrayToString(a []int64, delim string) string { _ = "STUB: not implemented"; return "" }

func MakeJSONArray(bytesArray [][]byte) []byte { _ = "STUB: not implemented"; return nil }

// insert '[' to the front

// append ']'

func SingleQuoteLiteralJoin(slice []string) string { _ = "STUB: not implemented"; return "" }

// TODO: use strings.Join() instead

type BufferedWriter struct {
	File   *os.File
	Writer *bufio.Writer
}

func CreateBufferedWriter(s string) (w BufferedWriter, err error) {
	_ = "STUB: not implemented"
	return *new(BufferedWriter), nil
}

func (b BufferedWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b BufferedWriter) GetFile() *os.File { _ = "STUB: not implemented"; return nil }

func (b BufferedWriter) Close() error { _ = "STUB: not implemented"; return nil }

type GZipWriter struct {
	File      *os.File
	GzWriter  *gzip.Writer
	BufWriter *bufio.Writer
}

func CreateGZ(s string) (w GZipWriter, err error) {
	_ = "STUB: not implemented"
	return *new(GZipWriter), nil
}

func (w GZipWriter) WriteGZ(s string) error { _ = "STUB: not implemented"; return nil }

func (w GZipWriter) Write(b []byte) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (GZipWriter) WriteRow(_ []any) error { _ = "STUB: not implemented"; return nil }

func (w GZipWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (w GZipWriter) GetLoadFile() *os.File { _ = "STUB: not implemented"; return nil }

func (w GZipWriter) CloseGZ() error { _ = "STUB: not implemented"; return nil }

func GetMacAddress() string {
	_ = "STUB: not implemented"
	// ----------------------
	// Get the local machine IP address
	// https://www.socketloop.com/tutorials/golang-how-do-I-get-the-local-ip-non-loopback-address
	// ----------------------
	return ""
}

// check the address type and if it is not a loopback then that's the current ip

// get all the system's or local machine's network interfaces

// only interested in the name with current IP address

// extract the hardware information base on the interface name captured above

/*
RunWithTimeout runs provided function f until provided timeout d.
If the timeout is reached, onTimeout callback will be called.
*/
func RunWithTimeout(f, onTimeout func(), d time.Duration) { _ = "STUB: not implemented"; return }

/*
IsValidUUID will check if provided string is a valid UUID
*/
func IsValidUUID(uuid string) bool { _ = "STUB: not implemented"; return false }

func FastUUID() uuid.UUID { _ = "STUB: not implemented"; return *new(uuid.UUID) }

func HasAWSRoleARNInConfig(configMap map[string]any) bool { _ = "STUB: not implemented"; return false }

func HasAWSKeysInConfig(config any) bool { _ = "STUB: not implemented"; return false }

func HasAWSRegionInConfig(config any) bool { _ = "STUB: not implemented"; return false }

func GetRudderObjectStorageAccessKeys() (accessKeyID, accessKey string) {
	_ = "STUB: not implemented"
	return "", ""
}

func GetRudderObjectStoragePrefix() (prefix string) { _ = "STUB: not implemented"; return "" }

func GetRegionHint() string { _ = "STUB: not implemented"; return "" }

func GetRudderObjectStorageConfig(prefixOverride string) (storageConfig map[string]any) {
	_ = "STUB: not implemented"
	// TODO: add error log if s3 keys are not available
	return nil
}

// set prefix from override for shared slave type nodes

func IsConfiguredToUseRudderObjectStorage(storageConfig map[string]any) bool {
	_ = "STUB: not implemented"
	return false
}

type ObjectStorageOptsT struct {
	Provider                    string
	Config                      any
	UseRudderStorage            bool
	RudderStoragePrefixOverride string
	WorkspaceID                 string
}

func GetObjectStorageConfig(opts ObjectStorageOptsT) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// GetParsedTimestamp returns the parsed timestamp
func GetParsedTimestamp(input any) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// GetTagName gets the tag name using a uuid and name
func GetTagName(id string, names ...string) string { _ = "STUB: not implemented"; return "" }

// UpdateJSONWithNewKeyVal enhances the json passed with key, val
func UpdateJSONWithNewKeyVal(params []byte, key string, val any) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ConcatErrors(givenErrors []error) error { _ = "STUB: not implemented"; return nil }

func isWarehouseMasterEnabled() bool { _ = "STUB: not implemented"; return false }

func GetWarehouseURL() (url string) { _ = "STUB: not implemented"; return "" }

type MapLookupError struct {
	SearchKey string // indicates the searchkey which is not present in the map
	Err       error  // contains the error occurred string while looking up the key in the map
	Level     int    // indicates the nesting level at which error has occurred
}

func (e *MapLookupError) Error() string { _ = "STUB: not implemented"; return "" }

// NestedMapLookup
// m:  a map from strings to other maps or values, of arbitrary depth
// ks: successive keys to reach an internal or leaf node (variadic)
// If an internal node is reached, will return the internal map
//
// Returns: (Exactly one of these will be nil)
// rval: the target node (if found)
// err:  an error created by fmt.Errorf
func NestedMapLookup(m map[string]any, ks ...string) (any, *MapLookupError) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// degenerate input

// we've reached the final key

// 1+ more keys

// SleepCtx sleeps for the given duration or until the context is canceled.
//
//	the context error is returned if context is canceled.
func SleepCtx(ctx context.Context, delay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func Unique(stringSlice []string) []string { _ = "STUB: not implemented"; return nil }

// MapLookup returns the value of the key in the map, or nil if the key is not present.
//
// If multiple keys are provided then it looks for nested maps recursively.
func MapLookup(mapToLookup map[string]any, keys ...string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func CopyStringMap(originalMap map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func GetDiskUsageOfFile(path string) (int64, error) {
	_ = "STUB: not implemented"
	// Notes
	// 1. stat.Blocks is the number of stat.Blksize blocks allocated to the file
	// 2. stat.Blksize is the filesystem block size for this filesystem
	// 3. We compute the actual disk usage of a (sparse) file by multiplying the number of blocks allocated to the file with the block size. This computes a different value than the one returned by stat.Size particularly for sparse files.
	return 0, nil
}

//nolint:unconvert // In amd64 architecture stat.Blksize is int64 whereas in arm64 it is int32

// DiskUsage calculates the path's disk usage recursively in bytes. If exts are provided, only files with matching extensions will be included in the result.
func DiskUsage(path string, ext ...string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func GetBadgerDBUsage(dir string) (int64, int64, int64, error) {
	_ = "STUB: not implemented"
	// Notes
	// Instead of using BadgerDB's internal function to get the disk usage, we are writing our own implementation because of the following reasons:
	// 1. BadgerDB internally creates a sparse memory backed file to store the data
	// 2. The size returned by the filepath.Walk used internally gives a misleading size because the file is mostly empty and doesn't consume any disk space
	return 0, 0, 0, nil
}

func GetInstanceID() string { _ = "STUB: not implemented"; return "" }

// This handles 2 kinds of server instances
// a) Processor OR Gateway running in non HA mod where the instance name ends with the index
// b) Gateway running in HA mode, where the instance name is of the form *-gw-ha-<index>-<statefulset-id>-<pod-id>

// explicitly using json-iter due to its unique behaviour with respect to handling invalid characters
var jsonfast jsonrs.JSON = jsonrs.NewWithLibrary(jsonrs.JsoniterLib)

// SanitizeJSON makes a json payload safe for writing into postgres.
// 1. Removes any \u0000 string from the payload
// ~2. Replaces any invalid utf8 characters using github.com/rudderlabs/rudder-go-kit/utf8~
// 3. unmarshals and marshals the payload to remove any extra keys
func SanitizeJSON(input json.RawMessage) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	// Remove null characters
	return *new(json.RawMessage), nil
}

// Validate JSON structure by unmarshaling and marshaling

func SanitizeString(input string) string { _ = "STUB: not implemented"; return "" }

// GetMurmurHash returns murmur3 hash of the input string with a default seed of 0
func GetMurmurHash(input string) uint64 { _ = "STUB: not implemented"; return 0 }

// GetMurmurHashWithSeed returns murmur3 hash of the input string with the provided seed
func GetMurmurHashWithSeed(input string, seed uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// TruncatedList returns a comma-separated string of items,
// showing at most maxShow items followed by "+N more" if truncated.
func TruncatedList(items []string, maxShow int) string { _ = "STUB: not implemented"; return "" }
