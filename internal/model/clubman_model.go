package model

import (
	"goclub/internal/model/basis"
)

type (
	ClubmanID basis.ID

	ClubmanName basis.Name

	ClubmanBirthday basis.Date

	ClubmanJoiningDate basis.Date

	ClubmanPhone basis.Phone

	ClubmanBase struct {
		Name        ClubmanName
		Birthday    ClubmanBirthday
		JoiningDate ClubmanJoiningDate
		Phone       ClubmanPhone
	}

	ClubmanExt struct{}

	ClubmanRefs struct{}

	Clubman basis.Table[ClubmanID, ClubmanBase, ClubmanExt, ClubmanRefs]
)
