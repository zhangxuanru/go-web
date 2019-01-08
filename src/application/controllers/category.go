package controllers

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"application/logic"
)

//栏目页
func Detail(w http.ResponseWriter, r *http.Request)  {
	catId := strings.Replace(r.URL.Path, "/entertainment/", "", -1)
	if len(catId) == 0{
		Redirect404(w,r)
		return
	}
	cId, _ := strconv.Atoi(catId)
	if cId == 0 {
		Redirect404(w,r)
		return
	}

	linkTags, picGeneralize, channelRecommend, err := logic.GetCateColDetail(cId)

	fmt.Fprintln(w,"hello---",cId)

	DisplayLayOut("category/index.html",nil,w)
}








