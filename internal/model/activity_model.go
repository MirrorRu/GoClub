package model

import "goclub/internal/model/basis"

type (
	ActivityID basis.ID

	ActivityName basis.Name

	ActivityStaffmanRole uint8

	ActivityBase struct {
		Name       ActivityName
		StaffRoles []ActivityStaffmanRole
	}

	ActivityExt struct{}

	ActivityRefs struct{}

	Activity basis.Table[ActivityID, ActivityBase, ActivityExt, ActivityExt]
)

const (
	ActivityStaffmanRoleUnknown ActivityStaffmanRole = iota
	ActivityStaffmanRoleMaster
	ActivityStaffmanRoleSuperviser
	ActivityStaffmanRoleSupplier
)
