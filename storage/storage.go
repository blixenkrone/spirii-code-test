package storage

import "context"

type TableProjectionReadWriter[T any] interface {
	Write(ctx context.Context, val T) error
	Read(ctx context.Context, id string) (T, error)
}
