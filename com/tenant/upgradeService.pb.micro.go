// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: upgradeService.proto

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

// Api Endpoints for UpgradeService service

func NewUpgradeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UpgradeService service

type UpgradeService interface {
	//升级店铺
	Store(ctx context.Context, in *UpgradeRequest, opts ...client.CallOption) (*UpgradeResponse, error)
	//升级平台库
	Master(ctx context.Context, in *UpgradeRequest, opts ...client.CallOption) (*UpgradeResponse, error)
	//重置数据
	Reset(ctx context.Context, in *UpgradeRequest, opts ...client.CallOption) (*UpgradeResponse, error)
}

type upgradeService struct {
	c    client.Client
	name string
}

func NewUpgradeService(name string, c client.Client) UpgradeService {
	return &upgradeService{
		c:    c,
		name: name,
	}
}

func (c *upgradeService) Store(ctx context.Context, in *UpgradeRequest, opts ...client.CallOption) (*UpgradeResponse, error) {
	req := c.c.NewRequest(c.name, "UpgradeService.Store", in)
	out := new(UpgradeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upgradeService) Master(ctx context.Context, in *UpgradeRequest, opts ...client.CallOption) (*UpgradeResponse, error) {
	req := c.c.NewRequest(c.name, "UpgradeService.Master", in)
	out := new(UpgradeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upgradeService) Reset(ctx context.Context, in *UpgradeRequest, opts ...client.CallOption) (*UpgradeResponse, error) {
	req := c.c.NewRequest(c.name, "UpgradeService.Reset", in)
	out := new(UpgradeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UpgradeService service

type UpgradeServiceHandler interface {
	//升级店铺
	Store(context.Context, *UpgradeRequest, *UpgradeResponse) error
	//升级平台库
	Master(context.Context, *UpgradeRequest, *UpgradeResponse) error
	//重置数据
	Reset(context.Context, *UpgradeRequest, *UpgradeResponse) error
}

func RegisterUpgradeServiceHandler(s server.Server, hdlr UpgradeServiceHandler, opts ...server.HandlerOption) error {
	type upgradeService interface {
		Store(ctx context.Context, in *UpgradeRequest, out *UpgradeResponse) error
		Master(ctx context.Context, in *UpgradeRequest, out *UpgradeResponse) error
		Reset(ctx context.Context, in *UpgradeRequest, out *UpgradeResponse) error
	}
	type UpgradeService struct {
		upgradeService
	}
	h := &upgradeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UpgradeService{h}, opts...))
}

type upgradeServiceHandler struct {
	UpgradeServiceHandler
}

func (h *upgradeServiceHandler) Store(ctx context.Context, in *UpgradeRequest, out *UpgradeResponse) error {
	return h.UpgradeServiceHandler.Store(ctx, in, out)
}

func (h *upgradeServiceHandler) Master(ctx context.Context, in *UpgradeRequest, out *UpgradeResponse) error {
	return h.UpgradeServiceHandler.Master(ctx, in, out)
}

func (h *upgradeServiceHandler) Reset(ctx context.Context, in *UpgradeRequest, out *UpgradeResponse) error {
	return h.UpgradeServiceHandler.Reset(ctx, in, out)
}
