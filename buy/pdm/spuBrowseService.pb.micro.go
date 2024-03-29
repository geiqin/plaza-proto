// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: spuBrowseService.proto

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

// Api Endpoints for SpuBrowseService service

func NewSpuBrowseServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SpuBrowseService service

type SpuBrowseService interface {
	Count(ctx context.Context, in *SpuBrowseRequest, opts ...client.CallOption) (*SpuBrowseResponse, error)
	Delete(ctx context.Context, in *SpuBrowseRequest, opts ...client.CallOption) (*SpuBrowseResponse, error)
	Search(ctx context.Context, in *SpuBrowseRequest, opts ...client.CallOption) (*SpuBrowseResponse, error)
}

type spuBrowseService struct {
	c    client.Client
	name string
}

func NewSpuBrowseService(name string, c client.Client) SpuBrowseService {
	return &spuBrowseService{
		c:    c,
		name: name,
	}
}

func (c *spuBrowseService) Count(ctx context.Context, in *SpuBrowseRequest, opts ...client.CallOption) (*SpuBrowseResponse, error) {
	req := c.c.NewRequest(c.name, "SpuBrowseService.Count", in)
	out := new(SpuBrowseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuBrowseService) Delete(ctx context.Context, in *SpuBrowseRequest, opts ...client.CallOption) (*SpuBrowseResponse, error) {
	req := c.c.NewRequest(c.name, "SpuBrowseService.Delete", in)
	out := new(SpuBrowseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuBrowseService) Search(ctx context.Context, in *SpuBrowseRequest, opts ...client.CallOption) (*SpuBrowseResponse, error) {
	req := c.c.NewRequest(c.name, "SpuBrowseService.Search", in)
	out := new(SpuBrowseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SpuBrowseService service

type SpuBrowseServiceHandler interface {
	Count(context.Context, *SpuBrowseRequest, *SpuBrowseResponse) error
	Delete(context.Context, *SpuBrowseRequest, *SpuBrowseResponse) error
	Search(context.Context, *SpuBrowseRequest, *SpuBrowseResponse) error
}

func RegisterSpuBrowseServiceHandler(s server.Server, hdlr SpuBrowseServiceHandler, opts ...server.HandlerOption) error {
	type spuBrowseService interface {
		Count(ctx context.Context, in *SpuBrowseRequest, out *SpuBrowseResponse) error
		Delete(ctx context.Context, in *SpuBrowseRequest, out *SpuBrowseResponse) error
		Search(ctx context.Context, in *SpuBrowseRequest, out *SpuBrowseResponse) error
	}
	type SpuBrowseService struct {
		spuBrowseService
	}
	h := &spuBrowseServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SpuBrowseService{h}, opts...))
}

type spuBrowseServiceHandler struct {
	SpuBrowseServiceHandler
}

func (h *spuBrowseServiceHandler) Count(ctx context.Context, in *SpuBrowseRequest, out *SpuBrowseResponse) error {
	return h.SpuBrowseServiceHandler.Count(ctx, in, out)
}

func (h *spuBrowseServiceHandler) Delete(ctx context.Context, in *SpuBrowseRequest, out *SpuBrowseResponse) error {
	return h.SpuBrowseServiceHandler.Delete(ctx, in, out)
}

func (h *spuBrowseServiceHandler) Search(ctx context.Context, in *SpuBrowseRequest, out *SpuBrowseResponse) error {
	return h.SpuBrowseServiceHandler.Search(ctx, in, out)
}
