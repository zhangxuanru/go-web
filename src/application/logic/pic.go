package logic

import (
	"fmt"
	"application/models"
	"strconv"
	"config"
	"libary/redis"
)

func GetImgUrl(imgId string) (imgUrl string) {
	id, _ := strconv.Atoi(imgId)
	if id == 0{
		 return ""
	}
	key := fmt.Sprintf(config.REDIS_IMAGEURL_ID,id)
	val := redis.Get(key)
	if len(val) > 0{
		 return val
	}
	where := fmt.Sprintf("id=%d",id)
	rs, err := models.GePicList(where, 0, 1)
	if err!=nil{
		 return  ""
	}
	if fileName,ok:= rs[0]["file_name"];ok{
		 imgUrl = config.STATICDOMAIN+"/"+fileName
		 redis.Set(key,imgUrl,"",0)
         return  imgUrl
	}
	redis.Set(key,"","",0)
    return ""
}




