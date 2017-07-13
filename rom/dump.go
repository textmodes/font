// +build ignore

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	mark := flag.Int("mark", 8, "mark")
	flag.Parse()

	for _, arg := range flag.Args() {
		b, err := ioutil.ReadFile(arg)
		if err != nil {
			panic(err)
		}

		for i, c := range b {
			fmt.Printf("%#04x:%04d:%08b\n", i, i+1, c)
			if *mark > 0 && i%(*mark) == (*mark)-1 {
				fmt.Println("")
			}
		}
	}
}
