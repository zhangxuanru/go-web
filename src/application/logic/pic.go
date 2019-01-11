package logic

import (
	"fmt"
	"application/models"
	"strconv"
	"config"
	"libary/redis"
	"libary/util"
)

func GetImgUrl(imgId string) (imgUrl string) {
	id, _ := strconv.Atoi(imgId)
	if id == 0{
		 return ""
	}
	key := fmt.Sprintf(config.REDIS_IMAGEURL_ID,id)
	val := redis.Get(key)
	if len(val) > 0{
		 return config.STATICDOMAIN+"/"+val
	}
	where := fmt.Sprintf("id=%d",id)
	rs, err := models.GePicList(where, 0, 1)
	if err!=nil{
		 return  ""
	}
	if fileName,ok:= rs[0]["file_name"];ok{
		 imgUrl = config.STATICDOMAIN+"/"+fileName
		 redis.Set(key,fileName,"",0)
         return  imgUrl
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















