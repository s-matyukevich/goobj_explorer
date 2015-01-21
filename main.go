package main

import (
	"flag"
	"fmt"
	"github.com/golang/go/src/cmd/internal/goobj"
	"github.com/kr/pretty"
	"os"
)

func main() {
	var file = flag.String("o", "", "go object file")
	var pkgpath = flag.String("p", "main", "package path")

	flag.Parse()

	if *file == "" {
		fmt.Println("Usage of goobj_explorer:")
		flag.PrintDefaults()
		//fmt.Printf("Ussage: goobj_explorer -o <path to go object file> [-p <package path(default if main)>]\n")
		return
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Printf("%# v\n", err)
		return
	}
	obj, err := goobj.Parse(f, *pkgpath)
	pretty.Printf("%# v\n", obj)
	f.Close()
	if err != nil {
		fmt.Printf("error reading %s: %v\n", file, err)
		return
	}
}
