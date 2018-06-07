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

	for _, pkg := range info.SplitPackages() {
		fmt.Printf("%s-%s: %s\n", pkg.Pkgname, info.Version(), pkg.Pkgdesc)
	}
}
