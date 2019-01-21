package Service

import (
	"libary/ES"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
	"libary/logger"
	"fmt"
	"libary/redis"
	"strings"
	"strconv"
)

//根据categoryId去ES中搜索group数据
func GetCategoryGroupList(categoryId,start,limit int) (ret map[int]map[string]interface{},count string) {
	cacheKey := fmt.Sprintf("group_category_%d_page_%d_limit_%d",categoryId,start,limit)
	totalKey := fmt.Sprintf("group_category_%d",categoryId)
	val := redis.Get(cacheKey)
	count = redis.Get(totalKey)
	if len(val) > 0 && len(count) > 0{
		ret = GetGroupDataByGroupIdLine(val)
		return ret,count
	}
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("category_id",categoryId))
	include := elastic.NewFetchSourceContext(true).Include("group_id")
	result, e := ES.GetEs().Search().Index("group_category").Type("_doc").Query(query).FetchSourceContext(include).From((start-1)*limit).Size(limit).Sort("img_date",false).Do(context.Background())
    if e != nil{
    	logger.ErrorLog.Println("GetCategoryGroupList ES ERROR:",e)
    	return
	}
	total := strconv.FormatInt( result.TotalHits(),10)
	groupIdLine := ""
	for _, hit := range result.Hits.Hits {
		item := make(map[string]interface{})
		e := json.Unmarshal(*hit.Source, &item)
		if e != nil{
			continue
		}
		groupId:=  fmt.Sprintf("%.0f",item["group_id"])
		item["group_id"] = groupId
		groupIdLine += groupId+","
	}
	if len(groupIdLine) == 0{
		return
	}
	redis.Set(cacheKey,groupIdLine,"EX",600)
	redis.Set(totalKey,total,"EX",600)
	ret = GetGroupDataByGroupIdLine(groupIdLine)
    return ret,total
}

//根据groupid string 去索引中找具体的group数据
func GetGroupDataByGroupIdLine(groupIdLine string)(ret map[int]map[string]interface{}) {
	      groupIdLine = strings.Trim(groupIdLine,",")
		  groupIdList := strings.Split(groupIdLine,",")
	 	  groupIdSl := make([]interface{},len(groupIdList))
		  for k,val := range groupIdList{
			   groupIdSl[k] =  val
		  }
	     query := elastic.NewBoolQuery().Must(elastic.NewTermsQuery("group_id", groupIdSl...))
	     result, e := ES.GetEs().Search().Index("group").Type("_doc").Size(len(groupIdList)).Query(query).Sort("img_date",false).Do(context.Background())
	     if e != nil{
	          return
		 }
	     ret = make(map[int]map[string]interface{},len(groupIdList))
	     for k,hit := range result.Hits.Hits{
			 item := make(map[string]interface{})
			 err := json.Unmarshal(*hit.Source, &item)
			 if err!=nil{
			 	 continue
			 }
			 ret[k] = item
		 }
	     return
}










