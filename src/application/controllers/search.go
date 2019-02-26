package controllers

import (
	"net/http"
	"application/logic"
	"strings"
	"strconv"
	"fmt"
	"math"
	"config"
	"encoding/json"
)

var (
	assign  map[string]interface{}
	search  logic.Search
	result  map[int]map[string]interface{}
	totalHit int64
	searchGroupList  map[int]map[string]interface{}
	groupTotalHit int64
	topicTotalHit int64
	groupPageCount int
	keyWord string
	Phrase  bool
	searchType string
	link string
)

func initSearch(keyWord string,start,size int,Phrase bool) (logic.Search)  {
	return  logic.Search{
		Keyword:keyWord,
		Phrase:Phrase,
		UniqueTopic:true,
		Start:start,
		Size:size,
	}
}


//根据关键字搜索
func Search(w http.ResponseWriter, r *http.Request) {
	assign = make(map[string]interface{})
	keyWord = strings.TrimSpace(r.FormValue("keyword"))
	Phrase = false
	if phrase := strings.TrimSpace(r.FormValue("phrase"));phrase == "1"{
		Phrase = true
	}
	if searchType = strings.TrimSpace(r.FormValue("searchType"));searchType==""{
		searchType = "topic"
	}
	if page,_ = strconv.Atoi(r.FormValue("page"));page == 0{
		page = 1
	}
	 service := initSearch(keyWord, (page-1)*size, size,Phrase)
	 //topic
	 result, topicTotalHit = service.GetTopicSearch()
	 //group
	 groupList,groupTotalHit = service.GroupSearch()
	 total = topicTotalHit
	 link = "topic"
	 totalHit = groupTotalHit + topicTotalHit
	 if searchType != "topic"{
		  result = groupList
		  total = groupTotalHit
		  link = "group"
	 }
	if total > 0 && int(total) > size{
		sumPage := fmt.Sprintf("%.0f",math.Ceil(float64(total)/float64(size)))
		pageCount,_ = strconv.Atoi(sumPage)
	}
	assign["List"] = result
    assign["total"] = total
    assign["topicTotalHit"] = topicTotalHit
    assign["groupTotalHit"] = groupTotalHit
    assign["link"] = link
    assign["totalHit"] = totalHit
    assign["page"] = page
    assign["prevPage"] = page-1
    assign["nextPage"] = page+1
    assign["searchType"] = searchType
    assign["pageCount"] = pageCount
    assign["keyWord"] = keyWord
	assign["title"] = keyWord +" 搜索结果 - 编辑图片"
	assign["keywords"] = keyWord + "  " + config.KEYWORDS
	assign["description"] = keyWord +" " + config.DESCRIPTION
    DisplayLayOut("search/index.html",assign,w)
}


//搜索框关键字自动补全
func Completion(w http.ResponseWriter,r *http.Request)  {
	keyWord = strings.TrimSpace(r.FormValue("keyword"))
	list := make(map[string]interface{})
	list["code"] = "200"
	list["data"] = make(map[string]string)
	if len(keyWord) > 0{
	    list["data"] = logic.CompletionData(keyWord)
	}
	data,_ := json.Marshal(list)
	fmt.Fprintln(w,string(data))
}




