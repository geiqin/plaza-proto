// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: paymentService.proto

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

// Api Endpoints for PaymentService service

func NewPaymentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PaymentService service

type PaymentService interface {
	//获取支付方式列表
	GetMethodList(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*PaymentResponse, error)
	//获取支付方式信息
	GetMethodInfo(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*PaymentResponse, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) GetMethodList(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*PaymentResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentService.GetMethodList", in)
	out := new(PaymentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) GetMethodInfo(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*PaymentResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentService.GetMethodInfo", in)
	out := new(PaymentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaymentService service

type PaymentServiceHandler interface {
	//获取支付方式列表
	GetMethodList(context.Context, *PaymentRequest, *PaymentResponse) error
	//获取支付方式信息
	GetMethodInfo(context.Context, *PaymentRequest, *PaymentResponse) error
}

func RegisterPaymentServiceHandler(s server.Server, hdlr PaymentServiceHandler, opts ...server.HandlerOption) error {
	type paymentService interface {
		GetMethodList(ctx context.Context, in *PaymentRequest, out *PaymentResponse) error
		GetMethodInfo(ctx context.Context, in *PaymentRequest, out *PaymentResponse) error
	}
	type PaymentService struct {
		paymentService
	}
	h := &paymentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PaymentService{h}, opts...))
}

type paymentServiceHandler struct {
	PaymentServiceHandler
}

func (h *paymentServiceHandler) GetMethodList(ctx context.Context, in *PaymentRequest, out *PaymentResponse) error {
	return h.PaymentServiceHandler.GetMethodList(ctx, in, out)
}

func (h *paymentServiceHandler) GetMethodInfo(ctx context.Context, in *PaymentRequest, out *PaymentResponse) error {
	return h.PaymentServiceHandler.GetMethodInfo(ctx, in, out)
}
