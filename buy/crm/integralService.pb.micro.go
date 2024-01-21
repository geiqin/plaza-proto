// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: integralService.proto

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

// Api Endpoints for IntegralService service

func NewIntegralServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IntegralService service

type IntegralService interface {
	//积分增加【有效积分增加】
	IntegralAdd(ctx context.Context, in *IntegralInfo, opts ...client.CallOption) (*IntegralResponse, error)
	//积分使用【有效积分减少】
	IntegralUse(ctx context.Context, in *IntegralInfo, opts ...client.CallOption) (*IntegralResponse, error)
	//锁定积分增加
	LockingIntegralInsert(ctx context.Context, in *LockingIntegralInfo, opts ...client.CallOption) (*LockingIntegralResponse, error)
	//锁定积分释放
	LockingIntegralRollback(ctx context.Context, in *LockingIntegralRollbackInfo, opts ...client.CallOption) (*LockingIntegralResponse, error)
	//锁定积分生效
	LockingIntegralEffect(ctx context.Context, in *LockingIntegralEffectInfo, opts ...client.CallOption) (*LockingIntegralResponse, error)
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

func (c *integralService) IntegralAdd(ctx context.Context, in *IntegralInfo, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.IntegralAdd", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) IntegralUse(ctx context.Context, in *IntegralInfo, opts ...client.CallOption) (*IntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.IntegralUse", in)
	out := new(IntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) LockingIntegralInsert(ctx context.Context, in *LockingIntegralInfo, opts ...client.CallOption) (*LockingIntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.LockingIntegralInsert", in)
	out := new(LockingIntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) LockingIntegralRollback(ctx context.Context, in *LockingIntegralRollbackInfo, opts ...client.CallOption) (*LockingIntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.LockingIntegralRollback", in)
	out := new(LockingIntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralService) LockingIntegralEffect(ctx context.Context, in *LockingIntegralEffectInfo, opts ...client.CallOption) (*LockingIntegralResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralService.LockingIntegralEffect", in)
	out := new(LockingIntegralResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IntegralService service

type IntegralServiceHandler interface {
	//积分增加【有效积分增加】
	IntegralAdd(context.Context, *IntegralInfo, *IntegralResponse) error
	//积分使用【有效积分减少】
	IntegralUse(context.Context, *IntegralInfo, *IntegralResponse) error
	//锁定积分增加
	LockingIntegralInsert(context.Context, *LockingIntegralInfo, *LockingIntegralResponse) error
	//锁定积分释放
	LockingIntegralRollback(context.Context, *LockingIntegralRollbackInfo, *LockingIntegralResponse) error
	//锁定积分生效
	LockingIntegralEffect(context.Context, *LockingIntegralEffectInfo, *LockingIntegralResponse) error
}

func RegisterIntegralServiceHandler(s server.Server, hdlr IntegralServiceHandler, opts ...server.HandlerOption) error {
	type integralService interface {
		IntegralAdd(ctx context.Context, in *IntegralInfo, out *IntegralResponse) error
		IntegralUse(ctx context.Context, in *IntegralInfo, out *IntegralResponse) error
		LockingIntegralInsert(ctx context.Context, in *LockingIntegralInfo, out *LockingIntegralResponse) error
		LockingIntegralRollback(ctx context.Context, in *LockingIntegralRollbackInfo, out *LockingIntegralResponse) error
		LockingIntegralEffect(ctx context.Context, in *LockingIntegralEffectInfo, out *LockingIntegralResponse) error
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

func (h *integralServiceHandler) IntegralAdd(ctx context.Context, in *IntegralInfo, out *IntegralResponse) error {
	return h.IntegralServiceHandler.IntegralAdd(ctx, in, out)
}

func (h *integralServiceHandler) IntegralUse(ctx context.Context, in *IntegralInfo, out *IntegralResponse) error {
	return h.IntegralServiceHandler.IntegralUse(ctx, in, out)
}

func (h *integralServiceHandler) LockingIntegralInsert(ctx context.Context, in *LockingIntegralInfo, out *LockingIntegralResponse) error {
	return h.IntegralServiceHandler.LockingIntegralInsert(ctx, in, out)
}

func (h *integralServiceHandler) LockingIntegralRollback(ctx context.Context, in *LockingIntegralRollbackInfo, out *LockingIntegralResponse) error {
	return h.IntegralServiceHandler.LockingIntegralRollback(ctx, in, out)
}

func (h *integralServiceHandler) LockingIntegralEffect(ctx context.Context, in *LockingIntegralEffectInfo, out *LockingIntegralResponse) error {
	return h.IntegralServiceHandler.LockingIntegralEffect(ctx, in, out)
}
