package model

import (
	"goclub/internal/model/basis"
	"time"
)

type (
	EventID basis.ID

	EventName basis.Name

	EventStartMoment basis.Moment

	EventStaffItem struct {
		Role    ActivityStaffmanRole
		Staffer Staffman
	}

	EventStaff []EventStaffItem

	EventClubbers []Clubman

	EventStatus uint8

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

	Event basis.Table[EventID, EventBase, EventExt, EventRefs]
)

const (
	EventStatusUnknown EventStatus = iota
	EventStatusPlanned
	EventStatusCompleted
	EventStatusCancelled
)
