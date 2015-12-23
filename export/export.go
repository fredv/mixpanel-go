package export

import (
	"fmt"
	"net/url"
	"os"

	"github.com/fredv/mixpanel-go/client"
)

type ExportCommand struct {
	FromDate, ToDate, Where, Event string
}

func (cmd *ExportCommand) Run() error {
	mixpanelClient := client.NewMixpanelClient()
	return cmd.Export(mixpanelClient)
}

func (cmd *ExportCommand) Export(client *client.MixpanelClient) error {
	v := map[string]string{
		"from_date": cmd.FromDate,
		"to_date":   cmd.ToDate,
		"where":     cmd.Where,
		"event":     cmd.Event,
		"expire":    client.Timestamp(60),
	}

	values := client.Sign(v)

	url, err := url.Parse("https://data.mixpanel.com/api/2.0/export/?" + values.Encode())
	if err != nil {
		return err
	}
	body, err := client.Get(url)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, body)
	return nil
}
