package tarifsvc

import (
	"goclub/model/common"
	"goclub/model/tarifs"
)

type TarifsService interface {
	common.CRUDServiceExt[*tarifs.Tarif]
}
