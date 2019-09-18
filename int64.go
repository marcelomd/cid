package cid

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var ErrScanInt64 = errors.New("Incompatible type")

type Int64 int64

func (id Int64) String() string {
	return fmt.Sprintf("%d", int64(id))
}

func (id Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(Int64ToString(Int64Hash(int64(id))))
}

func (id *Int64) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	n, err := StringToInt64(s)
	if err != nil {
		return err
	}
	*id = Int64(Int64Unhash(n))
	return nil
}

func (id Int64) Value() (driver.Value, error) {
	i := int64(id)
	if i <= 0 {
		return nil, nil
	}
	return i, nil
}

func (id *Int64) Scan(value interface{}) error {
	if value == nil {
		*id = Int64(0)
		return nil
	}
	switch v := value.(type) {
	case int64:
		*id = Int64(v)
	default:
		return ErrScanInt64
	}
	return nil
}

func EncodeInt64(n int64) string {
	return Int64ToString(Int64Hash(n))
}

func DecodeInt64(s string) (int64, error) {
	n, err := StringToInt64(s)
	return Int64Unhash(n), err
}
