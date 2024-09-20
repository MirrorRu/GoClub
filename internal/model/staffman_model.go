package model

import (
	"goclub/internal/model/basis"
)

type (
	StaffmanID struct {
		basis.ID
	}

	StaffmanName struct {
		basis.Name
	}

	StaffmanBirthday struct {
		basis.Date
	}

	StaffmanJoiningDate struct {
		basis.Date
	}

	StaffmanPhone struct {
		basis.Phone
	}

	StaffmanBase struct {
		Name        StaffmanName
		Birthday    StaffmanBirthday
		JoiningDate StaffmanJoiningDate
		Phone       StaffmanPhone
	}

	StaffmanExt  struct{}
	StaffmanRefs struct{}

	Staffman struct {
		basis.TableRecord[StaffmanID, StaffmanBase, StaffmanExt, StaffmanRefs]
	}
)

func (id StaffmanID) NewValue() StaffmanID {
	id.Val++
	return id
}
