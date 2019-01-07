package controllers

import (
	"net/http"
	"strings"
	"application/logic"
	"strconv"
	"config"
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
	row, err := logic.GetGoodsDetailById(groupId, "")
	if err != nil{
		Redirect404(writer,request)
		return
	}
	result := make(map[string]interface{})
	result["row"] = row
	result["title"] = row["title"]
	result["description"] = config.DESCRIPTION
	result["keywords"] = config.KEYWORDS
	DisplayLayOut("group/index.html",result,writer)
}
