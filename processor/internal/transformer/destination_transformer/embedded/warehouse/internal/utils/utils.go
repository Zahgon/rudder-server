package utils

import (
	"regexp"
	"time"

	"github.com/araddon/dateparse"

	"github.com/rudderlabs/rudder-go-kit/jsonrs"

	"github.com/rudderlabs/rudder-server/processor/internal/transformer/destination_transformer/embedded/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/processor/types"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

var (
	rudderCreatedTables                      = sliceToMap([]string{"tracks", "pages", "screens", "aliases", "groups", "accounts"})
	rudderIsolatedTables                     = sliceToMap([]string{"users", "identifies"})
	sourceCategoriesToUseRecordID            = sliceToMap([]string{"cloud", "singer-protocol"})
	identityEnabledWarehouses                = sliceToMap([]string{whutils.SNOWFLAKE, whutils.BQ})
	destinationSupportJSONPathAsPartOfConfig = sliceToMap([]string{whutils.POSTGRES, whutils.RS, whutils.SNOWFLAKE, whutils.SnowpipeStreaming, whutils.BQ})

	supportedJSONPathPrefixes     = []string{"track.", "identify.", "page.", "screen.", "alias.", "group.", "extract."}
	fullEventColumnTypeByDestType = map[string]string{
		whutils.SNOWFLAKE:         model.JSONDataType,
		whutils.SnowpipeStreaming: model.JSONDataType,
		whutils.RS:                model.TextDataType,
		whutils.BQ:                model.StringDataType,
		whutils.POSTGRES:          model.JSONDataType,
		whutils.MSSQL:             model.JSONDataType,
		whutils.AzureSynapse:      model.JSONDataType,
		whutils.CLICKHOUSE:        model.StringDataType,
		whutils.S3Datalake:        model.StringDataType,
		whutils.DELTALAKE:         model.StringDataType,
		whutils.GCSDatalake:       model.StringDataType,
		whutils.AzureDatalake:     model.StringDataType,
	}

	reDateTime = regexp.MustCompile(
		`^([+-]?\d{4})((-)((0[1-9]|1[0-2])(-([12]\d|0[1-9]|3[01])))([T\s]((([01]\d|2[0-3])((:)[0-5]\d))(:\d+)?)?(:[0-5]\d([.]\d+)?)?([zZ]|([+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)$`,
	)
	maxTimestampFormat             = "2012-08-03 18:31:59.257000000 +00:00 UTC"
	validTimestampFormatsMaxLength = len(maxTimestampFormat)

	minTimeInMs = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	maxTimeInMs = time.Date(9999, 12, 31, 23, 59, 59, 999000000, time.UTC)

	jsonrsStd = jsonrs.NewWithLibrary(jsonrs.StdLib)
)

func init() {
	_ = dateparse.MustParse(maxTimestampFormat)
}

func sliceToMap(slice []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func IsDataLake(destType string) bool { _ = "STUB: not implemented"; return false }

func IsRudderSources(event map[string]any) bool { _ = "STUB: not implemented"; return false }

func IsRudderCreatedTable(tableName string) bool { _ = "STUB: not implemented"; return false }

func IsRudderIsolatedTable(tableName string) bool { _ = "STUB: not implemented"; return false }

func IsObject(val any) bool { _ = "STUB: not implemented"; return false }

func IsJSONCompatibleStructure(val any) bool { _ = "STUB: not implemented"; return false }

func ToJSONCompatible(structVal any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func IsArray(val any) bool { _ = "STUB: not implemented"; return false }

func IsIdentityEnabled(destType string) bool { _ = "STUB: not implemented"; return false }

func CanUseRecordID(sourceCategory string) bool { _ = "STUB: not implemented"; return false }

func HasJSONPathPrefix(jsonPath string) bool { _ = "STUB: not implemented"; return false }

func GetFullEventColumnTypeByDestType(destType string) string { _ = "STUB: not implemented"; return "" }

func ValidTimestamp(input string) bool { _ = "STUB: not implemented"; return false }

func ToTimestamp(val any) any { _ = "STUB: not implemented"; return *new(any) }

// parseTimestamp parses a timestamp string into time.Time.
// If it fails due to a "day out of range" error, it falls back to normalizing the date.
// JS automatically handles this https://www.programiz.com/online-compiler/4gfcMEByAur4q
// console.log(new Date('1988-04-31').toISOString()); // 1988-05-01T00:00:00.000Z
func parseTimestamp(input string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// ToString converts any value to a string representation.
// - If the value is nil, it returns an empty string.
// - If the value implements the fmt.Stringer interface, it returns the result of the String() method.
// - Otherwise, it returns a string representation using fmt.Sprintf.
func ToString(value any) string { _ = "STUB: not implemented"; return "" }

// IsEmptyString checks if the given value is considered "blank."
// - A value is considered blank if its string representation is an empty string.
// - The function first converts the value to its string representation using ToString and checks if its length is zero.
//
// Corresponding lodash implementation: https://playcode.io/2371369
//
//	const isBlank = (value) => {
//	 try {
//	   return _.isEmpty(_.toString(value));
//	 } catch (e) {
//	   console.error(`Error in isBlank: ${e.message}`);
//	   return false;
//	 }
//	};
func IsEmptyString(value any) bool { _ = "STUB: not implemented"; return false }

func IsJSONPathSupportedAsPartOfConfig(destType string) bool {
	_ = "STUB: not implemented"
	return false
}

func ExtractMessageID(event *types.TransformerEvent, uuidGenerator func() string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func ExtractReceivedAt(event *types.TransformerEvent, now func() time.Time) string {
	_ = "STUB: not implemented"
	return ""
}

// MarshalJSON marshals the input to JSON. It escapes HTML characters (e.g. &, <, and > from \u0026, \u003c, and \u003e) by default.
// It also trims the output to avoid trailing spaces.
func MarshalJSON(input any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UTF16RuneCountInString returns the UTF-16 code unit count of the string.
func UTF16RuneCountInString(s string) int { _ = "STUB: not implemented"; return 0 }
