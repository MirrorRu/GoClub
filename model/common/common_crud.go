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

type CRUDRepo[DataType any, KeyType any] interface {
	MemberCreate(ctx context.Context, data DataType) (result CRUDResult[DataType])
	MemberUpdate(ctx context.Context, data DataType) (result CRUDResult[DataType])
	MemberDelete(ctx context.Context, key KeyType) (result CRUDResult[struct{}])
	MemberRead(ctx context.Context, key KeyType) (result CRUDResult[DataType])
	MemberList(ctx context.Context, filter any) (result CRUDResult[[]DataType])
}

type CRUDService[DataType any, KeyType any] interface {
	Read(ctx context.Context, key KeyType) (result CRUDResult[DataType])
	Create(ctx context.Context, data DataType) (result CRUDResult[DataType])
	Update(ctx context.Context, data DataType) (result CRUDResult[DataType])
	Delete(ctx context.Context, key KeyType) (result CRUDResult[struct{}])
	Listing(ctx context.Context, filter any) (result CRUDResult[[]DataType])
}
