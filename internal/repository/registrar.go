package repository

import (
	"errors"
	"sync"
)

type (
	StorageName       string
	StoragesRegistrar interface {
		Register(StorageName, any) error
		Unregister(StorageName, any) error
	}

	StorageRegistry struct {
		mx       sync.Mutex
		registry map[StorageName]any
	}
)

var (
	ErrStorageNameAlreadyRegistered            = errors.New("storageNameAlreadyRegistered")
	ErrStorageNameNotRegistered                = errors.New("storageNameNotRegistered")
	ErrStorageNameRegisteredWithAnotherStorage = errors.New("storageNameRegisteredWithAnotherStorage")
)

func NewStorageRegistry() *StorageRegistry {
	r := new(StorageRegistry)
	r.registry = make(map[StorageName]any)
	return r
}

func (r *StorageRegistry) Register(name StorageName, store any) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	_, ok := r.registry[name]
	if ok {
		return ErrStorageNameAlreadyRegistered
	}
	r.registry[name] = store
	return nil
}

func (r *StorageRegistry) Unregister(name StorageName, store any) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	s, ok := r.registry[name]
	if !ok {
		return ErrStorageNameNotRegistered
	}
	if s != store {
		return ErrStorageNameRegisteredWithAnotherStorage
	}
	delete(r.registry, name)
	return nil
}
