// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: calculatorService.proto

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

// Api Endpoints for CalculatorService service

func NewCalculatorServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CalculatorService service

type CalculatorService interface {
	//商品显示级核算
	Display(ctx context.Context, in *DisplayRequest, opts ...client.CallOption) (*CalculatorResponse, error)
	//下单购买级核算
	Purchase(ctx context.Context, in *PurchaseRequest, opts ...client.CallOption) (*CalculatorResponse, error)
	//支付完成级核算
	Payment(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*CalculatorResponse, error)
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

func (c *calculatorService) Display(ctx context.Context, in *DisplayRequest, opts ...client.CallOption) (*CalculatorResponse, error) {
	req := c.c.NewRequest(c.name, "CalculatorService.Display", in)
	out := new(CalculatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorService) Purchase(ctx context.Context, in *PurchaseRequest, opts ...client.CallOption) (*CalculatorResponse, error) {
	req := c.c.NewRequest(c.name, "CalculatorService.Purchase", in)
	out := new(CalculatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorService) Payment(ctx context.Context, in *PaymentRequest, opts ...client.CallOption) (*CalculatorResponse, error) {
	req := c.c.NewRequest(c.name, "CalculatorService.Payment", in)
	out := new(CalculatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalculatorService service

type CalculatorServiceHandler interface {
	//商品显示级核算
	Display(context.Context, *DisplayRequest, *CalculatorResponse) error
	//下单购买级核算
	Purchase(context.Context, *PurchaseRequest, *CalculatorResponse) error
	//支付完成级核算
	Payment(context.Context, *PaymentRequest, *CalculatorResponse) error
}

func RegisterCalculatorServiceHandler(s server.Server, hdlr CalculatorServiceHandler, opts ...server.HandlerOption) error {
	type calculatorService interface {
		Display(ctx context.Context, in *DisplayRequest, out *CalculatorResponse) error
		Purchase(ctx context.Context, in *PurchaseRequest, out *CalculatorResponse) error
		Payment(ctx context.Context, in *PaymentRequest, out *CalculatorResponse) error
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

func (h *calculatorServiceHandler) Display(ctx context.Context, in *DisplayRequest, out *CalculatorResponse) error {
	return h.CalculatorServiceHandler.Display(ctx, in, out)
}

func (h *calculatorServiceHandler) Purchase(ctx context.Context, in *PurchaseRequest, out *CalculatorResponse) error {
	return h.CalculatorServiceHandler.Purchase(ctx, in, out)
}

func (h *calculatorServiceHandler) Payment(ctx context.Context, in *PaymentRequest, out *CalculatorResponse) error {
	return h.CalculatorServiceHandler.Payment(ctx, in, out)
}