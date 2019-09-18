# CID

Coded/obfuscated IDs for when you don't want to expose your database keys when marshaling to JSON. Supports int64 and uint64.

It works by "hashing" the original identifier using a modular multiplicative inverse operation then encoding with Crockford's base32 when marshaling to JSON. This is reversible, so when you get a valid string, it is unmarshaled to the original id.

## Install:
```
go get github.com/marcelomd/cid
```

## Thanks
This borrows heavily from [github.com/c2h5oh/hide](https://github.com/c2h5oh/hide) and [github.com/richardlehane/crock32](https://github.com/richardlehane/crock32). Thanks guys!

## Set your own primes
It is recommended to set your own primes, so you don't end up with the same obfuscated IDs as everyone else. Xor is used to further scramble the resulting ID.
```
err := cid.Int64SetPrime(myPrime, myXor)
```

## Example
Just use `cid.Int64` or `cid.Uint64` in place of the corresponding types you want obfuscated.
```
package main

import (
    "encoding/json"
    "fmt"
    "github.com/marcelomd/cid"
)

type User struct {
    Id   cid.Int64 `db:"id" json:"id"`
    Name string    `db:"name" json:"name"`
}

var (
    myPrime int64 = 0x1337133713373317
    myXor   int64 = 0x7b1ab1ab1ab1ab1a
)

func main() {
    fmt.Println("CID")
    err := cid.Int64SetPrime(myPrime, myXor)
    if err != nil {
        panic(err)
    }

    x := int64(123123123)
    y := Cid.EncodeInt64(x)
    z, err := Cid.DecodeInt64(y)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Direct obfuscation: %d -> %s -> %d\n", x, y, z)

    fmt.Println("JSON marshal/unmarshal:")
    for i := 0; i < 10; i++ {
        toJson := User{Id: cid.Int64(i), Name: "Bla"}
        j, err := json.Marshal(toJson)
        if err != nil {
            panic(err)
        }

        fromJson := User{}
        err = json.Unmarshal(j, &fromJson)
        if err != nil {
            panic(err)
        }

        fmt.Printf("User: %s -> %s -> %s\n", toJson, string(j), fromJson)
    }

```

```
CID
Direct obfuscation: 123123123 -> 2T1F6M9DW0NRF -> 123123123
JSON marshal/unmarshal:
User: {0 Bla} -> {"id":"7P6NHNCDB3ART","name":"Bla"} -> {0 Bla}
User: {1 Bla} -> {"id":"6GBD2KG4RD60D","name":"Bla"} -> {1 Bla}
User: {2 Bla} -> {"id":"5TX4QRMYDZK9M","name":"Bla"} -> {2 Bla}
User: {3 Bla} -> {"id":"45FW81RHH8CJZ","name":"Bla"} -> {3 Bla}
User: {4 Bla} -> {"id":"3FHQXEXB6TST6","name":"Bla"} -> {4 Bla}
User: {5 Bla} -> {"id":"1P2EHQ1XA4N39","name":"Bla"} -> {5 Bla}
User: {6 Bla} -> {"id":"0GM62W5MZN6CG","name":"Bla"} -> {6 Bla}
User: {7 Bla} -> {"id":"7V6SQ5AE37KNV","name":"Bla"} -> {7 Bla}
User: {8 Bla} -> {"id":"658H82E1GGCX2","name":"Bla"} -> {8 Bla}
User: {9 Bla} -> {"id":"5FX8X8JV42R6N","name":"Bla"} -> {9 Bla}
```

