package testhelper

import (
	"database/sql"
	"testing"
)

const (
	TestKeyPairUnencrypted = "SNOWPIPE_STREAMING_KEYPAIR_UNENCRYPTED_INTEGRATION_TEST_CREDENTIALS"
	TestKeyPairEncrypted   = "SNOWPIPE_STREAMING_KEYPAIR_ENCRYPTED_INTEGRATION_TEST_CREDENTIALS"
)

type TestCredentials struct {
	Account              string `json:"account"`
	Warehouse            string `json:"warehouse"`
	User                 string `json:"user"`
	Role                 string `json:"role"`
	Database             string `json:"database"`
	PrivateKey           string `json:"privateKey"`
	PrivateKeyPassphrase string `json:"privateKeyPassphrase"`
}

func GetSnowpipeTestCredentials(key string) (*TestCredentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RandSchema() string { _ = "STUB: not implemented"; return "" }

func DropSchema(t testing.TB, db *sql.DB, namespace string) { _ = "STUB: not implemented"; return }
