package redis

import (
	"github.com/garyburd/redigo/redis"
	"libary/logger"
	"config"
)

var Conn redis.Conn

var err error

func init()  {
	Conn, err = redis.Dial("tcp", config.REDIS_HOST)
	if err != nil{
		logger.ErrorLog.Println("redis connect error",err)
	}
}

