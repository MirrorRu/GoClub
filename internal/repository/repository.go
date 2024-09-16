package repository

import (
	"fmt"
	"goclub/internal/model/basis"
	"sync"
)

type (
	StringFilter string

	TableReadOption uint32

	TableStructTyper[T comparable] interface {
		SetID(T)
		GetID() T
	}

	TableStorager[IDType basis.ComparableIDer[any], TableStructType TableStructTyper[IDType], FilteringType any] interface {
		Add(TableStructType) (IDType, error)
		Read(IDType) (TableStructType, error)
		Update(TableStructType) error
		Delete(IDType) error
		List(FilteringType, TableReadOption) ([]TableStructType, error)
	}

	TableMemoryStore[IDType basis.ComparableIDer[IDType], TableStructType TableStructTyper[IDType]] struct {
		locker  sync.Mutex
		lastID  IDType
		storage map[IDType]TableStructType
	}
)

const (
	TableReadWithAll TableReadOption = 1 << iota
	TableReadWithIDs
	TableReadWithBase
	TableReadWithExt
	TableReadWithRefs
)

func (s *TableMemoryStore[IDType, TableStructType]) Add(t TableStructType) (IDType, error) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.lastID.Inc()
	t.SetID(s.lastID)
	s.storage[s.lastID] = t
	return s.lastID, nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Read(id IDType) (TableStructType, error) {
	s.locker.Lock()
	defer s.locker.Unlock()

	val, ok := s.storage[id]
	if !ok {
		return val, fmt.Errorf("Key %v not found", id)
	}
	return val, nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Update(data TableStructType) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	id := data.GetID()
	_, ok := s.storage[id]
	if !ok {
		return fmt.Errorf("Key %v not found", id)
	}
	s.storage[id] = data

	return nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Delete(id IDType) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.storage, id)

	return nil
}

/*
	List(FilteringType, TableReadOption) ([]TableStructType, error)
*/
