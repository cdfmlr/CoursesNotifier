package data

import "log"

func init(){
	log.SetPrefix("[data] ")
	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Lshortfile|log.LUTC)
}