package controllers

import (
	"net/http"
	"application/logic"
	"strings"
	"strconv"
	"fmt"
	"math"
)

var (
	assign  map[string]interface{}
	search  logic.Search
	result  map[int]map[string]interface{}
	totalHit int64
	searchGroupList  map[int]map[string]interface{}
	groupTotalHit int64
	groupPageCount int
	keyWord string
)

func initSearch(keyWord string,start,size int) (logic.Search)  {
	return  logic.Search{
		Keyword:keyWord,
		Start:start,
		Size:size,
	}
}


//根据关键字搜索
func Search(w http.ResponseWriter, r *http.Request) {
	assign = make(map[string]interface{})
	keyWord = strings.TrimSpace(r.FormValue("keyword"))
	if page,_ = strconv.Atoi(r.FormValue("page"));page == 0{
		page = 1
	}
	service := initSearch(keyWord, (page-1)*size, size)
	//topic
	result, total = service.TopicGroupSearch()
	//group
	searchGroupList,groupTotalHit = service.GroupSearch()

	if total > 0 && int(total) > size{
		sumPage := fmt.Sprintf("%.0f",math.Ceil(float64(total)/float64(size)))
		pageCount,_ = strconv.Atoi(sumPage)
	}
	if groupTotalHit > 0 && int(groupTotalHit) > size{
		sumPage := fmt.Sprintf("%.0f",math.Ceil(float64(groupTotalHit)/float64(size)))
		groupPageCount,_ = strconv.Atoi(sumPage)
	}
	assign["groupList"] = searchGroupList
	assign["groupTotalHit"] = groupTotalHit
	assign["topicList"] = result
    assign["topicTotal"] = total
    assign["page"] = page
    assign["topicPageCount"] = pageCount
    assign["groupPageCount"] = groupPageCount
    DisplayLayOut("search/index.html",assign,w)
}





