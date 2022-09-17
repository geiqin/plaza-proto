// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: payChannelService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/micro-kit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PayChannelService service

func NewPayChannelServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PayChannelService service

type PayChannelService interface {
	Bind(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error)
	Cancel(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error)
	SetStatus(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error)
	Get(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error)
	List(ctx context.Context, in *PayChannelRequest, opts ...client.CallOption) (*PayChannelResponse, error)
	Search(ctx context.Context, in *PayChannelRequest, opts ...client.CallOption) (*PayChannelResponse, error)
}

type payChannelService struct {
	c    client.Client
	name string
}

func NewPayChannelService(name string, c client.Client) PayChannelService {
	return &payChannelService{
		c:    c,
		name: name,
	}
}

func (c *payChannelService) Bind(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PayChannelService.Bind", in)
	out := new(PayChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelService) Cancel(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PayChannelService.Cancel", in)
	out := new(PayChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelService) SetStatus(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PayChannelService.SetStatus", in)
	out := new(PayChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelService) Get(ctx context.Context, in *PayChannel, opts ...client.CallOption) (*PayChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PayChannelService.Get", in)
	out := new(PayChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelService) List(ctx context.Context, in *PayChannelRequest, opts ...client.CallOption) (*PayChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PayChannelService.List", in)
	out := new(PayChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payChannelService) Search(ctx context.Context, in *PayChannelRequest, opts ...client.CallOption) (*PayChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PayChannelService.Search", in)
	out := new(PayChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PayChannelService service

type PayChannelServiceHandler interface {
	Bind(context.Context, *PayChannel, *PayChannelResponse) error
	Cancel(context.Context, *PayChannel, *PayChannelResponse) error
	SetStatus(context.Context, *PayChannel, *PayChannelResponse) error
	Get(context.Context, *PayChannel, *PayChannelResponse) error
	List(context.Context, *PayChannelRequest, *PayChannelResponse) error
	Search(context.Context, *PayChannelRequest, *PayChannelResponse) error
}

func RegisterPayChannelServiceHandler(s server.Server, hdlr PayChannelServiceHandler, opts ...server.HandlerOption) error {
	type payChannelService interface {
		Bind(ctx context.Context, in *PayChannel, out *PayChannelResponse) error
		Cancel(ctx context.Context, in *PayChannel, out *PayChannelResponse) error
		SetStatus(ctx context.Context, in *PayChannel, out *PayChannelResponse) error
		Get(ctx context.Context, in *PayChannel, out *PayChannelResponse) error
		List(ctx context.Context, in *PayChannelRequest, out *PayChannelResponse) error
		Search(ctx context.Context, in *PayChannelRequest, out *PayChannelResponse) error
	}
	type PayChannelService struct {
		payChannelService
	}
	h := &payChannelServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PayChannelService{h}, opts...))
}

type payChannelServiceHandler struct {
	PayChannelServiceHandler
}

func (h *payChannelServiceHandler) Bind(ctx context.Context, in *PayChannel, out *PayChannelResponse) error {
	return h.PayChannelServiceHandler.Bind(ctx, in, out)
}

func (h *payChannelServiceHandler) Cancel(ctx context.Context, in *PayChannel, out *PayChannelResponse) error {
	return h.PayChannelServiceHandler.Cancel(ctx, in, out)
}

func (h *payChannelServiceHandler) SetStatus(ctx context.Context, in *PayChannel, out *PayChannelResponse) error {
	return h.PayChannelServiceHandler.SetStatus(ctx, in, out)
}

func (h *payChannelServiceHandler) Get(ctx context.Context, in *PayChannel, out *PayChannelResponse) error {
	return h.PayChannelServiceHandler.Get(ctx, in, out)
}

func (h *payChannelServiceHandler) List(ctx context.Context, in *PayChannelRequest, out *PayChannelResponse) error {
	return h.PayChannelServiceHandler.List(ctx, in, out)
}

func (h *payChannelServiceHandler) Search(ctx context.Context, in *PayChannelRequest, out *PayChannelResponse) error {
	return h.PayChannelServiceHandler.Search(ctx, in, out)
}
