package controllers

import (
	"net/http"
	"strings"
	"strconv"
	"fmt"
	"application/logic"
	"libary/logger"
	"config"
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
	if err!=nil{
		logger.ErrorLog.Println(r.URL.Path,"-catId:",catId,"-error:",err)
		Redirect404(w,r)
		return
	}
	catRow, _ := logic.GetCateGoryDataById(cId)
	catName := ""
	if len(catRow) > 0{
		catName = catRow["category_name"]
	}
	result := make(map[string]interface{})
	result["linkTags"] = linkTags
	result["picGeneralize"] = picGeneralize
	result["channelRecommend"] = channelRecommend
	result["catId"] = catId
	result["title"] = catName+"编辑图片 - 正版"+catName+"图片素材_"+catName+"图片素材下载 - 视觉中国"
	result["description"] = "视觉中国是全球性优质正版图片、视频等视觉内容平台型互联网上市公司(www.vcg.com)，与Getty Images深度合作，并拥有1300万用户的全球摄影创作社交平台(500px.com)和全球第三大图片公司Corbis 图库版权。为“版权视觉内容”的创作者和使用者提供了以大数据、人工智能和区块链等核心技术为基础的互联网交易平台。需求正版体育编辑图片,体育图片素材,体育图片素材下载,正版体育图片素材,体育编辑图片素材下载尽在视觉中国"
	result["keywords"] = "正版"+catName+"编辑图片,"+catName+"图片素材,"+catName+"图片素材下载,正版"+catName+"图片素材,"+catName+"编辑图片素材下载"
	result["baseUrl"] = config.BASEURL

	fmt.Printf("%+v",result)

	DisplayLayOut("category/index.html",result,w)
}









