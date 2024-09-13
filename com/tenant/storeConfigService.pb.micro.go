// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: storeConfigService.proto

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

// Api Endpoints for StoreConfigService service

func NewStoreConfigServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StoreConfigService service

type StoreConfigService interface {
	// 店铺配置新增
	Create(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error)
	// 店铺配置修改
	Update(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error)
	// 店铺配置删除
	Delete(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error)
	// 店铺配置获取
	Get(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error)
	// 店铺配置查询
	Search(ctx context.Context, in *StoreConfigRequest, opts ...client.CallOption) (*StoreConfigResponse, error)
	// 店铺配置列表
	List(ctx context.Context, in *StoreConfigRequest, opts ...client.CallOption) (*StoreConfigResponse, error)
}

type storeConfigService struct {
	c    client.Client
	name string
}

func NewStoreConfigService(name string, c client.Client) StoreConfigService {
	return &storeConfigService{
		c:    c,
		name: name,
	}
}

func (c *storeConfigService) Create(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error) {
	req := c.c.NewRequest(c.name, "StoreConfigService.Create", in)
	out := new(StoreConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeConfigService) Update(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error) {
	req := c.c.NewRequest(c.name, "StoreConfigService.Update", in)
	out := new(StoreConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeConfigService) Delete(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error) {
	req := c.c.NewRequest(c.name, "StoreConfigService.Delete", in)
	out := new(StoreConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeConfigService) Get(ctx context.Context, in *StoreConfig, opts ...client.CallOption) (*StoreConfigResponse, error) {
	req := c.c.NewRequest(c.name, "StoreConfigService.Get", in)
	out := new(StoreConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeConfigService) Search(ctx context.Context, in *StoreConfigRequest, opts ...client.CallOption) (*StoreConfigResponse, error) {
	req := c.c.NewRequest(c.name, "StoreConfigService.Search", in)
	out := new(StoreConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeConfigService) List(ctx context.Context, in *StoreConfigRequest, opts ...client.CallOption) (*StoreConfigResponse, error) {
	req := c.c.NewRequest(c.name, "StoreConfigService.List", in)
	out := new(StoreConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoreConfigService service

type StoreConfigServiceHandler interface {
	// 店铺配置新增
	Create(context.Context, *StoreConfig, *StoreConfigResponse) error
	// 店铺配置修改
	Update(context.Context, *StoreConfig, *StoreConfigResponse) error
	// 店铺配置删除
	Delete(context.Context, *StoreConfig, *StoreConfigResponse) error
	// 店铺配置获取
	Get(context.Context, *StoreConfig, *StoreConfigResponse) error
	// 店铺配置查询
	Search(context.Context, *StoreConfigRequest, *StoreConfigResponse) error
	// 店铺配置列表
	List(context.Context, *StoreConfigRequest, *StoreConfigResponse) error
}

func RegisterStoreConfigServiceHandler(s server.Server, hdlr StoreConfigServiceHandler, opts ...server.HandlerOption) error {
	type storeConfigService interface {
		Create(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error
		Update(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error
		Delete(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error
		Get(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error
		Search(ctx context.Context, in *StoreConfigRequest, out *StoreConfigResponse) error
		List(ctx context.Context, in *StoreConfigRequest, out *StoreConfigResponse) error
	}
	type StoreConfigService struct {
		storeConfigService
	}
	h := &storeConfigServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StoreConfigService{h}, opts...))
}

type storeConfigServiceHandler struct {
	StoreConfigServiceHandler
}

func (h *storeConfigServiceHandler) Create(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error {
	return h.StoreConfigServiceHandler.Create(ctx, in, out)
}

func (h *storeConfigServiceHandler) Update(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error {
	return h.StoreConfigServiceHandler.Update(ctx, in, out)
}

func (h *storeConfigServiceHandler) Delete(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error {
	return h.StoreConfigServiceHandler.Delete(ctx, in, out)
}

func (h *storeConfigServiceHandler) Get(ctx context.Context, in *StoreConfig, out *StoreConfigResponse) error {
	return h.StoreConfigServiceHandler.Get(ctx, in, out)
}

func (h *storeConfigServiceHandler) Search(ctx context.Context, in *StoreConfigRequest, out *StoreConfigResponse) error {
	return h.StoreConfigServiceHandler.Search(ctx, in, out)
}

func (h *storeConfigServiceHandler) List(ctx context.Context, in *StoreConfigRequest, out *StoreConfigResponse) error {
	return h.StoreConfigServiceHandler.List(ctx, in, out)
}
