// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: storageService.proto

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

// Api Endpoints for StorageService service

func NewStorageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StorageService service

type StorageService interface {
	// 存储新增
	Create(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error)
	// 存储修改
	Update(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error)
	// 存储删除
	Delete(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error)
	// 存储获取
	Get(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error)
	// 存储查询
	Search(ctx context.Context, in *StorageRequest, opts ...client.CallOption) (*StorageResponse, error)
	// 存储列表
	List(ctx context.Context, in *StorageRequest, opts ...client.CallOption) (*StorageResponse, error)
}

type storageService struct {
	c    client.Client
	name string
}

func NewStorageService(name string, c client.Client) StorageService {
	return &storageService{
		c:    c,
		name: name,
	}
}

func (c *storageService) Create(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error) {
	req := c.c.NewRequest(c.name, "StorageService.Create", in)
	out := new(StorageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) Update(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error) {
	req := c.c.NewRequest(c.name, "StorageService.Update", in)
	out := new(StorageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) Delete(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error) {
	req := c.c.NewRequest(c.name, "StorageService.Delete", in)
	out := new(StorageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) Get(ctx context.Context, in *Storage, opts ...client.CallOption) (*StorageResponse, error) {
	req := c.c.NewRequest(c.name, "StorageService.Get", in)
	out := new(StorageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) Search(ctx context.Context, in *StorageRequest, opts ...client.CallOption) (*StorageResponse, error) {
	req := c.c.NewRequest(c.name, "StorageService.Search", in)
	out := new(StorageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageService) List(ctx context.Context, in *StorageRequest, opts ...client.CallOption) (*StorageResponse, error) {
	req := c.c.NewRequest(c.name, "StorageService.List", in)
	out := new(StorageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StorageService service

type StorageServiceHandler interface {
	// 存储新增
	Create(context.Context, *Storage, *StorageResponse) error
	// 存储修改
	Update(context.Context, *Storage, *StorageResponse) error
	// 存储删除
	Delete(context.Context, *Storage, *StorageResponse) error
	// 存储获取
	Get(context.Context, *Storage, *StorageResponse) error
	// 存储查询
	Search(context.Context, *StorageRequest, *StorageResponse) error
	// 存储列表
	List(context.Context, *StorageRequest, *StorageResponse) error
}

func RegisterStorageServiceHandler(s server.Server, hdlr StorageServiceHandler, opts ...server.HandlerOption) error {
	type storageService interface {
		Create(ctx context.Context, in *Storage, out *StorageResponse) error
		Update(ctx context.Context, in *Storage, out *StorageResponse) error
		Delete(ctx context.Context, in *Storage, out *StorageResponse) error
		Get(ctx context.Context, in *Storage, out *StorageResponse) error
		Search(ctx context.Context, in *StorageRequest, out *StorageResponse) error
		List(ctx context.Context, in *StorageRequest, out *StorageResponse) error
	}
	type StorageService struct {
		storageService
	}
	h := &storageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StorageService{h}, opts...))
}

type storageServiceHandler struct {
	StorageServiceHandler
}

func (h *storageServiceHandler) Create(ctx context.Context, in *Storage, out *StorageResponse) error {
	return h.StorageServiceHandler.Create(ctx, in, out)
}

func (h *storageServiceHandler) Update(ctx context.Context, in *Storage, out *StorageResponse) error {
	return h.StorageServiceHandler.Update(ctx, in, out)
}

func (h *storageServiceHandler) Delete(ctx context.Context, in *Storage, out *StorageResponse) error {
	return h.StorageServiceHandler.Delete(ctx, in, out)
}

func (h *storageServiceHandler) Get(ctx context.Context, in *Storage, out *StorageResponse) error {
	return h.StorageServiceHandler.Get(ctx, in, out)
}

func (h *storageServiceHandler) Search(ctx context.Context, in *StorageRequest, out *StorageResponse) error {
	return h.StorageServiceHandler.Search(ctx, in, out)
}

func (h *storageServiceHandler) List(ctx context.Context, in *StorageRequest, out *StorageResponse) error {
	return h.StorageServiceHandler.List(ctx, in, out)
}
