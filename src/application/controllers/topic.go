package controllers

import (
	"net/http"
	"config"
	"application/logic"
	"strconv"
	"fmt"
)

func TopicList(w http.ResponseWriter, r *http.Request)  {
	result := make(map[string]interface{})
	cateId := r.FormValue("categoryId")
	categoryId,_ := strconv.Atoi(cateId)
	page,_ := strconv.Atoi(r.FormValue("page"))
	if page == 0{
		page = 1
	}
	service := logic.Topic{
			CategoryId:categoryId,
			Start:1,
			Limit:1,
	}
	rs, total := logic.GetESTopicListByCategoryId(&service)

	fmt.Println(total)
	fmt.Println()
	fmt.Println(len(rs))
	fmt.Printf("%+v",rs)



	//顶级分类
	topCategoryList, _ := logic.GetAllTopCategoryList( 0,13)

	result["topCategoryList"] = topCategoryList
	result["cateId"] = cateId
	result["categoryId"] = categoryId
	result["page"] = page
	result["title"] = config.TITLE
	result["keywords"] = "编辑图片专题,正版图片素材专题,图片素材专题下载,图片专题下载,编辑图片专题下载"
	result["description"] = config.DESCRIPTION
	DisplayLayOut("topic/list.html",result,w)
}



func TopicDetail(w http.ResponseWriter,r *http.Request)  {
	result := make(map[string]interface{})

	DisplayLayOut("topic/detail.html",result,w)
}


