package repository

import (
	"fmt"
	"goclub/internal/model/basis"
	"sync"
)

type (
	//StringFilter string

	TableReadOption uint32

	TableIdentier[T comparable] interface {
		SetID(T)
		GetID() T
	}

	TableStorager[IDType comparable, TableStructType TableIdentier[IDType], FilteringType TableFilter[TableStructType]] interface {
		Add(*TableStructType) (IDType, error)
		Read(IDType) (*TableStructType, error)
		Update(*TableStructType) error
		Delete(IDType) error
		List(FilteringType, TableReadOption) ([]*TableStructType, error)
	}

	IDTyper[T comparable] interface {
		comparable
		basis.NewValuer[T]
	}

	TableMemoryStore[IDType IDTyper[IDType], TableStructType TableIdentier[IDType]] struct {
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
	s.lastID = s.lastID.NewValue()
	t.SetID(s.lastID)
	s.storage[s.lastID] = t
	return s.lastID, nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Read(id IDType) (TableStructType, error) {
	s.locker.Lock()
	defer s.locker.Unlock()

	val, ok := s.storage[id]
	if !ok {
		return val, fmt.Errorf("key %v not found", id)
	}
	return val, nil
}

func (s *TableMemoryStore[IDType, TableStructType]) Update(data TableStructType) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	id := data.GetID()
	_, ok := s.storage[id]
	if !ok {
		return fmt.Errorf("key %v not found", id)
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

func (s *TableMemoryStore[IDType, TableStructType]) List(filter TableFilter[TableStructType], readOpt TableReadOption) ([]TableStructType, error) {
	result := make([]TableStructType, 0)
	for _, v := range s.storage {
		if filter.Filter(v) {
			result = append(result, v)
		}
	}
	return result, nil
}
