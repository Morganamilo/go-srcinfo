package main

import (
	"fmt"
	"github.com/morganamilo/go-srcinfo"
)

func main() {
	info, err := srcinfo.ParseFile("SRCINFO")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info)
}
