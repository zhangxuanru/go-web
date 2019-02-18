package logic

import (
	"application/Service"
	"fmt"
	"libary/util"
)

type Topic struct {
	CategoryId int
	Start int
	Limit int
}

func initTopic(topic *Topic) ( Service.Topic) {
	service := Service.Topic{
		CategoryId:topic.CategoryId,
		Start:topic.Start,
		Limit:topic.Limit,
	}
	return service
}

//根据分类获取专题列表
func (topic *Topic) GetESTopicListByCategoryId() (r map[int]map[string]interface{},total int64)  {
	service := initTopic(topic)
	ret, total:= service.GetESTopicListByCategoryId()
	ret = processEsTopicList(ret)
	return ret,total
}

//处理从ES中获取的group list 列表
func processEsTopicList(ret map[int]map[string]interface{}) ( map[int]map[string]interface{}) {
   for _,val := range ret{
	   val["img_url"] = val["equalw_url"]
       imageId,ok := val["image_id"]
       if ok{
			  imageId := fmt.Sprintf("%.0f",imageId.(float64))
			  imgUrl := GetImgUrl(imageId)
			  if len(imgUrl) > 0{
				  val["img_url"] = imgUrl
			  }
		  }
	   publishTime,ok := val["created_time"]
	   if ok{
		   pubTime := fmt.Sprintf("%.0f",publishTime.(float64))
		   val["publish_date"] =  util.FormattingTimeRubbing(pubTime)
	   }
	   topicId,ok := val["topic_id"]
	   if ok{
		   topicId = fmt.Sprintf("%.0f",topicId.(float64))
		   val["topic_id"] = topicId
	   }
   }
	return ret
}





