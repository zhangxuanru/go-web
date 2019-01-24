package Service

import (
	"libary/ES"
     "context"
	"encoding/json"
)
type GroupEs struct {
}

//去ES中获取group list
func (group *GroupEs) GetGroupList(start int,limit int) (ret map[int]map[string]interface{},count int64) {
	result, e := ES.GetEs().Search().Index("group").Type("_doc").From(start).Size(limit).Sort("img_date",false).Do(context.Background())
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








