// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: statisticsSerivce.proto

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

// Api Endpoints for StatisticsService service

func NewStatisticsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StatisticsService service

type StatisticsService interface {
	//数据总览【后台】
	Overview(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error)
	//支付订单统计【后台】
	PayOrderTotal(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error)
	//订单发货统计【后台】
	ShippingTotal(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error)
	//订单售后统计【后台】
	ComplaintTotal(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error)
}

type statisticsService struct {
	c    client.Client
	name string
}

func NewStatisticsService(name string, c client.Client) StatisticsService {
	return &statisticsService{
		c:    c,
		name: name,
	}
}

func (c *statisticsService) Overview(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error) {
	req := c.c.NewRequest(c.name, "StatisticsService.Overview", in)
	out := new(StatisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsService) PayOrderTotal(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error) {
	req := c.c.NewRequest(c.name, "StatisticsService.PayOrderTotal", in)
	out := new(StatisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsService) ShippingTotal(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error) {
	req := c.c.NewRequest(c.name, "StatisticsService.ShippingTotal", in)
	out := new(StatisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsService) ComplaintTotal(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error) {
	req := c.c.NewRequest(c.name, "StatisticsService.ComplaintTotal", in)
	out := new(StatisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StatisticsService service

type StatisticsServiceHandler interface {
	//数据总览【后台】
	Overview(context.Context, *StatisticsRequest, *StatisticsResponse) error
	//支付订单统计【后台】
	PayOrderTotal(context.Context, *StatisticsRequest, *StatisticsResponse) error
	//订单发货统计【后台】
	ShippingTotal(context.Context, *StatisticsRequest, *StatisticsResponse) error
	//订单售后统计【后台】
	ComplaintTotal(context.Context, *StatisticsRequest, *StatisticsResponse) error
}

func RegisterStatisticsServiceHandler(s server.Server, hdlr StatisticsServiceHandler, opts ...server.HandlerOption) error {
	type statisticsService interface {
		Overview(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error
		PayOrderTotal(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error
		ShippingTotal(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error
		ComplaintTotal(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error
	}
	type StatisticsService struct {
		statisticsService
	}
	h := &statisticsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StatisticsService{h}, opts...))
}

type statisticsServiceHandler struct {
	StatisticsServiceHandler
}

func (h *statisticsServiceHandler) Overview(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error {
	return h.StatisticsServiceHandler.Overview(ctx, in, out)
}

func (h *statisticsServiceHandler) PayOrderTotal(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error {
	return h.StatisticsServiceHandler.PayOrderTotal(ctx, in, out)
}

func (h *statisticsServiceHandler) ShippingTotal(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error {
	return h.StatisticsServiceHandler.ShippingTotal(ctx, in, out)
}

func (h *statisticsServiceHandler) ComplaintTotal(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error {
	return h.StatisticsServiceHandler.ComplaintTotal(ctx, in, out)
}
