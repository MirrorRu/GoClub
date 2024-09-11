package model

import (
	base_flag "goclub/internal/model/base_types/flag"
	base_id "goclub/internal/model/base_types/id"
	base_moment "goclub/internal/model/base_types/moment"
	base_name "goclub/internal/model/base_types/name"
)

type (
	EventID struct {
		base_id.Value
	}

	EventName struct {
		base_name.Value
	}

	EventStartMoment struct {
		base_moment.Value
	}

	EventIsMasterStaffer struct {
		base_flag.Value
	}

	EventStaffer struct {
		Staffer  Staffman
		IsMaster EventIsMasterStaffer
	}

	EventStaff []EventStaffer

	EventClubbers []Clubman

	Event struct {
		ID       EventID
		Name     EventName
		Start    EventStartMoment
		Staff    EventStaff
		Clubbers EventClubbers
	}
)
