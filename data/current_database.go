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
	"fmt"
	"log"
	"strings"
	"time"
)

type CurrentDatabase struct {
	dataSourceName string
}

func NewCurrentDatabase(dataSourceName string) *CurrentDatabase {
	// 加入时区与解析时间的请求 which are required by scan termbegin into time.Time
	if strings.Contains(dataSourceName, "?") {
		dataSourceName = fmt.Sprintf("%s&loc=Asia%%2FShanghai&parseTime=true", dataSourceName)
	} else {
		dataSourceName = fmt.Sprintf("%s?charset=utf8&loc=Asia%%2FShanghai&parseTime=true", dataSourceName)
	}
	return &CurrentDatabase{dataSourceName: dataSourceName}
}

// GetCurrentTermBeginDate 返回当前学期开始时间
func (crtdb *CurrentDatabase) GetCurrentTermBeginDate() (time.Time, error) {
	db, err := sql.Open("mysql", crtdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return time.Time{}, err
	}
	defer db.Close()
	return getCurrentTermBeginDate(db)
}

// updateCurrentTermBeginDate 用来在给定数据库连接中 termbegin 的值更新
// 传入 currentTermBegin = "2020-02-17"
func (crtdb *CurrentDatabase) UpdateCurrentTermBeginDate(currentTermBegin string) (rowsAffected int64, err error) {
	db, err := sql.Open("mysql", crtdb.dataSourceName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer db.Close()
	return updateCurrentTermBeginDate(db, currentTermBegin)
}

/****************************************************/
/* 👇以下为实际数据库操作，需给定 Open 了的 *DB 进行操作👇  */
/**************************************************/

// getCurrentTermBeginDate 返回当前学期开始时间
func getCurrentTermBeginDate(db *sql.DB) (time.Time, error) {
	var tb time.Time

	rows, err := db.Query("SELECT termbegin FROM current")

	for rows.Next() {
		err = rows.Scan(&tb)
	}
	if err != nil {
		log.Println(err)
		return tb, err
	}
	return tb, nil
}

// updateCurrentTermBeginDate 用来在给定数据库连接中 termbegin 的值更新
// 传入 currentTermBegin = "2020-02-17"
func updateCurrentTermBeginDate(db *sql.DB, currentTermBegin string) (rowsAffected int64, err error) {
	stmt, err := db.Prepare("UPDATE current SET termbegin=?")
	res, err := stmt.Exec(currentTermBegin)
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return rowsAffected, nil
}
