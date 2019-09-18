package cid

import (
	"errors"
)

// Crockford's base32
// Adapted from github.com/richardlehane/crock32
const maxlen64 = 13 //len("FZZZZZZZZZZZZ") -> 0xffffffffffffffff
var (
	digitsUpper      = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	digitsLower      = "0123456789abcdefghjkmnpqrstvwxyz"
	digits           = digitsUpper
	ErrInvalidString = errors.New("Invalid character in string")
	ErrTooBig        = errors.New("String too big")
)

func SetUpperCase() {
	digits = digitsUpper
}

func SetLowerCase() {
	digits = digitsLower
}

func toString(n uint64) string {
	var i int
	var a [maxlen64]byte
	for i = range a {
		a[i] = '0'
	}
	i = maxlen64
	for n >= 32 {
		i--
		a[i] = digits[n%32]
		n /= 32
	}
	i--
	a[i] = digits[n]
	return string(a[:])
}

func Int64ToString(n int64) string {
	return toString(uint64(n))
}

func Uint64ToString(n uint64) string {
	return toString(n)
}

func fromString(s string) (uint64, error) {
	var n uint64
	for i := 0; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case d == 'O', d == 'o':
			v = '0'
		case d == 'L', d == 'l', d == 'I', d == 'i':
			v = '1'
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'h':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'H':
			v = d - 'A' + 10
		case 'j' <= d && d <= 'k':
			v = d - 'a' + 9
		case 'J' <= d && d <= 'K':
			v = d - 'A' + 9
		case 'm' <= d && d <= 'n':
			v = d - 'a' + 8
		case 'M' <= d && d <= 'N':
			v = d - 'A' + 8
		case 'p' <= d && d <= 't':
			v = d - 'a' + 7
		case 'P' <= d && d <= 'T':
			v = d - 'A' + 7
		case 'v' <= d && d <= 'z':
			v = d - 'a' + 6
		case 'V' <= d && d <= 'Z':
			v = d - 'A' + 6
		default:
			return 0, ErrInvalidString
		}
		n = n*32 + uint64(v)
	}
	return n, nil
}

func StringToInt64(s string) (int64, error) {
	if len(s) > maxlen64 {
		return 0, ErrTooBig
	}
	if len(s) == maxlen64 && s[0] > '7' {
		return 0, ErrTooBig
	}
	n, err := fromString(s)
	return int64(n), err
}

func StringToUint64(s string) (uint64, error) {
	if len(s) > maxlen64 {
		return 0, ErrTooBig
	}
	if len(s) == maxlen64 && (s[0] > 'f' || s[0] > 'F') {
		return 0, ErrTooBig
	}
	return fromString(s)
}
