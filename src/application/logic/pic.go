package logic

import (
	"fmt"
	"application/models"
	"strconv"
	"config"
	"libary/redis"
	"libary/util"
	"strings"
)

func GetImgUrl(imgId string) (imgUrl string) {
	if len(imgId) == 0{
		return ""
	}
	id, _ := strconv.Atoi(imgId)
	if id == 0{
		 return ""
	}
	key := fmt.Sprintf(config.REDIS_IMAGEURL_ID,id)
	val := redis.Get(key)
	if len(val) > 0{
		if strings.Contains(val,"http"){
             return val
		}
		 return config.STATICDOMAIN+"/"+val
	}
	where := fmt.Sprintf("id=%d",id)
	rs, err := models.GePicList(where, 0, 1)
	if err!=nil{
		 return  ""
	}
	if fileName,ok:= rs[0]["file_name"];ok && len(fileName) > 0{
		 imgUrl = config.STATICDOMAIN+"/"+fileName
		 redis.Set(key,fileName,"",0)
         return  imgUrl
	}
   if len(rs[0]["src_url"]) > 0{
	   redis.Set(key,rs[0]["src_url"],"",0)
	   return  rs[0]["src_url"]
   }
    return ""
 }



//获取pic 列表
func GetPicListByGroupId(groupId int,where string,start,limit int) (r map[int]map[string]string, err error) {
	r, err = models.GetPicListByGroupId(groupId,where, start,limit)
	if err != nil{
		return
	}
	for _,v := range r{
		imgId := v["img_id"]
        if imgId != "0" && len(imgId) > 0 {
			v["img_url"] = GetImgUrl(imgId)
		 }
		 v["title"] = util.SecurityString(v["title"])
		 v["img_date"] = util.FormattingTimeRubbing(v["img_date"])
	}
   return
}

func GetPicCountByGroupId(groupId int) (int) {
	r, err := models.GetPicCountByGroupId(groupId)
	if len(r) == 0 || err !=nil{
		return 0
	}
	i, e := strconv.Atoi(r["c"])
	if e!=nil{
		return 0
	}
	return i
}




