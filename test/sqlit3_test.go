package test

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestA(t *testing.T) {
	db, err := sql.Open("sqlite3", "./hsql.sqlite3")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
