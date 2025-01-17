package repack

import (
	"goclub/engine/internal/api"
	"goclub/model/common"
	"goclub/model/tarifs"
)

type tarif int

const Tarif tarif = 0

func (tarif) FromInfo(info *api.TarifInfo) *tarifs.Tarif {

	return &tarifs.Tarif{
		ID:        tarifs.ID(info.GetId()),
		Name:      tarifs.Name(info.GetName()),
		StartDate: tarifs.Date{Date: common.NewDate(info.GetStartDate().GetYyyyMmDd())},
		EndDate:   tarifs.Date{Date: common.NewDate(info.EndDate.GetYyyyMmDd())},
	}
}

func (tarif) ToInfo(tarif *tarifs.Tarif) *api.TarifInfo {
	if tarif == nil {
		return nil
	}
	return &api.TarifInfo{
		Id:        int64(tarif.ID),
		Name:      string(tarif.Name),
		StartDate: &api.Date{YyyyMmDd: int32(tarif.StartDate.Date)},
		EndDate:   &api.Date{YyyyMmDd: int32(tarif.EndDate.Date)},
	}
}

func (tarif) ToInfos(tarifs []*tarifs.Tarif) []*api.TarifInfo {
	if tarifs == nil {
		return nil
	}
	res := make([]*api.TarifInfo, len(tarifs))
	for idx := range tarifs {
		res[idx] = Tarif.ToInfo(tarifs[idx])
	}
	return res
}
