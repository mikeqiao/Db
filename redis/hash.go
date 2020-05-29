package redis

import (
	"github.com/gomodule/redigo/redis"
)

func (r *CRedis) Hash_GetAllData(table string) (map[string]string, error) {
	c := r.Pool.Get()
	value, err := redis.StringMap(c.Do("hgetall", table))
	if nil != err {
		//	log.Error("table:%v, error:%v", table, err)
	}

	c.Close()
	return value, err
}
