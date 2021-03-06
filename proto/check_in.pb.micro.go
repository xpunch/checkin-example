// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/check_in.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CheckInService service

func NewCheckInServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CheckInService service

type CheckInService interface {
	CheckIn(ctx context.Context, in *CheckInRequest, opts ...client.CallOption) (*CheckInReply, error)
}

type checkInService struct {
	c    client.Client
	name string
}

func NewCheckInService(name string, c client.Client) CheckInService {
	return &checkInService{
		c:    c,
		name: name,
	}
}

func (c *checkInService) CheckIn(ctx context.Context, in *CheckInRequest, opts ...client.CallOption) (*CheckInReply, error) {
	req := c.c.NewRequest(c.name, "CheckInService.CheckIn", in)
	out := new(CheckInReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CheckInService service

type CheckInServiceHandler interface {
	CheckIn(context.Context, *CheckInRequest, *CheckInReply) error
}

func RegisterCheckInServiceHandler(s server.Server, hdlr CheckInServiceHandler, opts ...server.HandlerOption) error {
	type checkInService interface {
		CheckIn(ctx context.Context, in *CheckInRequest, out *CheckInReply) error
	}
	type CheckInService struct {
		checkInService
	}
	h := &checkInServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CheckInService{h}, opts...))
}

type checkInServiceHandler struct {
	CheckInServiceHandler
}

func (h *checkInServiceHandler) CheckIn(ctx context.Context, in *CheckInRequest, out *CheckInReply) error {
	return h.CheckInServiceHandler.CheckIn(ctx, in, out)
}
