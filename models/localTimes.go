package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const localDataTimeFormat string = "2006-01-02 15:04:05"

type LocalTime struct {
	time.Time
}

// MarshalJSON 重写time.Time的MarshalJSON方法，格式化 2021-08-09T11:42:32+08:00 这个输出
func (t LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format(localDataTimeFormat))), nil
}

// Value insert timestamp into mysql need this function.
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value time.Time
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
