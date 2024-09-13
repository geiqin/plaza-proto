// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: storePluginService.proto

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

// Api Endpoints for StorePluginService service

func NewStorePluginServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StorePluginService service

type StorePluginService interface {
	//店铺应用安装
	Install(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error)
	//店铺应用升级
	Upgrade(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error)
	//店铺应用移除
	Remove(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error)
	//店铺应用开关（开启/关闭）
	Switch(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error)
	//店铺应用详情信息
	Detail(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error)
	//店铺应用基本信息
	Base(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error)
	//店铺应用列表
	List(ctx context.Context, in *StorePluginRequest, opts ...client.CallOption) (*StorePluginResponse, error)
	//店铺应用查询
	Search(ctx context.Context, in *StorePluginRequest, opts ...client.CallOption) (*StorePluginResponse, error)
	//应用配置保存
	ConfigSave(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error)
	//有效应用名称列表
	PluginsValidCodes(ctx context.Context, in *StorePluginRequest, opts ...client.CallOption) (*StorePluginResponse, error)
}

type storePluginService struct {
	c    client.Client
	name string
}

func NewStorePluginService(name string, c client.Client) StorePluginService {
	return &storePluginService{
		c:    c,
		name: name,
	}
}

func (c *storePluginService) Install(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.Install", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) Upgrade(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.Upgrade", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) Remove(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.Remove", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) Switch(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.Switch", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) Detail(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.Detail", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) Base(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.Base", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) List(ctx context.Context, in *StorePluginRequest, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.List", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) Search(ctx context.Context, in *StorePluginRequest, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.Search", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) ConfigSave(ctx context.Context, in *StorePlugin, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.ConfigSave", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storePluginService) PluginsValidCodes(ctx context.Context, in *StorePluginRequest, opts ...client.CallOption) (*StorePluginResponse, error) {
	req := c.c.NewRequest(c.name, "StorePluginService.PluginsValidCodes", in)
	out := new(StorePluginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StorePluginService service

type StorePluginServiceHandler interface {
	//店铺应用安装
	Install(context.Context, *StorePlugin, *StorePluginResponse) error
	//店铺应用升级
	Upgrade(context.Context, *StorePlugin, *StorePluginResponse) error
	//店铺应用移除
	Remove(context.Context, *StorePlugin, *StorePluginResponse) error
	//店铺应用开关（开启/关闭）
	Switch(context.Context, *StorePlugin, *StorePluginResponse) error
	//店铺应用详情信息
	Detail(context.Context, *StorePlugin, *StorePluginResponse) error
	//店铺应用基本信息
	Base(context.Context, *StorePlugin, *StorePluginResponse) error
	//店铺应用列表
	List(context.Context, *StorePluginRequest, *StorePluginResponse) error
	//店铺应用查询
	Search(context.Context, *StorePluginRequest, *StorePluginResponse) error
	//应用配置保存
	ConfigSave(context.Context, *StorePlugin, *StorePluginResponse) error
	//有效应用名称列表
	PluginsValidCodes(context.Context, *StorePluginRequest, *StorePluginResponse) error
}

func RegisterStorePluginServiceHandler(s server.Server, hdlr StorePluginServiceHandler, opts ...server.HandlerOption) error {
	type storePluginService interface {
		Install(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error
		Upgrade(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error
		Remove(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error
		Switch(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error
		Detail(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error
		Base(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error
		List(ctx context.Context, in *StorePluginRequest, out *StorePluginResponse) error
		Search(ctx context.Context, in *StorePluginRequest, out *StorePluginResponse) error
		ConfigSave(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error
		PluginsValidCodes(ctx context.Context, in *StorePluginRequest, out *StorePluginResponse) error
	}
	type StorePluginService struct {
		storePluginService
	}
	h := &storePluginServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StorePluginService{h}, opts...))
}

type storePluginServiceHandler struct {
	StorePluginServiceHandler
}

func (h *storePluginServiceHandler) Install(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.Install(ctx, in, out)
}

func (h *storePluginServiceHandler) Upgrade(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.Upgrade(ctx, in, out)
}

func (h *storePluginServiceHandler) Remove(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.Remove(ctx, in, out)
}

func (h *storePluginServiceHandler) Switch(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.Switch(ctx, in, out)
}

func (h *storePluginServiceHandler) Detail(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.Detail(ctx, in, out)
}

func (h *storePluginServiceHandler) Base(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.Base(ctx, in, out)
}

func (h *storePluginServiceHandler) List(ctx context.Context, in *StorePluginRequest, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.List(ctx, in, out)
}

func (h *storePluginServiceHandler) Search(ctx context.Context, in *StorePluginRequest, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.Search(ctx, in, out)
}

func (h *storePluginServiceHandler) ConfigSave(ctx context.Context, in *StorePlugin, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.ConfigSave(ctx, in, out)
}

func (h *storePluginServiceHandler) PluginsValidCodes(ctx context.Context, in *StorePluginRequest, out *StorePluginResponse) error {
	return h.StorePluginServiceHandler.PluginsValidCodes(ctx, in, out)
}
