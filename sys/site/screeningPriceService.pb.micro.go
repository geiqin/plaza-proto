// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: screeningPriceService.proto

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

// Api Endpoints for ScreeningPriceService service

func NewScreeningPriceServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ScreeningPriceService service

type ScreeningPriceService interface {
	Create(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error)
	Update(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error)
	Delete(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error)
	Get(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error)
	Search(ctx context.Context, in *ScreeningPriceRequest, opts ...client.CallOption) (*ScreeningPriceResponse, error)
	List(ctx context.Context, in *ScreeningPriceRequest, opts ...client.CallOption) (*ScreeningPriceResponse, error)
}

type screeningPriceService struct {
	c    client.Client
	name string
}

func NewScreeningPriceService(name string, c client.Client) ScreeningPriceService {
	return &screeningPriceService{
		c:    c,
		name: name,
	}
}

func (c *screeningPriceService) Create(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ScreeningPriceService.Create", in)
	out := new(ScreeningPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screeningPriceService) Update(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ScreeningPriceService.Update", in)
	out := new(ScreeningPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screeningPriceService) Delete(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ScreeningPriceService.Delete", in)
	out := new(ScreeningPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screeningPriceService) Get(ctx context.Context, in *ScreeningPrice, opts ...client.CallOption) (*ScreeningPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ScreeningPriceService.Get", in)
	out := new(ScreeningPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screeningPriceService) Search(ctx context.Context, in *ScreeningPriceRequest, opts ...client.CallOption) (*ScreeningPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ScreeningPriceService.Search", in)
	out := new(ScreeningPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screeningPriceService) List(ctx context.Context, in *ScreeningPriceRequest, opts ...client.CallOption) (*ScreeningPriceResponse, error) {
	req := c.c.NewRequest(c.name, "ScreeningPriceService.List", in)
	out := new(ScreeningPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScreeningPriceService service

type ScreeningPriceServiceHandler interface {
	Create(context.Context, *ScreeningPrice, *ScreeningPriceResponse) error
	Update(context.Context, *ScreeningPrice, *ScreeningPriceResponse) error
	Delete(context.Context, *ScreeningPrice, *ScreeningPriceResponse) error
	Get(context.Context, *ScreeningPrice, *ScreeningPriceResponse) error
	Search(context.Context, *ScreeningPriceRequest, *ScreeningPriceResponse) error
	List(context.Context, *ScreeningPriceRequest, *ScreeningPriceResponse) error
}

func RegisterScreeningPriceServiceHandler(s server.Server, hdlr ScreeningPriceServiceHandler, opts ...server.HandlerOption) error {
	type screeningPriceService interface {
		Create(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error
		Update(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error
		Delete(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error
		Get(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error
		Search(ctx context.Context, in *ScreeningPriceRequest, out *ScreeningPriceResponse) error
		List(ctx context.Context, in *ScreeningPriceRequest, out *ScreeningPriceResponse) error
	}
	type ScreeningPriceService struct {
		screeningPriceService
	}
	h := &screeningPriceServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ScreeningPriceService{h}, opts...))
}

type screeningPriceServiceHandler struct {
	ScreeningPriceServiceHandler
}

func (h *screeningPriceServiceHandler) Create(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error {
	return h.ScreeningPriceServiceHandler.Create(ctx, in, out)
}

func (h *screeningPriceServiceHandler) Update(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error {
	return h.ScreeningPriceServiceHandler.Update(ctx, in, out)
}

func (h *screeningPriceServiceHandler) Delete(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error {
	return h.ScreeningPriceServiceHandler.Delete(ctx, in, out)
}

func (h *screeningPriceServiceHandler) Get(ctx context.Context, in *ScreeningPrice, out *ScreeningPriceResponse) error {
	return h.ScreeningPriceServiceHandler.Get(ctx, in, out)
}

func (h *screeningPriceServiceHandler) Search(ctx context.Context, in *ScreeningPriceRequest, out *ScreeningPriceResponse) error {
	return h.ScreeningPriceServiceHandler.Search(ctx, in, out)
}

func (h *screeningPriceServiceHandler) List(ctx context.Context, in *ScreeningPriceRequest, out *ScreeningPriceResponse) error {
	return h.ScreeningPriceServiceHandler.List(ctx, in, out)
}
