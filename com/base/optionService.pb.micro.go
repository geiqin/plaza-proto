// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: optionService.proto

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

// Api Endpoints for OptionService service

func NewOptionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OptionService service

type OptionService interface {
	// 全局配置新增
	Create(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error)
	// 全局配置修改
	Update(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error)
	// 全局配置删除
	Delete(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error)
	// 全局配置获取
	Get(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error)
	// 全局配置查询
	Search(ctx context.Context, in *OptionRequest, opts ...client.CallOption) (*OptionResponse, error)
	// 全局配置列表
	List(ctx context.Context, in *OptionRequest, opts ...client.CallOption) (*OptionResponse, error)
}

type optionService struct {
	c    client.Client
	name string
}

func NewOptionService(name string, c client.Client) OptionService {
	return &optionService{
		c:    c,
		name: name,
	}
}

func (c *optionService) Create(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error) {
	req := c.c.NewRequest(c.name, "OptionService.Create", in)
	out := new(OptionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionService) Update(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error) {
	req := c.c.NewRequest(c.name, "OptionService.Update", in)
	out := new(OptionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionService) Delete(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error) {
	req := c.c.NewRequest(c.name, "OptionService.Delete", in)
	out := new(OptionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionService) Get(ctx context.Context, in *Option, opts ...client.CallOption) (*OptionResponse, error) {
	req := c.c.NewRequest(c.name, "OptionService.Get", in)
	out := new(OptionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionService) Search(ctx context.Context, in *OptionRequest, opts ...client.CallOption) (*OptionResponse, error) {
	req := c.c.NewRequest(c.name, "OptionService.Search", in)
	out := new(OptionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionService) List(ctx context.Context, in *OptionRequest, opts ...client.CallOption) (*OptionResponse, error) {
	req := c.c.NewRequest(c.name, "OptionService.List", in)
	out := new(OptionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OptionService service

type OptionServiceHandler interface {
	// 全局配置新增
	Create(context.Context, *Option, *OptionResponse) error
	// 全局配置修改
	Update(context.Context, *Option, *OptionResponse) error
	// 全局配置删除
	Delete(context.Context, *Option, *OptionResponse) error
	// 全局配置获取
	Get(context.Context, *Option, *OptionResponse) error
	// 全局配置查询
	Search(context.Context, *OptionRequest, *OptionResponse) error
	// 全局配置列表
	List(context.Context, *OptionRequest, *OptionResponse) error
}

func RegisterOptionServiceHandler(s server.Server, hdlr OptionServiceHandler, opts ...server.HandlerOption) error {
	type optionService interface {
		Create(ctx context.Context, in *Option, out *OptionResponse) error
		Update(ctx context.Context, in *Option, out *OptionResponse) error
		Delete(ctx context.Context, in *Option, out *OptionResponse) error
		Get(ctx context.Context, in *Option, out *OptionResponse) error
		Search(ctx context.Context, in *OptionRequest, out *OptionResponse) error
		List(ctx context.Context, in *OptionRequest, out *OptionResponse) error
	}
	type OptionService struct {
		optionService
	}
	h := &optionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OptionService{h}, opts...))
}

type optionServiceHandler struct {
	OptionServiceHandler
}

func (h *optionServiceHandler) Create(ctx context.Context, in *Option, out *OptionResponse) error {
	return h.OptionServiceHandler.Create(ctx, in, out)
}

func (h *optionServiceHandler) Update(ctx context.Context, in *Option, out *OptionResponse) error {
	return h.OptionServiceHandler.Update(ctx, in, out)
}

func (h *optionServiceHandler) Delete(ctx context.Context, in *Option, out *OptionResponse) error {
	return h.OptionServiceHandler.Delete(ctx, in, out)
}

func (h *optionServiceHandler) Get(ctx context.Context, in *Option, out *OptionResponse) error {
	return h.OptionServiceHandler.Get(ctx, in, out)
}

func (h *optionServiceHandler) Search(ctx context.Context, in *OptionRequest, out *OptionResponse) error {
	return h.OptionServiceHandler.Search(ctx, in, out)
}

func (h *optionServiceHandler) List(ctx context.Context, in *OptionRequest, out *OptionResponse) error {
	return h.OptionServiceHandler.List(ctx, in, out)
}
