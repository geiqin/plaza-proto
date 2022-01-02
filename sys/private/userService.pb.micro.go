// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: userService.proto

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

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Create(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	UpdateStatus(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Get(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Search(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Update", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateStatus(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateStatus", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Get", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Search(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Search", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *UserRequest, *UserResponse) error
	Update(context.Context, *UserRequest, *UserResponse) error
	UpdateStatus(context.Context, *UserRequest, *UserResponse) error
	Get(context.Context, *UserRequest, *UserResponse) error
	Delete(context.Context, *UserRequest, *UserResponse) error
	Search(context.Context, *UserRequest, *UserResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Create(ctx context.Context, in *UserRequest, out *UserResponse) error
		Update(ctx context.Context, in *UserRequest, out *UserResponse) error
		UpdateStatus(ctx context.Context, in *UserRequest, out *UserResponse) error
		Get(ctx context.Context, in *UserRequest, out *UserResponse) error
		Delete(ctx context.Context, in *UserRequest, out *UserResponse) error
		Search(ctx context.Context, in *UserRequest, out *UserResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Create(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) Update(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *userServiceHandler) UpdateStatus(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *userServiceHandler) Get(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *userServiceHandler) Search(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.Search(ctx, in, out)
}
