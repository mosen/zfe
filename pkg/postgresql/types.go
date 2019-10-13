package postgresql

import (
	"database/sql/driver"
	"time"
)

type ZabbixTimestamp time.Time

func (ts ZabbixTimestamp) Value() (driver.Value, error) {
	return time.Time(ts).Unix(), nil
}

func (ts ZabbixTimestamp) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	intValue, ok := value.(int64)
	if !ok {
		return nil
	}
	ts = ZabbixTimestamp(time.Unix(intValue, 0))
	return nil
}
