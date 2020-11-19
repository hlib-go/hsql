package hsql

import (
	"database/sql"
	"encoding/json"
	"errors"
)

var (
	ERR_ROW_NIL = errors.New("查询结果为空")
)

// QueryMap 查询单行map
func QueryMap(db *sql.DB, query string, args ...interface{}) (m map[string]interface{}, err error) {
	list, err := QueryList(db, query, args)
	if err != nil {
		return
	}
	if list == nil || len(list) == 0 {
		err = ERR_ROW_NIL
		return
	}
	m = list[0]
	return
}

// QueryList 查询分片map
func QueryList(db *sql.DB, query string, args ...interface{}) (list []map[string]interface{}, err error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	return RowsConvertSliceMap(rows)
}

// QueryStruct 查询结果转为结构体
func QueryStruct(rowResult interface{}, db *sql.DB, query string, args ...interface{}) (err error) {
	m, err := QueryMap(db, query, args)
	bytes, err := json.Marshal(m)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, rowResult)
	if err != nil {
		return
	}
	return
}

// QuerySlice 查询结构转为分片结构体
func QuerySlice(rowsResult interface{}, db *sql.DB, query string, args ...interface{}) (err error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}
	err = RowsConvertSliceStruct(rows, rowsResult)
	if err != nil {
		return
	}
	return
}
