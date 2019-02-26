package logic

import (
	"application/Service"
)

type Search struct {
	Keyword string
	UniqueTopic bool
	Phrase bool
	TopicId int
	Start int
	Size int
}

func initSearch(search *Search) (service  Service.Search) {
	service = Service.Search{
		 Keyword:search.Keyword,
		 UniqueTopic:search.UniqueTopic,
		 Phrase:search.Phrase,
		 Start:search.Start,
		 Size:search.Size,
		 TopicId:search.TopicId,
	}
	return service
}


//在指定的topic内搜索group
func (search *Search) TopicGroupSearch() (result map[int]map[string]interface{},total int64){
	service := initSearch(search)
	result, total = service.TopicGroupSearch()
	result = processEsTopicGroup(result)
	return result,total
}

//根据关键字 搜索topic，
func (search *Search) GetTopicSearch() (result map[int]map[string]interface{},total int64){
	service := initSearch(search)
	result,total = service.GetTopicSearch()
	result = processEsTopicList(result)
	return
}


//搜索group
func (search *Search) GroupSearch() (result map[int]map[string]interface{},total int64) {
	service := initSearch(search)
	result,total = service.GroupSearch()
	result = processEsGroupList(result)
	return
}


//搜索框关键字自动补全
func CompletionData(keyWord string) (result map[string]string) {
	 search := Search{
	 	Keyword:keyWord,
	 	Size:10,
	 	Start:0,
	 }
	service := initSearch(&search)
	result = service.CompletionData()
	return
}



