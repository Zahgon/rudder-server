package v2

import (
	"sync"
)

// NewOauthTokenCache returns a new cache for storing OAuth tokens.
func NewOauthTokenCache() OauthTokenCache { _ = "STUB: not implemented"; return *new(OauthTokenCache) }

// OauthTokenCache is an interface for a cache that stores OAuth tokens.
type OauthTokenCache interface {
	// Load retrieves the OAuth token associated with the given key.
	Load(key string) (OAuthToken, bool)
	// Store saves the OAuth token with the associated key.
	Store(key string, value OAuthToken)
	// Delete removes the OAuth token associated with the given key.
	Delete(key string)
}

type syncMapCache[T any] struct {
	m sync.Map
}

func (c *syncMapCache[T]) Load(key string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (c *syncMapCache[T]) Store(key string, value T) { _ = "STUB: not implemented"; return }

func (c *syncMapCache[T]) Delete(key string) { _ = "STUB: not implemented"; return }
