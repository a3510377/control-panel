package permission

import "database/sql/driver"

func (p *Permission) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		return nil
	case string:
		if src == "" {
			return nil
		}

		*p = StringToPermission(src)
	case int:
		*p = Permission(src)
	}

	return nil
}

func (f *Permission) Value() (driver.Value, error) {
	return f.Int64(), nil
}
