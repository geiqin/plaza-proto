// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: integralSettingService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for IntegralSettingService service

func NewIntegralSettingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IntegralSettingService service

type IntegralSettingService interface {
	// 获取积分通用配置
	Get(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error)
	// 设置积分通用有效期
	SetEffective(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error)
	// 设置积分获取上限
	SetLimit(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error)
	// 设置积分保护期
	SetProtect(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error)
	// 自定义积分名称
	SetName(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error)
	// 设置积分兑换比例
	SetRatio(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error)
	// 设置积分抵现
	SetCashExchange(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error)
}

type integralSettingService struct {
	c    client.Client
	name string
}

func NewIntegralSettingService(name string, c client.Client) IntegralSettingService {
	return &integralSettingService{
		c:    c,
		name: name,
	}
}

func (c *integralSettingService) Get(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralSettingService.Get", in)
	out := new(IntegralSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralSettingService) SetEffective(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralSettingService.SetEffective", in)
	out := new(IntegralSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralSettingService) SetLimit(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralSettingService.SetLimit", in)
	out := new(IntegralSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralSettingService) SetProtect(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralSettingService.SetProtect", in)
	out := new(IntegralSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralSettingService) SetName(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralSettingService.SetName", in)
	out := new(IntegralSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralSettingService) SetRatio(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralSettingService.SetRatio", in)
	out := new(IntegralSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralSettingService) SetCashExchange(ctx context.Context, in *IntegralSetting, opts ...client.CallOption) (*IntegralSettingResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralSettingService.SetCashExchange", in)
	out := new(IntegralSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IntegralSettingService service

type IntegralSettingServiceHandler interface {
	// 获取积分通用配置
	Get(context.Context, *IntegralSetting, *IntegralSettingResponse) error
	// 设置积分通用有效期
	SetEffective(context.Context, *IntegralSetting, *IntegralSettingResponse) error
	// 设置积分获取上限
	SetLimit(context.Context, *IntegralSetting, *IntegralSettingResponse) error
	// 设置积分保护期
	SetProtect(context.Context, *IntegralSetting, *IntegralSettingResponse) error
	// 自定义积分名称
	SetName(context.Context, *IntegralSetting, *IntegralSettingResponse) error
	// 设置积分兑换比例
	SetRatio(context.Context, *IntegralSetting, *IntegralSettingResponse) error
	// 设置积分抵现
	SetCashExchange(context.Context, *IntegralSetting, *IntegralSettingResponse) error
}

func RegisterIntegralSettingServiceHandler(s server.Server, hdlr IntegralSettingServiceHandler, opts ...server.HandlerOption) error {
	type integralSettingService interface {
		Get(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error
		SetEffective(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error
		SetLimit(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error
		SetProtect(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error
		SetName(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error
		SetRatio(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error
		SetCashExchange(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error
	}
	type IntegralSettingService struct {
		integralSettingService
	}
	h := &integralSettingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&IntegralSettingService{h}, opts...))
}

type integralSettingServiceHandler struct {
	IntegralSettingServiceHandler
}

func (h *integralSettingServiceHandler) Get(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error {
	return h.IntegralSettingServiceHandler.Get(ctx, in, out)
}

func (h *integralSettingServiceHandler) SetEffective(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error {
	return h.IntegralSettingServiceHandler.SetEffective(ctx, in, out)
}

func (h *integralSettingServiceHandler) SetLimit(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error {
	return h.IntegralSettingServiceHandler.SetLimit(ctx, in, out)
}

func (h *integralSettingServiceHandler) SetProtect(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error {
	return h.IntegralSettingServiceHandler.SetProtect(ctx, in, out)
}

func (h *integralSettingServiceHandler) SetName(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error {
	return h.IntegralSettingServiceHandler.SetName(ctx, in, out)
}

func (h *integralSettingServiceHandler) SetRatio(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error {
	return h.IntegralSettingServiceHandler.SetRatio(ctx, in, out)
}

func (h *integralSettingServiceHandler) SetCashExchange(ctx context.Context, in *IntegralSetting, out *IntegralSettingResponse) error {
	return h.IntegralSettingServiceHandler.SetCashExchange(ctx, in, out)
}