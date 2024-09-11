package model

import (
	"database/sql/driver"
	"time"
)

func (x MemberBirthday) Value() (driver.Value, error) {
	return time.Time(x), nil
}

type (
	MemberID       ID
	MemberFullName string
	MemberBirthday time.Time

	MemberBalance struct {
		Current   AccrualValue   `json:"current" gorm:"not null;default:0"`
		Withdrawn WithdrawnValue `json:"withdrawn" gorm:"not null;default:0"`
	}

	MemberInfo struct {
		ID       MemberID       `gorm:"primarykey"`
		FullName MemberFullName `gorm:"not null;size:100;check:full_name!=''"`
		Birthday MemberBirthday `gorm:"not null;type:timestamptz`
		Balance  MemberBalance  `gorm:"embedded"`
	}

	MemberCreateReq struct {
		MemberInfo
	}

	MemberCreateResp struct {
		CRUDResult
		MemberID
	}

	MemberUpdateReq struct {
		MemberInfo
	}

	MemberUpdateResp struct {
		CRUDResult
	}

	MemberDeleteReq struct {
		MemberID
	}

	MemberDeleteResp struct {
		CRUDResult
	}

	MemberReadReq struct {
		MemberID
	}

	MemberReadResp struct {
		MemberInfo
		CRUDResult
	}

	MemberListingReq struct {
	}

	MemberListingResp struct {
		CRUDResult
		Items []MemberInfo
	}
)
