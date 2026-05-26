package etcdkeys

import (
	"github.com/rudderlabs/rudder-go-kit/config"
)

// MigrationRequestKeyPrefix returns the etcd key prefix for migration requests
func MigrationRequestKeyPrefix(config *config.Config) string { _ = "STUB: not implemented"; return "" }

// MigrationJobKeyPrefix returns the etcd key prefix for migration jobs
func MigrationJobKeyPrefix(config *config.Config) string { _ = "STUB: not implemented"; return "" }

// ReloadGatewayRequestKeyPrefix returns the etcd key prefix for gateway reload requests
func ReloadGatewayRequestKeyPrefix(config *config.Config) string {
	_ = "STUB: not implemented"
	return ""
}

// getEtcdNamespace retrieves the key namespace from the configuration
func getEtcdNamespace(config *config.Config) string { _ = "STUB: not implemented"; return "" }
