package main

import (
	"flag"
	"fmt"
	s "lsize/size"
)

var (
	path       string
	decreasing bool
	raw        bool
)

func main() {
	var f s.Files
	flag.StringVar(&path, "path", "./", "Path to list")
	flag.BoolVar(&decreasing, "order", true, "Default: true, order of the listing")
	flag.BoolVar(&raw, "raw", false, "Displays the raw data if user wants to parse it ")
	flag.Parse()

	err := f.Init(path)
	if err != nil {
		fmt.Println("Unexpected error or a file in path")
		fmt.Println(err)
		return
	}
	f.GetSizes()
	if !raw {
		f.List(decreasing)
	} else {
		f.ListRaw(decreasing)
	}
}
