// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: designService.proto

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

// Api Endpoints for DesignService service

func NewDesignServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DesignService service

type DesignService interface {
	Create(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error)
	Update(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error)
	Delete(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error)
	Get(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error)
	Search(ctx context.Context, in *DesignRequest, opts ...client.CallOption) (*DesignResponse, error)
	List(ctx context.Context, in *DesignRequest, opts ...client.CallOption) (*DesignResponse, error)
}

type designService struct {
	c    client.Client
	name string
}

func NewDesignService(name string, c client.Client) DesignService {
	return &designService{
		c:    c,
		name: name,
	}
}

func (c *designService) Create(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error) {
	req := c.c.NewRequest(c.name, "DesignService.Create", in)
	out := new(DesignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *designService) Update(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error) {
	req := c.c.NewRequest(c.name, "DesignService.Update", in)
	out := new(DesignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *designService) Delete(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error) {
	req := c.c.NewRequest(c.name, "DesignService.Delete", in)
	out := new(DesignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *designService) Get(ctx context.Context, in *Design, opts ...client.CallOption) (*DesignResponse, error) {
	req := c.c.NewRequest(c.name, "DesignService.Get", in)
	out := new(DesignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *designService) Search(ctx context.Context, in *DesignRequest, opts ...client.CallOption) (*DesignResponse, error) {
	req := c.c.NewRequest(c.name, "DesignService.Search", in)
	out := new(DesignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *designService) List(ctx context.Context, in *DesignRequest, opts ...client.CallOption) (*DesignResponse, error) {
	req := c.c.NewRequest(c.name, "DesignService.List", in)
	out := new(DesignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DesignService service

type DesignServiceHandler interface {
	Create(context.Context, *Design, *DesignResponse) error
	Update(context.Context, *Design, *DesignResponse) error
	Delete(context.Context, *Design, *DesignResponse) error
	Get(context.Context, *Design, *DesignResponse) error
	Search(context.Context, *DesignRequest, *DesignResponse) error
	List(context.Context, *DesignRequest, *DesignResponse) error
}

func RegisterDesignServiceHandler(s server.Server, hdlr DesignServiceHandler, opts ...server.HandlerOption) error {
	type designService interface {
		Create(ctx context.Context, in *Design, out *DesignResponse) error
		Update(ctx context.Context, in *Design, out *DesignResponse) error
		Delete(ctx context.Context, in *Design, out *DesignResponse) error
		Get(ctx context.Context, in *Design, out *DesignResponse) error
		Search(ctx context.Context, in *DesignRequest, out *DesignResponse) error
		List(ctx context.Context, in *DesignRequest, out *DesignResponse) error
	}
	type DesignService struct {
		designService
	}
	h := &designServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DesignService{h}, opts...))
}

type designServiceHandler struct {
	DesignServiceHandler
}

func (h *designServiceHandler) Create(ctx context.Context, in *Design, out *DesignResponse) error {
	return h.DesignServiceHandler.Create(ctx, in, out)
}

func (h *designServiceHandler) Update(ctx context.Context, in *Design, out *DesignResponse) error {
	return h.DesignServiceHandler.Update(ctx, in, out)
}

func (h *designServiceHandler) Delete(ctx context.Context, in *Design, out *DesignResponse) error {
	return h.DesignServiceHandler.Delete(ctx, in, out)
}

func (h *designServiceHandler) Get(ctx context.Context, in *Design, out *DesignResponse) error {
	return h.DesignServiceHandler.Get(ctx, in, out)
}

func (h *designServiceHandler) Search(ctx context.Context, in *DesignRequest, out *DesignResponse) error {
	return h.DesignServiceHandler.Search(ctx, in, out)
}

func (h *designServiceHandler) List(ctx context.Context, in *DesignRequest, out *DesignResponse) error {
	return h.DesignServiceHandler.List(ctx, in, out)
}
