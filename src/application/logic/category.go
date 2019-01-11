package logic

import (
	"application/models"
	"strconv"
	"libary/util"
	"strings"
)

func GetCateGoryDataById(categoryId int) (r map[string]string, err error) {
	 return  models.GetCateGoryDataById(categoryId)
}

func GetCateGoryById(categoryId int) (r map[string]string, err error) {
	r, err = GetCateGoryDataById(categoryId)
	if len(r) > 0{
		 replacer := strings.NewReplacer("新闻滚动", "","滚动","")
		 r["category_name"] = replacer.Replace(r["category_name"])
	}
	return
}

//栏目详情页整理数据
func GetCateColDetail(categoryId int)  (linkTags,picGeneralize,channelRecommend map[int]map[string]string,err error) {
	linkTags, err = GetLinkDataByCategoryId(categoryId)
 	picGeneralize, _ = GetRecommendData(categoryId, " data_type='1' ",0,4)
	channelRecommend,_ = GetRecommendData(categoryId, " data_type='2' ",0,12)
	for _,item :=  range channelRecommend{
         if item["image_id"] != "0"{
			 item["src"] = GetImgUrl(item["image_id"])
		 }
		 if item["item_id"] != "0"{
		 	 groupId,_ := strconv.Atoi(item["item_id"])
			 groupRow, err := models.GetGoodsDetailById(groupId,"")
			 if err != nil || len(groupRow) == 0{
			 	 continue
			 }
			 item["group_id"] = groupRow["group_id"]
			 item["group_pics_num"] = groupRow["group_pics_num"]
			 item["img_date"] = util.FormattingTimeRubbing(groupRow["img_date"])
		 }
	 }
    for _,val := range picGeneralize{
		if val["image_id"] != "0"{
			val["src"] = GetImgUrl(val["image_id"])
		}
	}
	return
}



