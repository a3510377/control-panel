package JTime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

func (t *Time) Scan(src any) error {
	if value, ok := src.(time.Time); ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", src)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := t.Time()
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}
