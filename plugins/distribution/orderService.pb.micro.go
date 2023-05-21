// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderService.proto

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

// Api Endpoints for OrderService service

func NewOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderService service

type OrderService interface {
	// 获得佣金详情
	Detail(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error)
	// 查询佣金列表
	Search(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Detail(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Detail", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Search(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Search", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	// 获得佣金详情
	Detail(context.Context, *Order, *OrderResponse) error
	// 查询佣金列表
	Search(context.Context, *OrderRequest, *OrderResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		Detail(ctx context.Context, in *Order, out *OrderResponse) error
		Search(ctx context.Context, in *OrderRequest, out *OrderResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) Detail(ctx context.Context, in *Order, out *OrderResponse) error {
	return h.OrderServiceHandler.Detail(ctx, in, out)
}

func (h *orderServiceHandler) Search(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.Search(ctx, in, out)
}
