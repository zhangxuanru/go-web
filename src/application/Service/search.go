package Service

import (
	"libary/ES"
	"github.com/olivere/elastic"
	"context"
	"libary/logger"
	"encoding/json"
)

type Search struct {
     Keyword string
     TopicId int
     Start int
     Size int
}

var(
	searchResult *elastic.SearchResult
	result map[int]map[string]interface{}
	err error
	totalHits int64
)

func (search *Search) TopicSearch()  {

}

//在指定的topic内搜索group，根据topic_id
func (search *Search) TopicGroupSearch() (result map[int]map[string]interface{},total int64) {
	query:= elastic.NewBoolQuery().Must(elastic.NewTermQuery("topic_id",search.TopicId),elastic.NewMatchQuery("title",search.Keyword))
	include := elastic.NewFetchSourceContext(true).Include("group_id","title","equalw_url","equalw_url_imageid","group_pics_num","img_date")
	searchResult,err = ES.GetEs().Search().Index("topic_group").Type("_doc").Query(query).FetchSourceContext(include).From(search.Start).Size(search.Size).Sort("img_date",false).Do(context.Background())
	if err != nil{
		logger.ErrorLog.Println("TopicGroupSearch ES ERROR:",err)
		return
	}
	totalHits = searchResult.TotalHits()
	result = make(map[int]map[string]interface{},search.Size)
	for k, hit := range searchResult.Hits.Hits {
		item := make(map[string]interface{})
		if e := json.Unmarshal(*hit.Source, &item);e != nil{
			continue
		}
		result[k] = item
	}
     return result,totalHits
}




func (search *Search) Search()  {

}



