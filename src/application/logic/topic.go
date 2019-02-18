package logic

import (
	"application/Service"
	"fmt"
	"libary/util"
)

type Topic struct {
	CategoryId int
	TopicId int
	Start int
	Limit int
}

func initTopic(topic *Topic) ( Service.Topic) {
	service := Service.Topic{
		CategoryId:topic.CategoryId,
		TopicId:topic.TopicId,
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

//根据topicId获取group list 数据
func (topic *Topic) GetTopicGroupList() (r map[int]map[string]interface{},total int64) {
	service := initTopic(topic)
	ret, count := service.GetTopicGroupList()
	ret = processEsTopicGroup(ret)
	return ret,count
}


//根据topicId获取topic详情
func (topic *Topic) GetTopicDetail() (ret map[string]interface{}) {
	service := initTopic(topic)
	return service.GetTopicDetail()
}




//处理从ES中获取的topic  list 列表
func processEsTopicList(ret map[int]map[string]interface{}) ( map[int]map[string]interface{}) {
   for _,val := range ret{
	   val["img_url"] = val["equalw_url"]
       if imageId,ok := val["image_id"];ok{
			  imageId := fmt.Sprintf("%.0f",imageId.(float64))
			  if  imgUrl := GetImgUrl(imageId); len(imgUrl) > 0{
				   val["img_url"] = imgUrl
			  }
		  }
	   if publishTime,ok := val["created_time"];ok{
		   pubTime := fmt.Sprintf("%.0f",publishTime.(float64))
		   val["publish_date"] =  util.FormattingTimeRubbing(pubTime)
	   }
	   if topicId,ok := val["topic_id"];ok{
		   topicId = fmt.Sprintf("%.0f",topicId.(float64))
		   val["topic_id"] = topicId
	   }
   }
	return ret
}


//处理从ES中获取的topic group数据
func processEsTopicGroup(ret map[int]map[string]interface{}) ( map[int]map[string]interface{}) {
	for _,val := range ret{
		val["img_url"] = val["equalw_url"]
		if imageId,ok := val["equalw_url_imageid"]; ok{
			imageId := fmt.Sprintf("%.0f",imageId.(float64))
			if  imgUrl := GetImgUrl(imageId);len(imgUrl) > 0{
				val["img_url"] = imgUrl
			}
		}
		if publishTime,ok := val["img_date"];ok{
			pubTime := fmt.Sprintf("%.0f",publishTime.(float64))
			val["publish_date"] = util.FormattingTimeRubbing(pubTime)
		}
		if groupId,ok := val["group_id"];ok{
			groupId = fmt.Sprintf("%.0f",groupId.(float64))
			val["group_id"] = groupId
		}
	}
	return  ret
}




