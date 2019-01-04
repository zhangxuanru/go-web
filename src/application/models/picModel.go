package models

import (
	"config"
	"fmt"
	"libary/db"
)

const PIC_TABLE  = config.DB_PREFIX+"pic_images"

func GePicList(where string,start int,limit int) (r map[int]map[string]string, err error) {
	if len(where) > 0{
   	   sql = fmt.Sprintf("SELECT id,file_name from %s where %s LIMIT %d,%d",PIC_TABLE,where,start,limit)
	}else{
	   sql = fmt.Sprintf("SELECT id,file_name from %s LIMIT %d,%d",PIC_TABLE,start,limit)
	}
	return db.GetList(sql)
}



