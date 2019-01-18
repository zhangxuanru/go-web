package controllers

import (
	"net/http"
	"strings"
	"strconv"
	"application/logic"
	"libary/logger"
	"config"
	"math"
	"fmt"
)

//栏目页
var linkTags,picGeneralize,channelRecommend map[int]map[string]string
var catRow map[string]string
var subCategoryList  map[int]map[string]string
var err error
var size = 100

//页面分页还没有做
func Detail(w http.ResponseWriter, r *http.Request)  {
	catId := strings.Replace(r.URL.Path, "/entertainment/", "", -1)
	pid := r.FormValue("pid")
	cId, _ := strconv.Atoi(catId)
	pCid,_ := strconv.Atoi(pid)
	page,_ := strconv.Atoi(r.FormValue("page"))
	if cId == 0 {
		Redirect404(w,r)
		return
	}
	//推荐group，链接TAG
	if pCid > 0{
		linkTags, picGeneralize, channelRecommend, err = logic.GetCateColDetail(pCid)
	}else{
	    linkTags, picGeneralize, channelRecommend, err = logic.GetCateColDetail(cId)
	}
	if err!=nil{
		logger.ErrorLog.Println(r.URL.Path,"-catId:",catId,"-error:",err)
		Redirect404(w,r)
		return
	}
	//子分类
	if pCid > 0{
		 subCategoryList, _ = logic.GetSubCateGoryData(pCid,0,12)
	}else{
		 subCategoryList, _ = logic.GetSubCateGoryData(cId,0,12)
	}
	//当前分类信息
   if pCid > 0{
	   catRow, _ = logic.GetCateGoryById(pCid)
   }else{
	   catRow, _ = logic.GetCateGoryById(cId)
   }
	if len(pid) == 0{
		pid = catId
	}
   //当前分类group list
   if page == 0{
	    page = 1
   }
   groupList,total := logic.GetCateGoryGroupList(cId, page, size)
   countPage:="0"
   if len(total) > 0{
	   count, _ := strconv.Atoi(total)
	   maxPage := math.Ceil(float64(count/size))
	   countPage = fmt.Sprintf("%.0f",maxPage)
   }
    catName := catRow["category_name"]
	result := make(map[string]interface{})
	result["groupList"] = groupList
	result["countPage"] = countPage
	result["total"] = total
	result["linkTags"] = linkTags
	result["linkTagsLen"] = len(linkTags)
	result["picGeneralize"] = picGeneralize
	result["channelRecommend"] = channelRecommend
	result["subCategoryList"] = subCategoryList
	result["catId"] = catId
	result["pid"] = pid
	result["catName"] = catName
	result["title"] = catName+"编辑图片 - 正版"+catName+"图片素材_"+catName+"图片素材下载 - 视觉中国"
	result["description"] = "视觉中国是全球性优质正版图片、视频等视觉内容平台型互联网上市公司(www.vcg.com)，与Getty Images深度合作，并拥有1300万用户的全球摄影创作社交平台(500px.com)和全球第三大图片公司Corbis 图库版权。为“版权视觉内容”的创作者和使用者提供了以大数据、人工智能和区块链等核心技术为基础的互联网交易平台。需求正版体育编辑图片,体育图片素材,体育图片素材下载,正版体育图片素材,体育编辑图片素材下载尽在视觉中国"
	result["keywords"] = "正版"+catName+"编辑图片,"+catName+"图片素材,"+catName+"图片素材下载,正版"+catName+"图片素材,"+catName+"编辑图片素材下载"
	result["baseUrl"] = config.BASEURL
	DisplayLayOut("category/index.html",result,w)
}









