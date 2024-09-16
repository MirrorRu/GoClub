package model

import (
	"goclub/internal/model/basis"
)

type (
	EventStatus uint8

	EventID struct {
		basis.ID
	}

	EventName struct {
		basis.Name
	}

	EventStartMoment struct {
		basis.Moment
	}

/*
	EventStaffItem struct {
		Role    ActivityStaffmanRole
		Staffer Staffman
	}

	EventStaff []EventStaffItem

	EventClubbers []Clubman

	EventBase struct {
		Activity Activity
		Status   EventStatus
		Start    EventStartMoment
		Dration  time.Duration
	}

	EventExt struct{}

	EventRefs struct {
		Staff    EventStaff
		Clubbers EventClubbers
	}
*/

// Event basis.TableRecord[EventID, EventBase, EventExt, EventRefs]
)

const (
	EventStatusUnknown EventStatus = iota
	EventStatusPlanned
	EventStatusCompleted
	EventStatusCancelled
)
