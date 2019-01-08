package logic

import "application/models"

func GetCateGoryDataById(categoryId int) (r map[string]string, err error) {
	 return  models.GetCateGoryDataById(categoryId)
}
