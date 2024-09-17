package repository

import (
	"fmt"
	"goclub/internal/model/basis"
	"strings"
	"sync"
)

type (
	StringFilter string

	TableReadOption uint32

	TableStructTyper[T comparable] interface {
		SetID(T)
		GetID() T
	}

	TableFilter[T any] interface {
		Filter(T) (acceptable bool)
	}

	TableStorager[IDType basis.ComparableIDer[any], TableStructType TableStructTyper[IDType], FilteringType TableFilter[TableStructType]] interface {
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

	TableSubstrFilter[T any] struct {
		subStr2find string
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

func (s *TableMemoryStore[IDType, TableStructType]) List(filter TableFilter[TableStructType], readOpt TableReadOption) ([]TableStructType, error) {
	result := make([]TableStructType, 0)
	for _, v := range s.storage {
		if filter.Filter(v) {
			result = append(result, v)
		}
	}
	return result, nil
}

func NewTableSubstrFilter[T any](subStr2find string) TableSubstrFilter[T] {
	return TableSubstrFilter[T]{subStr2find: subStr2find}
}

func (s TableSubstrFilter[T]) Filter(x T) (acceptable bool) {
	return strings.Contains(fmt.Sprintf("%v", x), s.subStr2find)
}
