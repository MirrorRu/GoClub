package model

import (
	"goclub/internal/model/basis"
	"goclub/internal/repository"
)

type (
	ClubmanID struct {
		basis.ID
	}

	ClubmanName struct {
		basis.Name
	}

	ClubmanBirthday struct {
		basis.Date
	}

	ClubmanJoiningDate struct {
		basis.Date
	}

	ClubmanPhone struct {
		basis.Phone
	}

	ClubmanBase struct {
		Name        ClubmanName
		Birthday    ClubmanBirthday
		JoiningDate ClubmanJoiningDate
		Phone       ClubmanPhone
	}

	ClubmanExt struct{}

	ClubmanRefs struct{}

	Clubman struct {
		basis.TableRecord[ClubmanID, ClubmanBase, ClubmanExt, ClubmanRefs]
	}
)

func (id ClubmanID) NewValue() repository.NewValuer {
	id.Val++
	return id
}

type ClubmanInmemStore struct {
	repository.TableMemoryStore[ClubmanID, *Clubman]
}

func NewClubmanInmemStore(registrar repository.StoragesRegistrar) *ClubmanInmemStore {
	s := new(ClubmanInmemStore)
	s.TableMemoryStore.Init(registrar)
	return s
}
