package logic

import (
	"application/models"
	"strconv"
	"time"
)

//获取group列表
func GetGroupList(where string,start int,limit int)  (r map[int]map[string]string,err error){
	list, e := models.GetGroupList(where, start, limit)
	if e!=nil{
		 return list, e
	}
	for _,val := range list{
		imgUrl := GetImgUrl(val["equalh_image_id"])
		val["equalhImgUrl"] = val["equalh_url"]
		if len(imgUrl) > 0{
			val["equalhImgUrl"] = imgUrl
		}
		i, _ := strconv.ParseInt(val["img_date"], 10, 64)
		val["img_date_format"] = time.Unix(i,0).Format("2006-01-02 03:04:05")
	}
	return list, e
}

//获取group总数
func GetGroupCount(where string) (total int) {
	r, err := models.GetGroupCount(where)
	if err !=nil{
		 return 0
	}
	i, e := strconv.Atoi(r["c"])
	if e!=nil{
		return 0
	}
	return i
}








