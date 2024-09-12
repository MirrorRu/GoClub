package model

import (
	"goclub/internal/model/basis"
)

type (
	StaffmanID basis.ID

	StaffmanName basis.Name

	StaffmanBirthday basis.Date

	StaffmanJoiningDate basis.Date

	StaffmanPhone basis.Phone

	StaffmanBase struct {
		Name        StaffmanName
		Birthday    StaffmanBirthday
		JoiningDate StaffmanJoiningDate
		Phone       StaffmanPhone
	}

	StaffmanExt  struct{}
	StaffmanRefs struct{}

	Staffman basis.Table[StaffmanID, StaffmanBase, StaffmanExt, StaffmanRefs]
)
