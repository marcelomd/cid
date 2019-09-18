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
	fmt.Println("cid")
	err := cid.Int64SetPrime(myPrime, myXor)
	if err != nil {
		panic(err)
	}

	x := int64(123123123)
	y := cid.EncodeInt64(x)
	z, err := cid.DecodeInt64(y)
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
}
