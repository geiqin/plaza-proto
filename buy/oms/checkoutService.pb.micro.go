// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: checkoutService.proto

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

// Api Endpoints for CheckoutService service

func NewCheckoutServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CheckoutService service

type CheckoutService interface {
	//填写订单
	Write(ctx context.Context, in *CheckoutRequest, opts ...client.CallOption) (*CheckoutResponse, error)
	//提交订单
	Submit(ctx context.Context, in *CheckoutRequest, opts ...client.CallOption) (*CheckoutResponse, error)
	//充值下单
	Recharge(ctx context.Context, in *CheckoutRequest, opts ...client.CallOption) (*CheckoutResponse, error)
}

type checkoutService struct {
	c    client.Client
	name string
}

func NewCheckoutService(name string, c client.Client) CheckoutService {
	return &checkoutService{
		c:    c,
		name: name,
	}
}

func (c *checkoutService) Write(ctx context.Context, in *CheckoutRequest, opts ...client.CallOption) (*CheckoutResponse, error) {
	req := c.c.NewRequest(c.name, "CheckoutService.Write", in)
	out := new(CheckoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutService) Submit(ctx context.Context, in *CheckoutRequest, opts ...client.CallOption) (*CheckoutResponse, error) {
	req := c.c.NewRequest(c.name, "CheckoutService.Submit", in)
	out := new(CheckoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutService) Recharge(ctx context.Context, in *CheckoutRequest, opts ...client.CallOption) (*CheckoutResponse, error) {
	req := c.c.NewRequest(c.name, "CheckoutService.Recharge", in)
	out := new(CheckoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CheckoutService service

type CheckoutServiceHandler interface {
	//填写订单
	Write(context.Context, *CheckoutRequest, *CheckoutResponse) error
	//提交订单
	Submit(context.Context, *CheckoutRequest, *CheckoutResponse) error
	//充值下单
	Recharge(context.Context, *CheckoutRequest, *CheckoutResponse) error
}

func RegisterCheckoutServiceHandler(s server.Server, hdlr CheckoutServiceHandler, opts ...server.HandlerOption) error {
	type checkoutService interface {
		Write(ctx context.Context, in *CheckoutRequest, out *CheckoutResponse) error
		Submit(ctx context.Context, in *CheckoutRequest, out *CheckoutResponse) error
		Recharge(ctx context.Context, in *CheckoutRequest, out *CheckoutResponse) error
	}
	type CheckoutService struct {
		checkoutService
	}
	h := &checkoutServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CheckoutService{h}, opts...))
}

type checkoutServiceHandler struct {
	CheckoutServiceHandler
}

func (h *checkoutServiceHandler) Write(ctx context.Context, in *CheckoutRequest, out *CheckoutResponse) error {
	return h.CheckoutServiceHandler.Write(ctx, in, out)
}

func (h *checkoutServiceHandler) Submit(ctx context.Context, in *CheckoutRequest, out *CheckoutResponse) error {
	return h.CheckoutServiceHandler.Submit(ctx, in, out)
}

func (h *checkoutServiceHandler) Recharge(ctx context.Context, in *CheckoutRequest, out *CheckoutResponse) error {
	return h.CheckoutServiceHandler.Recharge(ctx, in, out)
}
