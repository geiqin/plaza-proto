// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: posterPositionService.proto

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

// Api Endpoints for PosterPositionService service

func NewPosterPositionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PosterPositionService service

type PosterPositionService interface {
	Create(ctx context.Context, in *PosterPosition, opts ...client.CallOption) (*PosterPositionResponse, error)
	Update(ctx context.Context, in *PosterPosition, opts ...client.CallOption) (*PosterPositionResponse, error)
	Delete(ctx context.Context, in *PosterPositionRequest, opts ...client.CallOption) (*PosterPositionResponse, error)
	Get(ctx context.Context, in *PosterPosition, opts ...client.CallOption) (*PosterPositionResponse, error)
	Search(ctx context.Context, in *PosterPositionRequest, opts ...client.CallOption) (*PosterPositionResponse, error)
}

type posterPositionService struct {
	c    client.Client
	name string
}

func NewPosterPositionService(name string, c client.Client) PosterPositionService {
	return &posterPositionService{
		c:    c,
		name: name,
	}
}

func (c *posterPositionService) Create(ctx context.Context, in *PosterPosition, opts ...client.CallOption) (*PosterPositionResponse, error) {
	req := c.c.NewRequest(c.name, "PosterPositionService.Create", in)
	out := new(PosterPositionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterPositionService) Update(ctx context.Context, in *PosterPosition, opts ...client.CallOption) (*PosterPositionResponse, error) {
	req := c.c.NewRequest(c.name, "PosterPositionService.Update", in)
	out := new(PosterPositionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterPositionService) Delete(ctx context.Context, in *PosterPositionRequest, opts ...client.CallOption) (*PosterPositionResponse, error) {
	req := c.c.NewRequest(c.name, "PosterPositionService.Delete", in)
	out := new(PosterPositionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterPositionService) Get(ctx context.Context, in *PosterPosition, opts ...client.CallOption) (*PosterPositionResponse, error) {
	req := c.c.NewRequest(c.name, "PosterPositionService.Get", in)
	out := new(PosterPositionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterPositionService) Search(ctx context.Context, in *PosterPositionRequest, opts ...client.CallOption) (*PosterPositionResponse, error) {
	req := c.c.NewRequest(c.name, "PosterPositionService.Search", in)
	out := new(PosterPositionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PosterPositionService service

type PosterPositionServiceHandler interface {
	Create(context.Context, *PosterPosition, *PosterPositionResponse) error
	Update(context.Context, *PosterPosition, *PosterPositionResponse) error
	Delete(context.Context, *PosterPositionRequest, *PosterPositionResponse) error
	Get(context.Context, *PosterPosition, *PosterPositionResponse) error
	Search(context.Context, *PosterPositionRequest, *PosterPositionResponse) error
}

func RegisterPosterPositionServiceHandler(s server.Server, hdlr PosterPositionServiceHandler, opts ...server.HandlerOption) error {
	type posterPositionService interface {
		Create(ctx context.Context, in *PosterPosition, out *PosterPositionResponse) error
		Update(ctx context.Context, in *PosterPosition, out *PosterPositionResponse) error
		Delete(ctx context.Context, in *PosterPositionRequest, out *PosterPositionResponse) error
		Get(ctx context.Context, in *PosterPosition, out *PosterPositionResponse) error
		Search(ctx context.Context, in *PosterPositionRequest, out *PosterPositionResponse) error
	}
	type PosterPositionService struct {
		posterPositionService
	}
	h := &posterPositionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PosterPositionService{h}, opts...))
}

type posterPositionServiceHandler struct {
	PosterPositionServiceHandler
}

func (h *posterPositionServiceHandler) Create(ctx context.Context, in *PosterPosition, out *PosterPositionResponse) error {
	return h.PosterPositionServiceHandler.Create(ctx, in, out)
}

func (h *posterPositionServiceHandler) Update(ctx context.Context, in *PosterPosition, out *PosterPositionResponse) error {
	return h.PosterPositionServiceHandler.Update(ctx, in, out)
}

func (h *posterPositionServiceHandler) Delete(ctx context.Context, in *PosterPositionRequest, out *PosterPositionResponse) error {
	return h.PosterPositionServiceHandler.Delete(ctx, in, out)
}

func (h *posterPositionServiceHandler) Get(ctx context.Context, in *PosterPosition, out *PosterPositionResponse) error {
	return h.PosterPositionServiceHandler.Get(ctx, in, out)
}

func (h *posterPositionServiceHandler) Search(ctx context.Context, in *PosterPositionRequest, out *PosterPositionResponse) error {
	return h.PosterPositionServiceHandler.Search(ctx, in, out)
}
