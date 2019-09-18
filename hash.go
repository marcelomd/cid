package cid

import (
	"errors"
	"math/big"
)

// Knuth multiplicative hash
// Adapted from  github.com/c2h5oh/hide
var (
	int64prime         = new(big.Int).SetInt64(0x7b1ab1ab1ab1ab03)
	int64inverse       = new(big.Int).SetInt64(0x2f26820f2e7c97ab)
	int64xor           = new(big.Int).SetInt64(0x7badabadabadaba7)
	int64max           = new(big.Int).SetInt64(0x7fffffffffffffff)
	uint64prime        = new(big.Int).SetUint64(0x7b1ab1ab1ab1ab03)
	uint64inverse      = new(big.Int).SetUint64(0x2f26820f2e7c97ab)
	uint64xor          = new(big.Int).SetUint64(0x7badabadabadaba7)
	uint64max          = new(big.Int).SetUint64(0xffffffffffffffff)
	bigOne             = big.NewInt(1)
	ErrOutOfRange      = errors.New("Prime is greater than max value for the type")
	ErrNotAPrime       = errors.New("It is not a prime number")
	ErrCannotModInvert = errors.New("Cannot mod invert")
)

func getModInverse(prime, max *big.Int) (*big.Int, error) {
	if prime.Cmp(max) > 0 {
		return nil, ErrOutOfRange
	}
	if !prime.ProbablyPrime(100) {
		return nil, ErrNotAPrime
	}
	plusOne := new(big.Int)
	plusOne.Add(max, bigOne)
	i := new(big.Int)
	i = i.ModInverse(prime, plusOne)
	if i == nil {
		return nil, ErrCannotModInvert
	}
	return i, nil
}

func mmi(n, prime, max *big.Int) {
	n.Mul(n, prime)
	n.And(n, max)
}

func Int64SetPrime(prime, xor int64) error {
	p := new(big.Int).SetInt64(prime)
	i, err := getModInverse(p, int64max)
	if err != nil {
		return err
	}
	int64prime = p
	int64inverse = i
	int64xor = new(big.Int).SetInt64(xor)
	return nil
}

func Int64Hash(n int64) int64 {
	x := new(big.Int).SetInt64(n)
	mmi(x, int64prime, int64max)
	x.Xor(x, int64xor)
	return x.Int64()
}

func Int64Unhash(n int64) int64 {
	x := new(big.Int).SetInt64(n)
	x.Xor(x, int64xor)
	mmi(x, int64inverse, int64max)
	return x.Int64()
}

func Uint64SetPrime(prime, xor uint64) error {
	p := new(big.Int).SetUint64(prime)
	i, err := getModInverse(p, uint64max)
	if err != nil {
		return err
	}
	uint64prime = p
	uint64inverse = i
	uint64xor = new(big.Int).SetUint64(xor)
	return nil
}

func Uint64Hash(n uint64) uint64 {
	x := new(big.Int).SetUint64(n)
	mmi(x, uint64prime, uint64max)
	x.Xor(x, uint64xor)
	return x.Uint64()
}

func Uint64Unhash(n uint64) uint64 {
	x := new(big.Int).SetUint64(n)
	x.Xor(x, uint64xor)
	mmi(x, uint64inverse, uint64max)
	return x.Uint64()
}
