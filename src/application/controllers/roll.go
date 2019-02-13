package controllers

import (
	"net/http"
	"config"
	"application/logic"
	"strconv"
	"fmt"
	"math"
)

func Roll(w http.ResponseWriter, r *http.Request)  {
	result := make(map[string]interface{})
	cateId := r.FormValue("categoryId")
	categoryId,_ := strconv.Atoi(cateId)
	page,_ := strconv.Atoi(r.FormValue("page"))
	countPage := "0"
	if page == 0{
		page = 1
	}
	groupList, total,_ := logic.GetESGroupListByCategory(categoryId, (page-1)*size, size)
	if total > 0{
		countPage = fmt.Sprintf("%.0f",math.Ceil(float64(float64(total)/float64(size))))
	}
	topCategoryList, _ := logic.GetAllTopCategoryList( 0,13)
	result["groupList"] = groupList
	result["topCategoryList"] = topCategoryList
	result["categoryId"] = categoryId
	result["cateId"] = cateId
	result["total"] = total
	result["countPage"] = countPage
	result["page"] = page
	result["baseUrl"] = config.BASEURL
	result["title"] = config.TITLE
	result["keywords"] = config.KEYWORDS
	result["description"] = config.DESCRIPTION
	DisplayLayOut("roll/index.html",result,w)
}


