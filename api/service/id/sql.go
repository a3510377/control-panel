package id

import "database/sql/driver"

func (f *ID) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		return nil
	case string:
		if src == "" {
			return nil
		}

		*f = StringToID(src)
	case int:
		*f = ID(src)
	}

	return nil
}

func (f *ID) Value() (driver.Value, error) {
	return f.Int64(), nil
}
