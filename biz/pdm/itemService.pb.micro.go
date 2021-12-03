// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: itemService.proto

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

// Api Endpoints for MyItemService service

func NewMyItemServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MyItemService service

type MyItemService interface {
	// 查询商品
	Search(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	// 获取商品信息
	Get(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	// 增加商品销量
	AddSoldNum(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
}

type myItemService struct {
	c    client.Client
	name string
}

func NewMyItemService(name string, c client.Client) MyItemService {
	return &myItemService{
		c:    c,
		name: name,
	}
}

func (c *myItemService) Search(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "MyItemService.Search", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myItemService) Get(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "MyItemService.Get", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myItemService) AddSoldNum(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "MyItemService.AddSoldNum", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyItemService service

type MyItemServiceHandler interface {
	// 查询商品
	Search(context.Context, *ItemRequest, *ItemResponse) error
	// 获取商品信息
	Get(context.Context, *ItemRequest, *ItemResponse) error
	// 增加商品销量
	AddSoldNum(context.Context, *ItemRequest, *ItemResponse) error
}

func RegisterMyItemServiceHandler(s server.Server, hdlr MyItemServiceHandler, opts ...server.HandlerOption) error {
	type myItemService interface {
		Search(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Get(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		AddSoldNum(ctx context.Context, in *ItemRequest, out *ItemResponse) error
	}
	type MyItemService struct {
		myItemService
	}
	h := &myItemServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyItemService{h}, opts...))
}

type myItemServiceHandler struct {
	MyItemServiceHandler
}

func (h *myItemServiceHandler) Search(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.MyItemServiceHandler.Search(ctx, in, out)
}

func (h *myItemServiceHandler) Get(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.MyItemServiceHandler.Get(ctx, in, out)
}

func (h *myItemServiceHandler) AddSoldNum(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.MyItemServiceHandler.AddSoldNum(ctx, in, out)
}

// Api Endpoints for ItemService service

func NewItemServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ItemService service

type ItemService interface {
	//添加商品
	Create(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error)
	//修改商品
	Update(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error)
	//删除商品
	Delete(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//商品上架
	Listing(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//商品下架
	UnListing(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//恢复商品
	Recovery(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//永久删除商品
	Destroy(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//商品基本信息
	Get(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//商品基本信息列表(SRV专用)
	GoodsList(ctx context.Context, in *GoodsRequest, opts ...client.CallOption) (*ItemResponse, error)
	//商品列表（来源基本信息）
	List(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//查询商品
	Search(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	FrontSearch(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	Display(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error)
	//商品规格详情
	SkuDetail(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	//查询已删除商品
	DeletedSearch(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	// 商品排序
	Sort(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
}

type itemService struct {
	c    client.Client
	name string
}

func NewItemService(name string, c client.Client) ItemService {
	return &itemService{
		c:    c,
		name: name,
	}
}

func (c *itemService) Create(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Create", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Update(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Update", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Delete(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Delete", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Listing(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Listing", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) UnListing(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.UnListing", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Recovery(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Recovery", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Destroy(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Destroy", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Get(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Get", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) GoodsList(ctx context.Context, in *GoodsRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.GoodsList", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) List(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.List", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Search(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Search", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) FrontSearch(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.FrontSearch", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Display(ctx context.Context, in *Item, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Display", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) SkuDetail(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.SkuDetail", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) DeletedSearch(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.DeletedSearch", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemService) Sort(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "ItemService.Sort", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ItemService service

type ItemServiceHandler interface {
	//添加商品
	Create(context.Context, *Item, *ItemResponse) error
	//修改商品
	Update(context.Context, *Item, *ItemResponse) error
	//删除商品
	Delete(context.Context, *ItemRequest, *ItemResponse) error
	//商品上架
	Listing(context.Context, *ItemRequest, *ItemResponse) error
	//商品下架
	UnListing(context.Context, *ItemRequest, *ItemResponse) error
	//恢复商品
	Recovery(context.Context, *ItemRequest, *ItemResponse) error
	//永久删除商品
	Destroy(context.Context, *ItemRequest, *ItemResponse) error
	//商品基本信息
	Get(context.Context, *ItemRequest, *ItemResponse) error
	//商品基本信息列表(SRV专用)
	GoodsList(context.Context, *GoodsRequest, *ItemResponse) error
	//商品列表（来源基本信息）
	List(context.Context, *ItemRequest, *ItemResponse) error
	//查询商品
	Search(context.Context, *ItemRequest, *ItemResponse) error
	FrontSearch(context.Context, *ItemRequest, *ItemResponse) error
	Display(context.Context, *Item, *ItemResponse) error
	//商品规格详情
	SkuDetail(context.Context, *ItemRequest, *ItemResponse) error
	//查询已删除商品
	DeletedSearch(context.Context, *ItemRequest, *ItemResponse) error
	// 商品排序
	Sort(context.Context, *ItemRequest, *ItemResponse) error
}

func RegisterItemServiceHandler(s server.Server, hdlr ItemServiceHandler, opts ...server.HandlerOption) error {
	type itemService interface {
		Create(ctx context.Context, in *Item, out *ItemResponse) error
		Update(ctx context.Context, in *Item, out *ItemResponse) error
		Delete(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Listing(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		UnListing(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Recovery(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Destroy(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Get(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		GoodsList(ctx context.Context, in *GoodsRequest, out *ItemResponse) error
		List(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Search(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		FrontSearch(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Display(ctx context.Context, in *Item, out *ItemResponse) error
		SkuDetail(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		DeletedSearch(ctx context.Context, in *ItemRequest, out *ItemResponse) error
		Sort(ctx context.Context, in *ItemRequest, out *ItemResponse) error
	}
	type ItemService struct {
		itemService
	}
	h := &itemServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ItemService{h}, opts...))
}

type itemServiceHandler struct {
	ItemServiceHandler
}

func (h *itemServiceHandler) Create(ctx context.Context, in *Item, out *ItemResponse) error {
	return h.ItemServiceHandler.Create(ctx, in, out)
}

func (h *itemServiceHandler) Update(ctx context.Context, in *Item, out *ItemResponse) error {
	return h.ItemServiceHandler.Update(ctx, in, out)
}

func (h *itemServiceHandler) Delete(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.Delete(ctx, in, out)
}

func (h *itemServiceHandler) Listing(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.Listing(ctx, in, out)
}

func (h *itemServiceHandler) UnListing(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.UnListing(ctx, in, out)
}

func (h *itemServiceHandler) Recovery(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.Recovery(ctx, in, out)
}

func (h *itemServiceHandler) Destroy(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.Destroy(ctx, in, out)
}

func (h *itemServiceHandler) Get(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.Get(ctx, in, out)
}

func (h *itemServiceHandler) GoodsList(ctx context.Context, in *GoodsRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.GoodsList(ctx, in, out)
}

func (h *itemServiceHandler) List(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.List(ctx, in, out)
}

func (h *itemServiceHandler) Search(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.Search(ctx, in, out)
}

func (h *itemServiceHandler) FrontSearch(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.FrontSearch(ctx, in, out)
}

func (h *itemServiceHandler) Display(ctx context.Context, in *Item, out *ItemResponse) error {
	return h.ItemServiceHandler.Display(ctx, in, out)
}

func (h *itemServiceHandler) SkuDetail(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.SkuDetail(ctx, in, out)
}

func (h *itemServiceHandler) DeletedSearch(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.DeletedSearch(ctx, in, out)
}

func (h *itemServiceHandler) Sort(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ItemServiceHandler.Sort(ctx, in, out)
}
