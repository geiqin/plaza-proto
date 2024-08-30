// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: storeService.proto

package services

import (
	_ "../common"
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

// Api Endpoints for StoreService service

func NewStoreServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StoreService service

type StoreService interface {
	//创建店铺
	Create(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//修改店铺版本（超级管理员可用）
	UpdateVersion(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//修改店铺信息
	UpdateProfile(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//设置状态
	UpdateStatus(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//删除店铺
	Delete(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error)
	//获取店铺信息
	Get(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//根据名称获取店铺信息
	GetByName(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error)
	//查询店铺
	Search(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error)
	//获取店铺列表
	List(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error)
	//获取店铺功能模块
	GetModules(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error)
}

type storeService struct {
	c    client.Client
	name string
}

func NewStoreService(name string, c client.Client) StoreService {
	return &storeService{
		c:    c,
		name: name,
	}
}

func (c *storeService) Create(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Create", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) UpdateVersion(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.UpdateVersion", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) UpdateProfile(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.UpdateProfile", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) UpdateStatus(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.UpdateStatus", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Delete(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Delete", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Get(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Get", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) GetByName(ctx context.Context, in *Store, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.GetByName", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Search(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.Search", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) List(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.List", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) GetModules(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "StoreService.GetModules", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoreService service

type StoreServiceHandler interface {
	//创建店铺
	Create(context.Context, *Store, *StoreResponse) error
	//修改店铺版本（超级管理员可用）
	UpdateVersion(context.Context, *Store, *StoreResponse) error
	//修改店铺信息
	UpdateProfile(context.Context, *Store, *StoreResponse) error
	//设置状态
	UpdateStatus(context.Context, *Store, *StoreResponse) error
	//删除店铺
	Delete(context.Context, *StoreRequest, *StoreResponse) error
	//获取店铺信息
	Get(context.Context, *Store, *StoreResponse) error
	//根据名称获取店铺信息
	GetByName(context.Context, *Store, *StoreResponse) error
	//查询店铺
	Search(context.Context, *StoreRequest, *StoreResponse) error
	//获取店铺列表
	List(context.Context, *StoreRequest, *StoreResponse) error
	//获取店铺功能模块
	GetModules(context.Context, *StoreRequest, *StoreResponse) error
}

func RegisterStoreServiceHandler(s server.Server, hdlr StoreServiceHandler, opts ...server.HandlerOption) error {
	type storeService interface {
		Create(ctx context.Context, in *Store, out *StoreResponse) error
		UpdateVersion(ctx context.Context, in *Store, out *StoreResponse) error
		UpdateProfile(ctx context.Context, in *Store, out *StoreResponse) error
		UpdateStatus(ctx context.Context, in *Store, out *StoreResponse) error
		Delete(ctx context.Context, in *StoreRequest, out *StoreResponse) error
		Get(ctx context.Context, in *Store, out *StoreResponse) error
		GetByName(ctx context.Context, in *Store, out *StoreResponse) error
		Search(ctx context.Context, in *StoreRequest, out *StoreResponse) error
		List(ctx context.Context, in *StoreRequest, out *StoreResponse) error
		GetModules(ctx context.Context, in *StoreRequest, out *StoreResponse) error
	}
	type StoreService struct {
		storeService
	}
	h := &storeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StoreService{h}, opts...))
}

type storeServiceHandler struct {
	StoreServiceHandler
}

func (h *storeServiceHandler) Create(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Create(ctx, in, out)
}

func (h *storeServiceHandler) UpdateVersion(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.UpdateVersion(ctx, in, out)
}

func (h *storeServiceHandler) UpdateProfile(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.UpdateProfile(ctx, in, out)
}

func (h *storeServiceHandler) UpdateStatus(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *storeServiceHandler) Delete(ctx context.Context, in *StoreRequest, out *StoreResponse) error {
	return h.StoreServiceHandler.Delete(ctx, in, out)
}

func (h *storeServiceHandler) Get(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.Get(ctx, in, out)
}

func (h *storeServiceHandler) GetByName(ctx context.Context, in *Store, out *StoreResponse) error {
	return h.StoreServiceHandler.GetByName(ctx, in, out)
}

func (h *storeServiceHandler) Search(ctx context.Context, in *StoreRequest, out *StoreResponse) error {
	return h.StoreServiceHandler.Search(ctx, in, out)
}

func (h *storeServiceHandler) List(ctx context.Context, in *StoreRequest, out *StoreResponse) error {
	return h.StoreServiceHandler.List(ctx, in, out)
}

func (h *storeServiceHandler) GetModules(ctx context.Context, in *StoreRequest, out *StoreResponse) error {
	return h.StoreServiceHandler.GetModules(ctx, in, out)
}
