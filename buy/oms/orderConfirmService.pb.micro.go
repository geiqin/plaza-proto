// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderConfirmService.proto

package services

import (
	fmt "fmt"
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

// Api Endpoints for OrderConfirmService service

func NewOrderConfirmServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderConfirmService service

type OrderConfirmService interface {
	//确认首页
	Index(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error)
	//获取确认订单页面是否展示快递配送和到店自提
	CheckShipping(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error)
	//计算订单金额
	ComputedOrder(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error)
	//可选优惠劵列表
	CouponList(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error)
	//订单确认创建
	Create(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error)
}

type orderConfirmService struct {
	c    client.Client
	name string
}

func NewOrderConfirmService(name string, c client.Client) OrderConfirmService {
	return &orderConfirmService{
		c:    c,
		name: name,
	}
}

func (c *orderConfirmService) Index(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "OrderConfirmService.Index", in)
	out := new(OrderConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderConfirmService) CheckShipping(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "OrderConfirmService.CheckShipping", in)
	out := new(OrderConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderConfirmService) ComputedOrder(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "OrderConfirmService.ComputedOrder", in)
	out := new(OrderConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderConfirmService) CouponList(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "OrderConfirmService.CouponList", in)
	out := new(OrderConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderConfirmService) Create(ctx context.Context, in *OrderConfirmRequest, opts ...client.CallOption) (*OrderConfirmResponse, error) {
	req := c.c.NewRequest(c.name, "OrderConfirmService.Create", in)
	out := new(OrderConfirmResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderConfirmService service

type OrderConfirmServiceHandler interface {
	//确认首页
	Index(context.Context, *OrderConfirmRequest, *OrderConfirmResponse) error
	//获取确认订单页面是否展示快递配送和到店自提
	CheckShipping(context.Context, *OrderConfirmRequest, *OrderConfirmResponse) error
	//计算订单金额
	ComputedOrder(context.Context, *OrderConfirmRequest, *OrderConfirmResponse) error
	//可选优惠劵列表
	CouponList(context.Context, *OrderConfirmRequest, *OrderConfirmResponse) error
	//订单确认创建
	Create(context.Context, *OrderConfirmRequest, *OrderConfirmResponse) error
}

func RegisterOrderConfirmServiceHandler(s server.Server, hdlr OrderConfirmServiceHandler, opts ...server.HandlerOption) error {
	type orderConfirmService interface {
		Index(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error
		CheckShipping(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error
		ComputedOrder(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error
		CouponList(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error
		Create(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error
	}
	type OrderConfirmService struct {
		orderConfirmService
	}
	h := &orderConfirmServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderConfirmService{h}, opts...))
}

type orderConfirmServiceHandler struct {
	OrderConfirmServiceHandler
}

func (h *orderConfirmServiceHandler) Index(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error {
	return h.OrderConfirmServiceHandler.Index(ctx, in, out)
}

func (h *orderConfirmServiceHandler) CheckShipping(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error {
	return h.OrderConfirmServiceHandler.CheckShipping(ctx, in, out)
}

func (h *orderConfirmServiceHandler) ComputedOrder(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error {
	return h.OrderConfirmServiceHandler.ComputedOrder(ctx, in, out)
}

func (h *orderConfirmServiceHandler) CouponList(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error {
	return h.OrderConfirmServiceHandler.CouponList(ctx, in, out)
}

func (h *orderConfirmServiceHandler) Create(ctx context.Context, in *OrderConfirmRequest, out *OrderConfirmResponse) error {
	return h.OrderConfirmServiceHandler.Create(ctx, in, out)
}
