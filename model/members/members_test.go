package members

import (
	"goclub/model/common"
	"testing"
)

func TestSetValue(t *testing.T) {
	var m1 = map[string]common.DataFieldValue{
		fieldID.Name:   ID(15),
		fieldName.Name: "ssss",
	}
	var member Member
	member.SetValues(m1)
	t.Log(member.ID)
}
