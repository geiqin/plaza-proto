// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authorizeService.proto

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

// Api Endpoints for AuthorizeService service

func NewAuthorizeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthorizeService service

type AuthorizeService interface {
	Info(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
	UserInfo(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
	PermissionInfo(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
	ModifyMobile(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error)
	ModifyAvatar(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error)
	ModifyPwd(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*UserResponse, error)
}

type authorizeService struct {
	c    client.Client
	name string
}

func NewAuthorizeService(name string, c client.Client) AuthorizeService {
	return &authorizeService{
		c:    c,
		name: name,
	}
}

func (c *authorizeService) Info(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Info", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) UserInfo(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.UserInfo", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) PermissionInfo(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.PermissionInfo", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) ModifyMobile(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.ModifyMobile", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) ModifyAvatar(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.ModifyAvatar", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) ModifyPwd(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.ModifyPwd", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthorizeService service

type AuthorizeServiceHandler interface {
	Info(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
	UserInfo(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
	PermissionInfo(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
	ModifyMobile(context.Context, *User, *UserResponse) error
	ModifyAvatar(context.Context, *User, *UserResponse) error
	ModifyPwd(context.Context, *AuthorizeRequest, *UserResponse) error
}

func RegisterAuthorizeServiceHandler(s server.Server, hdlr AuthorizeServiceHandler, opts ...server.HandlerOption) error {
	type authorizeService interface {
		Info(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
		UserInfo(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
		PermissionInfo(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
		ModifyMobile(ctx context.Context, in *User, out *UserResponse) error
		ModifyAvatar(ctx context.Context, in *User, out *UserResponse) error
		ModifyPwd(ctx context.Context, in *AuthorizeRequest, out *UserResponse) error
	}
	type AuthorizeService struct {
		authorizeService
	}
	h := &authorizeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthorizeService{h}, opts...))
}

type authorizeServiceHandler struct {
	AuthorizeServiceHandler
}

func (h *authorizeServiceHandler) Info(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.Info(ctx, in, out)
}

func (h *authorizeServiceHandler) UserInfo(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.UserInfo(ctx, in, out)
}

func (h *authorizeServiceHandler) PermissionInfo(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthorizeServiceHandler.PermissionInfo(ctx, in, out)
}

func (h *authorizeServiceHandler) ModifyMobile(ctx context.Context, in *User, out *UserResponse) error {
	return h.AuthorizeServiceHandler.ModifyMobile(ctx, in, out)
}

func (h *authorizeServiceHandler) ModifyAvatar(ctx context.Context, in *User, out *UserResponse) error {
	return h.AuthorizeServiceHandler.ModifyAvatar(ctx, in, out)
}

func (h *authorizeServiceHandler) ModifyPwd(ctx context.Context, in *AuthorizeRequest, out *UserResponse) error {
	return h.AuthorizeServiceHandler.ModifyPwd(ctx, in, out)
}
