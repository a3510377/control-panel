package id

import (
	"database/sql/driver"
	"errors"
)

func (f *ID) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		return nil
	case string:
		if src == "" {
			return nil
		}
		*f = StringToID(src)
		return nil
	case int:
		*f = ID(src)
		return nil
	case int64:
		*f = ID(src)
		return nil
	}

	return errors.New("failed to scan ID")
}

func (f *ID) Value() (driver.Value, error) {
	return f.Int64(), nil
}
