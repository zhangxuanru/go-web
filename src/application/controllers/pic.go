package controllers

import (
	"net/http"
	"strings"
	"application/logic"
	"strconv"
	"html/template"
)

//图片详情页
func PicPageDetail(w http.ResponseWriter, r *http.Request)  {
	picId := strings.Replace(r.URL.Path, "/editorialPic/", "", -1)
	groupId,_ := strconv.Atoi(r.FormValue("groupid"))
	if len(picId) == 0 || groupId == 0{
		Redirect404(w,r)
		return
	}
	row, err2 := logic.GetPicListByGroupId(groupId, "pic_id="+picId, 0, 1,"")
	if err2!=nil || len(row) == 0{
		Redirect404(w,r)
		return
	}
	prevPicId := row[0]["prevPicId"]
	nextPicId := row[0]["nextPicId"]
	if len(prevPicId) <= 1{
		picRow, _ := logic.GetPicListByGroupId(groupId, "pic_id<"+row[0]["pic_id"]+" AND prevPicId=0", 0, 1,"pic_id DESC")
		prevPicId = picRow[0]["pic_id"]
     }
	if len(nextPicId) <= 1{
		picRow, _ := logic.GetPicListByGroupId(groupId, "pic_id>"+row[0]["pic_id"]+" AND nextPicId=0", 0, 1,"pic_id ASC")
		nextPicId = picRow[0]["pic_id"]
	 }
	var caption template.HTML
	rowDetail, _ := logic.GetPicDetailByPicId(picId)
	if len(rowDetail) > 0 {
		caption = template.HTML(rowDetail[0]["caption"])
	}
	if strings.Contains(row[0]["title_old"],"<br/>"){
		split := strings.Split(row[0]["title_old"], "<br/>")
		row[0]["title"] = split[0]
	}

	//group 信息
	groupInfo, _ := logic.GetGoodsDetailById(groupId, "")
	result := make(map[string]interface{})
	result["title"] = row[0]["title"]
	result["description"] =  row[0]["title"]
	result["keywords"] = row[0]["title"]
	result["groupId"] = groupId
	result["prevPicId"] = prevPicId
	result["nextPicId"] = nextPicId
	result["row"]  = row[0]
	result["rowDetail"] = rowDetail[0]
	result["groupInfo"] = groupInfo
	result["caption"] = caption
	DisplayLayOut("pic/index.html",result,w)
}
