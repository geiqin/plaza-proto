// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: linkAddressService.proto

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

// Api Endpoints for LinkAddressService service

func NewLinkAddressServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LinkAddressService service

type LinkAddressService interface {
	// 页面链接新增
	Create(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error)
	// 页面链接修改
	Update(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error)
	// 页面链接删除
	Delete(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error)
	// 页面链接获取
	Get(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error)
	// 页面链接查询
	Search(ctx context.Context, in *LinkAddressRequest, opts ...client.CallOption) (*LinkAddressResponse, error)
	// 页面链接列表
	List(ctx context.Context, in *LinkAddressRequest, opts ...client.CallOption) (*LinkAddressResponse, error)
}

type linkAddressService struct {
	c    client.Client
	name string
}

func NewLinkAddressService(name string, c client.Client) LinkAddressService {
	return &linkAddressService{
		c:    c,
		name: name,
	}
}

func (c *linkAddressService) Create(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error) {
	req := c.c.NewRequest(c.name, "LinkAddressService.Create", in)
	out := new(LinkAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAddressService) Update(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error) {
	req := c.c.NewRequest(c.name, "LinkAddressService.Update", in)
	out := new(LinkAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAddressService) Delete(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error) {
	req := c.c.NewRequest(c.name, "LinkAddressService.Delete", in)
	out := new(LinkAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAddressService) Get(ctx context.Context, in *LinkAddress, opts ...client.CallOption) (*LinkAddressResponse, error) {
	req := c.c.NewRequest(c.name, "LinkAddressService.Get", in)
	out := new(LinkAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAddressService) Search(ctx context.Context, in *LinkAddressRequest, opts ...client.CallOption) (*LinkAddressResponse, error) {
	req := c.c.NewRequest(c.name, "LinkAddressService.Search", in)
	out := new(LinkAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAddressService) List(ctx context.Context, in *LinkAddressRequest, opts ...client.CallOption) (*LinkAddressResponse, error) {
	req := c.c.NewRequest(c.name, "LinkAddressService.List", in)
	out := new(LinkAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LinkAddressService service

type LinkAddressServiceHandler interface {
	// 页面链接新增
	Create(context.Context, *LinkAddress, *LinkAddressResponse) error
	// 页面链接修改
	Update(context.Context, *LinkAddress, *LinkAddressResponse) error
	// 页面链接删除
	Delete(context.Context, *LinkAddress, *LinkAddressResponse) error
	// 页面链接获取
	Get(context.Context, *LinkAddress, *LinkAddressResponse) error
	// 页面链接查询
	Search(context.Context, *LinkAddressRequest, *LinkAddressResponse) error
	// 页面链接列表
	List(context.Context, *LinkAddressRequest, *LinkAddressResponse) error
}

func RegisterLinkAddressServiceHandler(s server.Server, hdlr LinkAddressServiceHandler, opts ...server.HandlerOption) error {
	type linkAddressService interface {
		Create(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error
		Update(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error
		Delete(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error
		Get(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error
		Search(ctx context.Context, in *LinkAddressRequest, out *LinkAddressResponse) error
		List(ctx context.Context, in *LinkAddressRequest, out *LinkAddressResponse) error
	}
	type LinkAddressService struct {
		linkAddressService
	}
	h := &linkAddressServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LinkAddressService{h}, opts...))
}

type linkAddressServiceHandler struct {
	LinkAddressServiceHandler
}

func (h *linkAddressServiceHandler) Create(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error {
	return h.LinkAddressServiceHandler.Create(ctx, in, out)
}

func (h *linkAddressServiceHandler) Update(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error {
	return h.LinkAddressServiceHandler.Update(ctx, in, out)
}

func (h *linkAddressServiceHandler) Delete(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error {
	return h.LinkAddressServiceHandler.Delete(ctx, in, out)
}

func (h *linkAddressServiceHandler) Get(ctx context.Context, in *LinkAddress, out *LinkAddressResponse) error {
	return h.LinkAddressServiceHandler.Get(ctx, in, out)
}

func (h *linkAddressServiceHandler) Search(ctx context.Context, in *LinkAddressRequest, out *LinkAddressResponse) error {
	return h.LinkAddressServiceHandler.Search(ctx, in, out)
}

func (h *linkAddressServiceHandler) List(ctx context.Context, in *LinkAddressRequest, out *LinkAddressResponse) error {
	return h.LinkAddressServiceHandler.List(ctx, in, out)
}
