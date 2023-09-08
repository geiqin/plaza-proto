// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: termService.proto

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

// Api Endpoints for TermService service

func NewTermServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TermService service

type TermService interface {
	Create(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error)
	Update(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error)
	Delete(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error)
	Get(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error)
	Tree(ctx context.Context, in *TermRequest, opts ...client.CallOption) (*TermResponse, error)
	Search(ctx context.Context, in *TermRequest, opts ...client.CallOption) (*TermResponse, error)
}

type termService struct {
	c    client.Client
	name string
}

func NewTermService(name string, c client.Client) TermService {
	return &termService{
		c:    c,
		name: name,
	}
}

func (c *termService) Create(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error) {
	req := c.c.NewRequest(c.name, "TermService.Create", in)
	out := new(TermResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termService) Update(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error) {
	req := c.c.NewRequest(c.name, "TermService.Update", in)
	out := new(TermResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termService) Delete(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error) {
	req := c.c.NewRequest(c.name, "TermService.Delete", in)
	out := new(TermResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termService) Get(ctx context.Context, in *Term, opts ...client.CallOption) (*TermResponse, error) {
	req := c.c.NewRequest(c.name, "TermService.Get", in)
	out := new(TermResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termService) Tree(ctx context.Context, in *TermRequest, opts ...client.CallOption) (*TermResponse, error) {
	req := c.c.NewRequest(c.name, "TermService.Tree", in)
	out := new(TermResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termService) Search(ctx context.Context, in *TermRequest, opts ...client.CallOption) (*TermResponse, error) {
	req := c.c.NewRequest(c.name, "TermService.Search", in)
	out := new(TermResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TermService service

type TermServiceHandler interface {
	Create(context.Context, *Term, *TermResponse) error
	Update(context.Context, *Term, *TermResponse) error
	Delete(context.Context, *Term, *TermResponse) error
	Get(context.Context, *Term, *TermResponse) error
	Tree(context.Context, *TermRequest, *TermResponse) error
	Search(context.Context, *TermRequest, *TermResponse) error
}

func RegisterTermServiceHandler(s server.Server, hdlr TermServiceHandler, opts ...server.HandlerOption) error {
	type termService interface {
		Create(ctx context.Context, in *Term, out *TermResponse) error
		Update(ctx context.Context, in *Term, out *TermResponse) error
		Delete(ctx context.Context, in *Term, out *TermResponse) error
		Get(ctx context.Context, in *Term, out *TermResponse) error
		Tree(ctx context.Context, in *TermRequest, out *TermResponse) error
		Search(ctx context.Context, in *TermRequest, out *TermResponse) error
	}
	type TermService struct {
		termService
	}
	h := &termServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TermService{h}, opts...))
}

type termServiceHandler struct {
	TermServiceHandler
}

func (h *termServiceHandler) Create(ctx context.Context, in *Term, out *TermResponse) error {
	return h.TermServiceHandler.Create(ctx, in, out)
}

func (h *termServiceHandler) Update(ctx context.Context, in *Term, out *TermResponse) error {
	return h.TermServiceHandler.Update(ctx, in, out)
}

func (h *termServiceHandler) Delete(ctx context.Context, in *Term, out *TermResponse) error {
	return h.TermServiceHandler.Delete(ctx, in, out)
}

func (h *termServiceHandler) Get(ctx context.Context, in *Term, out *TermResponse) error {
	return h.TermServiceHandler.Get(ctx, in, out)
}

func (h *termServiceHandler) Tree(ctx context.Context, in *TermRequest, out *TermResponse) error {
	return h.TermServiceHandler.Tree(ctx, in, out)
}

func (h *termServiceHandler) Search(ctx context.Context, in *TermRequest, out *TermResponse) error {
	return h.TermServiceHandler.Search(ctx, in, out)
}
