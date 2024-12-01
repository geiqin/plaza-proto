// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: iconService.proto

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

// Api Endpoints for IconService service

func NewIconServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IconService service

type IconService interface {
	// 公共图标新增
	Create(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error)
	// 公共图标修改
	Update(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error)
	// 公共图标删除
	Delete(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error)
	// 公共图标获取
	Get(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error)
	// 公共图标查询
	Search(ctx context.Context, in *IconRequest, opts ...client.CallOption) (*IconResponse, error)
	// 公共图标列表
	List(ctx context.Context, in *IconRequest, opts ...client.CallOption) (*IconResponse, error)
}

type iconService struct {
	c    client.Client
	name string
}

func NewIconService(name string, c client.Client) IconService {
	return &iconService{
		c:    c,
		name: name,
	}
}

func (c *iconService) Create(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error) {
	req := c.c.NewRequest(c.name, "IconService.Create", in)
	out := new(IconResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iconService) Update(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error) {
	req := c.c.NewRequest(c.name, "IconService.Update", in)
	out := new(IconResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iconService) Delete(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error) {
	req := c.c.NewRequest(c.name, "IconService.Delete", in)
	out := new(IconResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iconService) Get(ctx context.Context, in *Icon, opts ...client.CallOption) (*IconResponse, error) {
	req := c.c.NewRequest(c.name, "IconService.Get", in)
	out := new(IconResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iconService) Search(ctx context.Context, in *IconRequest, opts ...client.CallOption) (*IconResponse, error) {
	req := c.c.NewRequest(c.name, "IconService.Search", in)
	out := new(IconResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iconService) List(ctx context.Context, in *IconRequest, opts ...client.CallOption) (*IconResponse, error) {
	req := c.c.NewRequest(c.name, "IconService.List", in)
	out := new(IconResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IconService service

type IconServiceHandler interface {
	// 公共图标新增
	Create(context.Context, *Icon, *IconResponse) error
	// 公共图标修改
	Update(context.Context, *Icon, *IconResponse) error
	// 公共图标删除
	Delete(context.Context, *Icon, *IconResponse) error
	// 公共图标获取
	Get(context.Context, *Icon, *IconResponse) error
	// 公共图标查询
	Search(context.Context, *IconRequest, *IconResponse) error
	// 公共图标列表
	List(context.Context, *IconRequest, *IconResponse) error
}

func RegisterIconServiceHandler(s server.Server, hdlr IconServiceHandler, opts ...server.HandlerOption) error {
	type iconService interface {
		Create(ctx context.Context, in *Icon, out *IconResponse) error
		Update(ctx context.Context, in *Icon, out *IconResponse) error
		Delete(ctx context.Context, in *Icon, out *IconResponse) error
		Get(ctx context.Context, in *Icon, out *IconResponse) error
		Search(ctx context.Context, in *IconRequest, out *IconResponse) error
		List(ctx context.Context, in *IconRequest, out *IconResponse) error
	}
	type IconService struct {
		iconService
	}
	h := &iconServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&IconService{h}, opts...))
}

type iconServiceHandler struct {
	IconServiceHandler
}

func (h *iconServiceHandler) Create(ctx context.Context, in *Icon, out *IconResponse) error {
	return h.IconServiceHandler.Create(ctx, in, out)
}

func (h *iconServiceHandler) Update(ctx context.Context, in *Icon, out *IconResponse) error {
	return h.IconServiceHandler.Update(ctx, in, out)
}

func (h *iconServiceHandler) Delete(ctx context.Context, in *Icon, out *IconResponse) error {
	return h.IconServiceHandler.Delete(ctx, in, out)
}

func (h *iconServiceHandler) Get(ctx context.Context, in *Icon, out *IconResponse) error {
	return h.IconServiceHandler.Get(ctx, in, out)
}

func (h *iconServiceHandler) Search(ctx context.Context, in *IconRequest, out *IconResponse) error {
	return h.IconServiceHandler.Search(ctx, in, out)
}

func (h *iconServiceHandler) List(ctx context.Context, in *IconRequest, out *IconResponse) error {
	return h.IconServiceHandler.List(ctx, in, out)
}
