package courseTicker

import (
	"example.com/CoursesNotifier/models"
)

type Notifier interface {
	Notify(student *models.Student, course *models.Course)
}
