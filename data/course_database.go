package data

import (
	"database/sql"
	"example.com/CoursesNotifier/models"
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
