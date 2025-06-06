// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: integralService.proto

package services

import (
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

// Api Endpoints for IntegralService service

func NewIntegralServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IntegralService service

type IntegralService interface {
	// 积分设置修改
	Save(ctx context.Context, in *Integral, opts ...client.CallOption) (*IntegralResponse, error)
	// 积分设置获取
	Get(ctx context.Context, in *Integral, opts ...client.CallOption) (*IntegralResponse, error)
	//增加积分值
	AddIntegralValue(ctx context.Context, in *IntegralValue, opts ...client.CallOption) (*IntegralResponse, error)
	//使用积分值
	UseIntegralValue(ctx context.Context, in *IntegralValue, opts ...client.CallOption) (*IntegralResponse, error)
	//回滚积分值
	RollbackIntegralValue(ctx context.Context, in *IntegralValue, opts ...client.CallOption) (*IntegralResponse, error)
}

type integralService struct {
	c    client.Client
	name string
}

func NewIntegralService(name string, c client.Client) IntegralService {
	return &integralService{
		c:    c,
		name: name,
	}
}

func (c *integralService) Save(ctx context.Context, in *Integral, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.Save", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) Get(ctx context.Context, in *Integral, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.Get", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) AddIntegralValue(ctx context.Context, in *IntegralValue, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.AddIntegralValue", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) UseIntegralValue(ctx context.Context, in *IntegralValue, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.UseIntegralValue", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) RollbackIntegralValue(ctx context.Context, in *IntegralValue, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.RollbackIntegralValue", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IntegralService service

type IntegralServiceHandler interface {
	// 积分设置修改
	Save(context.Context, *Integral, *IntegralResponse) error
	// 积分设置获取
	Get(context.Context, *Integral, *IntegralResponse) error
	//增加积分值
	AddIntegralValue(context.Context, *IntegralValue, *IntegralResponse) error
	//使用积分值
	UseIntegralValue(context.Context, *IntegralValue, *IntegralResponse) error
	//回滚积分值
	RollbackIntegralValue(context.Context, *IntegralValue, *IntegralResponse) error
}

func RegisterIntegralServiceHandler(s server.Server, hdlr IntegralServiceHandler, opts ...server.HandlerOption) error {
	type integralService interface {
		Save(ctx context.Context, in *Integral, out *IntegralResponse) error
		Get(ctx context.Context, in *Integral, out *IntegralResponse) error
		AddIntegralValue(ctx context.Context, in *IntegralValue, out *IntegralResponse) error
		UseIntegralValue(ctx context.Context, in *IntegralValue, out *IntegralResponse) error
		RollbackIntegralValue(ctx context.Context, in *IntegralValue, out *IntegralResponse) error
	}
	type IntegralService struct {
		integralService
	}
	h := &integralServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&IntegralService{h}, opts...))
}

type integralServiceHandler struct {
	IntegralServiceHandler
}

func (h *integralServiceHandler) Save(ctx context.Context, in *Integral, out *IntegralResponse) error {
	return h.IntegralServiceHandler.Save(ctx, in, out)
}

func (h *integralServiceHandler) Get(ctx context.Context, in *Integral, out *IntegralResponse) error {
	return h.IntegralServiceHandler.Get(ctx, in, out)
}

func (h *integralServiceHandler) AddIntegralValue(ctx context.Context, in *IntegralValue, out *IntegralResponse) error {
	return h.IntegralServiceHandler.AddIntegralValue(ctx, in, out)
}

func (h *integralServiceHandler) UseIntegralValue(ctx context.Context, in *IntegralValue, out *IntegralResponse) error {
	return h.IntegralServiceHandler.UseIntegralValue(ctx, in, out)
}

func (h *integralServiceHandler) RollbackIntegralValue(ctx context.Context, in *IntegralValue, out *IntegralResponse) error {
	return h.IntegralServiceHandler.RollbackIntegralValue(ctx, in, out)
}
