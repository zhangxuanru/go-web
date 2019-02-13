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
	groupList,total := logic.GetCateGoryGroupList(categoryId, page, size)
	if len(total) > 0{
		count, _ := strconv.Atoi(total)
		countPage = fmt.Sprintf("%.0f",math.Ceil(float64(count/size)))
	}
	topCategoryList, _ := logic.GetAllTopCategoryList( 0,13)

	fmt.Println("-------------------")
	fmt.Println(len(groupList))
	fmt.Printf("%+v",groupList)
	fmt.Println()

	result["groupList"] = groupList
	result["topCategoryList"] = topCategoryList
	result["categoryId"] = categoryId
	result["cateId"] = cateId
	result["total"] = total
	result["countPage"] = countPage
	result["baseUrl"] = config.BASEURL
	DisplayLayOut("roll/index.html",result,w)
}


