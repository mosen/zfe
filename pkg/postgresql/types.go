package postgresql

import (
	"database/sql/driver"
	"time"
)

// The ZabbixTimestamp type decodes/encodes to and from the SQL Unix Epoch timestamp to time.Time
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

// The ZabbixBoolean type takes care of decoding 0/1 integer values into booleans
type ZabbixBoolean bool

func (zb ZabbixBoolean) Value() (driver.Value, error) {
	if zb {
		return 1, nil
	} else {
		return 0, nil
	}
}

func (zb ZabbixBoolean) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	intValue, ok := value.(int64)
	if !ok {
		return nil
	}
	zb = intValue == 1
	return nil
}
