package controllers

import (
	"net/http"
	"application/logic"
	"math"
	"strconv"
)

var Limit = 100
var r  map[int]map[string]string
var e error

func Index(writer http.ResponseWriter, request *http.Request)  {
	result := make(map[string]interface{})
	strPage := request.FormValue("page")
	if strPage == ""{
		 strPage = "1"
	}
	page, _ := strconv.Atoi(strPage)
	if page <= 0 || page > 1000{
		page = 1
	}
	start := (page-1) * Limit
	//group
	total := logic.GetGroupCount("")
	if total > 0{
	    r,e = logic.GetGroupList("", start, Limit)
		if e != nil{
			SaveErrorLog("index.Index",e)
			Redirect404(writer,request)
			return
		}
	}
	//banner
	 lBanner,_ := logic.GetLeftBannerList(0,8)
	 rBanner,_ := logic.GetRightBanner(0,4)

     totalPage := math.Ceil(float64(total)/float64(Limit))
	 result["groupList"] = r
 	 result["lBanner"] = lBanner
	 result["rBanner"] = rBanner
	 result["total"] = total
	 result["totalPage"] = totalPage
	 result["page"] = page
	 DisplayLayOut("index/index.html",result,writer)
}



