package models

import (
	"fmt"
	"config"
	"libary/db"
)

const CATEGORY_TABLE  = config.DB_PREFIX+"category"

func GetCateGoryDataById(categoryId int) (r map[string]string, err error)  {
       sql = fmt.Sprintf("SELECT category_id,category_name FROM %s WHERE category_id=%d",CATEGORY_TABLE,categoryId)
       return db.GetRow(sql)
}

//获取子分类
func GetSubCateGoryData(categoryId,start,limit int) (result map[int]map[string]string,err error) {
	sql := fmt.Sprintf("SELECT category_id,category_name FROM %s WHERE category_pid=%d  ORDER BY id ASC LIMIT %d,%d",CATEGORY_TABLE,categoryId,start,limit)
	return db.GetList(sql)
}

