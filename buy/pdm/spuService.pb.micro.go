// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: spuService.proto

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

// Api Endpoints for SpuService service

func NewSpuServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SpuService service

type SpuService interface {
	//创建商品
	Create(ctx context.Context, in *FormSpu, opts ...client.CallOption) (*SpuResponse, error)
	//修改商品
	Update(ctx context.Context, in *FormSpu, opts ...client.CallOption) (*SpuResponse, error)
	//删除商品
	Delete(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error)
	//获取商品（仅有Spu和skus信息）
	Get(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error)
	//商品详情（仅有Spu信息）
	GetBase(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error)
	//商品详情（后台编辑显示）
	Detail(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error)
	//规格详情
	SpecDetailInfo(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SkuResponse, error)
	//规格详情列表
	SpecDetailList(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SkuResponse, error)
	//商品上下架
	SetSale(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error)
	//商品排序
	SetSort(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error)
	//获取列表(后台服务)
	List(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error)
	//商品查询
	Search(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error)
}

type spuService struct {
	c    client.Client
	name string
}

func NewSpuService(name string, c client.Client) SpuService {
	return &spuService{
		c:    c,
		name: name,
	}
}

func (c *spuService) Create(ctx context.Context, in *FormSpu, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Create", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Update(ctx context.Context, in *FormSpu, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Update", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Delete(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Delete", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Get(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Get", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) GetBase(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.GetBase", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Detail(ctx context.Context, in *Spu, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Detail", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) SpecDetailInfo(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.SpecDetailInfo", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) SpecDetailList(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.SpecDetailList", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) SetSale(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.SetSale", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) SetSort(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.SetSort", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) List(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.List", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Search(ctx context.Context, in *SpuRequest, opts ...client.CallOption) (*SpuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Search", in)
	out := new(SpuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SpuService service

type SpuServiceHandler interface {
	//创建商品
	Create(context.Context, *FormSpu, *SpuResponse) error
	//修改商品
	Update(context.Context, *FormSpu, *SpuResponse) error
	//删除商品
	Delete(context.Context, *Spu, *SpuResponse) error
	//获取商品（仅有Spu和skus信息）
	Get(context.Context, *Spu, *SpuResponse) error
	//商品详情（仅有Spu信息）
	GetBase(context.Context, *Spu, *SpuResponse) error
	//商品详情（后台编辑显示）
	Detail(context.Context, *Spu, *SpuResponse) error
	//规格详情
	SpecDetailInfo(context.Context, *SpuRequest, *SkuResponse) error
	//规格详情列表
	SpecDetailList(context.Context, *SpuRequest, *SkuResponse) error
	//商品上下架
	SetSale(context.Context, *SpuRequest, *SpuResponse) error
	//商品排序
	SetSort(context.Context, *SpuRequest, *SpuResponse) error
	//获取列表(后台服务)
	List(context.Context, *SpuRequest, *SpuResponse) error
	//商品查询
	Search(context.Context, *SpuRequest, *SpuResponse) error
}

func RegisterSpuServiceHandler(s server.Server, hdlr SpuServiceHandler, opts ...server.HandlerOption) error {
	type spuService interface {
		Create(ctx context.Context, in *FormSpu, out *SpuResponse) error
		Update(ctx context.Context, in *FormSpu, out *SpuResponse) error
		Delete(ctx context.Context, in *Spu, out *SpuResponse) error
		Get(ctx context.Context, in *Spu, out *SpuResponse) error
		GetBase(ctx context.Context, in *Spu, out *SpuResponse) error
		Detail(ctx context.Context, in *Spu, out *SpuResponse) error
		SpecDetailInfo(ctx context.Context, in *SpuRequest, out *SkuResponse) error
		SpecDetailList(ctx context.Context, in *SpuRequest, out *SkuResponse) error
		SetSale(ctx context.Context, in *SpuRequest, out *SpuResponse) error
		SetSort(ctx context.Context, in *SpuRequest, out *SpuResponse) error
		List(ctx context.Context, in *SpuRequest, out *SpuResponse) error
		Search(ctx context.Context, in *SpuRequest, out *SpuResponse) error
	}
	type SpuService struct {
		spuService
	}
	h := &spuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SpuService{h}, opts...))
}

type spuServiceHandler struct {
	SpuServiceHandler
}

func (h *spuServiceHandler) Create(ctx context.Context, in *FormSpu, out *SpuResponse) error {
	return h.SpuServiceHandler.Create(ctx, in, out)
}

func (h *spuServiceHandler) Update(ctx context.Context, in *FormSpu, out *SpuResponse) error {
	return h.SpuServiceHandler.Update(ctx, in, out)
}

func (h *spuServiceHandler) Delete(ctx context.Context, in *Spu, out *SpuResponse) error {
	return h.SpuServiceHandler.Delete(ctx, in, out)
}

func (h *spuServiceHandler) Get(ctx context.Context, in *Spu, out *SpuResponse) error {
	return h.SpuServiceHandler.Get(ctx, in, out)
}

func (h *spuServiceHandler) GetBase(ctx context.Context, in *Spu, out *SpuResponse) error {
	return h.SpuServiceHandler.GetBase(ctx, in, out)
}

func (h *spuServiceHandler) Detail(ctx context.Context, in *Spu, out *SpuResponse) error {
	return h.SpuServiceHandler.Detail(ctx, in, out)
}

func (h *spuServiceHandler) SpecDetailInfo(ctx context.Context, in *SpuRequest, out *SkuResponse) error {
	return h.SpuServiceHandler.SpecDetailInfo(ctx, in, out)
}

func (h *spuServiceHandler) SpecDetailList(ctx context.Context, in *SpuRequest, out *SkuResponse) error {
	return h.SpuServiceHandler.SpecDetailList(ctx, in, out)
}

func (h *spuServiceHandler) SetSale(ctx context.Context, in *SpuRequest, out *SpuResponse) error {
	return h.SpuServiceHandler.SetSale(ctx, in, out)
}

func (h *spuServiceHandler) SetSort(ctx context.Context, in *SpuRequest, out *SpuResponse) error {
	return h.SpuServiceHandler.SetSort(ctx, in, out)
}

func (h *spuServiceHandler) List(ctx context.Context, in *SpuRequest, out *SpuResponse) error {
	return h.SpuServiceHandler.List(ctx, in, out)
}

func (h *spuServiceHandler) Search(ctx context.Context, in *SpuRequest, out *SpuResponse) error {
	return h.SpuServiceHandler.Search(ctx, in, out)
}
