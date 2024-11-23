package tarifs

import (
	"goclub/model/common"
)

type (
	ID   common.ID
	Name common.Name
	Date struct {
		common.Date
	}

	Tarif struct {
		ID        ID
		Name      Name
		StartDate Date
		EndDate   Date
	}
)

// func (d Date) ToDate() time.Time {
// 	//return d, nil
// 	year := d / 10000
// 	d = d % 10000
// 	month := d / 100
// 	day := d % 100
// 	val := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
// 	return val
// }

// func (d *Date) FromDate(t time.Time) {
// 	//return d, nil
// 	*d = Date(t.Year()*10000 + int(t.Month())*100 + t.Day())
// }

// // Value implements the [driver.Valuer] interface.
// func (d Date) Value() (driver.Value, error) {
// 	return d.ToDate(), nil
// }

// // Scan implements the [Scanner] interface.
// func (d *Date) Scan(value any) error {
// 	switch x := value.(type) {
// 	case int:
// 	case int32:
// 	case int64:
// 	case uint:
// 	case uint32:
// 	case uint64:
// 		*d = Date(x)
// 	case time.Time:
// 		d.FromDate(x)
// 	default:
// 		fmt.Println(reflect.TypeOf(value))
// 	}
// 	return nil
// }
