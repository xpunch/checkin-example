package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"github.com/xpunch/checkin-example/proto"
)

func main() {
	rc := redis.NewClient(&redis.Options{})
	srv := micro.NewService(micro.Name("checkin.srv"))
	hdl := NewCheckInServiceHandler(rc)
	if err := proto.RegisterCheckInServiceHandler(srv.Server(), hdl); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
