package controllers

import (
	"net/http"
	"config"
	"application/logic"
	"strconv"
	"fmt"
	"math"
)

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



func TopicDetail(w http.ResponseWriter,r *http.Request)  {
	result := make(map[string]interface{})

	DisplayLayOut("topic/detail.html",result,w)
}


