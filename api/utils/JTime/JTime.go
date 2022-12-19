package JTime

import (
	"fmt"
	"time"
)

type Time time.Time

const TimeFormat = "2006-01-02 15:04:05"

func (t *Time) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Time().Format(TimeFormat))), nil
}

func (t Time) Time() time.Time { return time.Time(t) }
func (t Time) String() string  { return t.Time().Format(TimeFormat) }
