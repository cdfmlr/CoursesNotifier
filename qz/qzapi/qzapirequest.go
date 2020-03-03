/*
 * Copyright 2020 CDFMLR
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

	// fmt.Println(urlPath)

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

	err = json.Unmarshal(body, res)
	if err != nil {
		log.Println(err, "\nresp.body:\n" , string(body))
		return err
	}
	return nil
}