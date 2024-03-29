// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: navTypeService.proto

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

// Api Endpoints for NavTypeService service

func NewNavTypeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NavTypeService service

type NavTypeService interface {
	Get(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error)
	Create(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error)
	Update(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error)
	Delete(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error)
	Search(ctx context.Context, in *NavTypeRequest, opts ...client.CallOption) (*NavTypeResponse, error)
	List(ctx context.Context, in *NavTypeRequest, opts ...client.CallOption) (*NavTypeResponse, error)
	StoreNavs(ctx context.Context, in *NavTypeRequest, opts ...client.CallOption) (*StoreNavsResponse, error)
}

type navTypeService struct {
	c    client.Client
	name string
}

func NewNavTypeService(name string, c client.Client) NavTypeService {
	return &navTypeService{
		c:    c,
		name: name,
	}
}

func (c *navTypeService) Get(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error) {
	req := c.c.NewRequest(c.name, "NavTypeService.Get", in)
	out := new(NavTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navTypeService) Create(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error) {
	req := c.c.NewRequest(c.name, "NavTypeService.Create", in)
	out := new(NavTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navTypeService) Update(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error) {
	req := c.c.NewRequest(c.name, "NavTypeService.Update", in)
	out := new(NavTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navTypeService) Delete(ctx context.Context, in *NavType, opts ...client.CallOption) (*NavTypeResponse, error) {
	req := c.c.NewRequest(c.name, "NavTypeService.Delete", in)
	out := new(NavTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navTypeService) Search(ctx context.Context, in *NavTypeRequest, opts ...client.CallOption) (*NavTypeResponse, error) {
	req := c.c.NewRequest(c.name, "NavTypeService.Search", in)
	out := new(NavTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navTypeService) List(ctx context.Context, in *NavTypeRequest, opts ...client.CallOption) (*NavTypeResponse, error) {
	req := c.c.NewRequest(c.name, "NavTypeService.List", in)
	out := new(NavTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navTypeService) StoreNavs(ctx context.Context, in *NavTypeRequest, opts ...client.CallOption) (*StoreNavsResponse, error) {
	req := c.c.NewRequest(c.name, "NavTypeService.StoreNavs", in)
	out := new(StoreNavsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NavTypeService service

type NavTypeServiceHandler interface {
	Get(context.Context, *NavType, *NavTypeResponse) error
	Create(context.Context, *NavType, *NavTypeResponse) error
	Update(context.Context, *NavType, *NavTypeResponse) error
	Delete(context.Context, *NavType, *NavTypeResponse) error
	Search(context.Context, *NavTypeRequest, *NavTypeResponse) error
	List(context.Context, *NavTypeRequest, *NavTypeResponse) error
	StoreNavs(context.Context, *NavTypeRequest, *StoreNavsResponse) error
}

func RegisterNavTypeServiceHandler(s server.Server, hdlr NavTypeServiceHandler, opts ...server.HandlerOption) error {
	type navTypeService interface {
		Get(ctx context.Context, in *NavType, out *NavTypeResponse) error
		Create(ctx context.Context, in *NavType, out *NavTypeResponse) error
		Update(ctx context.Context, in *NavType, out *NavTypeResponse) error
		Delete(ctx context.Context, in *NavType, out *NavTypeResponse) error
		Search(ctx context.Context, in *NavTypeRequest, out *NavTypeResponse) error
		List(ctx context.Context, in *NavTypeRequest, out *NavTypeResponse) error
		StoreNavs(ctx context.Context, in *NavTypeRequest, out *StoreNavsResponse) error
	}
	type NavTypeService struct {
		navTypeService
	}
	h := &navTypeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&NavTypeService{h}, opts...))
}

type navTypeServiceHandler struct {
	NavTypeServiceHandler
}

func (h *navTypeServiceHandler) Get(ctx context.Context, in *NavType, out *NavTypeResponse) error {
	return h.NavTypeServiceHandler.Get(ctx, in, out)
}

func (h *navTypeServiceHandler) Create(ctx context.Context, in *NavType, out *NavTypeResponse) error {
	return h.NavTypeServiceHandler.Create(ctx, in, out)
}

func (h *navTypeServiceHandler) Update(ctx context.Context, in *NavType, out *NavTypeResponse) error {
	return h.NavTypeServiceHandler.Update(ctx, in, out)
}

func (h *navTypeServiceHandler) Delete(ctx context.Context, in *NavType, out *NavTypeResponse) error {
	return h.NavTypeServiceHandler.Delete(ctx, in, out)
}

func (h *navTypeServiceHandler) Search(ctx context.Context, in *NavTypeRequest, out *NavTypeResponse) error {
	return h.NavTypeServiceHandler.Search(ctx, in, out)
}

func (h *navTypeServiceHandler) List(ctx context.Context, in *NavTypeRequest, out *NavTypeResponse) error {
	return h.NavTypeServiceHandler.List(ctx, in, out)
}

func (h *navTypeServiceHandler) StoreNavs(ctx context.Context, in *NavTypeRequest, out *StoreNavsResponse) error {
	return h.NavTypeServiceHandler.StoreNavs(ctx, in, out)
}
