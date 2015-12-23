package main

import "github.com/fredv/mixpanel-go/export"

func main() {
	cmd := &export.ExportCommand{
		FromDate: "2015-12-12",
		ToDate:   "2015-12-13",
	}
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
