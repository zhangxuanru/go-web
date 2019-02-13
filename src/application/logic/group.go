package logic

import (
	"application/models"
	"strconv"
	"time"
	"libary/util"
	"strings"
	"application/Service"
	"fmt"
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


//从ES中获取group列表
func GetESGroupList(start int,limit int)  (r map[int]map[string]interface{},err error) {
    group := Service.GroupEs{}
	ret, _ := group.GetGroupList(start, limit)
	ret = processEsGroupList(ret)
    return ret,nil
}

//从ES中根据one_category获取group列表
func GetESGroupListByCategory(oneCategoryId,start int,limit int) (r map[int]map[string]interface{},total int64,err error)  {
	group := Service.GroupEs{
		 OneCategoryId:oneCategoryId,
	}
	ret, total:= group.GetGroupList(start, limit)
	ret = processEsGroupList(ret)
	return  ret,total,nil
}

//处理从ES中获取的group list 列表
func processEsGroupList(ret map[int]map[string]interface{}) ( map[int]map[string]interface{}) {
	for _,val := range ret{
		equalHImageId,ok := val["equalh_image_id"]
		if ok{
			equalHImageId = equalHImageId.(float64)
			imageId := fmt.Sprintf("%.0f",equalHImageId)
			imgUrl := GetImgUrl(imageId)
			val["equalhImgUrl"] = val["equalh_url"]
			if len(imgUrl) > 0{
				val["equalhImgUrl"] = imgUrl
			}
		}
		imgDate,ok := val["img_date"]
		if ok{
			imgTime := fmt.Sprintf("%.0f",imgDate.(float64))
			i, _ := strconv.ParseInt(imgTime, 10, 64)
			val["img_date_format"] = time.Unix(i,0).Format("2006-01-02 03:04:05")
		}
		groupId,ok := val["group_id"]
		if ok{
			groupId = fmt.Sprintf("%.0f",groupId.(float64))
			val["group_id"] = groupId
		}
	}
	return ret
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








