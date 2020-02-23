package data

import (
	"database/sql"
	"example.com/CoursesNotifier/models"
	"log"
)

type StudentCourseRelationshipDatabase struct {
	dataSourceName string
}

func NewStudentCourseRelationshipDatabase(dataSourceName string) *StudentCourseRelationshipDatabase {
	return &StudentCourseRelationshipDatabase{dataSourceName: dataSourceName}
}

/****************************************************************************************/
/*   该文件中所有CRUD方法、函数只在数据库执行报错（比如试图插入已存在的主键）时返回不为 nil 的 err     */
/**************************************************************************************/

// Insert 连接数据库，将给定的一条 Relationship 插入数据库。
// 给定的 Relationship 必须指定 sid, cid；
// 插入数据不会检测是否重复！
// 返回 Rows Affected
func (rdb *StudentCourseRelationshipDatabase) Insert(relationship models.Relationship) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", rdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return insertRelationship(db, relationship)
}

// GetAllRelationships 返回库中所有 Relationship 记录
func (rdb *StudentCourseRelationshipDatabase) GetAllRelationships() ([]models.Relationship, error) {
	db, err := sql.Open("mysql", rdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return []models.Relationship{}, err
	}
	defer db.Close()
	return getAllRelationships(db)
}

// GetRelationshipsOfStudent 返回库中与给定 sid 相关的 Relationship 记录
// 若指定关系记录不存在将返回 (&models.Relationship{}, nil)
func (rdb *StudentCourseRelationshipDatabase) GetRelationshipsOfStudent(sid string) (*models.Relationship, error) {
	db, err := sql.Open("mysql", rdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return &models.Relationship{}, err
	}
	defer db.Close()
	return getRelationshipOfStudent(db, sid)
}

// Update 用来在数据库中将 sid 标识的记录更新为传入的 relationship
// 若给定 sid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func (rdb *StudentCourseRelationshipDatabase) Update(sid string, relationship models.Relationship) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", rdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return updateRelationship(db, sid, relationship)
}

// Delete 尝试删除数据库中给定 relationship 记录
// 若给定 relationship 不存在库中，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func (rdb *StudentCourseRelationshipDatabase) Delete(relationship models.Relationship) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", rdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return deleteRelationship(db, relationship)
}

/****************************************************/
/* 👇以下为实际数据库操作，需给定 Open 了的 *DB 进行操作👇  */
/**************************************************/

// insertRelationship 负责将给定的一条 Relationship 插入给定数据库连接。
// 给定的 Relationship 必须指定 sid, cid；
// 插入数据不会检测是否重复！
// 返回 Rows Affected
func insertRelationship(db *sql.DB, relationship models.Relationship) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("INSERT INTO coursetaking SET sid=?,cid=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(relationship.Sid, relationship.Cid)
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

// getAllRelationships 返回给定数据库连接中所有 Relationship 记录
func getAllRelationships(db *sql.DB) ([]models.Relationship, error) {
	var relationships []models.Relationship
	rows, err := db.Query("SELECT sid,cid FROM coursetaking")
	if err != nil {
		log.Println(err)
		return relationships, err
	}
	for rows.Next() {
		var r models.Relationship
		err = rows.Scan(&r.Sid, &r.Cid)
		if err != nil {
			log.Println(err)
			return relationships, err
		}
		relationships = append(relationships, r)
	}
	return relationships, nil
}

// getRelationshipOfStudent 返回给定数据库连接中与给定 sid 相关的 Relationship 记录
// 若指定关系记录不存在将返回 (&models.Relationship{}, nil)
func getRelationshipOfStudent(db *sql.DB, sid string) (*models.Relationship, error) {
	var relationship models.Relationship
	rows, err := db.Query("SELECT sid,cid FROM coursetaking WHERE sid=?", sid)
	if err != nil {
		log.Println(err)
		return &relationship, err
	}
	for rows.Next() {
		var r models.Relationship
		err = rows.Scan(&r.Sid, &r.Cid)
		if err != nil {
			log.Println(err)
			return &relationship, err
		}
		relationship = r
		break
	}
	return &relationship, nil
}

// updateRelationship 用来在给定数据库连接中将 sid 标识的记录更新为传入的 relationship
// 若给定 sid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func updateRelationship(db *sql.DB, sid string, relationship models.Relationship) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("UPDATE coursetaking SET sid=?,cid=? WHERE sid=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(relationship.Sid, relationship.Cid, sid)
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

// deleteRelationship 尝试删除给定数据库连接中给定 sid 对应的 relationship 记录
// 若给定 sid 不存在，会得到 rowsAffected=0 err=nil,没有数据库不会被更改，也不会有错误产生
// 返回 Rows Affected
func deleteRelationship(db *sql.DB, relationship models.Relationship) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("DELETE FROM coursetaking WHERE sid=? AND cid=?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(relationship.Sid, relationship.Cid)
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
