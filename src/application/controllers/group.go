package controllers

import (
	"net/http"
	"strings"
	"application/logic"
	"strconv"
	"config"
	"math"
)

func GroupDetail(writer http.ResponseWriter, request *http.Request)  {
	pathId := strings.Replace(request.URL.Path, "/group/", "", -1)
	if len(pathId) == 0{
		Redirect404(writer,request)
		return
	}
	groupId, _ := strconv.Atoi(pathId)
	if groupId == 0 {
		Redirect404(writer,request)
		return
	}
	strPage := request.FormValue("page")
	if strPage == ""{
		strPage = "1"
	}
	page, _ := strconv.Atoi(strPage)
	if page <= 0 || page > 1000{
		page = 1
	}
	start := (page-1) * Limit
	//group detail
	row, err := logic.GetGoodsDetailById(groupId, "")
	if err != nil{
		Redirect404(writer,request)
		return
	}
	//get pic list
	picList, e := logic.GetPicListByGroupId(groupId,"",start,Limit)
	if e!= nil{
		Redirect404(writer,request)
		return
	}
	totalPage := float64(0)
	groupPicsNum,ok := row["group_pics_num"]
	if ok{
		num, _ := strconv.ParseFloat(groupPicsNum, 64)
		totalPage = math.Ceil(num/float64(Limit))
	}
	result := make(map[string]interface{})
	result["row"] = row
	result["totalPage"] = totalPage
	result["page"] = page
	result["title"] = row["title"]
	result["height"] = row["height"]
	result["width"] = row["width"]
	result["groupId"] = groupId
	result["picList"] = picList
	result["description"] = config.DESCRIPTION
	result["keywords"] = config.KEYWORDS
	result["baseUrl"] = config.BASEURL
	DisplayLayOut("group/index.html",result,writer)
}







