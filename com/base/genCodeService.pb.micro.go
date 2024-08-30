// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: genCodeService.proto

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

// Api Endpoints for GenCodeService service

func NewGenCodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GenCodeService service

type GenCodeService interface {
	Create(ctx context.Context, in *GenCode, opts ...client.CallOption) (*GenCodeResponse, error)
	Update(ctx context.Context, in *GenCode, opts ...client.CallOption) (*GenCodeResponse, error)
	Delete(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error)
	Get(ctx context.Context, in *GenCode, opts ...client.CallOption) (*GenCodeResponse, error)
	Search(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error)
	//数据库列表
	Databases(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error)
	//表列表
	Tables(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeTableResponse, error)
	//同步
	Sync(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error)
	//预览代码
	Preview(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error)
	//生成代码
	Generate(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error)
}

type genCodeService struct {
	c    client.Client
	name string
}

func NewGenCodeService(name string, c client.Client) GenCodeService {
	return &genCodeService{
		c:    c,
		name: name,
	}
}

func (c *genCodeService) Create(ctx context.Context, in *GenCode, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Create", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Update(ctx context.Context, in *GenCode, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Update", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Delete(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Delete", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Get(ctx context.Context, in *GenCode, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Get", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Search(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Search", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Databases(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Databases", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Tables(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeTableResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Tables", in)
	out := new(GenCodeTableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Sync(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.sync", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Preview(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Preview", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genCodeService) Generate(ctx context.Context, in *GenCodeRequest, opts ...client.CallOption) (*GenCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GenCodeService.Generate", in)
	out := new(GenCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GenCodeService service

type GenCodeServiceHandler interface {
	Create(context.Context, *GenCode, *GenCodeResponse) error
	Update(context.Context, *GenCode, *GenCodeResponse) error
	Delete(context.Context, *GenCodeRequest, *GenCodeResponse) error
	Get(context.Context, *GenCode, *GenCodeResponse) error
	Search(context.Context, *GenCodeRequest, *GenCodeResponse) error
	//数据库列表
	Databases(context.Context, *GenCodeRequest, *GenCodeResponse) error
	//表列表
	Tables(context.Context, *GenCodeRequest, *GenCodeTableResponse) error
	//同步
	Sync(context.Context, *GenCodeRequest, *GenCodeResponse) error
	//预览代码
	Preview(context.Context, *GenCodeRequest, *GenCodeResponse) error
	//生成代码
	Generate(context.Context, *GenCodeRequest, *GenCodeResponse) error
}

func RegisterGenCodeServiceHandler(s server.Server, hdlr GenCodeServiceHandler, opts ...server.HandlerOption) error {
	type genCodeService interface {
		Create(ctx context.Context, in *GenCode, out *GenCodeResponse) error
		Update(ctx context.Context, in *GenCode, out *GenCodeResponse) error
		Delete(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error
		Get(ctx context.Context, in *GenCode, out *GenCodeResponse) error
		Search(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error
		Databases(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error
		Tables(ctx context.Context, in *GenCodeRequest, out *GenCodeTableResponse) error
		Sync(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error
		Preview(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error
		Generate(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error
	}
	type GenCodeService struct {
		genCodeService
	}
	h := &genCodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GenCodeService{h}, opts...))
}

type genCodeServiceHandler struct {
	GenCodeServiceHandler
}

func (h *genCodeServiceHandler) Create(ctx context.Context, in *GenCode, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Create(ctx, in, out)
}

func (h *genCodeServiceHandler) Update(ctx context.Context, in *GenCode, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Update(ctx, in, out)
}

func (h *genCodeServiceHandler) Delete(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Delete(ctx, in, out)
}

func (h *genCodeServiceHandler) Get(ctx context.Context, in *GenCode, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Get(ctx, in, out)
}

func (h *genCodeServiceHandler) Search(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Search(ctx, in, out)
}

func (h *genCodeServiceHandler) Databases(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Databases(ctx, in, out)
}

func (h *genCodeServiceHandler) Tables(ctx context.Context, in *GenCodeRequest, out *GenCodeTableResponse) error {
	return h.GenCodeServiceHandler.Tables(ctx, in, out)
}

func (h *genCodeServiceHandler) Sync(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Sync(ctx, in, out)
}

func (h *genCodeServiceHandler) Preview(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Preview(ctx, in, out)
}

func (h *genCodeServiceHandler) Generate(ctx context.Context, in *GenCodeRequest, out *GenCodeResponse) error {
	return h.GenCodeServiceHandler.Generate(ctx, in, out)
}
