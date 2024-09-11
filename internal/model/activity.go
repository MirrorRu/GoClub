package model

import (
	base_id "goclub/internal/model/base_types/id"
	base_name "goclub/internal/model/base_types/name"
)

type (
	ActivityID struct {
		base_id.Value
	}

	ActivityName struct {
		base_name.Value
	}

	Activity struct {
		ID   ActivityID
		Name ActivityName
	}
)
