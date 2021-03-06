// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: timesService.proto

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

// Api Endpoints for TimesService service

func NewTimesServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TimesService service

type TimesService interface {
	Create(ctx context.Context, in *Times, opts ...client.CallOption) (*TimesResponse, error)
	Update(ctx context.Context, in *Times, opts ...client.CallOption) (*TimesResponse, error)
	Delete(ctx context.Context, in *TimesWhere, opts ...client.CallOption) (*TimesResponse, error)
	List(ctx context.Context, in *TimesWhere, opts ...client.CallOption) (*TimesResponse, error)
}

type timesService struct {
	c    client.Client
	name string
}

func NewTimesService(name string, c client.Client) TimesService {
	return &timesService{
		c:    c,
		name: name,
	}
}

func (c *timesService) Create(ctx context.Context, in *Times, opts ...client.CallOption) (*TimesResponse, error) {
	req := c.c.NewRequest(c.name, "TimesService.Create", in)
	out := new(TimesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesService) Update(ctx context.Context, in *Times, opts ...client.CallOption) (*TimesResponse, error) {
	req := c.c.NewRequest(c.name, "TimesService.Update", in)
	out := new(TimesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesService) Delete(ctx context.Context, in *TimesWhere, opts ...client.CallOption) (*TimesResponse, error) {
	req := c.c.NewRequest(c.name, "TimesService.Delete", in)
	out := new(TimesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timesService) List(ctx context.Context, in *TimesWhere, opts ...client.CallOption) (*TimesResponse, error) {
	req := c.c.NewRequest(c.name, "TimesService.List", in)
	out := new(TimesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TimesService service

type TimesServiceHandler interface {
	Create(context.Context, *Times, *TimesResponse) error
	Update(context.Context, *Times, *TimesResponse) error
	Delete(context.Context, *TimesWhere, *TimesResponse) error
	List(context.Context, *TimesWhere, *TimesResponse) error
}

func RegisterTimesServiceHandler(s server.Server, hdlr TimesServiceHandler, opts ...server.HandlerOption) error {
	type timesService interface {
		Create(ctx context.Context, in *Times, out *TimesResponse) error
		Update(ctx context.Context, in *Times, out *TimesResponse) error
		Delete(ctx context.Context, in *TimesWhere, out *TimesResponse) error
		List(ctx context.Context, in *TimesWhere, out *TimesResponse) error
	}
	type TimesService struct {
		timesService
	}
	h := &timesServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TimesService{h}, opts...))
}

type timesServiceHandler struct {
	TimesServiceHandler
}

func (h *timesServiceHandler) Create(ctx context.Context, in *Times, out *TimesResponse) error {
	return h.TimesServiceHandler.Create(ctx, in, out)
}

func (h *timesServiceHandler) Update(ctx context.Context, in *Times, out *TimesResponse) error {
	return h.TimesServiceHandler.Update(ctx, in, out)
}

func (h *timesServiceHandler) Delete(ctx context.Context, in *TimesWhere, out *TimesResponse) error {
	return h.TimesServiceHandler.Delete(ctx, in, out)
}

func (h *timesServiceHandler) List(ctx context.Context, in *TimesWhere, out *TimesResponse) error {
	return h.TimesServiceHandler.List(ctx, in, out)
}
