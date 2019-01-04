package controllers

import "net/http"

//404 page
func NotFound(writer http.ResponseWriter, request *http.Request)  {
	DisplayLayOut("public/404.html",nil,writer)
}



