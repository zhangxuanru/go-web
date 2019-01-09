package logic

import "application/models"

func GetLinkDataByCategoryId(categoryId int)  (result map[int]map[string]string,err error)  {
	return models.GetLinkDataByCategoryId(categoryId)
}

func GetRecommendData(categoryId int,where string, start int,limit int) (result map[int]map[string]string,err error) {
	return models.GetRecommendData(categoryId,where,start,limit)
}


