package model_test

import (
	"fmt"
	"goclub/internal/model"
	"goclub/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	func NewActivityInmemStore(registrar StoragesRegistrar) *ActivityInmemStore {
		s := new(ActivityInmemStore)
		s.TableMemoryStore.Init(registrar)
		return s
	}

	func NewClubmanInmemStore(registrar StoragesRegistrar) *ClubmanInmemStore {
		s := new(ClubmanInmemStore)
		s.TableMemoryStore.Init(registrar)
		return s
	}

	func NewEventIInmemStore(registrar StoragesRegistrar) *EventInmemStore {
		s := new(EventInmemStore)
		s.TableMemoryStore.Init(registrar)
		return s
	}

	func NewStaffmanInmemStore(registrar StoragesRegistrar) *StaffmanInmemStore {
		s := new(StaffmanInmemStore)
		s.TableMemoryStore.Init(registrar)
		return s
	}
*/
func TestNewClubmanInmemStore(t *testing.T) {
	reger := repository.NewStorageRegistry()
	s := model.NewClubmanInmemStore(reger)
	assert.NotNil(t, s, "Проверка создания объекта хранилища")
	{
		ss, ok := interface{}(s).(repository.ITableStorage)
		assert.True(t, ok, "Проверка реализации интерфейса")
		assert.NotNil(t, ss, "Проверка сслыки интерфейса")
		//ss.Add(&model.Activity{})
	}

}

func TestClubmanInmemStoreCRUD(t *testing.T) {
	reger := repository.NewStorageRegistry()
	s := model.NewClubmanInmemStore(reger)
	cm := model.Clubman{}
	cm.Base.Name.Val = ""
	id, err := s.Add(&cm)
	t.Helper()
	//println("AddResult1", id, err)
	fmt.Println("AddResult2", id, err)
	assert.NoError(t, err)
	//t.Log("AddResult3", id, err)
	t.Error("AddResult4", id, err)

}
