package redis_mod

import (
	"context"
	"time"
)

func (c *RedisClient) SetCache(key string, value string, duration time.Duration) error {
	ok, err := c.Rds.Set(context.Background(), key, value, duration).Result()
	if ok == OK {
		return nil
	}
	return err
}

//
// ZRangeByScoreAsc
//  @-Description: 从有序集合里按照score正序取数据
//  @-param key
//  @-param start 从0开始
//  @-param pageSize 每页数据
//
func (c *RedisClient) ZRangeByScoreAsc(key string, start, pageSize int64) ([]string, error) {
	end := start + pageSize - 1
	return c.Rds.ZRange(context.Background(), key, start, end).Result()
}

//
// ZRangeByScoreDesc
//  @-Description: 从有序集合里按照score正序取数据
//  @-param key
//  @-param start 从0开始
//  @-param pageSize 每页数据
//  @-return []string
//  @-return error
//
func (c *RedisClient) ZRangeByScoreDesc(key string, start, pageSize int64) ([]string, error) {
	end := start + pageSize - 1
	return c.Rds.ZRevRange(context.Background(), key, start, end).Result()
}

//
// ZCount 获取集合统计
//  @-Description:
//  @-param key
//  @-return int
//  @-return error
//
func (c *RedisClient) ZCount(key string) (int64, error) {
	return c.Rds.ZCard(context.Background(), key).Result()
}
