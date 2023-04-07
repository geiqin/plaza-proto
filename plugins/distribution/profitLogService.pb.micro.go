// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: profitLogService.proto

package services

import (
	_ "../common"
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

// Api Endpoints for ProfitLogService service

func NewProfitLogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProfitLogService service

type ProfitLogService interface {
	// 获得佣金详情
	Get(ctx context.Context, in *ProfitLog, opts ...client.CallOption) (*ProfitLogResponse, error)
	// 查询佣金列表
	Search(ctx context.Context, in *ProfitLogRequest, opts ...client.CallOption) (*ProfitLogResponse, error)
}

type profitLogService struct {
	c    client.Client
	name string
}

func NewProfitLogService(name string, c client.Client) ProfitLogService {
	return &profitLogService{
		c:    c,
		name: name,
	}
}

func (c *profitLogService) Get(ctx context.Context, in *ProfitLog, opts ...client.CallOption) (*ProfitLogResponse, error) {
	req := c.c.NewRequest(c.name, "ProfitLogService.Get", in)
	out := new(ProfitLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profitLogService) Search(ctx context.Context, in *ProfitLogRequest, opts ...client.CallOption) (*ProfitLogResponse, error) {
	req := c.c.NewRequest(c.name, "ProfitLogService.Search", in)
	out := new(ProfitLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProfitLogService service

type ProfitLogServiceHandler interface {
	// 获得佣金详情
	Get(context.Context, *ProfitLog, *ProfitLogResponse) error
	// 查询佣金列表
	Search(context.Context, *ProfitLogRequest, *ProfitLogResponse) error
}

func RegisterProfitLogServiceHandler(s server.Server, hdlr ProfitLogServiceHandler, opts ...server.HandlerOption) error {
	type profitLogService interface {
		Get(ctx context.Context, in *ProfitLog, out *ProfitLogResponse) error
		Search(ctx context.Context, in *ProfitLogRequest, out *ProfitLogResponse) error
	}
	type ProfitLogService struct {
		profitLogService
	}
	h := &profitLogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProfitLogService{h}, opts...))
}

type profitLogServiceHandler struct {
	ProfitLogServiceHandler
}

func (h *profitLogServiceHandler) Get(ctx context.Context, in *ProfitLog, out *ProfitLogResponse) error {
	return h.ProfitLogServiceHandler.Get(ctx, in, out)
}

func (h *profitLogServiceHandler) Search(ctx context.Context, in *ProfitLogRequest, out *ProfitLogResponse) error {
	return h.ProfitLogServiceHandler.Search(ctx, in, out)
}