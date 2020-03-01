package data

import (
	"database/sql"
	"example.com/CoursesNotifier/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type StudentDatabase struct {
	dataSourceName string
}

func NewStudentDatabase(dataSourceName string) *StudentDatabase {
	return &StudentDatabase{dataSourceName: dataSourceName}
}

/****************************************************************************************/
/*   该文件中所有CRUD方法、函数只在数据库执行报错（比如试图插入已存在的主键）时返回不为 nil 的 err     */
/**************************************************************************************/

// Insert 连接数据库，将给定的一条 Student 插入数据库。
// 给定的 Student 必须指定 sid, pwd, wxuser, createtime；
// 若给定学生 sid 已存在，数据库不会被更改，并返回一个错误（err!=nil）
// 返回 Rows Affected
func (sdb *StudentDatabase) Insert(student models.Student) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return insertStudent(db, student)
}

// GetStudents 返回库中所有 Student 记录
func (sdb *StudentDatabase) GetStudents() ([]models.Student, error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return []models.Student{}, err
	}
	defer db.Close()
	return getStudents(db)
}

// GetStudent 返回库中给定 sid 为标识的 Student 记录
// 若指定学生记录不存在将返回 (&models.Student{}, nil)
func (sdb *StudentDatabase) GetStudent(sid string) (*models.Student, error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return &models.Student{}, err
	}
	defer db.Close()
	return getStudent(db, sid)
}

// GetByWxUser 返回给定数据库连接中给定 wxUser 为标识的 Student 记录
// 若指定学生记录不存在将返回 (&models.Student{}, nil)
func (sdb *StudentDatabase) GetByWxUser(wxUser string) (*models.Student, error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return &models.Student{}, err
	}
	defer db.Close()
	return getStudentByWxUser(db, wxUser)
}

// Update 用来在数据库中将 sid 标识的记录更新为传入的 student
// 若给定 sid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func (sdb *StudentDatabase) Update(sid string, student models.Student) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return updateStudent(db, sid, student)
}

// Delete 尝试删除数据库中给定 sid 对应的 student 记录
// 若给定 sid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func (sdb *StudentDatabase) Delete(sid string) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", sdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return deleteStudent(db, sid)
}

/****************************************************/
/* 👇以下为实际数据库操作，需给定 Open 了的 *DB 进行操作👇  */
/**************************************************/

// insertStudent 负责将给定的一条 Student 插入给定数据库连接。
// 给定的 Student 必须指定 sid, pwd, wxuser, createtime；
// 若给定学生 sid 已存在，数据库不会被更改，并返回一个错误（err!=nil）
// 返回 Rows Affected
func insertStudent(db *sql.DB, student models.Student) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("INSERT INTO student SET sid=?,pwd=?,wxuser=?,createtime=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(student.Sid, student.Pwd, student.WxUser, student.CreateTime)
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

// getStudents 返回给定数据库连接中所有 Student 记录
func getStudents(db *sql.DB) ([]models.Student, error) {
	var students []models.Student
	rows, err := db.Query("SELECT sid,pwd,wxuser,createtime FROM student")
	if err != nil {
		log.Println(err)
		return students, err
	}
	for rows.Next() {
		var s models.Student
		err = rows.Scan(&s.Sid, &s.Pwd, &s.WxUser, &s.CreateTime)
		if err != nil {
			log.Println(err)
			return students, err
		}
		students = append(students, s)
	}
	return students, nil
}

// getStudent 返回给定数据库连接中给定 sid 为标识的 Student 记录
// 若指定学生记录不存在将返回 (&models.Student{}, nil)
func getStudent(db *sql.DB, sid string) (*models.Student, error) {
	var student models.Student
	rows, err := db.Query("SELECT sid,pwd,wxuser,createtime FROM student WHERE sid=?", sid)
	if err != nil {
		log.Println(err)
		return &student, err
	}
	for rows.Next() {
		var s models.Student
		err = rows.Scan(&s.Sid, &s.Pwd, &s.WxUser, &s.CreateTime)
		if err != nil {
			log.Println(err)
			return &student, err
		}
		student = s
		break
	}
	return &student, nil
}

// getStudentByWxUser 返回给定数据库连接中给定 wxUser 为标识的 Student 记录
// 若指定学生记录不存在将返回 (&models.Student{}, nil)
func getStudentByWxUser(db *sql.DB, wxUser string) (*models.Student, error) {
	var student models.Student
	rows, err := db.Query("SELECT sid,pwd,wxuser,createtime FROM student WHERE wxuser=?", wxUser)
	if err != nil {
		log.Println(err)
		return &student, err
	}
	for rows.Next() {
		var s models.Student
		err = rows.Scan(&s.Sid, &s.Pwd, &s.WxUser, &s.CreateTime)
		if err != nil {
			log.Println(err)
			return &student, err
		}
		student = s
		break
	}
	return &student, nil
}

// updateStudent 用来在给定数据库连接中将 sid 标识的记录更新为传入的 student
// 若给定 sid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func updateStudent(db *sql.DB, sid string, student models.Student) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("UPDATE student SET pwd=?,wxuser=?,createtime=? WHERE sid=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(student.Pwd, student.WxUser, student.CreateTime, sid)
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

// deleteStudent 尝试删除给定数据库连接中给定 sid 对应的 student 记录
// 若给定 sid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func deleteStudent(db *sql.DB, sid string) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("DELETE FROM student WHERE sid=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(sid)
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
