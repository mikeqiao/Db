package redis

func (r *CRedis) DelTable(table string) {
	c := r.Pool.Get()
	_, err := c.Do("del", table)
	if nil != err {
		//	log.Error("table:%v, error:%v", table, err)
	}
	c.Close()
}
