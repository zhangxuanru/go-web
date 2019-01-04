package models

import (
	"fmt"
	"config"
	"strings"
	"libary/db"
)

const BANNER_TABLE  = config.DB_PREFIX+"banner"
var Fields = []string{"id","type","text","origin_imgurl","link","img_id","link_id","img_width","recommend_type"}

func GetBannerList(where string,start int,limit int) (list map[int]map[string]string,err error) {
	  field := strings.Join(Fields,",")
	  if len(where) > 0{
           sql = fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY id DESC LIMIT %d,%d",field,BANNER_TABLE,where,start,limit)
	  }else{
           sql = fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d,%d",field,BANNER_TABLE,start,limit)
	  }
	  list,err = db.GetList(sql)
      return
}



