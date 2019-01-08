package logic

import (
	"application/models"
	"strconv"
	"time"
	"libary/util"
	"strings"
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

//group详情
func GetGoodsDetailById(groupId int,where string) (list map[string]string, err error) {
	list, err = models.GetGoodsDetailById(groupId,where)
	if len(list) == 0 || err !=nil{
		 return
	}
	//获取分类
	list["cateGory"] = ""
	if catId,ok := list["oneCategory"];ok{
		id, _ := strconv.Atoi(catId)
	 	cateRow, _ := GetCateGoryDataById(id)
	 	if len(cateRow) > 0{
			list["cateGory"] = cateRow["category_name"]
		}
	}
	if len(list["cateGory"]) > 0{
		replacer := strings.NewReplacer("新闻滚动", "","滚动","")
		list["cateGory"] = replacer.Replace(list["cateGory"])
	}
	list["caption"]  = util.SecurityString(list["caption"])
	list["img_date"] = util.FormattingTimeRubbing(list["img_date"])
	if err != nil{
		 return
	}
	 return
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








