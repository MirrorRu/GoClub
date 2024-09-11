package model

import (
	base_date "goclub/internal/model/base_types/date"
	base_id "goclub/internal/model/base_types/id"
	base_name "goclub/internal/model/base_types/name"
	base_phone "goclub/internal/model/base_types/phone"
)

type (
	StaffmanID struct {
		base_id.Value
	}

	StaffmanName struct {
		base_name.Value
	}

	StaffmanBirthday struct {
		base_date.Value
	}

	StaffmanJoiningDate struct {
		base_date.Value
	}

	StaffmanPhone struct {
		base_phone.Value
	}

	Staffman struct {
		ID          StaffmanID
		Name        StaffmanName
		Birthday    StaffmanBirthday
		JoiningDate StaffmanJoiningDate
		Phone       StaffmanPhone
	}
)
