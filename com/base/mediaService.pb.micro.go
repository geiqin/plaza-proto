// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: mediaService.proto

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

// Api Endpoints for MediaService service

func NewMediaServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MediaService service

type MediaService interface {
	// 公共图标新增
	Create(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error)
	// 公共图标修改
	Update(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error)
	// 公共图标删除
	Delete(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error)
	// 公共图标获取
	Get(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error)
	// 公共图标查询
	Search(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error)
	// 公共图标列表
	List(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error)
	//素材统计
	Statistics(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error)
	//上传凭证生成
	UploadToken(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error)
	//上传回调处理
	UploadCallback(ctx context.Context, in *CallbackInfo, opts ...client.CallOption) (*MediaResponse, error)
}

type mediaService struct {
	c    client.Client
	name string
}

func NewMediaService(name string, c client.Client) MediaService {
	return &mediaService{
		c:    c,
		name: name,
	}
}

func (c *mediaService) Create(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.Create", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) Update(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.Update", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) Delete(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.Delete", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) Get(ctx context.Context, in *Media, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.Get", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) Search(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.Search", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) List(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.List", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) Statistics(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.Statistics", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) UploadToken(ctx context.Context, in *MediaRequest, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.UploadToken", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaService) UploadCallback(ctx context.Context, in *CallbackInfo, opts ...client.CallOption) (*MediaResponse, error) {
	req := c.c.NewRequest(c.name, "MediaService.UploadCallback", in)
	out := new(MediaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MediaService service

type MediaServiceHandler interface {
	// 公共图标新增
	Create(context.Context, *Media, *MediaResponse) error
	// 公共图标修改
	Update(context.Context, *Media, *MediaResponse) error
	// 公共图标删除
	Delete(context.Context, *Media, *MediaResponse) error
	// 公共图标获取
	Get(context.Context, *Media, *MediaResponse) error
	// 公共图标查询
	Search(context.Context, *MediaRequest, *MediaResponse) error
	// 公共图标列表
	List(context.Context, *MediaRequest, *MediaResponse) error
	//素材统计
	Statistics(context.Context, *MediaRequest, *MediaResponse) error
	//上传凭证生成
	UploadToken(context.Context, *MediaRequest, *MediaResponse) error
	//上传回调处理
	UploadCallback(context.Context, *CallbackInfo, *MediaResponse) error
}

func RegisterMediaServiceHandler(s server.Server, hdlr MediaServiceHandler, opts ...server.HandlerOption) error {
	type mediaService interface {
		Create(ctx context.Context, in *Media, out *MediaResponse) error
		Update(ctx context.Context, in *Media, out *MediaResponse) error
		Delete(ctx context.Context, in *Media, out *MediaResponse) error
		Get(ctx context.Context, in *Media, out *MediaResponse) error
		Search(ctx context.Context, in *MediaRequest, out *MediaResponse) error
		List(ctx context.Context, in *MediaRequest, out *MediaResponse) error
		Statistics(ctx context.Context, in *MediaRequest, out *MediaResponse) error
		UploadToken(ctx context.Context, in *MediaRequest, out *MediaResponse) error
		UploadCallback(ctx context.Context, in *CallbackInfo, out *MediaResponse) error
	}
	type MediaService struct {
		mediaService
	}
	h := &mediaServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MediaService{h}, opts...))
}

type mediaServiceHandler struct {
	MediaServiceHandler
}

func (h *mediaServiceHandler) Create(ctx context.Context, in *Media, out *MediaResponse) error {
	return h.MediaServiceHandler.Create(ctx, in, out)
}

func (h *mediaServiceHandler) Update(ctx context.Context, in *Media, out *MediaResponse) error {
	return h.MediaServiceHandler.Update(ctx, in, out)
}

func (h *mediaServiceHandler) Delete(ctx context.Context, in *Media, out *MediaResponse) error {
	return h.MediaServiceHandler.Delete(ctx, in, out)
}

func (h *mediaServiceHandler) Get(ctx context.Context, in *Media, out *MediaResponse) error {
	return h.MediaServiceHandler.Get(ctx, in, out)
}

func (h *mediaServiceHandler) Search(ctx context.Context, in *MediaRequest, out *MediaResponse) error {
	return h.MediaServiceHandler.Search(ctx, in, out)
}

func (h *mediaServiceHandler) List(ctx context.Context, in *MediaRequest, out *MediaResponse) error {
	return h.MediaServiceHandler.List(ctx, in, out)
}

func (h *mediaServiceHandler) Statistics(ctx context.Context, in *MediaRequest, out *MediaResponse) error {
	return h.MediaServiceHandler.Statistics(ctx, in, out)
}

func (h *mediaServiceHandler) UploadToken(ctx context.Context, in *MediaRequest, out *MediaResponse) error {
	return h.MediaServiceHandler.UploadToken(ctx, in, out)
}

func (h *mediaServiceHandler) UploadCallback(ctx context.Context, in *CallbackInfo, out *MediaResponse) error {
	return h.MediaServiceHandler.UploadCallback(ctx, in, out)
}
