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
      http.HandleFunc("/editorialPic/",Call(controllers.PicPageDetail))
      http.HandleFunc("/editorial/all-update/",Call(controllers.Roll))
	  http.HandleFunc("/editorial-topics/",Call(controllers.TopicList))
	  http.HandleFunc("/topic/",Call(controllers.TopicDetail))
	  http.HandleFunc("/search/",Call(controllers.Search))
      http.HandleFunc("/completion/",Call(controllers.Completion))
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
