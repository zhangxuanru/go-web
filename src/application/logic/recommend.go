package logic

import "application/models"

func GetLinkDataByCategoryId(categoryId int)  (result map[int]map[string]string,err error)  {
	return models.GetLinkDataByCategoryId(categoryId)
}

func GetRemmendData(categoryId int,where string) (result map[int]map[string]string,err error) {
	return models.GetRemmendData(categoryId,where)
}


