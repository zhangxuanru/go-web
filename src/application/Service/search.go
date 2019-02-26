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
	 Phrase bool
     TopicId int
     Start int
     Size int
}

var(
	searchResult *elastic.SearchResult
	searchService  *elastic.SearchService
	searchFields []string
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
	querys = querys[:0]
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
	querys = querys[:0]
	if utf8.RuneCountInString(search.Keyword) < 4 || search.Phrase == true{
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



//根据关键字在topic_group中搜索
func (search *Search) GetTopicSearch() (result map[int]map[string]interface{},total int64) {
	querys = querys[:0]
	searchFields = []string{"topic_id","group_id","title","equalw_url","equalw_url_imageid","group_pics_num","img_date"}
	searchService = ES.GetEs().Search().Index("topic_group").Type("_doc").From(search.Start).Size(search.Size).
		Sort("img_date",false)
	if utf8.RuneCountInString(search.Keyword) < 4 || search.Phrase == true{
		matchPhraseQuery = elastic.NewMatchPhraseQuery("title",search.Keyword)
		querys = append(querys,matchPhraseQuery)
	}else{
		matchQuery = elastic.NewMatchQuery("title",search.Keyword)
		querys = append(querys,matchQuery)
	}
	query = elastic.NewBoolQuery().Must(querys...)
	include := elastic.NewFetchSourceContext(true).Include(searchFields...)
	searchResult,err = searchService.Query(query).FetchSourceContext(include).Collapse(elastic.NewCollapseBuilder("topic_id")).Do(context.Background())
	if err != nil{
		logger.ErrorLog.Println("GetTopicSearch ES ERROR:",err)
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
	if len(result) < int(totalHits){
		 totalHits = int64(len(result))
	}
	return result,totalHits
}


//搜索框关键字自动补全
func (search *Search) CompletionData() (result map[string]string) {
	groupSuggestName := "group_completion"
	topicSuggestName := "topic_completion"
	result = search.GetCompletionData("group",groupSuggestName,"title_completion")
	topicResult := search.GetCompletionData("topic",topicSuggestName,"title_completion")
	for id,text := range topicResult{
		result[id] = text
	}
	if len(result) > search.Size{
		ret,i := make(map[string]string),0
		for key,val := range result{
			if i >= search.Size{
				break
			}
			ret[key] = val
			i++
		}
        return  ret
	}
	return result
}


//获取自动补全数据
func (search *Search) GetCompletionData(index,suggestName,Field string)(result map[string]string) {
	result = make(map[string]string)
	searchResult,err = search.getCompletionService(index,suggestName,Field)
	if err != nil{
		logger.ErrorLog.Println("CompletionData ES ERROR:",err)
		return
	}
	searchSuggest := searchResult.Suggest
	if len(searchSuggest) == 0{
		return
	}
	for _,v := range searchSuggest[suggestName]{
		for _,val := range v.Options{
			result[val.Id] = val.Text
		}
	}
	return result
}



//获取自动补全服务
func (search *Search) getCompletionService(index,suggestName,Field string) (searchResult *elastic.SearchResult,err error) {
	completionSuggester := elastic.NewCompletionSuggester(suggestName).Text(search.Keyword).Field(Field)
	searchResult,err = ES.GetEs().Search().Index(index).Suggester(completionSuggester).From(search.Start).Size(search.Size).Do(context.Background())
	if err != nil{
		logger.ErrorLog.Println("CompletionData ES ERROR:",err)
		return
	}
	return  searchResult,err
}

