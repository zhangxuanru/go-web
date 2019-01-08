package logic

import (
	"application/models"
	"fmt"
)

func GetCateGoryDataById(categoryId int) (r map[string]string, err error) {
	 return  models.GetCateGoryDataById(categoryId)
}

//栏目详情页整理数据
func GetCateColDetail(categoryId int)  (linkTags,picGeneralize,channelRecommend map[int]map[string]string,err error) {
	linkTags, err = GetLinkDataByCategoryId(categoryId)
	fmt.Printf("%+v",linkTags)
	picGeneralize = make(map[int]map[string]string)
	channelRecommend = make(map[int]map[string]string)
	list, _ := GetRemmendData(categoryId, "")
	if len(list) > 0{
	   for key,item := range list{
	   	    if item["image_id"] != "0" {
	   	    	item["src"] = GetImgUrl(item["image_id"])
			}
            if item["data_type"] == "1"{
				 picGeneralize[key] = item
			}else{
				 channelRecommend[key] = item
			}
	   }
	}
    list = nil
	return
}



