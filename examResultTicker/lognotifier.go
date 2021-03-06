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

package examResultTicker

import (
	"example.com/CoursesNotifier/models"
	"github.com/cdfmlr/qzgo"
	"log"
)

type LogNotifier string

func (l LogNotifier) Notify(student *models.Student, result qzgo.GetCjcxRespBodyItem) {
	log.Printf("(LogNotifier %s) Exam Result Notify:\n\t|--> student: %s\n\t|--> results: %#v", l, student.Sid, result)
}
