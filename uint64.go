package cid

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var ErrScanUint64 = errors.New("Incompatible type")

type Uint64 uint64

func (id Uint64) String() string {
	return fmt.Sprintf("%d", uint64(id))
}

func (id Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(Uint64ToString(Uint64Hash(uint64(id))))
}

func (id *Uint64) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	n, err := StringToUint64(s)
	if err != nil {
		return err
	}
	*id = Uint64(Uint64Unhash(n))
	return nil
}

func (id Uint64) Value() (driver.Value, error) {
	i := uint64(id)
	if i <= 0 {
		return nil, nil
	}
	return i, nil
}

func (id *Uint64) Scan(value interface{}) error {
	if value == nil {
		*id = Uint64(0)
		return nil
	}
	switch v := value.(type) {
	case uint64:
		*id = Uint64(v)
	default:
		return ErrScanUint64
	}
	return nil
}

func EncodeUint64(n uint64) string {
	return Uint64ToString(Uint64Hash(n))
}

func DecodeUint64(s string) (uint64, error) {
	n, err := StringToUint64(s)
	return Uint64Unhash(n), err
}
