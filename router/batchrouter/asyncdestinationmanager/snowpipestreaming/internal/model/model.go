package model

import (
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type (
	CreateChannelRequest struct {
		RudderIdentifier string        `json:"rudderIdentifier"`
		Partition        string        `json:"partition"`
		AccountConfig    AccountConfig `json:"account"`
		TableConfig      TableConfig   `json:"table"`
	}
	AccountConfig struct {
		Account              string `json:"account"`
		User                 string `json:"user"`
		Role                 string `json:"role"`
		PrivateKey           string `json:"privateKey"`
		PrivateKeyPassphrase string `json:"privateKeyPassphrase"`
	}
	TableConfig struct {
		Database      string `json:"database"`
		Schema        string `json:"schema"`
		Table         string `json:"table"`
		EnableIceberg bool   `json:"enableIceberg"`
	}

	ChannelResponse struct {
		Success                bool                     `json:"success"`
		ChannelID              string                   `json:"channelId"`
		ChannelName            string                   `json:"channelName"`
		ClientName             string                   `json:"clientName"`
		Valid                  bool                     `json:"valid"`
		Deleted                bool                     `json:"deleted"`
		SnowpipeSchema         whutils.ModelTableSchema `json:"-"`
		Error                  string                   `json:"error"`
		Code                   string                   `json:"code"`
		SnowflakeSDKCode       string                   `json:"snowflakeSDKCode"`
		SnowflakeAPIHttpCode   int64                    `json:"snowflakeAPIHttpCode"`
		SnowflakeAPIStatusCode int64                    `json:"snowflakeAPIStatusCode"`
		SnowflakeAPIMessage    string                   `json:"snowflakeAPIMessage"`
	}

	ColumnInfo struct {
		Type  *string  `json:"type,omitempty"`
		Scale *float64 `json:"scale,omitempty"`
	}

	InsertRequest struct {
		Rows   []Row  `json:"rows"`
		Offset string `json:"offset"`
	}
	Row map[string]any

	InsertResponse struct {
		Success bool          `json:"success"`
		Errors  []InsertError `json:"errors"`
		Code    string        `json:"code"`
	}
	InsertError struct {
		RowIndex                    int64    `json:"rowIndex"`
		ExtraColNames               []string `json:"extraColNames"`
		MissingNotNullColNames      []string `json:"missingNotNullColNames"`
		NullValueForNotNullColNames []string `json:"nullValueForNotNullColNames"`
		Message                     string   `json:"message"`
	}

	StatusResponse struct {
		Success              bool   `json:"success"`
		Offset               string `json:"offset"`
		Valid                bool   `json:"valid"`
		LatestInsertedOffset string `json:"latestInsertedOffset"`
	}

	BulkStatusResponse struct {
		Success  bool                       `json:"success"`
		Statuses map[string]*StatusResponse `json:"statuses"`
		NotFound []string                   `json:"notFound"`
	}
)

func (c *ChannelResponse) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Prevent recursion

func generateSnowpipeSchema(tableSchema map[string]ColumnInfo) whutils.ModelTableSchema {
	_ = "STUB: not implemented"
	return *new(whutils.ModelTableSchema)
}

func cleanDataType(input string) string {
	_ = "STUB: not implemented"
	// Extract the portion before the first '('
	return ""
}

// Return as-is if no '(' is found
