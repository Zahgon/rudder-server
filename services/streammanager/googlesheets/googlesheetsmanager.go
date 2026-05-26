package googlesheets

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/tidwall/gjson"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/streammanager/common"
)

type Config struct {
	Credentials string              `mapstructure:"credentials"`
	SheetId     string              `mapstructure:"sheetId"`
	SheetName   string              `mapstructure:"sheetName"`
	EventKeyMap []map[string]string `mapstructure:"eventKeyMap"`
	TestConfig  TestConfig          `mapstructure:"testConfig"`
}

type TestConfig struct {
	Endpoint     string `mapstructure:"endpoint"`
	AccessToken  string `mapstructure:"accessToken"`
	RefreshToken string `mapstructure:"refreshToken"`
}

type Client struct {
	service *sheets.Service
	opts    common.Opts
}

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("streammanager").Child("googlesheets")
}

type GoogleSheetsProducer struct {
	config          Config
	client          *Client
	lock            sync.RWMutex
	isHeaderUpdated bool
}

// NewProducer creates a producer based on destination config
func NewProducer(destination *backendconfig.DestinationT, o common.Opts) (*GoogleSheetsProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If err is not nil then retrun

func (p *GoogleSheetsProducer) updateHeader() error { _ = "STUB: not implemented"; return nil }

// ** Preparing the Header Data **
// Creating the array of string which are then converted in to an array of interface which are to
// be added as header to each of the above spreadsheets.
// Example: | First Name | Last Name | Birth Day | Item Purchased | ..
// Here messageId is by default the first column

func (p *GoogleSheetsProducer) Produce(jsonData json.RawMessage, _ any) (statusCode int, respStatus, responseMessage string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

// insertHeaderDataToSheet inserts header data.
// Returns error for failure cases of API calls otherwise returns nil
func (p *GoogleSheetsProducer) insertHeaderDataToSheet(data []any) error {
	_ = "STUB: not implemented"
	// Creating value range for inserting row into sheet
	return nil
}

// insertRowDataToSheet appends row data list.
// Returns error for failure cases of API calls otherwise returns nil
func (p *GoogleSheetsProducer) insertRowDataToSheet(dataList [][]any) error {
	_ = "STUB: not implemented"
	// Creating value range for inserting row into sheet
	return nil
}

// parseTransformedData returns array of values from a json.
// source is the json object from transformer and we are iterating the json as a map
// and we are storing the data into designated position in array based on transformer
// mappings.
// Example payload we have from transformer without batching:
//
//	{
//			message:{
//				1: { attributeKey: "Product Purchased", attributeValue: "Realme C3" }
//				2: { attributeKey: "Product Value, attributeValue: "5900"}
//				..
//			}
//	}
//
// Example Payload we have from transformer with batching:
//
//	{
//			batch:[
//				{
//					message: {
//						1: { attributeKey: "Product Purchased", attributeValue: "Realme C3" }
//						2: { attributeKey: "Product Value, attributeValue: "5900"}
//						..
//					}
//				},
//				{
//					message: {
//						1: { attributeKey: "Product Purchased", attributeValue: "Realme C3" }
//						2: { attributeKey: "Product Value, attributeValue: "5900"}
//						..
//					}
//				}
//			]
//	}
func parseTransformedData(source gjson.Result) ([][]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Adding support for numeric type data

// getSheetsData is used to parse a string array to an interface array for compatibility
// with sheets-api
func getSheetsData(typedata []string) []any { _ = "STUB: not implemented"; return nil }

// handleServiceError is created for fail safety, if in any case when err type is not googleapi.Error
// server should not crash with a type error.
func handleServiceError(err error) (statusCode int, responseMessage string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func newOAuth2Client(config *Config) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func testClientOptions(config *Config) []option.ClientOption { _ = "STUB: not implemented"; return nil }

// skipcq: GO-S1020

// skipcq: GSC-G402

func realClientOptions(config *Config) ([]option.ClientOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareClientOptions(config *Config) ([]option.ClientOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// test configuration

func (*GoogleSheetsProducer) Close() error {
	_ = "STUB: not implemented"
	// no-op
	return nil
}
