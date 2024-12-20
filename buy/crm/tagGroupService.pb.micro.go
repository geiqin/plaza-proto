// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: tagGroupService.proto

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

// Api Endpoints for TagGroupService service

func NewTagGroupServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TagGroupService service

type TagGroupService interface {
	// 客户标签组新增
	Create(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error)
	// 客户标签组修改
	Update(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error)
	// 客户标签组删除
	Delete(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error)
	// 客户标签组获取
	Get(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error)
	// 客户标签组查询
	Search(ctx context.Context, in *TagGroupRequest, opts ...client.CallOption) (*TagGroupResponse, error)
	// 客户标签组列表
	List(ctx context.Context, in *TagGroupRequest, opts ...client.CallOption) (*TagGroupResponse, error)
}

type tagGroupService struct {
	c    client.Client
	name string
}

func NewTagGroupService(name string, c client.Client) TagGroupService {
	return &tagGroupService{
		c:    c,
		name: name,
	}
}

func (c *tagGroupService) Create(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error) {
	req := c.c.NewRequest(c.name, "TagGroupService.Create", in)
	out := new(TagGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagGroupService) Update(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error) {
	req := c.c.NewRequest(c.name, "TagGroupService.Update", in)
	out := new(TagGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagGroupService) Delete(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error) {
	req := c.c.NewRequest(c.name, "TagGroupService.Delete", in)
	out := new(TagGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagGroupService) Get(ctx context.Context, in *TagGroup, opts ...client.CallOption) (*TagGroupResponse, error) {
	req := c.c.NewRequest(c.name, "TagGroupService.Get", in)
	out := new(TagGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagGroupService) Search(ctx context.Context, in *TagGroupRequest, opts ...client.CallOption) (*TagGroupResponse, error) {
	req := c.c.NewRequest(c.name, "TagGroupService.Search", in)
	out := new(TagGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagGroupService) List(ctx context.Context, in *TagGroupRequest, opts ...client.CallOption) (*TagGroupResponse, error) {
	req := c.c.NewRequest(c.name, "TagGroupService.List", in)
	out := new(TagGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TagGroupService service

type TagGroupServiceHandler interface {
	// 客户标签组新增
	Create(context.Context, *TagGroup, *TagGroupResponse) error
	// 客户标签组修改
	Update(context.Context, *TagGroup, *TagGroupResponse) error
	// 客户标签组删除
	Delete(context.Context, *TagGroup, *TagGroupResponse) error
	// 客户标签组获取
	Get(context.Context, *TagGroup, *TagGroupResponse) error
	// 客户标签组查询
	Search(context.Context, *TagGroupRequest, *TagGroupResponse) error
	// 客户标签组列表
	List(context.Context, *TagGroupRequest, *TagGroupResponse) error
}

func RegisterTagGroupServiceHandler(s server.Server, hdlr TagGroupServiceHandler, opts ...server.HandlerOption) error {
	type tagGroupService interface {
		Create(ctx context.Context, in *TagGroup, out *TagGroupResponse) error
		Update(ctx context.Context, in *TagGroup, out *TagGroupResponse) error
		Delete(ctx context.Context, in *TagGroup, out *TagGroupResponse) error
		Get(ctx context.Context, in *TagGroup, out *TagGroupResponse) error
		Search(ctx context.Context, in *TagGroupRequest, out *TagGroupResponse) error
		List(ctx context.Context, in *TagGroupRequest, out *TagGroupResponse) error
	}
	type TagGroupService struct {
		tagGroupService
	}
	h := &tagGroupServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TagGroupService{h}, opts...))
}

type tagGroupServiceHandler struct {
	TagGroupServiceHandler
}

func (h *tagGroupServiceHandler) Create(ctx context.Context, in *TagGroup, out *TagGroupResponse) error {
	return h.TagGroupServiceHandler.Create(ctx, in, out)
}

func (h *tagGroupServiceHandler) Update(ctx context.Context, in *TagGroup, out *TagGroupResponse) error {
	return h.TagGroupServiceHandler.Update(ctx, in, out)
}

func (h *tagGroupServiceHandler) Delete(ctx context.Context, in *TagGroup, out *TagGroupResponse) error {
	return h.TagGroupServiceHandler.Delete(ctx, in, out)
}

func (h *tagGroupServiceHandler) Get(ctx context.Context, in *TagGroup, out *TagGroupResponse) error {
	return h.TagGroupServiceHandler.Get(ctx, in, out)
}

func (h *tagGroupServiceHandler) Search(ctx context.Context, in *TagGroupRequest, out *TagGroupResponse) error {
	return h.TagGroupServiceHandler.Search(ctx, in, out)
}

func (h *tagGroupServiceHandler) List(ctx context.Context, in *TagGroupRequest, out *TagGroupResponse) error {
	return h.TagGroupServiceHandler.List(ctx, in, out)
}
