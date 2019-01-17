package Service

import (
	"libary/ES"
	"github.com/olivere/elastic"
	"context"
	"fmt"
)

//根据categoryId去ES中搜索group数据
func GetCategoryGroupList(categoryId int)  {
	query := elastic.NewTermQuery("category_id",categoryId)

	result, e := ES.GetEs().Search().Index("group_category").Type("_doc").Query(query).Pretty(true).Do(context.Background())
    fmt.Println(e)

    fmt.Println(result.Hits.Hits)

	for _, hit := range result.Hits.Hits {
		//item := make(map[string]interface{})
	   // json.Unmarshal(*hit.Source, &item)
		fmt.Printf("%+v",hit)
		fmt.Println()
	}


 }








