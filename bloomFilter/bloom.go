package bloomfilter

import (
	"context"
	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	pool *redis.Pool
}

func NewRedisClient(pool *redis.Pool) *RedisClient {
	return &RedisClient{
		pool: pool,
	}
}

func (r *RedisClient) Eval(ctx context.Context, src string,
	keyCount int, keyAndArgs []interface{}) (interface{}, error) {

	args := make([]interface{}, 2+len(keyAndArgs))
	args[0] = src
	args[1] = keyCount

	copy(args[2:], keyAndArgs)
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return -1, err
	}

	defer conn.Close()

	return conn.Do("EVAL", args)
}
