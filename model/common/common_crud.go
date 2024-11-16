package common

import "context"

type CRUDInfo struct {
	RecordsAffected int64
}

type CRUDResult[T any] struct {
	CRUDInfo
	Error error
	Data  T
}

type CRUDRepo[DataType any] interface {
	Create(ctx context.Context, data DataType) (result CRUDResult[DataType])
	Update(ctx context.Context, data DataType) (result CRUDResult[DataType])
	Delete(ctx context.Context, keys ...any) (result CRUDResult[struct{}])
	Read(ctx context.Context, keys ...any) (result CRUDResult[DataType])
	List(ctx context.Context, filter any) (result CRUDResult[[]DataType])
}

type CRUDService[DataType any] interface {
	Create(ctx context.Context, data DataType) (result CRUDResult[DataType])
	Update(ctx context.Context, data DataType) (result CRUDResult[DataType])
	Delete(ctx context.Context, keys ...any) (result CRUDResult[struct{}])
	Read(ctx context.Context, keys ...any) (result CRUDResult[DataType])
	Listing(ctx context.Context, filter any) (result CRUDResult[[]DataType])
}
