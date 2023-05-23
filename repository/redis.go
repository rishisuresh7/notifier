package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisQueryer interface {
	GetString(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, timeOut time.Duration) error
	GetBytes(ctx context.Context, key string) ([]byte, error)
	GetDelString(ctx context.Context, key string) (string, error)
	Subscribe(ctx context.Context, channels ...string) pubSub
	PublishNotification(ctx context.Context, notification interface{}) error
}

type redisQueryer struct {
	client *redis.Client
}

func NewRedisQueryer(d *redis.Client) RedisQueryer {
	return &redisQueryer{
		client: d,
	}
}

func (r *redisQueryer) Set(ctx context.Context, key string, value interface{}, timeOut time.Duration) error {
	res := r.client.Set(ctx, key, value, timeOut)
	if err := res.Err(); err != nil {
		return fmt.Errorf("set: unable to set key in redis: %s", err)
	}

	return nil
}

func (r *redisQueryer) GetString(ctx context.Context, key string) (string, error) {
	res := r.client.Get(ctx, key)
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("getString: unable to get value from redis: %s", err)
	}

	return res.Val(), nil
}

func (r *redisQueryer) GetBytes(ctx context.Context, key string) ([]byte, error) {
	res := r.client.Get(ctx, key)
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("getBytes: unable to get value from redis: %s", err)
	}

	bytes, err := res.Bytes()
	if err != nil {
		return nil, fmt.Errorf("getBytes: unable to get bytes: %s", err)
	}

	return bytes, nil
}

func (r *redisQueryer) GetDelString(ctx context.Context, key string) (string, error) {
	res := r.client.GetDel(ctx, key)
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("getDelString: unable to get value from redis: %s", err)
	}

	return res.Val(), nil
}

func (r *redisQueryer) Subscribe(ctx context.Context, channels ...string) pubSub {
	return newPS(r.client.Subscribe(ctx, channels...))
}

func (r *redisQueryer) PublishNotification(ctx context.Context, notification interface{}) error {
	// notificationBytes, _ := json.Marshal(notification)
	// res := r.client.Publish(ctx, notificationChannel, notificationBytes)
	// if err := res.Err(); err != nil {
	// 	return fmt.Errorf("pushNotification: unable to push notification to redis: %s", err)
	// }

	return nil
}

type pubSub interface {
	Channel(channelSize int) <- chan *redis.Message
}

type ps struct {
	pubSub *redis.PubSub
}

func newPS(p *redis.PubSub) pubSub {
	return &ps{
		pubSub: p,
	}
}

func (p *ps) Channel(channelSize int) <- chan *redis.Message {
	return p.pubSub.Channel(redis.WithChannelSize(channelSize))
}