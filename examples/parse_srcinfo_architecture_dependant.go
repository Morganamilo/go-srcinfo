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

	for _, source := range info.Source {
		if source.Arch == "" {
			fmt.Printf("This source is for %s: %s\n", "any", source.Value)
		} else {
			fmt.Printf("This source is for %s: %s\n", source.Arch, source.Value)
		}
	}
}
