package test

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func init() {
	sql.Open("mysql", "")
}
