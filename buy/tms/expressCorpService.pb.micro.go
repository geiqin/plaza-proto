// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: expressCorpService.proto

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

// Api Endpoints for ExpressCorpService service

func NewExpressCorpServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ExpressCorpService service

type ExpressCorpService interface {
	// 快递公司新增
	Create(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error)
	// 快递公司新增添加
	AddTo(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error)
	// 快递公司修改
	Update(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error)
	// 快递公司删除
	Delete(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error)
	// 快递公司获取
	Get(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error)
	// 快递公司查询
	Search(ctx context.Context, in *ExpressCorpRequest, opts ...client.CallOption) (*ExpressCorpResponse, error)
	// 快递公司列表
	List(ctx context.Context, in *ExpressCorpRequest, opts ...client.CallOption) (*ExpressCorpResponse, error)
	// 快递公司选择器
	Selector(ctx context.Context, in *ExpressCorpRequest, opts ...client.CallOption) (*ExpressCorpResponse, error)
}

type expressCorpService struct {
	c    client.Client
	name string
}

func NewExpressCorpService(name string, c client.Client) ExpressCorpService {
	return &expressCorpService{
		c:    c,
		name: name,
	}
}

func (c *expressCorpService) Create(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.Create", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressCorpService) AddTo(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.AddTo", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressCorpService) Update(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.Update", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressCorpService) Delete(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.Delete", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressCorpService) Get(ctx context.Context, in *ExpressCorp, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.Get", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressCorpService) Search(ctx context.Context, in *ExpressCorpRequest, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.Search", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressCorpService) List(ctx context.Context, in *ExpressCorpRequest, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.List", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressCorpService) Selector(ctx context.Context, in *ExpressCorpRequest, opts ...client.CallOption) (*ExpressCorpResponse, error) {
	req := c.c.NewRequest(c.name, "ExpressCorpService.Selector", in)
	out := new(ExpressCorpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExpressCorpService service

type ExpressCorpServiceHandler interface {
	// 快递公司新增
	Create(context.Context, *ExpressCorp, *ExpressCorpResponse) error
	// 快递公司新增添加
	AddTo(context.Context, *ExpressCorp, *ExpressCorpResponse) error
	// 快递公司修改
	Update(context.Context, *ExpressCorp, *ExpressCorpResponse) error
	// 快递公司删除
	Delete(context.Context, *ExpressCorp, *ExpressCorpResponse) error
	// 快递公司获取
	Get(context.Context, *ExpressCorp, *ExpressCorpResponse) error
	// 快递公司查询
	Search(context.Context, *ExpressCorpRequest, *ExpressCorpResponse) error
	// 快递公司列表
	List(context.Context, *ExpressCorpRequest, *ExpressCorpResponse) error
	// 快递公司选择器
	Selector(context.Context, *ExpressCorpRequest, *ExpressCorpResponse) error
}

func RegisterExpressCorpServiceHandler(s server.Server, hdlr ExpressCorpServiceHandler, opts ...server.HandlerOption) error {
	type expressCorpService interface {
		Create(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error
		AddTo(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error
		Update(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error
		Delete(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error
		Get(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error
		Search(ctx context.Context, in *ExpressCorpRequest, out *ExpressCorpResponse) error
		List(ctx context.Context, in *ExpressCorpRequest, out *ExpressCorpResponse) error
		Selector(ctx context.Context, in *ExpressCorpRequest, out *ExpressCorpResponse) error
	}
	type ExpressCorpService struct {
		expressCorpService
	}
	h := &expressCorpServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ExpressCorpService{h}, opts...))
}

type expressCorpServiceHandler struct {
	ExpressCorpServiceHandler
}

func (h *expressCorpServiceHandler) Create(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.Create(ctx, in, out)
}

func (h *expressCorpServiceHandler) AddTo(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.AddTo(ctx, in, out)
}

func (h *expressCorpServiceHandler) Update(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.Update(ctx, in, out)
}

func (h *expressCorpServiceHandler) Delete(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.Delete(ctx, in, out)
}

func (h *expressCorpServiceHandler) Get(ctx context.Context, in *ExpressCorp, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.Get(ctx, in, out)
}

func (h *expressCorpServiceHandler) Search(ctx context.Context, in *ExpressCorpRequest, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.Search(ctx, in, out)
}

func (h *expressCorpServiceHandler) List(ctx context.Context, in *ExpressCorpRequest, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.List(ctx, in, out)
}

func (h *expressCorpServiceHandler) Selector(ctx context.Context, in *ExpressCorpRequest, out *ExpressCorpResponse) error {
	return h.ExpressCorpServiceHandler.Selector(ctx, in, out)
}
