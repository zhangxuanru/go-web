package Service

import (
	"libary/ES"
	"github.com/olivere/elastic"
	"context"
	"libary/logger"
	"encoding/json"
	"unicode/utf8"
)

type Search struct {
     Keyword string
	 UniqueTopic bool
     TopicId int
     Start int
     Size int
}

var(
	searchResult *elastic.SearchResult
	result map[int]map[string]interface{}
	querys []elastic.Query
	query elastic.Query
	matchQuery *elastic.MatchQuery
	matchPhraseQuery *elastic.MatchPhraseQuery
	termQuery *elastic.TermQuery
	err error
	totalHits int64
)

//在指定的topic内搜索group
func (search *Search) TopicGroupSearch() (result map[int]map[string]interface{},total int64) {
	if utf8.RuneCountInString(search.Keyword) < 4{
		matchPhraseQuery = elastic.NewMatchPhraseQuery("title",search.Keyword)
		querys = append(querys,matchPhraseQuery)
	}else{
		matchQuery = elastic.NewMatchQuery("title",search.Keyword)
		querys = append(querys,matchQuery)
	}
	if search.TopicId > 0{
		termQuery = elastic.NewTermQuery("topic_id",search.TopicId)
		querys = append(querys,termQuery)
	}
    query = elastic.NewBoolQuery().Must(querys...)
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


//在group中搜索
func (search *Search) GroupSearch()  (result map[int]map[string]interface{},total int64) {
	if utf8.RuneCountInString(search.Keyword) < 4{
		matchPhraseQuery = elastic.NewMatchPhraseQuery("title",search.Keyword)
		querys = append(querys,matchPhraseQuery)
	}else{
		matchQuery = elastic.NewMatchQuery("title",search.Keyword)
		querys = append(querys,matchQuery)
	}
	query = elastic.NewBoolQuery().Must(querys...)
	include := elastic.NewFetchSourceContext(true).Include("group_id","title","group_pics_num","img_date","equalw_url","equalw_url_imgid","equalh_url","equalh_image_id")
	searchResult,err = ES.GetEs().Search().Index("group").Type("_doc").Query(query).FetchSourceContext(include).From(search.Start).Size(search.Size).Sort("img_date",false).Do(context.Background())
	if err != nil{
		logger.ErrorLog.Println("GroupSearch ES ERROR:",err)
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


//在topic中搜索
func (search *Search) TopicSearch() (result map[int]map[string]interface{},total int64) {

	return
}



