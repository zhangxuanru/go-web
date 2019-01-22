package models

import (
	"config"
	"fmt"
	"libary/db"
)

const PIC_TABLE_IMAGES  = config.DB_PREFIX+"pic_images"
const PIC_TABLE  = config.DB_PREFIX+"pic"
const PIC_DETAIL_TABLE  = config.DB_PREFIX+"picdetail"

func GePicList(where string,start int,limit int) (r map[int]map[string]string, err error) {
	if len(where) > 0{
   	   sql = fmt.Sprintf("SELECT id,file_name,src_url from %s where %s LIMIT %d,%d",PIC_TABLE_IMAGES,where,start,limit)
	}else{
	   sql = fmt.Sprintf("SELECT id,file_name,src_url from %s LIMIT %d,%d",PIC_TABLE_IMAGES,start,limit)
	}
	return db.GetList(sql)
}


func GetPicListByGroupId(groupId int,where string,start,limit int,order string) (r map[int]map[string]string, err error) {
	if len(order) == 0{
		order = " ORDER BY id DESC "
	}else{
		tmpOrder := order
		order = "ORDER BY "+ tmpOrder
	}
	if len(where) > 0{
       sql= fmt.Sprintf("SELECT id,img_id,pic_id,group_id,res_id,img_url,title,img_date,prevPicId,nextPicId FROM %s WHERE group_id=%d AND %s %s LIMIT %d,%d",PIC_TABLE,groupId,where,order,start,limit)
	}else{
	    sql= fmt.Sprintf("SELECT id,img_id,pic_id,group_id,res_id,img_url,title,img_date,prevPicId,nextPicId FROM %s WHERE group_id=%d %s LIMIT %d,%d",PIC_TABLE,groupId,order,start,limit)
	}
   return db.GetList(sql)
}


func GetPicCountByGroupId(groupId int) (r map[string]string, err error) {
   sql = fmt.Sprintf("SELECT  COUNT(1) AS c  FROM %s WHERE group_id=%d ",PIC_TABLE,groupId)
   return db.GetRow(sql)
}


//获取PIC DETAIL表 详情
func GetPicDetailByWhere(where string) (r map[int]map[string]string, err error) {
	 fields:="pic_id,equalw_url,equalw_imgid,equalh_url,equalh_imgid,caption,file_type,size,store_size,specification,cameraman,brand,copyright,category"
     sql = fmt.Sprintf("SELECT %s FROM %s WHERE %s",fields,PIC_DETAIL_TABLE,where)
	 return db.GetList(sql)
}


