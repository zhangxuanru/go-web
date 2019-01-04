package logger

import (
	"log"
	"time"
	"fmt"
	"config"
	"os"
)

var AccessLog *log.Logger
var ErrorLog  * log.Logger
var InfoLog   *log.Logger

func init()  {
   year,month,day := time.Now().Year(),time.Now().Month(),time.Now().Day()
   accessLogFile := fmt.Sprintf(config.LOGDIR+"/access-%d-%d-%d.log",year,month,day)
   errorLogFile  := fmt.Sprintf(config.LOGDIR+"/error-%d-%d-%d.log",year,month,day)
   infoLogFile := fmt.Sprintf(config.LOGDIR+"/info-%d-%d-%d.log",year,month,day)
   accf, _ := os.OpenFile(accessLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
   errf, _ := os.OpenFile(errorLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
   info , _ := os.OpenFile(infoLogFile,os.O_CREATE|os.O_APPEND|os.O_WRONLY,os.ModePerm)
   AccessLog = log.New(accf,"access:",log.LstdFlags | log.Lshortfile | log.LUTC)
   ErrorLog = log.New(errf,"error",log.LstdFlags | log.Lshortfile | log.LUTC)
   InfoLog  = log.New(info,"info",log.LstdFlags|log.Lshortfile|log.LUTC)
}
