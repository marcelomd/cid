package main

import (
	"flag"
	"fmt"
	"github.com/marcelomd/cid"
	"strconv"
)

var (
	decode = flag.Bool("d", false, "Decode string")
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	for _, in := range args {
		if *decode {
			out, err := cid.DecodeInt64(in)
			if err != nil {
				fmt.Printf("%13s: %s\n", in, err.Error())
				continue
			}
			fmt.Printf("%13s: %20d\n", in, out)
		} else {
			i, err := strconv.ParseInt(in, 10, 64)
			if err != nil {
				fmt.Printf("%20s: %s\n", in, err.Error())
				continue
			}
			out := cid.EncodeInt64(i)
			fmt.Printf("%20d: %13s\n", i, out)
		}
	}

}
