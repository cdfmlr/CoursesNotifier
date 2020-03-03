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

package wxAccessToken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type accessTokenFetchJson struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func getAccessToken(appID, appSecret string) (*accessTokenFetchJson, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appID, appSecret)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return &accessTokenFetchJson{}, err
	}
	defer resp.Body.Close()
	accessed := &accessTokenFetchJson{}
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return &accessTokenFetchJson{}, err
		}
		err = json.Unmarshal(body, accessed)
		if err != nil {
			log.Println(err, "\n|->\tResp body:\n", body)
			return &accessTokenFetchJson{}, err
		}
	}
	return accessed, nil
}
