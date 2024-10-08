// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: menuService.proto

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

// Api Endpoints for MenuService service

func NewMenuServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MenuService service

type MenuService interface {
	Get(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Create(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Update(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	UpdateHide(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Delete(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error)
	Search(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error)
	List(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error)
	Tree(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error)
	//用户菜单路由
	Routes(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error)
	//插件导航创建
	PluginMenuCreate(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error)
	//插件导航删除
	PluginMenuDelete(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error)
}

type menuService struct {
	c    client.Client
	name string
}

func NewMenuService(name string, c client.Client) MenuService {
	return &menuService{
		c:    c,
		name: name,
	}
}

func (c *menuService) Get(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Get", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Create(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Create", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Update(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Update", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) UpdateHide(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.UpdateHide", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Delete(ctx context.Context, in *Menu, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Delete", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Search(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Search", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) List(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.List", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Tree(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Tree", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) Routes(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.Routes", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) PluginMenuCreate(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.PluginMenuCreate", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuService) PluginMenuDelete(ctx context.Context, in *MenuRequest, opts ...client.CallOption) (*MenuResponse, error) {
	req := c.c.NewRequest(c.name, "MenuService.PluginMenuDelete", in)
	out := new(MenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MenuService service

type MenuServiceHandler interface {
	Get(context.Context, *Menu, *MenuResponse) error
	Create(context.Context, *Menu, *MenuResponse) error
	Update(context.Context, *Menu, *MenuResponse) error
	UpdateHide(context.Context, *Menu, *MenuResponse) error
	Delete(context.Context, *Menu, *MenuResponse) error
	Search(context.Context, *MenuRequest, *MenuResponse) error
	List(context.Context, *MenuRequest, *MenuResponse) error
	Tree(context.Context, *MenuRequest, *MenuResponse) error
	//用户菜单路由
	Routes(context.Context, *MenuRequest, *MenuResponse) error
	//插件导航创建
	PluginMenuCreate(context.Context, *MenuRequest, *MenuResponse) error
	//插件导航删除
	PluginMenuDelete(context.Context, *MenuRequest, *MenuResponse) error
}

func RegisterMenuServiceHandler(s server.Server, hdlr MenuServiceHandler, opts ...server.HandlerOption) error {
	type menuService interface {
		Get(ctx context.Context, in *Menu, out *MenuResponse) error
		Create(ctx context.Context, in *Menu, out *MenuResponse) error
		Update(ctx context.Context, in *Menu, out *MenuResponse) error
		UpdateHide(ctx context.Context, in *Menu, out *MenuResponse) error
		Delete(ctx context.Context, in *Menu, out *MenuResponse) error
		Search(ctx context.Context, in *MenuRequest, out *MenuResponse) error
		List(ctx context.Context, in *MenuRequest, out *MenuResponse) error
		Tree(ctx context.Context, in *MenuRequest, out *MenuResponse) error
		Routes(ctx context.Context, in *MenuRequest, out *MenuResponse) error
		PluginMenuCreate(ctx context.Context, in *MenuRequest, out *MenuResponse) error
		PluginMenuDelete(ctx context.Context, in *MenuRequest, out *MenuResponse) error
	}
	type MenuService struct {
		menuService
	}
	h := &menuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MenuService{h}, opts...))
}

type menuServiceHandler struct {
	MenuServiceHandler
}

func (h *menuServiceHandler) Get(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Get(ctx, in, out)
}

func (h *menuServiceHandler) Create(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Create(ctx, in, out)
}

func (h *menuServiceHandler) Update(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Update(ctx, in, out)
}

func (h *menuServiceHandler) UpdateHide(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.UpdateHide(ctx, in, out)
}

func (h *menuServiceHandler) Delete(ctx context.Context, in *Menu, out *MenuResponse) error {
	return h.MenuServiceHandler.Delete(ctx, in, out)
}

func (h *menuServiceHandler) Search(ctx context.Context, in *MenuRequest, out *MenuResponse) error {
	return h.MenuServiceHandler.Search(ctx, in, out)
}

func (h *menuServiceHandler) List(ctx context.Context, in *MenuRequest, out *MenuResponse) error {
	return h.MenuServiceHandler.List(ctx, in, out)
}

func (h *menuServiceHandler) Tree(ctx context.Context, in *MenuRequest, out *MenuResponse) error {
	return h.MenuServiceHandler.Tree(ctx, in, out)
}

func (h *menuServiceHandler) Routes(ctx context.Context, in *MenuRequest, out *MenuResponse) error {
	return h.MenuServiceHandler.Routes(ctx, in, out)
}

func (h *menuServiceHandler) PluginMenuCreate(ctx context.Context, in *MenuRequest, out *MenuResponse) error {
	return h.MenuServiceHandler.PluginMenuCreate(ctx, in, out)
}

func (h *menuServiceHandler) PluginMenuDelete(ctx context.Context, in *MenuRequest, out *MenuResponse) error {
	return h.MenuServiceHandler.PluginMenuDelete(ctx, in, out)
}
