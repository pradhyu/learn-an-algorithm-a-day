package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	var str, separator string
	for index, entry := range os.Args[1:] {
		str += separator + strconv.Itoa(index) + ">" + entry
		separator = " "
	}
	fmt.Println(str)

}
