package controllers

import (
	"net/http"
	"config"
	"application/logic"
	"strconv"
	"fmt"
	"math"
	"strings"
)

//专题列表页
func TopicList(w http.ResponseWriter, r *http.Request)  {
	result := make(map[string]interface{})
	cateId := r.FormValue("categoryId")
	categoryId,_ := strconv.Atoi(cateId)
	page,_ := strconv.Atoi(r.FormValue("page"))
	pageCount := 0
	if page == 0{
		page = 1
	}
	service := logic.Topic{
			CategoryId:categoryId,
			Start:(page-1)*size,
			Limit:size,
	}
	topicList, total := service.GetESTopicListByCategoryId()
	if total > 0{
		sumPage := fmt.Sprintf("%.0f",math.Ceil(float64(total)/float64(size)))
		pageCount,_ = strconv.Atoi(sumPage)
	}
	//顶级分类
	topCategoryList, _ := logic.GetAllTopCategoryList( 0,13)

	result["topCategoryList"] = topCategoryList
	result["topicList"] = topicList
	result["pageCount"] = pageCount
	result["total"] = total
	result["cateId"] = cateId
	result["categoryId"] = categoryId
	result["page"] = page
	result["baseUrl"] = config.BASEURL
	result["title"] = config.TITLE
	result["keywords"] = "编辑图片专题,正版图片素材专题,图片素材专题下载,图片专题下载,编辑图片专题下载"
	result["description"] = config.DESCRIPTION
	DisplayLayOut("topic/list.html",result,w)
}



//专题详情页
func TopicDetail(w http.ResponseWriter,r *http.Request)  {
	id := strings.Replace(r.URL.Path, "/topic/", "", -1)
	if len(id) == 0{
		Redirect404(w,r)
		return
	}
	topicId, _ := strconv.Atoi(id)
	if topicId == 0 {
		Redirect404(w,r)
		return
	}
	page,_ := strconv.Atoi(r.FormValue("page"))
	if page == 0{
		page = 1
	}
	pageCount := 0
	service := logic.Topic{
         TopicId:topicId,
		 Start:(page-1)*size,
		 Limit:size,
	}
	groupList, total := service.GetTopicGroupList()
	if total > 0{
		sumPage := fmt.Sprintf("%.0f",math.Ceil(float64(total)/float64(size)))
		pageCount,_ = strconv.Atoi(sumPage)
	}
	detail := service.GetTopicDetail()

	fmt.Println()
	fmt.Printf("%+v",groupList)


	result := make(map[string]interface{})
	result["groupList"] = groupList
	result["detail"] = detail
	result["pageCount"] = pageCount
	result["total"] = total
	result["baseUrl"] = config.BASEURL
	result["title"] = detail["title"]
	result["keywords"] = "编辑图片专题,正版图片素材专题,图片素材专题下载,图片专题下载,编辑图片专题下载"
	result["description"] = config.DESCRIPTION
	DisplayLayOut("topic/detail.html",result,w)
}













