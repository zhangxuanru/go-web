package router

import (
	"net/http"
	"application/controllers"
	"libary/logger"
)

func InitRouter()  {
      http.HandleFunc("/",Call(controllers.Index))
      http.HandleFunc("/404",Call(controllers.NotFound))
      http.HandleFunc("/group/",Call(controllers.GroupDetail))
      http.HandleFunc("/entertainment/",Call(controllers.Detail))
 }

func Call( hand http.HandlerFunc) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
    	if request.RequestURI == "/images/favicon.ico"{
    		 return
		}
		defer func() {
			if err := recover(); err != nil{
				logger.ErrorLog.Println("URL",request.URL,"ERR:",err)
			}
		}()
		 logger.AccessLog.Println("URL:",request.URL,"RequestURI:",request.RequestURI,"cookie:",request.Cookies(),"user-agent:",request.UserAgent(),"Referer:",request.Referer())
		 hand(writer,request)
	}
}
