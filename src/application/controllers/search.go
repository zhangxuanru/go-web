package controllers

import (
	"net/http"
	"application/logic"
)

var (
	assign  map[string]interface{}
	search  logic.Search
	result  map[int]map[string]interface{}
	totalHit int64
	KeyWord string
)

func initSearch() (logic.Search)  {
	return  logic.Search{
		Keyword:KeyWord,
		TopicId:101301,
		Start:0,
		Size:10,
	}
}


func TopicSearch(w http.ResponseWriter, r *http.Request){

}


func Search(w http.ResponseWriter, r *http.Request) {

}



