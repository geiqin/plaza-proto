// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: statisticsService.proto

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

// Api Endpoints for Statistics service

func NewStatisticsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Statistics service

type StatisticsService interface {
	OrderDay(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error)
	//出库月统计
	StockOutMonth(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error)
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

func (c *statisticsService) OrderDay(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error) {
	req := c.c.NewRequest(c.name, "Statistics.OrderDay", in)
	out := new(StatisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statisticsService) StockOutMonth(ctx context.Context, in *StatisticsRequest, opts ...client.CallOption) (*StatisticsResponse, error) {
	req := c.c.NewRequest(c.name, "Statistics.StockOutMonth", in)
	out := new(StatisticsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Statistics service

type StatisticsHandler interface {
	OrderDay(context.Context, *StatisticsRequest, *StatisticsResponse) error
	//出库月统计
	StockOutMonth(context.Context, *StatisticsRequest, *StatisticsResponse) error
}

func RegisterStatisticsHandler(s server.Server, hdlr StatisticsHandler, opts ...server.HandlerOption) error {
	type statistics interface {
		OrderDay(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error
		StockOutMonth(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error
	}
	type Statistics struct {
		statistics
	}
	h := &statisticsHandler{hdlr}
	return s.Handle(s.NewHandler(&Statistics{h}, opts...))
}

type statisticsHandler struct {
	StatisticsHandler
}

func (h *statisticsHandler) OrderDay(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error {
	return h.StatisticsHandler.OrderDay(ctx, in, out)
}

func (h *statisticsHandler) StockOutMonth(ctx context.Context, in *StatisticsRequest, out *StatisticsResponse) error {
	return h.StatisticsHandler.StockOutMonth(ctx, in, out)
}
