package main

import (
	"context"
	"flag"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/xpunch/checkin-example/proto"
)

func main() {
	uid := flag.Uint64("uid", 0, "user id")
	flag.Parse()
	srv := micro.NewService(micro.Name("checkin.cli"))
	checkInService := proto.NewCheckInService("checkin.srv", srv.Client())
	reply, err := checkInService.CheckIn(context.TODO(), &proto.CheckInRequest{UserId: *uid})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(reply)
}
