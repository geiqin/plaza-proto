// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: regionService.proto

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

// Api Endpoints for RegionService service

func NewRegionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RegionService service

type RegionService interface {
	Create(ctx context.Context, in *Region, opts ...client.CallOption) (*RegionResponse, error)
	Update(ctx context.Context, in *Region, opts ...client.CallOption) (*RegionResponse, error)
	Delete(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error)
	Get(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error)
	Search(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error)
	List(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error)
}

type regionService struct {
	c    client.Client
	name string
}

func NewRegionService(name string, c client.Client) RegionService {
	return &regionService{
		c:    c,
		name: name,
	}
}

func (c *regionService) Create(ctx context.Context, in *Region, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "RegionService.Create", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionService) Update(ctx context.Context, in *Region, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "RegionService.Update", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionService) Delete(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "RegionService.Delete", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionService) Get(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "RegionService.Get", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionService) Search(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "RegionService.Search", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionService) List(ctx context.Context, in *RegionRequest, opts ...client.CallOption) (*RegionResponse, error) {
	req := c.c.NewRequest(c.name, "RegionService.List", in)
	out := new(RegionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegionService service

type RegionServiceHandler interface {
	Create(context.Context, *Region, *RegionResponse) error
	Update(context.Context, *Region, *RegionResponse) error
	Delete(context.Context, *RegionRequest, *RegionResponse) error
	Get(context.Context, *RegionRequest, *RegionResponse) error
	Search(context.Context, *RegionRequest, *RegionResponse) error
	List(context.Context, *RegionRequest, *RegionResponse) error
}

func RegisterRegionServiceHandler(s server.Server, hdlr RegionServiceHandler, opts ...server.HandlerOption) error {
	type regionService interface {
		Create(ctx context.Context, in *Region, out *RegionResponse) error
		Update(ctx context.Context, in *Region, out *RegionResponse) error
		Delete(ctx context.Context, in *RegionRequest, out *RegionResponse) error
		Get(ctx context.Context, in *RegionRequest, out *RegionResponse) error
		Search(ctx context.Context, in *RegionRequest, out *RegionResponse) error
		List(ctx context.Context, in *RegionRequest, out *RegionResponse) error
	}
	type RegionService struct {
		regionService
	}
	h := &regionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RegionService{h}, opts...))
}

type regionServiceHandler struct {
	RegionServiceHandler
}

func (h *regionServiceHandler) Create(ctx context.Context, in *Region, out *RegionResponse) error {
	return h.RegionServiceHandler.Create(ctx, in, out)
}

func (h *regionServiceHandler) Update(ctx context.Context, in *Region, out *RegionResponse) error {
	return h.RegionServiceHandler.Update(ctx, in, out)
}

func (h *regionServiceHandler) Delete(ctx context.Context, in *RegionRequest, out *RegionResponse) error {
	return h.RegionServiceHandler.Delete(ctx, in, out)
}

func (h *regionServiceHandler) Get(ctx context.Context, in *RegionRequest, out *RegionResponse) error {
	return h.RegionServiceHandler.Get(ctx, in, out)
}

func (h *regionServiceHandler) Search(ctx context.Context, in *RegionRequest, out *RegionResponse) error {
	return h.RegionServiceHandler.Search(ctx, in, out)
}

func (h *regionServiceHandler) List(ctx context.Context, in *RegionRequest, out *RegionResponse) error {
	return h.RegionServiceHandler.List(ctx, in, out)
}
