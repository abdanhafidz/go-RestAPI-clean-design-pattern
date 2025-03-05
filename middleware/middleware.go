package middleware

import (
	"math"
	"time"

	"gorm.io/gorm"
)

func RecordCheck(rows *gorm.DB) (string, error) {
	count := rows.RowsAffected
	err := rows.Error
	// fmt.Println(rows)
	// fmt.Println(count)
	if count == 0 {
		return "no-record", err
	} else if err != nil {
		return "query-error", err
	} else {
		return "ok", err
	}
}

func DiffTime(t1 time.Time, t2 time.Time) (int, int, int) {
	hs := t1.Sub(t2).Hours()
	hs, mf := math.Modf(hs)
	ms := mf * 60
	ms, sf := math.Modf(ms)
	ss := sf * 60
	return int(hs), int(ms), int(ss)
}
