// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: versionMenuService.proto

package services

import (
	_ "github.com/geiqin/micro-kit/protobuf/common"
	fmt "fmt"
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

// Api Endpoints for VersionMenuService service

func NewVersionMenuServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VersionMenuService service

type VersionMenuService interface {
	Get(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error)
	Create(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error)
	Update(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error)
	UpdateHide(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error)
	Delete(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error)
	Search(ctx context.Context, in *VersionMenuRequest, opts ...client.CallOption) (*VersionMenuResponse, error)
	List(ctx context.Context, in *VersionMenuRequest, opts ...client.CallOption) (*VersionMenuResponse, error)
	Tree(ctx context.Context, in *VersionMenuRequest, opts ...client.CallOption) (*VersionMenuResponse, error)
}

type versionMenuService struct {
	c    client.Client
	name string
}

func NewVersionMenuService(name string, c client.Client) VersionMenuService {
	return &versionMenuService{
		c:    c,
		name: name,
	}
}

func (c *versionMenuService) Get(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.Get", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionMenuService) Create(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.Create", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionMenuService) Update(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.Update", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionMenuService) UpdateHide(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.UpdateHide", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionMenuService) Delete(ctx context.Context, in *VersionMenu, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.Delete", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionMenuService) Search(ctx context.Context, in *VersionMenuRequest, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.Search", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionMenuService) List(ctx context.Context, in *VersionMenuRequest, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.List", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionMenuService) Tree(ctx context.Context, in *VersionMenuRequest, opts ...client.CallOption) (*VersionMenuResponse, error) {
	req := c.c.NewRequest(c.name, "VersionMenuService.Tree", in)
	out := new(VersionMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VersionMenuService service

type VersionMenuServiceHandler interface {
	Get(context.Context, *VersionMenu, *VersionMenuResponse) error
	Create(context.Context, *VersionMenu, *VersionMenuResponse) error
	Update(context.Context, *VersionMenu, *VersionMenuResponse) error
	UpdateHide(context.Context, *VersionMenu, *VersionMenuResponse) error
	Delete(context.Context, *VersionMenu, *VersionMenuResponse) error
	Search(context.Context, *VersionMenuRequest, *VersionMenuResponse) error
	List(context.Context, *VersionMenuRequest, *VersionMenuResponse) error
	Tree(context.Context, *VersionMenuRequest, *VersionMenuResponse) error
}

func RegisterVersionMenuServiceHandler(s server.Server, hdlr VersionMenuServiceHandler, opts ...server.HandlerOption) error {
	type versionMenuService interface {
		Get(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error
		Create(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error
		Update(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error
		UpdateHide(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error
		Delete(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error
		Search(ctx context.Context, in *VersionMenuRequest, out *VersionMenuResponse) error
		List(ctx context.Context, in *VersionMenuRequest, out *VersionMenuResponse) error
		Tree(ctx context.Context, in *VersionMenuRequest, out *VersionMenuResponse) error
	}
	type VersionMenuService struct {
		versionMenuService
	}
	h := &versionMenuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VersionMenuService{h}, opts...))
}

type versionMenuServiceHandler struct {
	VersionMenuServiceHandler
}

func (h *versionMenuServiceHandler) Get(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.Get(ctx, in, out)
}

func (h *versionMenuServiceHandler) Create(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.Create(ctx, in, out)
}

func (h *versionMenuServiceHandler) Update(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.Update(ctx, in, out)
}

func (h *versionMenuServiceHandler) UpdateHide(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.UpdateHide(ctx, in, out)
}

func (h *versionMenuServiceHandler) Delete(ctx context.Context, in *VersionMenu, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.Delete(ctx, in, out)
}

func (h *versionMenuServiceHandler) Search(ctx context.Context, in *VersionMenuRequest, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.Search(ctx, in, out)
}

func (h *versionMenuServiceHandler) List(ctx context.Context, in *VersionMenuRequest, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.List(ctx, in, out)
}

func (h *versionMenuServiceHandler) Tree(ctx context.Context, in *VersionMenuRequest, out *VersionMenuResponse) error {
	return h.VersionMenuServiceHandler.Tree(ctx, in, out)
}
