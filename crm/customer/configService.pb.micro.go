// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: configService.proto

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

// Api Endpoints for ConfigService service

func NewConfigServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ConfigService service

type ConfigService interface {
	//设置VIP配置
	SetVip(ctx context.Context, in *VipConfig, opts ...client.CallOption) (*VipConfigResponse, error)
	//获取VIP配置
	GetVip(ctx context.Context, in *VipConfig, opts ...client.CallOption) (*VipConfigResponse, error)
	//设置钱包配置
	SetWallet(ctx context.Context, in *WalletConfig, opts ...client.CallOption) (*WalletConfigResponse, error)
	//获取钱包配置
	GetWallet(ctx context.Context, in *WalletConfig, opts ...client.CallOption) (*WalletConfigResponse, error)
}

type configService struct {
	c    client.Client
	name string
}

func NewConfigService(name string, c client.Client) ConfigService {
	return &configService{
		c:    c,
		name: name,
	}
}

func (c *configService) SetVip(ctx context.Context, in *VipConfig, opts ...client.CallOption) (*VipConfigResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigService.SetVip", in)
	out := new(VipConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) GetVip(ctx context.Context, in *VipConfig, opts ...client.CallOption) (*VipConfigResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigService.GetVip", in)
	out := new(VipConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) SetWallet(ctx context.Context, in *WalletConfig, opts ...client.CallOption) (*WalletConfigResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigService.SetWallet", in)
	out := new(WalletConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) GetWallet(ctx context.Context, in *WalletConfig, opts ...client.CallOption) (*WalletConfigResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigService.GetWallet", in)
	out := new(WalletConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigService service

type ConfigServiceHandler interface {
	//设置VIP配置
	SetVip(context.Context, *VipConfig, *VipConfigResponse) error
	//获取VIP配置
	GetVip(context.Context, *VipConfig, *VipConfigResponse) error
	//设置钱包配置
	SetWallet(context.Context, *WalletConfig, *WalletConfigResponse) error
	//获取钱包配置
	GetWallet(context.Context, *WalletConfig, *WalletConfigResponse) error
}

func RegisterConfigServiceHandler(s server.Server, hdlr ConfigServiceHandler, opts ...server.HandlerOption) error {
	type configService interface {
		SetVip(ctx context.Context, in *VipConfig, out *VipConfigResponse) error
		GetVip(ctx context.Context, in *VipConfig, out *VipConfigResponse) error
		SetWallet(ctx context.Context, in *WalletConfig, out *WalletConfigResponse) error
		GetWallet(ctx context.Context, in *WalletConfig, out *WalletConfigResponse) error
	}
	type ConfigService struct {
		configService
	}
	h := &configServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConfigService{h}, opts...))
}

type configServiceHandler struct {
	ConfigServiceHandler
}

func (h *configServiceHandler) SetVip(ctx context.Context, in *VipConfig, out *VipConfigResponse) error {
	return h.ConfigServiceHandler.SetVip(ctx, in, out)
}

func (h *configServiceHandler) GetVip(ctx context.Context, in *VipConfig, out *VipConfigResponse) error {
	return h.ConfigServiceHandler.GetVip(ctx, in, out)
}

func (h *configServiceHandler) SetWallet(ctx context.Context, in *WalletConfig, out *WalletConfigResponse) error {
	return h.ConfigServiceHandler.SetWallet(ctx, in, out)
}

func (h *configServiceHandler) GetWallet(ctx context.Context, in *WalletConfig, out *WalletConfigResponse) error {
	return h.ConfigServiceHandler.GetWallet(ctx, in, out)
}
