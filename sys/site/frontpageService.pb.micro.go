// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: frontpageService.proto

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

// Api Endpoints for FrontpageService service

func NewFrontpageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FrontpageService service

type FrontpageService interface {
	//前端初始化
	Init(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageInitResponse, error)
	//首页数据
	Home(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageHomeResponse, error)
	//搜索页数据
	SearchIndex(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageResponse, error)
	//个人中心数据
	UserCenter(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageUserCenterResponse, error)
}

type frontpageService struct {
	c    client.Client
	name string
}

func NewFrontpageService(name string, c client.Client) FrontpageService {
	return &frontpageService{
		c:    c,
		name: name,
	}
}

func (c *frontpageService) Init(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageInitResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.Init", in)
	out := new(FrontpageInitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) Home(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageHomeResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.Home", in)
	out := new(FrontpageHomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) SearchIndex(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.SearchIndex", in)
	out := new(FrontpageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontpageService) UserCenter(ctx context.Context, in *FrontpageRequest, opts ...client.CallOption) (*FrontpageUserCenterResponse, error) {
	req := c.c.NewRequest(c.name, "FrontpageService.UserCenter", in)
	out := new(FrontpageUserCenterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontpageService service

type FrontpageServiceHandler interface {
	//前端初始化
	Init(context.Context, *FrontpageRequest, *FrontpageInitResponse) error
	//首页数据
	Home(context.Context, *FrontpageRequest, *FrontpageHomeResponse) error
	//搜索页数据
	SearchIndex(context.Context, *FrontpageRequest, *FrontpageResponse) error
	//个人中心数据
	UserCenter(context.Context, *FrontpageRequest, *FrontpageUserCenterResponse) error
}

func RegisterFrontpageServiceHandler(s server.Server, hdlr FrontpageServiceHandler, opts ...server.HandlerOption) error {
	type frontpageService interface {
		Init(ctx context.Context, in *FrontpageRequest, out *FrontpageInitResponse) error
		Home(ctx context.Context, in *FrontpageRequest, out *FrontpageHomeResponse) error
		SearchIndex(ctx context.Context, in *FrontpageRequest, out *FrontpageResponse) error
		UserCenter(ctx context.Context, in *FrontpageRequest, out *FrontpageUserCenterResponse) error
	}
	type FrontpageService struct {
		frontpageService
	}
	h := &frontpageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FrontpageService{h}, opts...))
}

type frontpageServiceHandler struct {
	FrontpageServiceHandler
}

func (h *frontpageServiceHandler) Init(ctx context.Context, in *FrontpageRequest, out *FrontpageInitResponse) error {
	return h.FrontpageServiceHandler.Init(ctx, in, out)
}

func (h *frontpageServiceHandler) Home(ctx context.Context, in *FrontpageRequest, out *FrontpageHomeResponse) error {
	return h.FrontpageServiceHandler.Home(ctx, in, out)
}

func (h *frontpageServiceHandler) SearchIndex(ctx context.Context, in *FrontpageRequest, out *FrontpageResponse) error {
	return h.FrontpageServiceHandler.SearchIndex(ctx, in, out)
}

func (h *frontpageServiceHandler) UserCenter(ctx context.Context, in *FrontpageRequest, out *FrontpageUserCenterResponse) error {
	return h.FrontpageServiceHandler.UserCenter(ctx, in, out)
}
