package yandexmetrica

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

var idClientMap = map[string]string{
	"ClientId": "CLIENT_ID",
	"Yclid":    "YCLID",
	"UserId":   "USER_ID",
}

type yandexMetricaMessageBody struct {
	ClientID any     `json:"ClientId"`
	YclID    any     `json:"Yclid"`
	UserID   any     `json:"UserId"`
	Target   string  `json:"Target"`
	DateTime string  `json:"DateTime"`
	Price    float64 `json:"Price"`
	Currency string  `json:"Currency"`
}

type yandexMetricaMessage struct {
	Message yandexMetricaMessageBody `json:"message"`
}

type idStruct struct {
	id         string
	clientType string
	headerName string
}

func (ym yandexMetricaMessageBody) ID() (idStruct, error) {
	_ = "STUB: not implemented"
	return *new(idStruct), nil
}

func getID(id any, headerName string) (idStruct, error) {
	_ = "STUB: not implemented"
	return *new(idStruct), nil
}

type YandexMetricaBulkUploader struct {
	logger       logger.Logger
	statsFactory stats.Stats
	Client       *http.Client
	destination  *backendconfig.DestinationT
}

func NewManager(conf *config.Config, logger logger.Logger, statsFactory stats.Stats, destination *backendconfig.DestinationT, backendConfig backendconfig.BackendConfig) (*YandexMetricaBulkUploader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This client is used for uploading data to yandex metrica

// Poll return a success response for the poll request every time by default
func (ym *YandexMetricaBulkUploader) Poll(_ context.Context, _ common.AsyncPoll) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

// GetUploadStats return a success response for the getUploadStats request every time by default
func (ym *YandexMetricaBulkUploader) GetUploadStats(_ common.GetUploadStatsInput) common.GetUploadStatsResponse {
	_ = "STUB: not implemented"
	return *new(common.GetUploadStatsResponse)
}

func generateCSVFromJSON(jsonData []byte, goalId string) (string, string, error) {
	_ = "STUB: not implemented"
	// Define an empty map to store the parsed JSON data
	return "", "", nil
}

// Open the CSV file for writing

// Create a CSV writer

// Define the header row based on key presence in "message" object

// Write the header row

// Extract and write data rows

// Flush the writer

// Return the chosen header

func copyDataIntoBuffer(csvFilePath string) (*bytes.Buffer, *multipart.Writer, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ym *YandexMetricaBulkUploader) uploadFileToDestination(uploadURL, csvFilePath, userIdType string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ym *YandexMetricaBulkUploader) generateErrorOutput(errorString string, err error, importingJobIds []int64) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

func (*YandexMetricaBulkUploader) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ym *YandexMetricaBulkUploader) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// extract counterId from destConfigJson as a string value

// We don't need to handle it, as we can receive a string response even before executing OAuth operations like Refresh Token.
// It's acceptable if the structure of respData doesn't match the oauthv2.TransportResponse struct.

// re-assign originalResponse

// error scenario
