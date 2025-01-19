// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: calculatorService.proto

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

// Api Endpoints for CalculatorService service

func NewCalculatorServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CalculatorService service

type CalculatorService interface {
	//快递运费计算
	ExpressFee(ctx context.Context, in *CalculatorRequest, opts ...client.CallOption) (*CalculatorResponse, error)
	//同城配送费计算
	DeliveryFee(ctx context.Context, in *CalculatorRequest, opts ...client.CallOption) (*CalculatorResponse, error)
}

type calculatorService struct {
	c    client.Client
	name string
}

func NewCalculatorService(name string, c client.Client) CalculatorService {
	return &calculatorService{
		c:    c,
		name: name,
	}
}

func (c *calculatorService) ExpressFee(ctx context.Context, in *CalculatorRequest, opts ...client.CallOption) (*CalculatorResponse, error) {
	req := c.c.NewRequest(c.name, "CalculatorService.ExpressFee", in)
	out := new(CalculatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorService) DeliveryFee(ctx context.Context, in *CalculatorRequest, opts ...client.CallOption) (*CalculatorResponse, error) {
	req := c.c.NewRequest(c.name, "CalculatorService.DeliveryFee", in)
	out := new(CalculatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalculatorService service

type CalculatorServiceHandler interface {
	//快递运费计算
	ExpressFee(context.Context, *CalculatorRequest, *CalculatorResponse) error
	//同城配送费计算
	DeliveryFee(context.Context, *CalculatorRequest, *CalculatorResponse) error
}

func RegisterCalculatorServiceHandler(s server.Server, hdlr CalculatorServiceHandler, opts ...server.HandlerOption) error {
	type calculatorService interface {
		ExpressFee(ctx context.Context, in *CalculatorRequest, out *CalculatorResponse) error
		DeliveryFee(ctx context.Context, in *CalculatorRequest, out *CalculatorResponse) error
	}
	type CalculatorService struct {
		calculatorService
	}
	h := &calculatorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CalculatorService{h}, opts...))
}

type calculatorServiceHandler struct {
	CalculatorServiceHandler
}

func (h *calculatorServiceHandler) ExpressFee(ctx context.Context, in *CalculatorRequest, out *CalculatorResponse) error {
	return h.CalculatorServiceHandler.ExpressFee(ctx, in, out)
}

func (h *calculatorServiceHandler) DeliveryFee(ctx context.Context, in *CalculatorRequest, out *CalculatorResponse) error {
	return h.CalculatorServiceHandler.DeliveryFee(ctx, in, out)
}