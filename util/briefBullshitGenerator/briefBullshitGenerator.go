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

package briefBullshitGenerator

import (
	"example.com/CoursesNotifier/util/jsonFileLoader"
	"log"
	"math/rand"
	"strings"
	"time"
)

// briefBullshitGenerator 一句废话生成器
// 生产一句废话，类似于如下格式：
// 		名人说过名言，这启示我们，关于问题怎么样，所以，建议行动。
// 用伪代码来表示就是：
//		// i 是随机数
// 		ff = Famous[i].Replace("a", Said[i]).Replace("b", Inspire[i])
// 		bb = Bosh[i].Replace("x", Questions[i])
// 		ss = So[i]
// 		dd = Do[i]
// 		briefBullshit = ff + bb + ss + dd

type briefBullshitGenerator struct {
	Famous []string `json:"famous"`
	Said []string `json:"said`
	Inspire []string `json:"inspire"`
	Bosh []string `json:"bosh"`
	Questions []string `json"questions"`
	So []string `json:"so"`
	Do []string `json:"do"`
}

func new(bullshitDataFile string) *briefBullshitGenerator {
	b := &briefBullshitGenerator{}
	if err := jsonFileLoader.Load(bullshitDataFile, b) ; err != nil {
		log.Println(err)
	}
	return b
}

func (b *briefBullshitGenerator) generate() string {
	rand.Seed(time.Now().Unix())
	ff := b.Famous[rand.Intn(len(b.Famous))]
	ff = strings.ReplaceAll(ff, "a", b.Said[rand.Intn(len(b.Said))])
	ff = strings.ReplaceAll(ff, "b", b.Inspire[rand.Intn(len(b.Inspire))])

	bb := b.Bosh[rand.Intn(len(b.Bosh))]
	bb = strings.ReplaceAll(bb, "x", b.Questions[rand.Intn(len(b.Questions))])

	ss := b.So[rand.Intn(len(b.So))]

	dd := b.Do[rand.Intn(len(b.Do))]

	return ff + bb + ss + dd
}

// Generate 生成一句废话
// 每次都会新建 briefBullshitGenerator 对象，可实现语料文件运行时更新。
func Generate(bullshitDataFile string) string {
	bg := new(bullshitDataFile)
	return bg.generate()
}