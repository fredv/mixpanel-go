package main

import (
	"fmt"

	"github.com/fredv/mixpanel-go/client"
	"github.com/fredv/mixpanel-go/export"
)

func main() {

	client := client.NewMixpanelClient()
	myMap, err := export.DistinctIDMap(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", myMap)

	cmd := &export.ExportCommand{
		FromDate: "2015-12-12",
		ToDate:   "2015-12-13",
	}
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
