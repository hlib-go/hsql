package hsql

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// RowsConvertSliceMap rows转换为切片map
func RowsConvertSliceMap(rows *sql.Rows) (list []map[string]interface{}, err error) {
	defer func() {
		if rows != nil {
			err = rows.Close()
			fmt.Println(err)
		}
	}()

	columns, err := rows.Columns()
	if err != nil {
		return
	}
	//临时存储每行数据,为每一列初始化一个指针
	row := make([]interface{}, len(columns))
	for index, _ := range row {
		var a interface{}
		row[index] = &a
	}
	//返回的切片
	for rows.Next() {
		err = rows.Scan(row...)
		if err != nil {
			return nil, err
		}
		item := make(map[string]interface{})
		for i, data := range row {
			item[columns[i]] = *data.(*interface{})
		}
		list = append(list, item)
	}
	return
}

// RowsConvertSliceStruct rows转换为结构体，注意result需要定义为切片，如 var rs []*strut ,  RowsConverStruct(rows,&rs)
func RowsConvertSliceStruct(rows *sql.Rows, result interface{}) (err error) {
	list, err := RowsConvertSliceMap(rows)
	if err != nil {
		return
	}
	bytes, err := json.Marshal(list)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return
	}
	return
}
