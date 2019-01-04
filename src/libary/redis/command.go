package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func IsExists(id interface{},Type string) (bool) {
	reply, _ := Conn.Do("GETBIT", Type, id)
	i,_ := reply.(int64)
	if i == 1 {
		return true
	}
	return false
}

func SaveBit(id interface{},Type string)  {
	reply, e := Conn.Do("SETBIT", Type, id, 1)
	if e != nil{
		fmt.Println("error:",e)
	}
	fmt.Println(reply)
}

func Set(key string,val string,exType string,exTime int) (bool) {
	var e error
    if len(exType) > 0{
		_, e = Conn.Do("SET", key, val, exType, exTime)
	}else{
		_, e = Conn.Do("SET", key, val)
	}
	if e !=nil{
		return false
	}
	return true
}

func Get(key string) (val string) {
	s, e := redis.String(Conn.Do("GET", key))
	if e != nil{
		return ""
	}
	return s
}


