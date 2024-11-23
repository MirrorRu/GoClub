package tarifsvc

import (
	"goclub/model/common"
	"goclub/model/tarifs"
)

type TarifsService interface {
	common.CRUDService[*tarifs.Tarif]
}
