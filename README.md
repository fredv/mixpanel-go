# WIP and experimental!

## mixpanel-go
Mixpanel API Go-Client for exporting event data. This is currently not meant for tracking, but rather solving an attribution problem we have encountered. We track user activity, but require certain funnels to work on an account (has many users) level. For this purpose we want to export historical events, map them to accounts and import them into a new project with an account-identifier as distinct_id.


### Exporting Data
See [https://mixpanel.com/docs/api-documentation/exporting-raw-data-you-inserted-into-mixpanel](https://mixpanel.com/docs/api-documentation/exporting-raw-data-you-inserted-into-mixpanel) for details on parameters.

Currently this is used to simply export all event data in a JSONL stream to Stdout.

```go
	apiKey := os.Getenv("MIXPANEL_API_KEY")
	apiSecret := os.Getenv("MIXPANEL_API_SECRET")
	cmd := &ExportCommand{
		FromDate:  "2015-12-12",
		ToDate:    "2015-12-13",
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		Where: "",
		Event: "",
	}
	err := cmd.Export()
```

### ToDos
* Extract Mixpanel client that can sign requests
* Make ExportCommand interface reusable
* Add CLI interface for ExportCommand
* Work on Import Command
