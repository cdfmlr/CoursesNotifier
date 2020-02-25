package qzapi

import "log"

func init(){
	log.SetPrefix("[CoursesNotifier] ")
	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Lshortfile|log.LUTC)
}