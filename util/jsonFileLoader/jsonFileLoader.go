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
