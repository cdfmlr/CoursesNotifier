package models

import "log"

func init(){
	log.SetPrefix("[models] ")
	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Lshortfile|log.LUTC)
}