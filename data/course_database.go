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

package data

import (
	"database/sql"
	"example.com/CoursesNotifier/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type CourseDatabase struct {
	dataSourceName string
}

func NewCourseDatabase(dataSourceName string) *CourseDatabase {
	return &CourseDatabase{dataSourceName: dataSourceName}
}

/****************************************************************************************/
/*   该文件中所有CRUD方法、函数只在数据库执行报错（比如试图插入已存在的主键）时返回不为 nil 的 err     */
/**************************************************************************************/

// Insert 连接数据库，将给定的一条 Course 插入数据库。
// 给定的 Course 必须指定 cid, name, teacher, location, begin, end, week, when；
// 若给定课程 cid 已存在，数据库不会被更改，并返回一个错误（err!=nil）
// 返回 Rows Affected
func (sdb *CourseDatabase) Insert(course models.Course) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return insertCourse(db, course)
}

// GetCourses 返回库中所有 Course 记录
func (sdb *CourseDatabase) GetCourses() ([]models.Course, error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return []models.Course{}, err
	}
	defer db.Close()
	return getCourses(db)
}

// GetCourse 返回库中给定 cid 为标识的 Course 记录
// 若指定课程记录不存在将返回 (&models.Course{}, nil)
func (sdb *CourseDatabase) GetCourse(cid string) (*models.Course, error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return &models.Course{}, err
	}
	defer db.Close()
	return getCourse(db, cid)
}

// GetCoursesOnTime 返回指定返回指定星期几、几点开始的所有课程（不分周次）
// e.g.
// 		sbd.GetCoursesOnTime(2, 3, "08:00")
// 表示 星期 3 的 08:00 开始的课程
func (sdb *CourseDatabase) GetCoursesOnTime(day int, begin string) ([]models.Course, error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return []models.Course{}, err
	}
	defer db.Close()
	return getCoursesOnTime(db, day, begin)
}

// GetCoursesBeginTime 获取所有可能的上课开始时间
func (sdb *CourseDatabase) GetCoursesBeginTime() ([]string, error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return []string{}, err
	}
	defer db.Close()
	return getCoursesBeginTime(db)
}

// Update 用来在数据库中将 cid 标识的记录更新为传入的 course
// 若给定 cid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func (sdb *CourseDatabase) Update(cid string, course models.Course) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return updateCourse(db, cid, course)
}

// Delete 尝试删除数据库中给定 cid 对应的 course 记录
// 若给定 cid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func (sdb *CourseDatabase) Delete(cid string) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return deleteCourse(db, cid)
}

/****************************************************/
/* 👇以下为实际数据库操作，需给定 Open 了的 *DB 进行操作👇  */
/**************************************************/

// insertCourse 负责将给定的一条 Course 插入给定数据库连接。
// 给定的 Course 必须指定 cid, name, teacher, location, begin, end, week, when；
// 若给定课程 cid 已存在，数据库不会被更改，并返回一个错误（err!=nil）
// 返回 Rows Affected
func insertCourse(db *sql.DB, course models.Course) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("INSERT INTO course SET cid=?,name=?,teacher=?,location=?,begin=?,end=?,week=?,time=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(course.Cid, course.Name, course.Teacher, course.Location, course.Begin, course.End, course.Week, course.When)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return rowsAffected, nil
}

// getCourses 返回给定数据库连接中所有 Course 记录
func getCourses(db *sql.DB) ([]models.Course, error) {
	var courses []models.Course
	rows, err := db.Query("SELECT cid,name,teacher,location,begin,end,week,time FROM course")
	if err != nil {
		log.Println(err)
		return courses, err
	}
	for rows.Next() {
		var c models.Course
		err = rows.Scan(&c.Cid, &c.Name, &c.Teacher, &c.Location, &c.Begin, &c.End, &c.Week, &c.When)
		if err != nil {
			log.Println(err)
			return courses, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

// getCourse 返回给定数据库连接中给定 cid 为标识的 Course 记录
// 若指定课程记录不存在将返回 (&models.Course{}, nil)
func getCourse(db *sql.DB, cid string) (*models.Course, error) {
	var course models.Course
	rows, err := db.Query("SELECT cid,name,teacher,location,begin,end,week,time FROM course WHERE cid=?", cid)
	if err != nil {
		log.Println(err)
		return &course, err
	}
	for rows.Next() {
		var c models.Course
		err = rows.Scan(&c.Cid, &c.Name, &c.Teacher, &c.Location, &c.Begin, &c.End, &c.Week, &c.When)
		if err != nil {
			log.Println(err)
			return &course, err
		}
		course = c
		break
	}
	return &course, nil
}

// getCourseOnTime 返回指定星期几、几点的所有课程（不分周次）
// e.g.
// 		GetCoursesOnTime(db, 2, 3, "08:00")
// 表示 星期 3 的 08:00 开始的课程
func getCoursesOnTime(db *sql.DB, day int, begin string) ([]models.Course, error) {
	var courses []models.Course
	timeLike := fmt.Sprintf("%d%%", day)
	rows, err := db.Query("SELECT cid,name,teacher,location,begin,end,week,time FROM course WHERE begin=? AND time LIKE ?", begin, timeLike)
	if err != nil {
		log.Println(err)
		return courses, err
	}
	for rows.Next() {
		var c models.Course
		err = rows.Scan(&c.Cid, &c.Name, &c.Teacher, &c.Location, &c.Begin, &c.End, &c.Week, &c.When)
		if err != nil {
			log.Println(err)
			return courses, err
		}
		courses = append(courses, c)
	}

	return courses, nil
}

// getCoursesBeginTime 获取所有可能的上课开始时间
func getCoursesBeginTime(db *sql.DB) ([]string, error) {
	var coursesBeginTimes []string
	rows, err := db.Query("SELECT DISTINCT begin FROM course")

	for rows.Next() {
		var cBTime string
		err = rows.Scan(&cBTime)
		coursesBeginTimes = append(coursesBeginTimes, cBTime)
	}

	if err != nil {
		log.Println(err)
		return coursesBeginTimes, err
	}
	return coursesBeginTimes, nil
}

// updateCourse 用来在给定数据库连接中将 cid 标识的记录更新为传入的 course
// 若给定 cid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func updateCourse(db *sql.DB, cid string, course models.Course) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("UPDATE course SET name=?,teacher=?,location=?,begin=?,end=?,week=?,time=? WHERE cid=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(course.Name, course.Teacher, course.Location, course.Begin, course.End, course.Week, course.When, cid)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return rowsAffected, nil
}

// deleteCourse 尝试删除给定数据库连接中给定 cid 对应的 course 记录
// 若给定 cid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func deleteCourse(db *sql.DB, cid string) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("DELETE FROM course WHERE cid=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(cid)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return rowsAffected, nil
}
