// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: versionPackageService.proto

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

// Api Endpoints for VersionPackageService service

func NewVersionPackageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VersionPackageService service

type VersionPackageService interface {
	// 版本套餐新增
	Create(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error)
	// 版本套餐修改
	Update(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error)
	// 版本套餐删除
	Delete(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error)
	// 版本套餐获取
	Get(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error)
	// 版本套餐查询
	Search(ctx context.Context, in *VersionPackageRequest, opts ...client.CallOption) (*VersionPackageResponse, error)
	// 版本套餐列表
	List(ctx context.Context, in *VersionPackageRequest, opts ...client.CallOption) (*VersionPackageResponse, error)
}

type versionPackageService struct {
	c    client.Client
	name string
}

func NewVersionPackageService(name string, c client.Client) VersionPackageService {
	return &versionPackageService{
		c:    c,
		name: name,
	}
}

func (c *versionPackageService) Create(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error) {
	req := c.c.NewRequest(c.name, "VersionPackageService.Create", in)
	out := new(VersionPackageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionPackageService) Update(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error) {
	req := c.c.NewRequest(c.name, "VersionPackageService.Update", in)
	out := new(VersionPackageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionPackageService) Delete(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error) {
	req := c.c.NewRequest(c.name, "VersionPackageService.Delete", in)
	out := new(VersionPackageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionPackageService) Get(ctx context.Context, in *VersionPackage, opts ...client.CallOption) (*VersionPackageResponse, error) {
	req := c.c.NewRequest(c.name, "VersionPackageService.Get", in)
	out := new(VersionPackageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionPackageService) Search(ctx context.Context, in *VersionPackageRequest, opts ...client.CallOption) (*VersionPackageResponse, error) {
	req := c.c.NewRequest(c.name, "VersionPackageService.Search", in)
	out := new(VersionPackageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionPackageService) List(ctx context.Context, in *VersionPackageRequest, opts ...client.CallOption) (*VersionPackageResponse, error) {
	req := c.c.NewRequest(c.name, "VersionPackageService.List", in)
	out := new(VersionPackageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VersionPackageService service

type VersionPackageServiceHandler interface {
	// 版本套餐新增
	Create(context.Context, *VersionPackage, *VersionPackageResponse) error
	// 版本套餐修改
	Update(context.Context, *VersionPackage, *VersionPackageResponse) error
	// 版本套餐删除
	Delete(context.Context, *VersionPackage, *VersionPackageResponse) error
	// 版本套餐获取
	Get(context.Context, *VersionPackage, *VersionPackageResponse) error
	// 版本套餐查询
	Search(context.Context, *VersionPackageRequest, *VersionPackageResponse) error
	// 版本套餐列表
	List(context.Context, *VersionPackageRequest, *VersionPackageResponse) error
}

func RegisterVersionPackageServiceHandler(s server.Server, hdlr VersionPackageServiceHandler, opts ...server.HandlerOption) error {
	type versionPackageService interface {
		Create(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error
		Update(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error
		Delete(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error
		Get(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error
		Search(ctx context.Context, in *VersionPackageRequest, out *VersionPackageResponse) error
		List(ctx context.Context, in *VersionPackageRequest, out *VersionPackageResponse) error
	}
	type VersionPackageService struct {
		versionPackageService
	}
	h := &versionPackageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VersionPackageService{h}, opts...))
}

type versionPackageServiceHandler struct {
	VersionPackageServiceHandler
}

func (h *versionPackageServiceHandler) Create(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error {
	return h.VersionPackageServiceHandler.Create(ctx, in, out)
}

func (h *versionPackageServiceHandler) Update(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error {
	return h.VersionPackageServiceHandler.Update(ctx, in, out)
}

func (h *versionPackageServiceHandler) Delete(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error {
	return h.VersionPackageServiceHandler.Delete(ctx, in, out)
}

func (h *versionPackageServiceHandler) Get(ctx context.Context, in *VersionPackage, out *VersionPackageResponse) error {
	return h.VersionPackageServiceHandler.Get(ctx, in, out)
}

func (h *versionPackageServiceHandler) Search(ctx context.Context, in *VersionPackageRequest, out *VersionPackageResponse) error {
	return h.VersionPackageServiceHandler.Search(ctx, in, out)
}

func (h *versionPackageServiceHandler) List(ctx context.Context, in *VersionPackageRequest, out *VersionPackageResponse) error {
	return h.VersionPackageServiceHandler.List(ctx, in, out)
}
