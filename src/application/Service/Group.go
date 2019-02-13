package Service

import (
	"libary/ES"
     "context"
	"encoding/json"
	"github.com/olivere/elastic"
)
type GroupEs struct {
	OneCategoryId  int
}

//去ES中获取group list
func (group *GroupEs) GetGroupList(start int,limit int) (ret map[int]map[string]interface{},count int64) {
	var result *elastic.SearchResult
	var e error
	if group.OneCategoryId > 0{
		query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("one_category",group.OneCategoryId))
		result, e = ES.GetEs().Search().Index("group").Type("_doc").Query(query).From(start).Size(limit).Sort("img_date",false).Do(context.Background())
	}else{
	    result, e = ES.GetEs().Search().Index("group").Type("_doc").From(start).Size(limit).Sort("img_date",false).Do(context.Background())
	}
    if e != nil{
    	  return
	}
	ret = make(map[int]map[string]interface{},limit)
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








