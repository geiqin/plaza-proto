// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: paymentChannelService.proto

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

// Api Endpoints for PaymentChannelService service

func NewPaymentChannelServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PaymentChannelService service

type PaymentChannelService interface {
	// 支付通道新增
	Create(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error)
	// 支付通道修改
	Update(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error)
	// 支付通道删除
	Delete(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error)
	// 支付通道获取
	Get(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error)
	// 支付通道查询
	Search(ctx context.Context, in *PaymentChannelRequest, opts ...client.CallOption) (*PaymentChannelResponse, error)
	// 支付通道列表
	List(ctx context.Context, in *PaymentChannelRequest, opts ...client.CallOption) (*PaymentChannelResponse, error)
}

type paymentChannelService struct {
	c    client.Client
	name string
}

func NewPaymentChannelService(name string, c client.Client) PaymentChannelService {
	return &paymentChannelService{
		c:    c,
		name: name,
	}
}

func (c *paymentChannelService) Create(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentChannelService.Create", in)
	out := new(PaymentChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentChannelService) Update(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentChannelService.Update", in)
	out := new(PaymentChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentChannelService) Delete(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentChannelService.Delete", in)
	out := new(PaymentChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentChannelService) Get(ctx context.Context, in *PaymentChannel, opts ...client.CallOption) (*PaymentChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentChannelService.Get", in)
	out := new(PaymentChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentChannelService) Search(ctx context.Context, in *PaymentChannelRequest, opts ...client.CallOption) (*PaymentChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentChannelService.Search", in)
	out := new(PaymentChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentChannelService) List(ctx context.Context, in *PaymentChannelRequest, opts ...client.CallOption) (*PaymentChannelResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentChannelService.List", in)
	out := new(PaymentChannelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaymentChannelService service

type PaymentChannelServiceHandler interface {
	// 支付通道新增
	Create(context.Context, *PaymentChannel, *PaymentChannelResponse) error
	// 支付通道修改
	Update(context.Context, *PaymentChannel, *PaymentChannelResponse) error
	// 支付通道删除
	Delete(context.Context, *PaymentChannel, *PaymentChannelResponse) error
	// 支付通道获取
	Get(context.Context, *PaymentChannel, *PaymentChannelResponse) error
	// 支付通道查询
	Search(context.Context, *PaymentChannelRequest, *PaymentChannelResponse) error
	// 支付通道列表
	List(context.Context, *PaymentChannelRequest, *PaymentChannelResponse) error
}

func RegisterPaymentChannelServiceHandler(s server.Server, hdlr PaymentChannelServiceHandler, opts ...server.HandlerOption) error {
	type paymentChannelService interface {
		Create(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error
		Update(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error
		Delete(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error
		Get(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error
		Search(ctx context.Context, in *PaymentChannelRequest, out *PaymentChannelResponse) error
		List(ctx context.Context, in *PaymentChannelRequest, out *PaymentChannelResponse) error
	}
	type PaymentChannelService struct {
		paymentChannelService
	}
	h := &paymentChannelServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PaymentChannelService{h}, opts...))
}

type paymentChannelServiceHandler struct {
	PaymentChannelServiceHandler
}

func (h *paymentChannelServiceHandler) Create(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error {
	return h.PaymentChannelServiceHandler.Create(ctx, in, out)
}

func (h *paymentChannelServiceHandler) Update(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error {
	return h.PaymentChannelServiceHandler.Update(ctx, in, out)
}

func (h *paymentChannelServiceHandler) Delete(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error {
	return h.PaymentChannelServiceHandler.Delete(ctx, in, out)
}

func (h *paymentChannelServiceHandler) Get(ctx context.Context, in *PaymentChannel, out *PaymentChannelResponse) error {
	return h.PaymentChannelServiceHandler.Get(ctx, in, out)
}

func (h *paymentChannelServiceHandler) Search(ctx context.Context, in *PaymentChannelRequest, out *PaymentChannelResponse) error {
	return h.PaymentChannelServiceHandler.Search(ctx, in, out)
}

func (h *paymentChannelServiceHandler) List(ctx context.Context, in *PaymentChannelRequest, out *PaymentChannelResponse) error {
	return h.PaymentChannelServiceHandler.List(ctx, in, out)
}
