package logic

import (
	"application/models"
)

func GetCateGoryDataById(categoryId int) (r map[string]string, err error) {
	 return  models.GetCateGoryDataById(categoryId)
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
	}

	for _,val := range picGeneralize{
		if val["image_id"] != "0"{
			val["src"] = GetImgUrl(val["image_id"])
		}
	}

	return
}



