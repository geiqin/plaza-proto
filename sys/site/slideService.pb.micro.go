// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: slideService.proto

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

// Api Endpoints for SlideService service

func NewSlideServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SlideService service

type SlideService interface {
	Create(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error)
	Update(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error)
	Delete(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error)
	Get(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error)
	Search(ctx context.Context, in *SlideRequest, opts ...client.CallOption) (*SlideResponse, error)
	List(ctx context.Context, in *SlideRequest, opts ...client.CallOption) (*SlideResponse, error)
}

type slideService struct {
	c    client.Client
	name string
}

func NewSlideService(name string, c client.Client) SlideService {
	return &slideService{
		c:    c,
		name: name,
	}
}

func (c *slideService) Create(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error) {
	req := c.c.NewRequest(c.name, "SlideService.Create", in)
	out := new(SlideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideService) Update(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error) {
	req := c.c.NewRequest(c.name, "SlideService.Update", in)
	out := new(SlideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideService) Delete(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error) {
	req := c.c.NewRequest(c.name, "SlideService.Delete", in)
	out := new(SlideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideService) Get(ctx context.Context, in *Slide, opts ...client.CallOption) (*SlideResponse, error) {
	req := c.c.NewRequest(c.name, "SlideService.Get", in)
	out := new(SlideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideService) Search(ctx context.Context, in *SlideRequest, opts ...client.CallOption) (*SlideResponse, error) {
	req := c.c.NewRequest(c.name, "SlideService.Search", in)
	out := new(SlideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slideService) List(ctx context.Context, in *SlideRequest, opts ...client.CallOption) (*SlideResponse, error) {
	req := c.c.NewRequest(c.name, "SlideService.List", in)
	out := new(SlideResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SlideService service

type SlideServiceHandler interface {
	Create(context.Context, *Slide, *SlideResponse) error
	Update(context.Context, *Slide, *SlideResponse) error
	Delete(context.Context, *Slide, *SlideResponse) error
	Get(context.Context, *Slide, *SlideResponse) error
	Search(context.Context, *SlideRequest, *SlideResponse) error
	List(context.Context, *SlideRequest, *SlideResponse) error
}

func RegisterSlideServiceHandler(s server.Server, hdlr SlideServiceHandler, opts ...server.HandlerOption) error {
	type slideService interface {
		Create(ctx context.Context, in *Slide, out *SlideResponse) error
		Update(ctx context.Context, in *Slide, out *SlideResponse) error
		Delete(ctx context.Context, in *Slide, out *SlideResponse) error
		Get(ctx context.Context, in *Slide, out *SlideResponse) error
		Search(ctx context.Context, in *SlideRequest, out *SlideResponse) error
		List(ctx context.Context, in *SlideRequest, out *SlideResponse) error
	}
	type SlideService struct {
		slideService
	}
	h := &slideServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SlideService{h}, opts...))
}

type slideServiceHandler struct {
	SlideServiceHandler
}

func (h *slideServiceHandler) Create(ctx context.Context, in *Slide, out *SlideResponse) error {
	return h.SlideServiceHandler.Create(ctx, in, out)
}

func (h *slideServiceHandler) Update(ctx context.Context, in *Slide, out *SlideResponse) error {
	return h.SlideServiceHandler.Update(ctx, in, out)
}

func (h *slideServiceHandler) Delete(ctx context.Context, in *Slide, out *SlideResponse) error {
	return h.SlideServiceHandler.Delete(ctx, in, out)
}

func (h *slideServiceHandler) Get(ctx context.Context, in *Slide, out *SlideResponse) error {
	return h.SlideServiceHandler.Get(ctx, in, out)
}

func (h *slideServiceHandler) Search(ctx context.Context, in *SlideRequest, out *SlideResponse) error {
	return h.SlideServiceHandler.Search(ctx, in, out)
}

func (h *slideServiceHandler) List(ctx context.Context, in *SlideRequest, out *SlideResponse) error {
	return h.SlideServiceHandler.List(ctx, in, out)
}
