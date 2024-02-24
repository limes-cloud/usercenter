// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.4
// source: user_center_field_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationServiceAddField = "/field.Service/AddField"
const OperationServiceAllFieldType = "/field.Service/AllFieldType"
const OperationServiceCurrentField = "/field.Service/CurrentField"
const OperationServiceDeleteField = "/field.Service/DeleteField"
const OperationServicePageField = "/field.Service/PageField"
const OperationServiceUpdateField = "/field.Service/UpdateField"

type ServiceHTTPServer interface {
	AddField(context.Context, *AddFieldRequest) (*AddFieldReply, error)
	AllFieldType(context.Context, *emptypb.Empty) (*AllFieldTypeReply, error)
	CurrentField(context.Context, *emptypb.Empty) (*CurrentFieldReply, error)
	DeleteField(context.Context, *DeleteFieldRequest) (*emptypb.Empty, error)
	PageField(context.Context, *PageFieldRequest) (*PageFieldReply, error)
	UpdateField(context.Context, *UpdateFieldRequest) (*emptypb.Empty, error)
}

func RegisterServiceHTTPServer(s *http.Server, srv ServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/user-center/admin/v1/field/types", _Service_AllFieldType0_HTTP_Handler(srv))
	r.GET("/user-center/admin/v1/fields", _Service_PageField0_HTTP_Handler(srv))
	r.POST("/user-center/admin/v1/field", _Service_AddField0_HTTP_Handler(srv))
	r.PUT("/user-center/admin/v1/field", _Service_UpdateField0_HTTP_Handler(srv))
	r.DELETE("/user-center/admin/v1/field", _Service_DeleteField0_HTTP_Handler(srv))
	r.GET("/user-center/client/v1/fields", _Service_CurrentField0_HTTP_Handler(srv))
}

func _Service_AllFieldType0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAllFieldType)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.AllFieldType(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AllFieldTypeReply)
		return ctx.Result(200, reply)
	}
}

func _Service_PageField0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageFieldRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServicePageField)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.PageField(ctx, req.(*PageFieldRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageFieldReply)
		return ctx.Result(200, reply)
	}
}

func _Service_AddField0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddFieldRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAddField)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.AddField(ctx, req.(*AddFieldRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddFieldReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateField0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateFieldRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateField)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateField(ctx, req.(*UpdateFieldRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Service_DeleteField0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteFieldRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceDeleteField)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteField(ctx, req.(*DeleteFieldRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Service_CurrentField0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceCurrentField)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CurrentField(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrentFieldReply)
		return ctx.Result(200, reply)
	}
}

type ServiceHTTPClient interface {
	AddField(ctx context.Context, req *AddFieldRequest, opts ...http.CallOption) (rsp *AddFieldReply, err error)
	AllFieldType(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *AllFieldTypeReply, err error)
	CurrentField(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *CurrentFieldReply, err error)
	DeleteField(ctx context.Context, req *DeleteFieldRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	PageField(ctx context.Context, req *PageFieldRequest, opts ...http.CallOption) (rsp *PageFieldReply, err error)
	UpdateField(ctx context.Context, req *UpdateFieldRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type ServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceHTTPClient(client *http.Client) ServiceHTTPClient {
	return &ServiceHTTPClientImpl{client}
}

func (c *ServiceHTTPClientImpl) AddField(ctx context.Context, in *AddFieldRequest, opts ...http.CallOption) (*AddFieldReply, error) {
	var out AddFieldReply
	pattern := "/user-center/admin/v1/field"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceAddField))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) AllFieldType(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*AllFieldTypeReply, error) {
	var out AllFieldTypeReply
	pattern := "/user-center/admin/v1/field/types"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceAllFieldType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) CurrentField(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*CurrentFieldReply, error) {
	var out CurrentFieldReply
	pattern := "/user-center/client/v1/fields"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceCurrentField))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) DeleteField(ctx context.Context, in *DeleteFieldRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user-center/admin/v1/field"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceDeleteField))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) PageField(ctx context.Context, in *PageFieldRequest, opts ...http.CallOption) (*PageFieldReply, error) {
	var out PageFieldReply
	pattern := "/user-center/admin/v1/fields"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServicePageField))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateField(ctx context.Context, in *UpdateFieldRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user-center/admin/v1/field"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateField))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
