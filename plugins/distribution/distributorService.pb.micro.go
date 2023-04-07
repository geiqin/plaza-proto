// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: distributorService.proto

package services

import (
	_ "../common"
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

// Api Endpoints for DistributorService service

func NewDistributorServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DistributorService service

type DistributorService interface {
	//检查用户是否是分销员【前端】
	Exists(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error)
	//编辑销售员信息
	Update(ctx context.Context, in *Distributor, opts ...client.CallOption) (*DistributorResponse, error)
	//批量设置销售员等级
	SetRank(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error)
	//清退销售员
	Remove(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error)
	//获取销售员信息
	Get(ctx context.Context, in *Distributor, opts ...client.CallOption) (*DistributorResponse, error)
	//分页查询销售员列表
	Search(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error)
	//获取销售员的团队数量
	TeamNum(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error)
	//查询销售员团队
	TeamSearch(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error)
}

type distributorService struct {
	c    client.Client
	name string
}

func NewDistributorService(name string, c client.Client) DistributorService {
	return &distributorService{
		c:    c,
		name: name,
	}
}

func (c *distributorService) Exists(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.Exists", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorService) Update(ctx context.Context, in *Distributor, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.Update", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorService) SetRank(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.SetRank", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorService) Remove(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.Remove", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorService) Get(ctx context.Context, in *Distributor, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.Get", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorService) Search(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.Search", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorService) TeamNum(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.TeamNum", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorService) TeamSearch(ctx context.Context, in *DistributorRequest, opts ...client.CallOption) (*DistributorResponse, error) {
	req := c.c.NewRequest(c.name, "DistributorService.TeamSearch", in)
	out := new(DistributorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DistributorService service

type DistributorServiceHandler interface {
	//检查用户是否是分销员【前端】
	Exists(context.Context, *DistributorRequest, *DistributorResponse) error
	//编辑销售员信息
	Update(context.Context, *Distributor, *DistributorResponse) error
	//批量设置销售员等级
	SetRank(context.Context, *DistributorRequest, *DistributorResponse) error
	//清退销售员
	Remove(context.Context, *DistributorRequest, *DistributorResponse) error
	//获取销售员信息
	Get(context.Context, *Distributor, *DistributorResponse) error
	//分页查询销售员列表
	Search(context.Context, *DistributorRequest, *DistributorResponse) error
	//获取销售员的团队数量
	TeamNum(context.Context, *DistributorRequest, *DistributorResponse) error
	//查询销售员团队
	TeamSearch(context.Context, *DistributorRequest, *DistributorResponse) error
}

func RegisterDistributorServiceHandler(s server.Server, hdlr DistributorServiceHandler, opts ...server.HandlerOption) error {
	type distributorService interface {
		Exists(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error
		Update(ctx context.Context, in *Distributor, out *DistributorResponse) error
		SetRank(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error
		Remove(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error
		Get(ctx context.Context, in *Distributor, out *DistributorResponse) error
		Search(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error
		TeamNum(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error
		TeamSearch(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error
	}
	type DistributorService struct {
		distributorService
	}
	h := &distributorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DistributorService{h}, opts...))
}

type distributorServiceHandler struct {
	DistributorServiceHandler
}

func (h *distributorServiceHandler) Exists(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error {
	return h.DistributorServiceHandler.Exists(ctx, in, out)
}

func (h *distributorServiceHandler) Update(ctx context.Context, in *Distributor, out *DistributorResponse) error {
	return h.DistributorServiceHandler.Update(ctx, in, out)
}

func (h *distributorServiceHandler) SetRank(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error {
	return h.DistributorServiceHandler.SetRank(ctx, in, out)
}

func (h *distributorServiceHandler) Remove(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error {
	return h.DistributorServiceHandler.Remove(ctx, in, out)
}

func (h *distributorServiceHandler) Get(ctx context.Context, in *Distributor, out *DistributorResponse) error {
	return h.DistributorServiceHandler.Get(ctx, in, out)
}

func (h *distributorServiceHandler) Search(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error {
	return h.DistributorServiceHandler.Search(ctx, in, out)
}

func (h *distributorServiceHandler) TeamNum(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error {
	return h.DistributorServiceHandler.TeamNum(ctx, in, out)
}

func (h *distributorServiceHandler) TeamSearch(ctx context.Context, in *DistributorRequest, out *DistributorResponse) error {
	return h.DistributorServiceHandler.TeamSearch(ctx, in, out)
}
