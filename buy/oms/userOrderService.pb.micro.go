// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: userOrderService.proto

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

// Api Endpoints for UserOrderService service

func NewUserOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserOrderService service

type UserOrderService interface {
	//订单总数
	Total(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error)
	//订单查询
	Search(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error)
	//订单详情
	Detail(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error)
	//订单取消
	Cancel(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error)
	//订单删除
	Delete(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error)
}

type userOrderService struct {
	c    client.Client
	name string
}

func NewUserOrderService(name string, c client.Client) UserOrderService {
	return &userOrderService{
		c:    c,
		name: name,
	}
}

func (c *userOrderService) Total(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error) {
	req := c.c.NewRequest(c.name, "UserOrderService.Total", in)
	out := new(UserOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) Search(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error) {
	req := c.c.NewRequest(c.name, "UserOrderService.Search", in)
	out := new(UserOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) Detail(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error) {
	req := c.c.NewRequest(c.name, "UserOrderService.Detail", in)
	out := new(UserOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) Cancel(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error) {
	req := c.c.NewRequest(c.name, "UserOrderService.Cancel", in)
	out := new(UserOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) Delete(ctx context.Context, in *UserOrderRequest, opts ...client.CallOption) (*UserOrderResponse, error) {
	req := c.c.NewRequest(c.name, "UserOrderService.Delete", in)
	out := new(UserOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserOrderService service

type UserOrderServiceHandler interface {
	//订单总数
	Total(context.Context, *UserOrderRequest, *UserOrderResponse) error
	//订单查询
	Search(context.Context, *UserOrderRequest, *UserOrderResponse) error
	//订单详情
	Detail(context.Context, *UserOrderRequest, *UserOrderResponse) error
	//订单取消
	Cancel(context.Context, *UserOrderRequest, *UserOrderResponse) error
	//订单删除
	Delete(context.Context, *UserOrderRequest, *UserOrderResponse) error
}

func RegisterUserOrderServiceHandler(s server.Server, hdlr UserOrderServiceHandler, opts ...server.HandlerOption) error {
	type userOrderService interface {
		Total(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error
		Search(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error
		Detail(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error
		Cancel(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error
		Delete(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error
	}
	type UserOrderService struct {
		userOrderService
	}
	h := &userOrderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserOrderService{h}, opts...))
}

type userOrderServiceHandler struct {
	UserOrderServiceHandler
}

func (h *userOrderServiceHandler) Total(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error {
	return h.UserOrderServiceHandler.Total(ctx, in, out)
}

func (h *userOrderServiceHandler) Search(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error {
	return h.UserOrderServiceHandler.Search(ctx, in, out)
}

func (h *userOrderServiceHandler) Detail(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error {
	return h.UserOrderServiceHandler.Detail(ctx, in, out)
}

func (h *userOrderServiceHandler) Cancel(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error {
	return h.UserOrderServiceHandler.Cancel(ctx, in, out)
}

func (h *userOrderServiceHandler) Delete(ctx context.Context, in *UserOrderRequest, out *UserOrderResponse) error {
	return h.UserOrderServiceHandler.Delete(ctx, in, out)
}
