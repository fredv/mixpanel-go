package client

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
)

type MixpanelClient struct {
	ApiKey, ApiSecret string
}

func NewMixpanelClient() *MixpanelClient {
	apiKey := os.Getenv("MIXPANEL_API_KEY")
	apiSecret := os.Getenv("MIXPANEL_API_SECRET")
	return &MixpanelClient{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (client *MixpanelClient) Get(url *url.URL) (io.ReadCloser, error) {
	rsp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	return rsp.Body, nil
}

func (client *MixpanelClient) Timestamp(offset int32) string {
	return fmt.Sprintf("%d", int32(time.Now().Unix())+offset)
}

func addToStrToSign(str, key, value string) string {
	return fmt.Sprintf("%s%s=%s", str, key, value)
}

func (client *MixpanelClient) Sign(valueMap map[string]string) (v url.Values) {
	valueMap["api_key"] = client.ApiKey

	keys := []string{}
	for key, _ := range valueMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	strToSign := ""
	v = url.Values{}
	v.Set(keys[0], valueMap[keys[0]])
	strToSign = addToStrToSign(strToSign, keys[0], valueMap[keys[0]])
	for i, key := range keys {
		if i > 0 && valueMap[key] != "" {
			v.Add(key, valueMap[key])
			strToSign = addToStrToSign(strToSign, key, valueMap[key])
		}
	}
	strToSign = strToSign + client.ApiSecret
	fmt.Printf("%s", strToSign)
	signature := fmt.Sprintf("%x", md5.Sum([]byte(strToSign)))
	fmt.Printf("%s", signature)
	v.Add("sig", signature)
	return v
}
