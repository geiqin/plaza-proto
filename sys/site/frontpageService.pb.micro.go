// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: frontpageService.proto

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

// Api Endpoints for FrontpageService service

func NewFrontpageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FrontpageService service

type FrontpageService interface {
	Create(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error)
	Update(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error)
	Delete(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error)
	Get(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error)
	Search(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageResponse, error)
	List(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageResponse, error)
}

type frontpageService struct {
	c    client.Client
	name string
}

func NewFrontpageService(name string, c client.Client) FrontpageService {
	return &frontpageService{
		c:    c,
		name: name,
	}
}

func (c *frontpageService) Create(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.Create", in)
	out := new(FrontpageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) Update(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.Update", in)
	out := new(FrontpageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) Delete(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.Delete", in)
	out := new(FrontpageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) Get(ctx context.Context, in *Frontpage, opts ...client.CallOption) (*FrontpageResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.Get", in)
	out := new(FrontpageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) Search(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.Search", in)
	out := new(FrontpageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) List(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.List", in)
	out := new(FrontpageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontpageService service

type FrontpageServiceHandler interface {
	Create(context.Context, *Frontpage, *FrontpageResponse) error
	Update(context.Context, *Frontpage, *FrontpageResponse) error
	Delete(context.Context, *Frontpage, *FrontpageResponse) error
	Get(context.Context, *Frontpage, *FrontpageResponse) error
	Search(context.Context, *FrontpageRequest, *FrontpageResponse) error
	List(context.Context, *FrontpageRequest, *FrontpageResponse) error
}

func RegisterFrontpageServiceHandler(s server.Server, hdlr FrontpageServiceHandler, opts ...server.HandlerOption) error {
	type frontpageService interface {
		Create(ctx context.Context, in *Frontpage, out *FrontpageResponse) error
		Update(ctx context.Context, in *Frontpage, out *FrontpageResponse) error
		Delete(ctx context.Context, in *Frontpage, out *FrontpageResponse) error
		Get(ctx context.Context, in *Frontpage, out *FrontpageResponse) error
		Search(ctx context.Context, in *FrontpageRequest, out *FrontpageResponse) error
		List(ctx context.Context, in *FrontpageRequest, out *FrontpageResponse) error
	}
	type FrontpageService struct {
		frontpageService
	}
	h := &frontpageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FrontpageService{h}, opts...))
}

type frontpageServiceHandler struct {
	FrontpageServiceHandler
}

func (h *frontpageServiceHandler) Create(ctx context.Context, in *Frontpage, out *FrontpageResponse) error {
	return h.FrontpageServiceHandler.Create(ctx, in, out)
}

func (h *frontpageServiceHandler) Update(ctx context.Context, in *Frontpage, out *FrontpageResponse) error {
	return h.FrontpageServiceHandler.Update(ctx, in, out)
}

func (h *frontpageServiceHandler) Delete(ctx context.Context, in *Frontpage, out *FrontpageResponse) error {
	return h.FrontpageServiceHandler.Delete(ctx, in, out)
}

func (h *frontpageServiceHandler) Get(ctx context.Context, in *Frontpage, out *FrontpageResponse) error {
	return h.FrontpageServiceHandler.Get(ctx, in, out)
}

func (h *frontpageServiceHandler) Search(ctx context.Context, in *FrontpageRequest, out *FrontpageResponse) error {
	return h.FrontpageServiceHandler.Search(ctx, in, out)
}

func (h *frontpageServiceHandler) List(ctx context.Context, in *FrontpageRequest, out *FrontpageResponse) error {
	return h.FrontpageServiceHandler.List(ctx, in, out)
}
