package jobsdb

import (
	"context"
)

func executeDbRequest[T any](ctx context.Context, jd *Handle, c *dbRequest[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// If priority pool is requested and configured, bypass the queue

type dbReqType int

const (
	undefinedReqType dbReqType = iota
	readReqType
	writeReqType
)

type dbRequest[T any] struct {
	reqType dbReqType
	name    string
	tags    *statTags
	command func() T
}

func newReadDbRequest[T any](name string, tags *statTags, command func() T) *dbRequest[T] {
	_ = "STUB: not implemented"
	return nil
}

func newWriteDbRequest[T any](name string, tags *statTags, command func() T) *dbRequest[T] {
	_ = "STUB: not implemented"
	return nil
}
