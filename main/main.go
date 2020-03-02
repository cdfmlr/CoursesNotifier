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
