package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"libary/logger"
     "config"
)


var err error
var DB  *sql.DB

func init(){
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",config.DB_USER,config.DB_PASSWORD,config.DB_HOST,config.DB_PORT, config.DB_DATABASE_NAME,config.DB_CHARSET)
	DB,err = sql.Open(config.DB_DRIVER_NAME, dataSourceName)
	if err != nil{
          logger.ErrorLog.Fatal("mysql connect error",err)
	}
}

func Insert(sql string)(id int64, err error) {
	logger.InfoLog.Println(sql)
	result, e := DB.Exec(sql)
	if e != nil{
		logger.ErrorLog.Println(sql,"insert data error:",e)
		logger.InfoLog.Println("isertid:0","error:",err)
        return 0,e
	}
	id, err = result.LastInsertId()
	logger.InfoLog.Println("isertid:",id,"error:",err)
	return
}

//返回单行
func GetRow(sql string) (r map[string]string, err error)  {
	logger.InfoLog.Println(sql)
	rows, e := DB.Query(sql)
	record := make(map[string]string)
	if e != nil{
		logger.InfoLog.Println("querydata:",e)
		return record, e
	}
	col, _ := rows.Columns()
	scanArgs := make([]interface{},len(col))
	valArgs  := make([]interface{},len(col))
	for k:= range valArgs{
		scanArgs[k] = &valArgs[k]
	}
	for rows.Next(){
		rows.Scan(scanArgs...)
		for i,cc := range valArgs{
			if cc != nil{
				record[col[i]] = string(cc.([]byte))
			}
		}
	}
	logger.InfoLog.Println(record)
    return record,nil
}

func GetList(sql string) (result map[int]map[string]string,err error) {
	logger.InfoLog.Println(sql)
	rows, e := DB.Query(sql)
	result = make(map[int]map[string]string)
	if e != nil{
		logger.InfoLog.Println("querydata:",e)
		 return result, e
	}
 	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	k:=0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		result[k] = record
		k++
	}
    return result,nil
}

func UpdateSql(sql string) (count int64, err error) {
	logger.InfoLog.Println(sql)
	result, e := DB.Exec(sql)
	if e != nil{
		logger.ErrorLog.Println(sql,"update data error:",e)
		logger.InfoLog.Println("update count:0","error:",err)
		return 0,e
	}
	count, err = result.RowsAffected()
	logger.InfoLog.Println("num:",count,"error:",err)
	return
}



