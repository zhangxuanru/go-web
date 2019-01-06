package controllers

import (
	"net/http"
	"fmt"
	"strings"
	"application/logic"
	"strconv"
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
	list, err := logic.GetGoodsDetailById(groupId, "")
	if err != nil{
		Redirect404(writer,request)
		return
	}

	fmt.Fprintln(writer,request.URL,request.URL.Query(),request.URL.Path)

	fmt.Fprintln(writer,groupId)

}
