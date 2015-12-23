package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
)

type ExportCommand struct {
	FromDate, ToDate, ApiKey, ApiSecret, Where, Event string
}

func main() {
	apiKey := os.Getenv("MIXPANEL_API_KEY")
	apiSecret := os.Getenv("MIXPANEL_API_SECRET")
	cmd := &ExportCommand{
		FromDate:  "2015-12-12",
		ToDate:    "2015-12-13",
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
	err := cmd.Export()
	if err != nil {
		panic(err)
	}
}

func addToStrToSign(str, key, value string) string {
	return fmt.Sprintf("%s%s=%s", str, key, value)
}

func (cmd *ExportCommand) Sign(valueMap map[string]string) (v url.Values) {
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
	strToSign = strToSign + cmd.ApiSecret
	fmt.Printf("%s", strToSign)
	signature := fmt.Sprintf("%x", md5.Sum([]byte(strToSign)))
	fmt.Printf("%s", signature)
	v.Add("sig", signature)
	return v
}

func timestamp(offset int32) string {
	return fmt.Sprintf("%d", int32(time.Now().Unix())+offset)
}

func (cmd *ExportCommand) Export() error {
	v := map[string]string{
		"api_key":   cmd.ApiKey,
		"from_date": cmd.FromDate,
		"to_date":   cmd.ToDate,
		"where":     cmd.Where,
		"event":     cmd.Event,
		"expire":    timestamp(60),
	}
	values := cmd.Sign(v)
	url, err := url.Parse("https://data.mixpanel.com/api/2.0/export/?" + values.Encode())
	if err != nil {
		return err
	}

	rsp, err := http.Get(url.String())
	defer rsp.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, string(body))
	return nil
}
