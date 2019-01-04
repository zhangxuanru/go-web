package models

import (
	"config"
	"fmt"
	"libary/db"
)

var sql string
var err error
const GROUP_TABLE  = config.DB_PREFIX+"group"
const GROUP_DETAIL_TABLE  = config.DB_PREFIX+"group_detail"

//获取group列表
func GetGroupList(where string,start int,limit int) (list map[int]map[string]string, err error) {
	if len(where) > 0{
   	   sql = fmt.Sprintf("SELECT g.group_id,title,img_date,g.group_pics_num, detail.equalw_url,detail.equalw_image_id,detail.equalh_url,detail.equalh_image_id,detail.url800,detail.url800_image_id from %s as g inner JOIN %s as detail  on g.group_id = detail.group_id where %s ORDER BY img_date DESC LIMIT %d,%d",GROUP_TABLE,GROUP_DETAIL_TABLE,where,start,limit)
	}else{
	   sql = fmt.Sprintf("SELECT g.group_id,title,img_date,g.group_pics_num,detail.equalw_url,detail.equalw_image_id,detail.equalh_url,detail.equalh_image_id,detail.url800,detail.url800_image_id from %s as g inner JOIN %s as detail  on g.group_id = detail.group_id   ORDER BY img_date DESC LIMIT %d,%d",GROUP_TABLE,GROUP_DETAIL_TABLE,start,limit)
	}
	list,err = db.GetList(sql)
	return  list,err
}


//获取group总数
func GetGroupCount(where string) (r map[string]string, err error) {
	if len(where) > 0{
		sql = fmt.Sprintf("SELECT COUNT(1) AS c  FROM %s WHERE %s",GROUP_TABLE,where)
	}else{
		sql = fmt.Sprintf("SELECT COUNT(1) AS c  FROM %s",GROUP_TABLE)
	}
	return db.GetRow(sql)
}
