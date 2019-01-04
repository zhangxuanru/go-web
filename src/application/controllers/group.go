package controllers

import (
	"net/http"
	"fmt"
	"strings"
)

func GroupDetail(writer http.ResponseWriter, request *http.Request)  {
	groupId := strings.Replace(request.URL.Path, "/group/", "", -1)

	fmt.Fprintln(writer,request.URL,request.URL.Query(),request.URL.Path)

	fmt.Fprintln(writer,groupId)

}
