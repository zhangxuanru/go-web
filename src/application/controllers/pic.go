package controllers

import (
	"net/http"
	"fmt"
	"strings"
)

//图片详情页
func PicPageDetail(w http.ResponseWriter, r *http.Request)  {
	picId := strings.Replace(r.URL.Path, "/editorialPic/", "", -1)
	groupId := r.FormValue("groupid")
	if len(picId) == 0 || len(groupId) == 0{
		Redirect404(w,r)
		return
	}


   fmt.Fprintln(w,"hello","---",picId,"-----",groupId)
}




