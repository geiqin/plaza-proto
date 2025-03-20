// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: memberAuthService.proto

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

// Api Endpoints for MemberAuthService service

func NewMemberAuthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MemberAuthService service

type MemberAuthService interface {
	//账号登录
	AccountLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error)
	//小程序登录页面自动加载，返回用户信息的缓存key，返回是否强制绑定手机号
	MiniAuthType(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error)
	//小程序授权登录，返回token
	MiniAuthLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error)
	//手机号直接登录
	PhoneLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error)
	////小程序授权后绑定手机号
	MiniBindingPhone(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error)
	//微信公众号授权登录，返回token
	WoaAuthLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error)
	//微信公众号授权绑定手机号
	WoaAuthBindingPhone(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error)
}

type memberAuthService struct {
	c    client.Client
	name string
}

func NewMemberAuthService(name string, c client.Client) MemberAuthService {
	return &memberAuthService{
		c:    c,
		name: name,
	}
}

func (c *memberAuthService) AccountLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error) {
	req := c.c.NewRequest(c.name, "MemberAuthService.AccountLogin", in)
	out := new(MemberAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberAuthService) MiniAuthType(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error) {
	req := c.c.NewRequest(c.name, "MemberAuthService.MiniAuthType", in)
	out := new(MemberAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberAuthService) MiniAuthLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error) {
	req := c.c.NewRequest(c.name, "MemberAuthService.MiniAuthLogin", in)
	out := new(MemberAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberAuthService) PhoneLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error) {
	req := c.c.NewRequest(c.name, "MemberAuthService.PhoneLogin", in)
	out := new(MemberAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberAuthService) MiniBindingPhone(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error) {
	req := c.c.NewRequest(c.name, "MemberAuthService.MiniBindingPhone", in)
	out := new(MemberAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberAuthService) WoaAuthLogin(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error) {
	req := c.c.NewRequest(c.name, "MemberAuthService.WoaAuthLogin", in)
	out := new(MemberAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberAuthService) WoaAuthBindingPhone(ctx context.Context, in *MemberAuthRequest, opts ...client.CallOption) (*MemberAuthResponse, error) {
	req := c.c.NewRequest(c.name, "MemberAuthService.WoaAuthBindingPhone", in)
	out := new(MemberAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MemberAuthService service

type MemberAuthServiceHandler interface {
	//账号登录
	AccountLogin(context.Context, *MemberAuthRequest, *MemberAuthResponse) error
	//小程序登录页面自动加载，返回用户信息的缓存key，返回是否强制绑定手机号
	MiniAuthType(context.Context, *MemberAuthRequest, *MemberAuthResponse) error
	//小程序授权登录，返回token
	MiniAuthLogin(context.Context, *MemberAuthRequest, *MemberAuthResponse) error
	//手机号直接登录
	PhoneLogin(context.Context, *MemberAuthRequest, *MemberAuthResponse) error
	////小程序授权后绑定手机号
	MiniBindingPhone(context.Context, *MemberAuthRequest, *MemberAuthResponse) error
	//微信公众号授权登录，返回token
	WoaAuthLogin(context.Context, *MemberAuthRequest, *MemberAuthResponse) error
	//微信公众号授权绑定手机号
	WoaAuthBindingPhone(context.Context, *MemberAuthRequest, *MemberAuthResponse) error
}

func RegisterMemberAuthServiceHandler(s server.Server, hdlr MemberAuthServiceHandler, opts ...server.HandlerOption) error {
	type memberAuthService interface {
		AccountLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error
		MiniAuthType(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error
		MiniAuthLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error
		PhoneLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error
		MiniBindingPhone(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error
		WoaAuthLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error
		WoaAuthBindingPhone(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error
	}
	type MemberAuthService struct {
		memberAuthService
	}
	h := &memberAuthServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MemberAuthService{h}, opts...))
}

type memberAuthServiceHandler struct {
	MemberAuthServiceHandler
}

func (h *memberAuthServiceHandler) AccountLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error {
	return h.MemberAuthServiceHandler.AccountLogin(ctx, in, out)
}

func (h *memberAuthServiceHandler) MiniAuthType(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error {
	return h.MemberAuthServiceHandler.MiniAuthType(ctx, in, out)
}

func (h *memberAuthServiceHandler) MiniAuthLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error {
	return h.MemberAuthServiceHandler.MiniAuthLogin(ctx, in, out)
}

func (h *memberAuthServiceHandler) PhoneLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error {
	return h.MemberAuthServiceHandler.PhoneLogin(ctx, in, out)
}

func (h *memberAuthServiceHandler) MiniBindingPhone(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error {
	return h.MemberAuthServiceHandler.MiniBindingPhone(ctx, in, out)
}

func (h *memberAuthServiceHandler) WoaAuthLogin(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error {
	return h.MemberAuthServiceHandler.WoaAuthLogin(ctx, in, out)
}

func (h *memberAuthServiceHandler) WoaAuthBindingPhone(ctx context.Context, in *MemberAuthRequest, out *MemberAuthResponse) error {
	return h.MemberAuthServiceHandler.WoaAuthBindingPhone(ctx, in, out)
}
