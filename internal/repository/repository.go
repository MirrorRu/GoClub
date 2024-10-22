package repository

import (
	"fmt"
	"sync"
)

type (
	TableReadOption int

	NewValuer interface {
		NewValue() NewValuer
	}

	IRowID interface {
		comparable
		NewValuer
	}

	ITableStruct interface {
		SetID(any)
		GetID() any
	}

	ITableFilter interface {
		Filter(ITableStruct) (acceptable bool)
		SetCriterias(FilterCriterias)
		Criterias() FilterCriterias
	}

	ITableStorage interface {
		Add(ITableStruct) (any, error)
		Read(any) (ITableStruct, error)
		Update(ITableStruct) error
		Delete(any) error
		// List(ITableFilter, TableReadOption) ([]ITableStruct, error)
	}

	TableMemoryStore[IDType IRowID, TableStruct ITableStruct] struct {
		locker    sync.Mutex
		lastID    IDType
		storage   map[IDType]TableStruct
		registrar StoragesRegistrar
	}
)

const (
	TableReadWithAll TableReadOption = 1 << iota
	TableReadWithIDs
	TableReadWithBase
	TableReadWithExt
	TableReadWithRefs
)

func (s *TableMemoryStore[IDType, TableStructType]) Init(registrar StoragesRegistrar) {
	s.storage = make(map[IDType]TableStructType)
	s.registrar = registrar
}

func (s *TableMemoryStore[IDType, TableStructType]) Add(t ITableStruct) (any, error) {
	s.locker.Lock()
	defer s.locker.Unlock()
	x, typeOK := s.lastID.NewValue().(IDType)
	if !typeOK {
		panic("bad type assertion in Add")
	}
	s.lastID = x
	t.SetID(s.lastID)
	s.storage[s.lastID] = t.(TableStructType)
	return s.lastID, nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Read(key any) (ITableStruct, error) {
	id, typeOK := key.(IDType)
	if !typeOK {
		panic("bad type assertion")
	}

	s.locker.Lock()
	defer s.locker.Unlock()
	val, ok := s.storage[id]
	if !ok {
		return val, fmt.Errorf("key %v not found", id)
	}
	return val, nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Update(data ITableStruct) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	id, typeOK := data.GetID().(IDType)
	if !typeOK {
		panic("bab type assertion")
	}
	_, inMap := s.storage[id]
	if !inMap {
		return fmt.Errorf("key %v not found", id)
	}
	s.storage[id] = data.(TableStructType)

	return nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Delete(key any) error {
	id, typeOK := key.(IDType)
	if !typeOK {
		panic("bad type assertion")
	}

	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.storage, id)

	return nil
}

func (s *TableMemoryStore[IDType, TableStructType]) List(filter TableFilter[TableStructType], readOpt TableReadOption) ([]TableStructType, error) {
	result := make([]TableStructType, 0)
	for _, v := range s.storage {
		if filter.Filter(v) {
			result = append(result, v)
		}
	}
	return result, nil
}
