package models

import (
	"fmt"
	"config"
	"libary/db"
)

const LINKTAG_TABLE  = config.DB_PREFIX+"link_tags"

const RECOMMEND_TABLE  = config.DB_PREFIX+"recommend"

func GetLinkDataByCategoryId(categoryId int) (result map[int]map[string]string,err error) {
   sql = fmt.Sprintf("SELECT title,link,item_id,item_type FROM %s FORCE INDEX(category_id) WHERE category_id=%d AND item_id>0",LINKTAG_TABLE,categoryId)
   return  db.GetList(sql)
}

func GetRemmendData(categoryId int,where string) (result map[int]map[string]string,err error) {
	field:="title,link,src,sort,category_id,image_id,item_id,item_type,data_type"
	if len(where) > 0{
		sql = fmt.Sprintf("SELECT %s FROM %s WHERE category_id=%d AND %s ORDER BY id DESC",field,RECOMMEND_TABLE,categoryId,where)
	}else{
        sql = fmt.Sprintf("SELECT  %s FROM %s WHERE category_id=%d ORDER BY id DESC  ",field,RECOMMEND_TABLE,categoryId)
  }
   return  db.GetList(sql)
}







