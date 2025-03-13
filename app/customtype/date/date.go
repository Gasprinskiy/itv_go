package customdate

import (
	"database/sql/driver"
	"strings"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	parsedTime, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(parsedTime)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d *Date) Scan(value interface{}) error {
	if t, ok := value.(time.Time); ok {
		*d = Date(t)
		return nil
	}
	return nil
}
