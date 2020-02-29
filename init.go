package CoursesNotifier

import "log"

func init(){
	log.SetPrefix("[CoursesNotifier] ")
	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Lshortfile|log.LUTC)
}
// TODO: Set log for each different subpackage