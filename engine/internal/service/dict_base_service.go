package service

import (
	"context"
	"goclub/model/common"
)

type dictBaseService[T comparable] struct {
	repo common.CRUDRepo[*T]
}

func NewDictBaseService[T comparable](repo common.CRUDRepo[*T]) *dictBaseService[T] {
	return &dictBaseService[T]{
		repo: repo,
	}
}

func (svc *dictBaseService[T]) Create(ctx context.Context, member *T) (result common.CRUDResult[*T]) {
	return svc.repo.Create(ctx, member)
}

func (svc *dictBaseService[T]) Update(ctx context.Context, member *T) (result common.CRUDResult[*T]) {
	return svc.repo.Update(ctx, member)
}

func (svc *dictBaseService[T]) Delete(ctx context.Context, keys ...any) (result common.CRUDResult[struct{}]) {
	return svc.repo.Delete(ctx, keys...)
}

func (svc *dictBaseService[T]) Read(ctx context.Context, keys ...any) (result common.CRUDResult[*T]) {
	return svc.repo.Read(ctx, keys...)
}

func (svc *dictBaseService[T]) Listing(ctx context.Context, filter any) (result common.CRUDResult[[]*T]) {
	return svc.repo.List(ctx, filter)
}
