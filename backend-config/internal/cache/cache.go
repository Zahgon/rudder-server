//go:generate mockgen -destination=./mock_cache.go -package=cache -source=./cache.go cache
package cache

import (
	"context"
	"crypto/cipher"
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/utils/pubsub"
)

var pkgLogger = logger.NewLogger().Child("backend-config-cache")

type Cache interface {
	Get(ctx context.Context) ([]byte, error)
}

type cacheStore struct {
	*sql.DB
	secret [32]byte
	key    string
}

// Start returns a new Cache instance, and starts a goroutine to cache the config
//
// secret is the secret key to encrypt the config with before storing it
//
// key is the key to use to store and fetch the config from the cache store
//
// ch is the channel to listen on for config updates and store them
func Start(ctx context.Context, secret [32]byte, key string, channelProvider func() pubsub.DataChannel) (Cache, error) {
	_ = "STUB: not implemented"
	return *new(Cache), nil
}

// setup db connection

// apply migrations

// clear config for other keys

// writeDebounce coalesces a burst of config updates into a single write

// persist writes the most recently received config to the database

// subscribe to config and debounce writes to db

// channel closed: flush the latest config before shutting down

// Encrypt and store the config to the database
func (db *cacheStore) set(ctx context.Context, config any) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip the write if the stored config is already identical. This avoids
// encrypting and transferring the (potentially large) config blob to the
// database when nothing has changed.

// encrypt

// write to config table

// Fetch the cached config when needed
func (db *cacheStore) Get(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	// read from database
	return nil, nil
}

// maybe fetch the config where workspaces = ''?

// decrypt and return

// setupDBConn sets up the database connection, creates the config table if it doesn't exist
func setupDBConn() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// clear config for all other keys
func (db *cacheStore) clear(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// apply config_cache migrations to the database
func migrate(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

func (db *cacheStore) encryptAES(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *cacheStore) decryptAES(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newGCM(secret [32]byte) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	// We need a 32-bytes key for AES-256
	// thus converting the secret to md5 (32-digit hexadecimal number)
	return *new(cipher.AEAD), nil
}
