package service

import (
	"context"
	"goclub/model/common"
)

type dictBaseServiceExt[T comparable] struct {
	repo common.CRUDRepoExt[*T]
}

func NewDictBaseServiceExt[T comparable](repo common.CRUDRepoExt[*T]) *dictBaseServiceExt[T] {
	return &dictBaseServiceExt[T]{
		repo: repo,
	}
}

func (svc *dictBaseServiceExt[T]) Create(ctx context.Context, member *T) (result common.CRUDResult[*T]) {
	return svc.repo.Create(ctx, member)
}

func (svc *dictBaseServiceExt[T]) Update(ctx context.Context, member *T) (result common.CRUDResult[*T]) {
	return svc.repo.Update(ctx, member)
}

func (svc *dictBaseServiceExt[T]) Delete(ctx context.Context, keys *T) (result common.CRUDResult[struct{}]) {
	return svc.repo.Delete(ctx, keys)
}

func (svc *dictBaseServiceExt[T]) Read(ctx context.Context, keys *T) (result common.CRUDResult[*T]) {
	return svc.repo.Read(ctx, keys)
}

func (svc *dictBaseServiceExt[T]) Listing(ctx context.Context, filter common.ListFilter) (result common.CRUDResult[[]*T]) {
	return svc.repo.List(ctx, filter)
}
