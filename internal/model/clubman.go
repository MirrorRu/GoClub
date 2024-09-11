package model

import (
	base_date "goclub/internal/model/base_types/date"
	base_id "goclub/internal/model/base_types/id"
	base_name "goclub/internal/model/base_types/name"
	base_phone "goclub/internal/model/base_types/phone"
)

type (
	ClubmanID struct {
		base_id.Value
	}

	ClubmanName struct {
		base_name.Value
	}

	ClubmanBirthday struct {
		base_date.Value
	}

	ClubmanJoiningDate struct {
		base_date.Value
	}

	ClubmanPhone struct {
		base_phone.Value
	}

	Clubman struct {
		ID          ClubmanID
		Name        ClubmanName
		Birthday    ClubmanBirthday
		JoiningDate ClubmanJoiningDate
		Phone       ClubmanPhone
	}
)
