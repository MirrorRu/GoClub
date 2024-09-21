package repository_test

import (
	"goclub/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	testStore struct {
		dummy int
	}

	storeSet struct {
		name  repository.StorageName
		store *testStore
	}
)

func TestNewStorageRegistry(t *testing.T) {
	t.Parallel()
	r := repository.NewStorageRegistry()
	assert.NotNil(t, r)
}

func TestStorageRegistryRegister(t *testing.T) {
	r := repository.NewStorageRegistry()

	ts1 := storeSet{name: "ts1", store: &testStore{1}}
	ts2 := storeSet{name: "ts2", store: &testStore{2}}
	assert.NoError(t, r.Register(ts1.name, ts1.store))
	assert.ErrorIs(t, r.Register(ts1.name, ts1.store), repository.ErrStorageNameAlreadyRegistered)
	assert.NoError(t, r.Register(ts2.name, ts2.store))
	assert.ErrorIs(t, r.Register(ts2.name, &testStore{}), repository.ErrStorageNameAlreadyRegistered)
}

func TestStorageRegistryUnregister(t *testing.T) {
	r := repository.NewStorageRegistry()

	ts1 := storeSet{name: "ts1", store: &testStore{1}}
	ts2 := storeSet{name: "ts2", store: &testStore{2}}

	assert.ErrorIs(t, r.Unregister(ts1.name, ts1.store), repository.ErrStorageNameNotRegistered, "Хранилище еще не зарегистрировано")
	require.NoError(t, r.Register(ts1.name, ts1.store))
	require.NoError(t, r.Register(ts2.name, ts2.store))

	assert.ErrorIs(t, r.Unregister(ts1.name, ts2.store), repository.ErrStorageNameRegisteredWithAnotherStorage, "Для этого имени зарегистрировано иное хранилище")
	assert.NoError(t, r.Unregister(ts1.name, ts1.store), "Ожидалась устешная резрегистрация 1-го хранилища")
	assert.NoError(t, r.Unregister(ts2.name, ts2.store), "Ожидалась устешная резрегистрация 2-го хранилища")
}
