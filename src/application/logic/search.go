package logic

import (
	"application/Service"
)

type Search struct {
	Keyword string
	TopicId int
	Start int
	Size int
}

func initSearch(search *Search) (service  Service.Search) {
	service = Service.Search{
		 Keyword:search.Keyword,
		 Start:search.Start,
		 Size:search.Size,
		 TopicId:search.TopicId,
	}
	return service
}


//在指定的topic内搜索group，根据topic_id
func (search *Search) TopicGroupSearch() (result map[int]map[string]interface{},total int64){
	service := initSearch(search)
	result, total = service.TopicGroupSearch()
	result = processEsTopicGroup(result)
	return result,total
}
