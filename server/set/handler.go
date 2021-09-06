package main

import (
	"context"
	"time"

	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"github.com/xpunch/checkin-example/proto"
)

var checkinScript = redis.NewScript(`
	if redis.call("SISMEMBER", KEYS[1], ARGV[1]) == 0 then
		if redis.call("SISMEMBER", KEYS[2], ARGV[1]) == 0 then
			redis.call("ZSET", KEYS[3], ARGV[1], 1)
			return 1
		else
			return redis.call("ZINCRBY", KEYS[3], ARGV[1])
		end
	else
		return redis.call("HGET", KEYS[3], ARGV[1])
	end
`)

func NewCheckInServiceHandler(redis *redis.Client) proto.CheckInServiceHandler {
	return &handler{redis: redis}
}

type handler struct {
	redis *redis.Client
}

func (h *handler) CheckIn(ctx context.Context, in *proto.CheckInRequest, out *proto.CheckInReply) error {
	now := time.Now()
	today, yesterday := now.Format("checkin.set@060102"), now.Add(-24*time.Hour).Format("checkin.set@060102")
	c, err := checkinScript.Run(ctx, h.redis, []string{today, yesterday}, in.UserId).Int64()
	if err != nil {
		logger.Error(err)
		return err
	}
	out.ContinuousDays = c
	return nil
}
