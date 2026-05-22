package pkg

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ISO8601Time time.Time

var iso8601Formater = "2006-01-02T15:04:05.999Z"

func (t *ISO8601Time) UnmarshalJSON(data []byte) error {
	var str string
	if str = string(data); str == "null" {
		return nil
	}
	var err error
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(iso8601Formater, timeStr)
	*t = ISO8601Time(t1)
	return err
}

func (t *ISO8601Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(*t).Format(iso8601Formater))
	return []byte(formatted), nil
}

func (t ISO8601Time) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format(iso8601Formater), nil
}

func (t *ISO8601Time) Scan(v any) error {
	switch vt := v.(type) {
	case time.Time:
		*t = ISO8601Time(vt)
	default:
		return errors.New("时间格式无法转换")
	}
	return nil
}

func (t ISO8601Time) String() string {
	t2 := time.Time(t).UTC()
	if t2.IsZero() {
		return "\"\""
	}
	return t2.Format(iso8601Formater)
}
