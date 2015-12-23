package export

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/fredv/mixpanel-go/client"
)

type MixpanelUser struct {
	DistinctID string                 `json:"$distinct_id"`
	Properties MixpanelUserProperties `json:"$properties"`
}

func (user *MixpanelUser) Account() string {
	return user.Properties.AccountCode
}

type MixpanelUserProperties struct {
	Created     string `json:"$created"`
	AccountCode string `json:"Account Code"`
}

type EngageResponse struct {
	Results   []MixpanelUser `json:"results"`
	Status    string         `json:"status"`
	SessionID string         `json:"session_id"`
}

func DistinctIDMap(client *client.MixpanelClient) (*map[string]string, error) {
	page := 0
	lastResultSet := 1000
	users := []MixpanelUser{}
	sessionID := ""
	for lastResultSet == 1000 {
		rsp, err := Engage(client, page, sessionID)
		if err != nil {
			return nil, err
		}
		for _, user := range rsp.Results {
			users = append(users, user)
		}
		fmt.Printf("\n%d %d Users", len(rsp.Results), len(users))
		page = page + 1
		sessionID = rsp.SessionID
		lastResultSet = len(rsp.Results)
	}

	distinctIDMap := map[string]string{}
	for _, user := range users {
		distinctIDMap[user.DistinctID] = user.Account()
	}
	fmt.Printf("\n%q", distinctIDMap)
	return &distinctIDMap, nil
}

func Engage(client *client.MixpanelClient, page int, sessionID string) (*EngageResponse, error) {
	v := map[string]string{
		"page":       fmt.Sprintf("%d", page),
		"session_id": sessionID,
		"expire":     client.Timestamp(60),
	}

	values := client.Sign(v)

	url, err := url.Parse("https://mixpanel.com/api/2.0/engage/?" + values.Encode())
	fmt.Printf("\n%s\n", url.String())
	if err != nil {
		return nil, err
	}
	reader, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	response := &EngageResponse{}
	err = json.NewDecoder(reader).Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
