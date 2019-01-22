package controllers

import (
	"config"
	"html/template"
	"net/http"
	"libary/logger"
)


func Display(file string,data interface{},writer http.ResponseWriter)  {
	file = config.VIEWDIR+file
	files := template.Must(template.ParseFiles(file))
	files.Execute(writer,data)
}

func DisplayLayOut(file string,data interface{},writer http.ResponseWriter)  {
	file = config.VIEWDIR + file
	head := config.VIEWDIR+"public/head.html"
	header := config.VIEWDIR+"public/header.html"
	footer := config.VIEWDIR+"public/footer.html"
	copyRight := config.VIEWDIR+"public/copyright.html"
	must := template.Must(template.ParseFiles(file,head,header,footer,copyRight))
	must.Funcs(template.FuncMap{"unescaped": unescaped})
	must.Execute(writer,data)
}

func unescaped (x string) interface{} {
	return template.HTML(x)
}

func Redirect404(w http.ResponseWriter,r *http.Request)  {
   http.Redirect(w,r,"/404",http.StatusFound)
}

func SaveErrorLog(url string,e error)  {
      logger.ErrorLog.Println("URL:",url,"ERR:",e)
}

