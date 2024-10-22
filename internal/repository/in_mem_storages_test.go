package repository_test

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
// func TestNewActivityInmemStore(t *testing.T) {
// 	reger := repository.NewStorageRegistry()
// 	s := repository.NewActivityInmemStore(reger)
// 	assert.NotNil(t, s, "Проверка создания объекта хранилища")
// 	{
// 		ss, ok := interface{}(s).(repository.TableStorager[model.ActivityID, *model.Activity, repository.TableFilter[*model.Activity]])
// 		assert.True(t, ok, "Проверка реализации интерфейса")
// 		assert.NotNil(t, ss, "Проверка сслыки интерфейса")
// 	}
// 	{
// 		ss, ok := interface{}(s).(repository.TableStorager2[model.ActivityID, *model.Activity, repository.TableFilter[*model.Activity]])
// 		assert.True(t, ok, "Проверка реализации интерфейса 2")
// 		assert.NotNil(t, ss, "Проверка сслыки интерфейса 2")
// 	}
// 	{
// 		ss, ok := interface{}(s).(repository.ITableStorage)
// 		assert.True(t, ok, "Проверка реализации интерфейса 3")
// 		assert.NotNil(t, ss, "Проверка сслыки интерфейса 3")
// 		//ss.Add(&model.Activity{})
// 	}

// }
