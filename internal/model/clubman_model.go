package model

import (
	"goclub/internal/model/basis"
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

func (id ClubmanID) NewValue() ClubmanID {
	id.Val++
	return id
}
