// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: navService.proto

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

// Api Endpoints for NavService service

func NewNavServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NavService service

type NavService interface {
	Get(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error)
	Create(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error)
	Update(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error)
	UpdateHidden(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error)
	Delete(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error)
	Search(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error)
	List(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error)
	Tree(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error)
	//插件导航创建
	PluginNavCreate(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error)
	//插件导航删除
	PluginNavDelete(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error)
}

type navService struct {
	c    client.Client
	name string
}

func NewNavService(name string, c client.Client) NavService {
	return &navService{
		c:    c,
		name: name,
	}
}

func (c *navService) Get(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.Get", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) Create(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.Create", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) Update(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.Update", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) UpdateHidden(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.UpdateHidden", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) Delete(ctx context.Context, in *Nav, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.Delete", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) Search(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.Search", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) List(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.List", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) Tree(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.Tree", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) PluginNavCreate(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.PluginNavCreate", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navService) PluginNavDelete(ctx context.Context, in *NavRequest, opts ...client.CallOption) (*NavResponse, error) {
	req := c.c.NewRequest(c.name, "NavService.PluginNavDelete", in)
	out := new(NavResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NavService service

type NavServiceHandler interface {
	Get(context.Context, *Nav, *NavResponse) error
	Create(context.Context, *Nav, *NavResponse) error
	Update(context.Context, *Nav, *NavResponse) error
	UpdateHidden(context.Context, *Nav, *NavResponse) error
	Delete(context.Context, *Nav, *NavResponse) error
	Search(context.Context, *NavRequest, *NavResponse) error
	List(context.Context, *NavRequest, *NavResponse) error
	Tree(context.Context, *NavRequest, *NavResponse) error
	//插件导航创建
	PluginNavCreate(context.Context, *NavRequest, *NavResponse) error
	//插件导航删除
	PluginNavDelete(context.Context, *NavRequest, *NavResponse) error
}

func RegisterNavServiceHandler(s server.Server, hdlr NavServiceHandler, opts ...server.HandlerOption) error {
	type navService interface {
		Get(ctx context.Context, in *Nav, out *NavResponse) error
		Create(ctx context.Context, in *Nav, out *NavResponse) error
		Update(ctx context.Context, in *Nav, out *NavResponse) error
		UpdateHidden(ctx context.Context, in *Nav, out *NavResponse) error
		Delete(ctx context.Context, in *Nav, out *NavResponse) error
		Search(ctx context.Context, in *NavRequest, out *NavResponse) error
		List(ctx context.Context, in *NavRequest, out *NavResponse) error
		Tree(ctx context.Context, in *NavRequest, out *NavResponse) error
		PluginNavCreate(ctx context.Context, in *NavRequest, out *NavResponse) error
		PluginNavDelete(ctx context.Context, in *NavRequest, out *NavResponse) error
	}
	type NavService struct {
		navService
	}
	h := &navServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&NavService{h}, opts...))
}

type navServiceHandler struct {
	NavServiceHandler
}

func (h *navServiceHandler) Get(ctx context.Context, in *Nav, out *NavResponse) error {
	return h.NavServiceHandler.Get(ctx, in, out)
}

func (h *navServiceHandler) Create(ctx context.Context, in *Nav, out *NavResponse) error {
	return h.NavServiceHandler.Create(ctx, in, out)
}

func (h *navServiceHandler) Update(ctx context.Context, in *Nav, out *NavResponse) error {
	return h.NavServiceHandler.Update(ctx, in, out)
}

func (h *navServiceHandler) UpdateHidden(ctx context.Context, in *Nav, out *NavResponse) error {
	return h.NavServiceHandler.UpdateHidden(ctx, in, out)
}

func (h *navServiceHandler) Delete(ctx context.Context, in *Nav, out *NavResponse) error {
	return h.NavServiceHandler.Delete(ctx, in, out)
}

func (h *navServiceHandler) Search(ctx context.Context, in *NavRequest, out *NavResponse) error {
	return h.NavServiceHandler.Search(ctx, in, out)
}

func (h *navServiceHandler) List(ctx context.Context, in *NavRequest, out *NavResponse) error {
	return h.NavServiceHandler.List(ctx, in, out)
}

func (h *navServiceHandler) Tree(ctx context.Context, in *NavRequest, out *NavResponse) error {
	return h.NavServiceHandler.Tree(ctx, in, out)
}

func (h *navServiceHandler) PluginNavCreate(ctx context.Context, in *NavRequest, out *NavResponse) error {
	return h.NavServiceHandler.PluginNavCreate(ctx, in, out)
}

func (h *navServiceHandler) PluginNavDelete(ctx context.Context, in *NavRequest, out *NavResponse) error {
	return h.NavServiceHandler.PluginNavDelete(ctx, in, out)
}
