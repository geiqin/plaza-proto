// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: wxService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for WxService service

func NewWxServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WxService service

type WxService interface {
	//小程序支付
	MiniPay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error)
	//H5支付
	H5Pay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error)
	//APP支付
	AppPay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error)
	//提交付款码支付
	MicroPay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error)
}

type wxService struct {
	c    client.Client
	name string
}

func NewWxService(name string, c client.Client) WxService {
	return &wxService{
		c:    c,
		name: name,
	}
}

func (c *wxService) MiniPay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error) {
	req := c.c.NewRequest(c.name, "WxService.MiniPay", in)
	out := new(WxResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxService) H5Pay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error) {
	req := c.c.NewRequest(c.name, "WxService.H5Pay", in)
	out := new(WxResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxService) AppPay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error) {
	req := c.c.NewRequest(c.name, "WxService.AppPay", in)
	out := new(WxResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wxService) MicroPay(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*WxResponse, error) {
	req := c.c.NewRequest(c.name, "WxService.MicroPay", in)
	out := new(WxResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WxService service

type WxServiceHandler interface {
	//小程序支付
	MiniPay(context.Context, *PaymentRequest, *WxResponse) error
	//H5支付
	H5Pay(context.Context, *PaymentRequest, *WxResponse) error
	//APP支付
	AppPay(context.Context, *PaymentRequest, *WxResponse) error
	//提交付款码支付
	MicroPay(context.Context, *PaymentRequest, *WxResponse) error
}

func RegisterWxServiceHandler(s server.Server, hdlr WxServiceHandler, opts ...server.HandlerOption) error {
	type wxService interface {
		MiniPay(ctx context.Context, in *PaymentRequest, out *WxResponse) error
		H5Pay(ctx context.Context, in *PaymentRequest, out *WxResponse) error
		AppPay(ctx context.Context, in *PaymentRequest, out *WxResponse) error
		MicroPay(ctx context.Context, in *PaymentRequest, out *WxResponse) error
	}
	type WxService struct {
		wxService
	}
	h := &wxServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WxService{h}, opts...))
}

type wxServiceHandler struct {
	WxServiceHandler
}

func (h *wxServiceHandler) MiniPay(ctx context.Context, in *PaymentRequest, out *WxResponse) error {
	return h.WxServiceHandler.MiniPay(ctx, in, out)
}

func (h *wxServiceHandler) H5Pay(ctx context.Context, in *PaymentRequest, out *WxResponse) error {
	return h.WxServiceHandler.H5Pay(ctx, in, out)
}

func (h *wxServiceHandler) AppPay(ctx context.Context, in *PaymentRequest, out *WxResponse) error {
	return h.WxServiceHandler.AppPay(ctx, in, out)
}

func (h *wxServiceHandler) MicroPay(ctx context.Context, in *PaymentRequest, out *WxResponse) error {
	return h.WxServiceHandler.MicroPay(ctx, in, out)
}
