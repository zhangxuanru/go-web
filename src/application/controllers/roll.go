package controllers

import "net/http"

func Roll(w http.ResponseWriter, r *http.Request)  {
	result := make(map[string]interface{})
	DisplayLayOut("roll/index.html",result,w)
}


