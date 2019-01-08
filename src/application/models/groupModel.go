package models

import (
	"config"
	"fmt"
	"libary/db"
	"strings"
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

//获取group详情
func GetGoodsDetailById(groupId int,where string) (list map[string]string, err error) {
	Fields = []string{"group_id","oneCategory","title","first_pic_id","group_pics_num","img_date"}
  if len(where) > 0{
      sql= fmt.Sprintf("SELECT %s  FROM %s WHERE group_id=%d AND %s",strings.Join(Fields,","), GROUP_TABLE,groupId,where)
	  }else{
	  sql= fmt.Sprintf("SELECT %s FROM %s WHERE group_id=%d ",strings.Join(Fields,","),GROUP_TABLE,groupId)
    }
	list,err = db.GetRow(sql)
	if len(list) == 0{
		return
	}
	lists, e := GetGroupDetailByIdData(groupId, "")
	if e !=nil{
		return
	}
	for k,v := range lists{
		list[k] = v
	}
	return
}


func GetGroupDetailByIdData(groupId int,where string)  (list map[string]string, err error) {
	Fields =[]string{"equalw_url","width","height","equalw_image_id","equalh_url","equalh_image_id","url800","url800_image_id","caption"}
    field := strings.Join(Fields,",")
	if len(where) > 0{
       sql= fmt.Sprintf("SELECT %s FROM %s WHERE group_id=%d AND  %s",field,GROUP_DETAIL_TABLE,groupId,where)
    }else {
       sql= fmt.Sprintf("SELECT %s FROM %s WHERE group_id=%d",field,GROUP_DETAIL_TABLE,groupId)
    }
	list,err = db.GetRow(sql)
	return
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
