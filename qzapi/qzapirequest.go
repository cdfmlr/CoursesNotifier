package qzapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func qzApiGet(school string, token string, res interface{}, a map[string]string) error {
	// Make URL
	rawUrl := fmt.Sprintf("http://jwxt.%s.edu.cn/app.do", school)

	Url, err := url.Parse(rawUrl)
	if err != nil {
		log.Println(err)
		return err
	}

	params := url.Values{}
	for k, v := range a {
		params.Set(k, v)
	}

	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	// Make Request and Header
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	if token != "" {
		req.Header.Add("token", token)
	}

	// GET and Parse Response
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(body))
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}