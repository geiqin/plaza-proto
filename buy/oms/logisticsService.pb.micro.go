// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: logisticsService.proto

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

// Api Endpoints for LogisticsService service

func NewLogisticsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LogisticsService service

type LogisticsService interface {
	// 订单发货(快递发货)
	SendExpress(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error)
	// 订单发货(同城配送)
	SendDelivery(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error)
	// 订单发货(上门自提)
	SendFetch(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error)
	//订单补发货
	SendRepair(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error)
	//订单换货发货
	SendExchange(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error)
	// 验证提货码(上门自提)
	VerifyFetchCode(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error)
}

type logisticsService struct {
	c    client.Client
	name string
}

func NewLogisticsService(name string, c client.Client) LogisticsService {
	return &logisticsService{
		c:    c,
		name: name,
	}
}

func (c *logisticsService) SendExpress(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error) {
	req := c.c.NewRequest(c.name, "LogisticsService.SendExpress", in)
	out := new(LogisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) SendDelivery(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error) {
	req := c.c.NewRequest(c.name, "LogisticsService.SendDelivery", in)
	out := new(LogisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) SendFetch(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error) {
	req := c.c.NewRequest(c.name, "LogisticsService.SendFetch", in)
	out := new(LogisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) SendRepair(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error) {
	req := c.c.NewRequest(c.name, "LogisticsService.SendRepair", in)
	out := new(LogisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) SendExchange(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error) {
	req := c.c.NewRequest(c.name, "LogisticsService.SendExchange", in)
	out := new(LogisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) VerifyFetchCode(ctx context.Context, in *LogisticsRequest, opts ...client.CallOption) (*LogisticsResponse, error) {
	req := c.c.NewRequest(c.name, "LogisticsService.VerifyFetchCode", in)
	out := new(LogisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogisticsService service

type LogisticsServiceHandler interface {
	// 订单发货(快递发货)
	SendExpress(context.Context, *LogisticsRequest, *LogisticsResponse) error
	// 订单发货(同城配送)
	SendDelivery(context.Context, *LogisticsRequest, *LogisticsResponse) error
	// 订单发货(上门自提)
	SendFetch(context.Context, *LogisticsRequest, *LogisticsResponse) error
	//订单补发货
	SendRepair(context.Context, *LogisticsRequest, *LogisticsResponse) error
	//订单换货发货
	SendExchange(context.Context, *LogisticsRequest, *LogisticsResponse) error
	// 验证提货码(上门自提)
	VerifyFetchCode(context.Context, *LogisticsRequest, *LogisticsResponse) error
}

func RegisterLogisticsServiceHandler(s server.Server, hdlr LogisticsServiceHandler, opts ...server.HandlerOption) error {
	type logisticsService interface {
		SendExpress(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error
		SendDelivery(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error
		SendFetch(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error
		SendRepair(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error
		SendExchange(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error
		VerifyFetchCode(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error
	}
	type LogisticsService struct {
		logisticsService
	}
	h := &logisticsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LogisticsService{h}, opts...))
}

type logisticsServiceHandler struct {
	LogisticsServiceHandler
}

func (h *logisticsServiceHandler) SendExpress(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error {
	return h.LogisticsServiceHandler.SendExpress(ctx, in, out)
}

func (h *logisticsServiceHandler) SendDelivery(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error {
	return h.LogisticsServiceHandler.SendDelivery(ctx, in, out)
}

func (h *logisticsServiceHandler) SendFetch(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error {
	return h.LogisticsServiceHandler.SendFetch(ctx, in, out)
}

func (h *logisticsServiceHandler) SendRepair(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error {
	return h.LogisticsServiceHandler.SendRepair(ctx, in, out)
}

func (h *logisticsServiceHandler) SendExchange(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error {
	return h.LogisticsServiceHandler.SendExchange(ctx, in, out)
}

func (h *logisticsServiceHandler) VerifyFetchCode(ctx context.Context, in *LogisticsRequest, out *LogisticsResponse) error {
	return h.LogisticsServiceHandler.VerifyFetchCode(ctx, in, out)
}
