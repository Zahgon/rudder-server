package cache

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
)

type CacheType int8

const (
	MemoryCacheType CacheType = iota
	BadgerCacheType
)

type Cache[T any] interface {
	Update(key string, value T) error
	Read(key string) ([]T, error)
	Stop() error
}

func New[T any](ct CacheType, origin string, l logger.Logger) (Cache[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
