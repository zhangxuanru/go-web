package util

import (
	"strings"
	"strconv"
	"time"
)

func SecurityString(str string)(string)  {
	replacer := strings.NewReplacer("\"","","<br/>","","<br>","","\r", "", "\n", "", " ", "")
	return  replacer.Replace(str)
}

func FormattingTimeRubbing(timeRub string) (string) {
	date, err := strconv.ParseInt(timeRub, 10, 64)
	if err != nil{
		return ""
	}
	return  time.Unix(date, 0).Format("2006-01-02 15:04:05")
}





