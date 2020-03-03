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

package main

import (
	"example.com/CoursesNotifier/app"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	flag.Usage = usage
	// 读取命令行参数
	confFile := flag.String("c", "", "set configuration `file`")
	flag.Parse()

	if *confFile == "" {
		fmt.Fprintln(os.Stderr, "Cannot run without configuration file given.")
		flag.Usage()
		return
	}

	coursesNotifier := app.New(*confFile)
	coursesNotifier.Run()

	log.Println("CoursesNotifier Running...")

	http.HandleFunc("/", greet)
	http.ListenAndServe(":9001", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func usage() {
	fmt.Fprintf(os.Stderr, `
CoursesNotifier v0.1.0 for NCEPU(Baoding)
All rights reserved (c) 2020 CDFMLR

Usage: CoursesNotifier [-c filename]

Options:
`)
	flag.PrintDefaults()
}
