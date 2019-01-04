package logic

import (
	"application/models"
	"strings"
	"fmt"
)

func GetLeftBannerList(start int,limit int)  (r map[int]map[string]string,err error) {
    list, e := models.GetBannerList("img_width > 300", start, limit)
    if e != nil{
    	 return  list, e
	}
	list = processBannerImgUrl(list)
    return list,e
}

func GetRightBanner(start int,limit int)(r map[int]map[string]string,err error)  {
	list, e := models.GetBannerList("img_width < 300 ", start, limit)
	list = processBannerImgUrl(list)
	return list,e
}

func processBannerImgUrl(list  map[int]map[string]string) (r map[int]map[string]string) {
	for _,v := range list{
		imgUrl := GetImgUrl(v["img_id"])
		if len(imgUrl) > 0{
			v["imgUrl"] = imgUrl
		}else{
			v["imgUrl"] = v["origin_imgurl"]
		}
	  v["href"] = fmt.Sprintf("/group/%s",v["link_id"])
	  if strings.Contains(v["link"],"topic"){
	  	   v["href"] = fmt.Sprintf("/topic/%s",v["link_id"])
	  }
    }
	return list
}






