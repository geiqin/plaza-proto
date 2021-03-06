// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderCouponService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/micro-kit/protobuf/common"
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

// Api Endpoints for OrderCouponService service

func NewOrderCouponServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderCouponService service

type OrderCouponService interface {
	Create(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error)
	Update(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error)
	Delete(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error)
	Get(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*OrderCouponResponse, error)
}

type orderCouponService struct {
	c    client.Client
	name string
}

func NewOrderCouponService(name string, c client.Client) OrderCouponService {
	return &orderCouponService{
		c:    c,
		name: name,
	}
}

func (c *orderCouponService) Create(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error) {
	req := c.c.NewRequest(c.name, "OrderCouponService.Create", in)
	out := new(OrderCouponResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCouponService) Update(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error) {
	req := c.c.NewRequest(c.name, "OrderCouponService.Update", in)
	out := new(OrderCouponResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCouponService) Delete(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error) {
	req := c.c.NewRequest(c.name, "OrderCouponService.Delete", in)
	out := new(OrderCouponResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCouponService) Get(ctx context.Context, in *OrderCoupon, opts ...client.CallOption) (*OrderCouponResponse, error) {
	req := c.c.NewRequest(c.name, "OrderCouponService.Get", in)
	out := new(OrderCouponResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCouponService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*OrderCouponResponse, error) {
	req := c.c.NewRequest(c.name, "OrderCouponService.Search", in)
	out := new(OrderCouponResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderCouponService service

type OrderCouponServiceHandler interface {
	Create(context.Context, *OrderCoupon, *OrderCouponResponse) error
	Update(context.Context, *OrderCoupon, *OrderCouponResponse) error
	Delete(context.Context, *OrderCoupon, *OrderCouponResponse) error
	Get(context.Context, *OrderCoupon, *OrderCouponResponse) error
	Search(context.Context, *common.BaseWhere, *OrderCouponResponse) error
}

func RegisterOrderCouponServiceHandler(s server.Server, hdlr OrderCouponServiceHandler, opts ...server.HandlerOption) error {
	type orderCouponService interface {
		Create(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error
		Update(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error
		Delete(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error
		Get(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *OrderCouponResponse) error
	}
	type OrderCouponService struct {
		orderCouponService
	}
	h := &orderCouponServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderCouponService{h}, opts...))
}

type orderCouponServiceHandler struct {
	OrderCouponServiceHandler
}

func (h *orderCouponServiceHandler) Create(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error {
	return h.OrderCouponServiceHandler.Create(ctx, in, out)
}

func (h *orderCouponServiceHandler) Update(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error {
	return h.OrderCouponServiceHandler.Update(ctx, in, out)
}

func (h *orderCouponServiceHandler) Delete(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error {
	return h.OrderCouponServiceHandler.Delete(ctx, in, out)
}

func (h *orderCouponServiceHandler) Get(ctx context.Context, in *OrderCoupon, out *OrderCouponResponse) error {
	return h.OrderCouponServiceHandler.Get(ctx, in, out)
}

func (h *orderCouponServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *OrderCouponResponse) error {
	return h.OrderCouponServiceHandler.Search(ctx, in, out)
}
