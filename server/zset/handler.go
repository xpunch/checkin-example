package main

import (
	"context"
	"time"

	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"github.com/xpunch/checkin-example/proto"
)

var checkinScript = redis.NewScript(`
	local val = redis.call("ZSCORE", KEYS[1], ARGV[1])
	if not val then
		local v2 = redis.call("ZSCORE", KEYS[2], ARGV[1])
		if not v2 then
			redis.call("ZADD", KEYS[1], 1, ARGV[1])
			return 1
		else
			redis.call("ZADD", KEYS[1], v2+1, ARGV[1])
			return v2+1
		end
	else
		return val
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
	today, yesterday := now.Format("checkin.zset@060102"), now.Add(-24*time.Hour).Format("checkin.zset@060102")
	c, err := checkinScript.Run(ctx, h.redis, []string{today, yesterday}, in.UserId).Int64()
	if err != nil {
		logger.Error(err)
		return err
	}
	out.ContinuousDays = c
	return nil
}
