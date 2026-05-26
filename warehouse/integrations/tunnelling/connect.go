package tunnelling

import (
	"database/sql"
	"errors"

	stunnel "github.com/rudderlabs/sql-tunnels/driver/ssh"
)

var (
	ErrMissingKey     = errors.New("missing mandatory key")
	ErrUnexpectedType = errors.New("unexpected type")
)

const (
	sshUser       = "sshUser"
	sshPort       = "sshPort"
	sshHost       = "sshHost"
	sshPrivateKey = "sshPrivateKey"
)

type (
	Config     map[string]any
	TunnelInfo struct {
		Config Config
	}
)

// ExtractTunnelInfoFromDestinationConfig extracts TunnelInfo from destination config if tunnel is enabled for the destination.
func ExtractTunnelInfoFromDestinationConfig(config Config) *TunnelInfo {
	_ = "STUB: not implemented"
	return nil
}

// Connect establishes a database connection over an SSH tunnel.
func Connect(dsn string, config Config) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractTunnelConfig(config Config) (*stunnel.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readString(key string, config Config) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
