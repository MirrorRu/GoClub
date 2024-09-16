package model

import "goclub/internal/model/basis"

type (
	ActivityID struct {
		basis.ID
	}

	ActivityName struct {
		basis.Name
	}

	ActivityStaffmanRole uint8

	ActivityBase struct {
		Name       ActivityName
		StaffRoles []ActivityStaffmanRole
	}

	ActivityExt struct{}

	ActivityRefs struct{}

	Activity basis.TableRecord[*ActivityID, ActivityBase, ActivityExt, ActivityExt]
)

const (
	ActivityStaffmanRoleUnknown ActivityStaffmanRole = iota
	ActivityStaffmanRoleMaster
	ActivityStaffmanRoleSuperviser
	ActivityStaffmanRoleSupplier
)
