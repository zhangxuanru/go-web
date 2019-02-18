package Service

import (
	"github.com/olivere/elastic"
	"libary/ES"
	"encoding/json"
	"context"
	"strconv"
	"libary/logger"
)

type Topic struct {
	CategoryId int
	TopicId int
	Start int
	Limit int
}

//去ES中获取topic list
func (topic *Topic) GetESTopicListByCategoryId() (ret map[int]map[string]interface{},count int64) {
	var result *elastic.SearchResult
	var e error
	if topic.CategoryId > 0{
		query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("category_id",topic.CategoryId))
		result, e = ES.GetEs().Search().Index("topic").Type("_doc").Query(query).From(topic.Start).Size(topic.Limit).Sort("publish_time",false).Do(context.Background())
	}else{
		result, e = ES.GetEs().Search().Index("topic").Type("_doc").From(topic.Start).Size(topic.Limit).Sort("publish_time",false).Do(context.Background())
	}
	if e != nil{
		return
	}
	ret = make(map[int]map[string]interface{},topic.Limit)
	for k, hit := range result.Hits.Hits {
		item := make(map[string]interface{})
		e := json.Unmarshal(*hit.Source, &item)
		if e!=nil{
			continue
		}
		ret[k] = item
	}
	return ret,result.TotalHits()
}


//在ES中根据topicId 获取group list
func (topic *Topic) GetTopicGroupList() (ret map[int]map[string]interface{},count int64) {
	var result *elastic.SearchResult
	var e error
	if topic.TopicId <= 0{
		return
	}
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("topic_id",topic.TopicId))
	result, e = ES.GetEs().Search().Index("topic_group").Type("_doc").Query(query).From(topic.Start).Size(topic.Limit).Sort("img_date",false).Do(context.Background())
	if e != nil{
		return
	}
	ret = make(map[int]map[string]interface{},topic.Limit)
	for k, hit := range result.Hits.Hits {
		item := make(map[string]interface{})
		e := json.Unmarshal(*hit.Source, &item)
		if e!=nil{
			continue
		}
		ret[k] = item
	}
	return ret,result.TotalHits()
}


//在ES中根据topicId获取topic 详情
func (topic *Topic) GetTopicDetail() (ret map[string]interface{}){
	   include := elastic.NewFetchSourceContext(true).Include("title")
	   topicId := strconv.Itoa(topic.TopicId)
	   result, e := ES.GetEs().Get().Index("topic").Type("_doc").Id(topicId).FetchSourceContext(include).Do(context.Background())
	   if e != nil{
		   logger.ErrorLog.Println("GetTopicDetail ES ERROR:",e)
		   return
	  }
	  item := make(map[string]interface{})
	  json.Unmarshal(*result.Source, &item)
	  return item
  }




