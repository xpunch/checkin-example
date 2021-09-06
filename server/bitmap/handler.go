package main

import (
	"context"
	"time"

	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"github.com/xpunch/checkin-example/proto"
)

var checkinScript = redis.NewScript(`
	if redis.call("GETBIT", KEYS[1], ARGV[1]) == 1 then
		return redis.call("HGET", KEYS[2], ARGV[1])
	else
		redis.call("SETBIT", KEYS[1], ARGV[1], 1)
		if redis.call("GETBIT", KEYS[3], ARGV[1]) == 1 then
			return redis.call("HINCRBY", KEYS[2], ARGV[1], 1)
		else
			redis.call("HSET", KEYS[2], ARGV[1], 1)
			return 1
		end
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
	today, yesterday := now.Format("checkin.bitmap@060102"), now.Add(-24*time.Hour).Format("checkin.bitmap@060102")
	c, err := checkinScript.Run(ctx, h.redis, []string{today, "checkindays", yesterday}, in.UserId).Int64()
	if err != nil {
		logger.Error(err)
		return err
	}
	out.ContinuousDays = c
	return nil
}
