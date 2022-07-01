package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"go.uber.org/zap"
)

type RedisClient struct {
	Client       *redis.Client
	influxWriter api.WriteAPIBlocking
	logger       *zap.Logger
}

func NewRedisClient(logger *zap.Logger, opts *redis.Options, influxClient influxdb2.Client, org string, bucket string) *RedisClient {
	w := influxClient.WriteAPIBlocking(org, bucket)
	return &RedisClient{
		Client:       redis.NewClient(opts),
		influxWriter: w,
		logger:       logger,
	}
}

func (c *RedisClient) Ping(ctx context.Context) error {
	c.logger.Debug("pinging redis")
	_, err := c.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	p := influxdb2.NewPointWithMeasurement("ping").
		AddTag("service", "redis").
		AddField("ping", 1).
		SetTime(time.Now())
	c.influxWriter.WritePoint(ctx, p)
	return nil
}

func (c *RedisClient) Set(ctx context.Context, key string, value string) error {
	c.logger.Debug("redis set", zap.String("key", key), zap.String("value", value))
	_, err := c.Client.Set(ctx, key, value, 0).Result()
	if err != nil {
		return err
	}

	p := influxdb2.NewPointWithMeasurement("set").
		AddTag("service", "redis").
		AddField("set", 1).
		SetTime(time.Now())
	c.influxWriter.WritePoint(ctx, p)
	return nil
}

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	c.logger.Debug("redis get", zap.String("key", key))
	value, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	p := influxdb2.NewPointWithMeasurement("get").
		AddTag("service", "redis").
		AddField("get", 1).
		SetTime(time.Now())
	c.influxWriter.WritePoint(ctx, p)
	return value, nil
}

func (c *RedisClient) Del(ctx context.Context, key string) error {
	c.logger.Debug("redis del", zap.String("key", key))
	_, err := c.Client.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	p := influxdb2.NewPointWithMeasurement("del").
		AddTag("service", "redis").
		AddField("del", 1).
		SetTime(time.Now())
	c.influxWriter.WritePoint(ctx, p)
	return nil
}

func (c *RedisClient) RPop(ctx context.Context, key string) (string, error) {
	c.logger.Debug("redis rpop", zap.String("key", key))
	value, err := c.Client.RPop(ctx, key).Result()
	if err != nil {
		return "", err
	}

	p := influxdb2.NewPointWithMeasurement("rpop").
		AddTag("service", "redis").
		AddField("rpop", 1).
		SetTime(time.Now())
	c.influxWriter.WritePoint(ctx, p)
	return value, nil
}

func (c *RedisClient) LPush(ctx context.Context, key string, value string) error {
	c.logger.Debug("redis lpush", zap.String("key", key), zap.String("value", value))
	_, err := c.Client.LPush(ctx, key, value).Result()
	if err != nil {
		return err
	}

	p := influxdb2.NewPointWithMeasurement("lpush").
		AddTag("service", "redis").
		AddField("lpush", 1).
		SetTime(time.Now())
	c.influxWriter.WritePoint(ctx, p)
	return nil
}

func (c *RedisClient) LLen(ctx context.Context, key string) (int64, error) {
	c.logger.Debug("redis llen", zap.String("key", key))
	value, err := c.Client.LLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	p := influxdb2.NewPointWithMeasurement("llen").
		AddTag("service", "redis").
		AddField("llen", 1).
		SetTime(time.Now())
	c.influxWriter.WritePoint(ctx, p)
	return value, nil
}
