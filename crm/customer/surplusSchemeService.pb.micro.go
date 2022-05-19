// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: surplusSchemeService.proto

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

// Api Endpoints for SurplusSchemeService service

func NewSurplusSchemeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SurplusSchemeService service

type SurplusSchemeService interface {
	//增加充值方案
	Create(ctx context.Context, in *SurplusScheme, opts ...client.CallOption) (*SurplusSchemeResponse, error)
	//修改充值方案
	Update(ctx context.Context, in *SurplusScheme, opts ...client.CallOption) (*SurplusSchemeResponse, error)
	//获得充值方案信息
	Get(ctx context.Context, in *SurplusScheme, opts ...client.CallOption) (*SurplusSchemeResponse, error)
	//查询充值方案信息
	Search(ctx context.Context, in *SurplusSchemeRequest, opts ...client.CallOption) (*SurplusSchemeResponse, error)
	//有效方案列表
	List(ctx context.Context, in *SurplusSchemeRequest, opts ...client.CallOption) (*SurplusSchemeResponse, error)
	//根据金额匹配最优方案
	Matching(ctx context.Context, in *SurplusSchemeRequest, opts ...client.CallOption) (*SurplusSchemeResponse, error)
}

type surplusSchemeService struct {
	c    client.Client
	name string
}

func NewSurplusSchemeService(name string, c client.Client) SurplusSchemeService {
	return &surplusSchemeService{
		c:    c,
		name: name,
	}
}

func (c *surplusSchemeService) Create(ctx context.Context, in *SurplusScheme, opts ...client.CallOption) (*SurplusSchemeResponse, error) {
	req := c.c.NewRequest(c.name, "SurplusSchemeService.Create", in)
	out := new(SurplusSchemeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surplusSchemeService) Update(ctx context.Context, in *SurplusScheme, opts ...client.CallOption) (*SurplusSchemeResponse, error) {
	req := c.c.NewRequest(c.name, "SurplusSchemeService.Update", in)
	out := new(SurplusSchemeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surplusSchemeService) Get(ctx context.Context, in *SurplusScheme, opts ...client.CallOption) (*SurplusSchemeResponse, error) {
	req := c.c.NewRequest(c.name, "SurplusSchemeService.Get", in)
	out := new(SurplusSchemeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surplusSchemeService) Search(ctx context.Context, in *SurplusSchemeRequest, opts ...client.CallOption) (*SurplusSchemeResponse, error) {
	req := c.c.NewRequest(c.name, "SurplusSchemeService.Search", in)
	out := new(SurplusSchemeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surplusSchemeService) List(ctx context.Context, in *SurplusSchemeRequest, opts ...client.CallOption) (*SurplusSchemeResponse, error) {
	req := c.c.NewRequest(c.name, "SurplusSchemeService.List", in)
	out := new(SurplusSchemeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surplusSchemeService) Matching(ctx context.Context, in *SurplusSchemeRequest, opts ...client.CallOption) (*SurplusSchemeResponse, error) {
	req := c.c.NewRequest(c.name, "SurplusSchemeService.Matching", in)
	out := new(SurplusSchemeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SurplusSchemeService service

type SurplusSchemeServiceHandler interface {
	//增加充值方案
	Create(context.Context, *SurplusScheme, *SurplusSchemeResponse) error
	//修改充值方案
	Update(context.Context, *SurplusScheme, *SurplusSchemeResponse) error
	//获得充值方案信息
	Get(context.Context, *SurplusScheme, *SurplusSchemeResponse) error
	//查询充值方案信息
	Search(context.Context, *SurplusSchemeRequest, *SurplusSchemeResponse) error
	//有效方案列表
	List(context.Context, *SurplusSchemeRequest, *SurplusSchemeResponse) error
	//根据金额匹配最优方案
	Matching(context.Context, *SurplusSchemeRequest, *SurplusSchemeResponse) error
}

func RegisterSurplusSchemeServiceHandler(s server.Server, hdlr SurplusSchemeServiceHandler, opts ...server.HandlerOption) error {
	type surplusSchemeService interface {
		Create(ctx context.Context, in *SurplusScheme, out *SurplusSchemeResponse) error
		Update(ctx context.Context, in *SurplusScheme, out *SurplusSchemeResponse) error
		Get(ctx context.Context, in *SurplusScheme, out *SurplusSchemeResponse) error
		Search(ctx context.Context, in *SurplusSchemeRequest, out *SurplusSchemeResponse) error
		List(ctx context.Context, in *SurplusSchemeRequest, out *SurplusSchemeResponse) error
		Matching(ctx context.Context, in *SurplusSchemeRequest, out *SurplusSchemeResponse) error
	}
	type SurplusSchemeService struct {
		surplusSchemeService
	}
	h := &surplusSchemeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SurplusSchemeService{h}, opts...))
}

type surplusSchemeServiceHandler struct {
	SurplusSchemeServiceHandler
}

func (h *surplusSchemeServiceHandler) Create(ctx context.Context, in *SurplusScheme, out *SurplusSchemeResponse) error {
	return h.SurplusSchemeServiceHandler.Create(ctx, in, out)
}

func (h *surplusSchemeServiceHandler) Update(ctx context.Context, in *SurplusScheme, out *SurplusSchemeResponse) error {
	return h.SurplusSchemeServiceHandler.Update(ctx, in, out)
}

func (h *surplusSchemeServiceHandler) Get(ctx context.Context, in *SurplusScheme, out *SurplusSchemeResponse) error {
	return h.SurplusSchemeServiceHandler.Get(ctx, in, out)
}

func (h *surplusSchemeServiceHandler) Search(ctx context.Context, in *SurplusSchemeRequest, out *SurplusSchemeResponse) error {
	return h.SurplusSchemeServiceHandler.Search(ctx, in, out)
}

func (h *surplusSchemeServiceHandler) List(ctx context.Context, in *SurplusSchemeRequest, out *SurplusSchemeResponse) error {
	return h.SurplusSchemeServiceHandler.List(ctx, in, out)
}

func (h *surplusSchemeServiceHandler) Matching(ctx context.Context, in *SurplusSchemeRequest, out *SurplusSchemeResponse) error {
	return h.SurplusSchemeServiceHandler.Matching(ctx, in, out)
}