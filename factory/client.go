package factory

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)


func (f *factory) redisDriver() (*redis.Client, error) {
	var err error
	redisSync.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:        fmt.Sprintf("%s:%d", f.config.RedisConfig.Host, f.config.RedisConfig.Port),
			Username:    f.config.RedisConfig.Username,
			DB:          f.config.RedisConfig.Database,
			DialTimeout: 1 * time.Minute,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		res := rdb.Ping(ctx)
		if res.Err() != nil {
			err = res.Err()
			return
		}

		f.redisConn = rdb
	})

	return f.redisConn, err
}
