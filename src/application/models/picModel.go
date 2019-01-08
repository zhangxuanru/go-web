package models

import (
	"config"
	"fmt"
	"libary/db"
)

const PIC_TABLE_IMAGES  = config.DB_PREFIX+"pic_images"
const PIC_TABLE  = config.DB_PREFIX+"pic"

func GePicList(where string,start int,limit int) (r map[int]map[string]string, err error) {
	if len(where) > 0{
   	   sql = fmt.Sprintf("SELECT id,file_name from %s where %s LIMIT %d,%d",PIC_TABLE_IMAGES,where,start,limit)
	}else{
	   sql = fmt.Sprintf("SELECT id,file_name from %s LIMIT %d,%d",PIC_TABLE_IMAGES,start,limit)
	}
	return db.GetList(sql)
}


func GetPicListByGroupId(groupId int,where string,start,limit int) (r map[int]map[string]string, err error) {
	if len(where) > 0{
      sql= fmt.Sprintf("SELECT img_id,pic_id,group_id,res_id,img_url,title,img_date,prevPicId,nextPicId FROM %s WHERE group_id=%d AND %s ORDER BY id DESC LIMIT %d,%d",PIC_TABLE,groupId,where,start,limit)
	}else{
	 sql= fmt.Sprintf("SELECT img_id,pic_id,group_id,res_id,img_url,title,img_date,prevPicId,nextPicId FROM %s WHERE group_id=%d ORDER BY id DESC LIMIT %d,%d",PIC_TABLE,groupId,start,limit)
	}
   return db.GetList(sql)
}


func GetPicCountByGroupId(groupId int) (r map[string]string, err error) {
   sql = fmt.Sprintf("SELECT  COUNT(1) AS c  FROM %s WHERE group_id=%d ",PIC_TABLE,groupId)
   return db.GetRow(sql)
}


