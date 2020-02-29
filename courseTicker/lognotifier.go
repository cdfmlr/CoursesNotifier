package courseTicker


import (
	"example.com/CoursesNotifier/models"
	"log"
)

type LogNotifier string

func (l LogNotifier) Notify(student *models.Student, course *models.Course) {
	log.Printf("(LogNotifier %s) Course Notify:\n\t|--> student: %s\n\t|--> course: %s (%s)", l, student.Sid, course.Cid, course.Name)
}

