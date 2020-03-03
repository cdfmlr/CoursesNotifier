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

package jsonFileLoader

import (
	"encoding/json"
	"io/ioutil"
)

// 从 filename 读取 JSON 文件，放入 v
// e.g.
//		conf := Conf{}
//		err := jsonFileLoader.Load("./config.json", &conf)
func Load(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(data, v)
	return err
}
