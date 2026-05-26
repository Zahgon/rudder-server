package eloqua

type EloquaServiceImpl struct {
	bulkApi string
}

func NewEloquaServiceImpl(version string) *EloquaServiceImpl { _ = "STUB: not implemented"; return nil }

func (e *EloquaServiceImpl) MakeHTTPRequest(data *HttpRequestData) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (e *EloquaServiceImpl) GetBaseEndpoint(data *HttpRequestData) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *EloquaServiceImpl) FetchFields(data *HttpRequestData) (*Fields, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EloquaServiceImpl) CreateImportDefinition(data *HttpRequestData, eventType string) (*ImportDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EloquaServiceImpl) UploadData(data *HttpRequestData, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *EloquaServiceImpl) UploadDataWithoutCSV(data *HttpRequestData, uploadData []map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *EloquaServiceImpl) RunSync(data *HttpRequestData) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *EloquaServiceImpl) CheckSyncStatus(data *HttpRequestData) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *EloquaServiceImpl) CheckRejectedData(data *HttpRequestData) (*RejectResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EloquaServiceImpl) DeleteImportDefinition(data *HttpRequestData) error {
	_ = "STUB: not implemented"
	return nil
}
