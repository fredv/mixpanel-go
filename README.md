# WIP and experimental!

## mixpanel-go
Mixpanel API Go-Client for exporting event data. This is currently not meant for tracking, but rather solving an attribution problem we have encountered. We track user activity, but require certain funnels to work on an account (has many users) level. For this purpose we want to export historical events, map them to accounts and import them into a new project with an account-identifier as distinct_id.


### Exporting Data
See [https://mixpanel.com/docs/api-documentation/exporting-raw-data-you-inserted-into-mixpanel](https://mixpanel.com/docs/api-documentation/exporting-raw-data-you-inserted-into-mixpanel) for details on parameters.

Currently this is used to simply export all event data in a JSONL stream to Stdout.

You can test the export using:
```sh
	MIXPANEL_API_KEY={YOUR_KEY} MIXPANEL_API_SECRET={YOUR_SECRET} go run main.go
```

`client/client.go` contains the setup of the Mixpanel client and request signing.

The Export Command uses the API v2 export for JSONL events of the specified date range with an expire set to 60 seconds per default.

```go
	cmd := &ExportCommand{
		FromDate:  "2015-12-12",
		ToDate:    "2015-12-13",
		Where: "",
		Event: "",
	}
	err := cmd.Run()
```

### ToDos
* Make ExportCommand interface reusable
* Add CLI interface for ExportCommand
* Work on Import Command
