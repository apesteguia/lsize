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
	flag.BoolVar(&raw, "raw", true, "Displays the raw data if user wants to parse it ")
	flag.Parse()

	fmt.Println("Reading files...")
	f.Init(path)
	f.GetSizes()
	if raw {
		f.List(decreasing)
	} else {
        f.ListRaw(decreasing)
	}
}