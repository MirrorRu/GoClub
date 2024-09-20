package basis

import (
	"time"
)

type (
	SingleBased[T any] struct {
		Val T
	}

	ID struct {
		SingleBased[int64]
	}
	Name struct {
		SingleBased[string]
	}
	Flag struct {
		SingleBased[bool]
	}
	Qty struct {
		SingleBased[float64]
	}
	Date struct {
		SingleBased[time.Time]
	}
	Moment struct {
		SingleBased[time.Time]
	}
	Amount struct {
		SingleBased[float64]
	}
	Phone struct {
		SingleBased[string]
	}

	TableRecord[IDFields any, BaseFields any, ExtFields any, RefFields any] struct {
		ID   IDFields
		Base BaseFields
		Ext  ExtFields
		Refs RefFields
	}

	NewValuer[T any] interface {
		NewValue() T
	}
)

func (sb SingleBased[T]) Base() T {
	return sb.Val
}

func (t *TableRecord[IDFields, BaseFields, ExtFields, RefFields]) SetID(id IDFields) {
	t.ID = id
}

func (t *TableRecord[IDFields, BaseFields, ExtFields, RefFields]) GetID() IDFields {
	return t.ID
}

/*

type (
	TableStructer[T comparable] interface {
		SetID(T)
		GetID() T
	}

	TableStorager[IDType comparable, TableStructPtrType TableStructer[IDType]] interface {
		Add(TableStructPtrType) (IDType, error)
		Read(IDType) (TableStructPtrType, error)
		Update(TableStructPtrType) error
		Delete(IDType) error
	}

	IDTypeConstraint[T comparable] interface {
		comparable
		NewValuer[T]
	}

	TableMemoryStore[IDType IDTypeConstraint[IDType], TableStructPtrType TableStructer[IDType]] struct {
		locker  sync.Mutex
		lastID  IDType
		storage map[IDType](TableStructPtrType)
	}
)

func (s *TableMemoryStore[IDType, TableStructPtrType]) Add(t TableStructPtrType) (IDType, error) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.lastID = s.lastID.NewValue()
	t.SetID(s.lastID)
	s.storage[s.lastID] = t
	return s.lastID, nil
}

func (s *TableMemoryStore[IDType, TableStructPtrType]) Read(id IDType) (TableStructPtrType, error) {
	s.locker.Lock()
	defer s.locker.Unlock()

	val, ok := s.storage[id]
	if !ok {
		return val, fmt.Errorf("key %v not found", id)
	}
	return val, nil
}

func (s *TableMemoryStore[IDType, TableStructPtrType]) Update(data TableStructPtrType) error {
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

func (s *TableMemoryStore[IDType, TableStructPtrType]) Delete(id IDType) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.storage, id)

	return nil
}
*/
