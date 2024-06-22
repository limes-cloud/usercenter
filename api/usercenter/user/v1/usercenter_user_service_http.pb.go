// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/usercenter/user/usercenter_user_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserCreateUser = "/usercenter.api.usercenter.user.v1.User/CreateUser"
const OperationUserDeleteTrashUser = "/usercenter.api.usercenter.user.v1.User/DeleteTrashUser"
const OperationUserDeleteUser = "/usercenter.api.usercenter.user.v1.User/DeleteUser"
const OperationUserExportUser = "/usercenter.api.usercenter.user.v1.User/ExportUser"
const OperationUserGetCurrentUser = "/usercenter.api.usercenter.user.v1.User/GetCurrentUser"
const OperationUserGetTrashUser = "/usercenter.api.usercenter.user.v1.User/GetTrashUser"
const OperationUserGetUser = "/usercenter.api.usercenter.user.v1.User/GetUser"
const OperationUserImportUser = "/usercenter.api.usercenter.user.v1.User/ImportUser"
const OperationUserListTrashUser = "/usercenter.api.usercenter.user.v1.User/ListTrashUser"
const OperationUserListUser = "/usercenter.api.usercenter.user.v1.User/ListUser"
const OperationUserRevertTrashUser = "/usercenter.api.usercenter.user.v1.User/RevertTrashUser"
const OperationUserUpdateCurrentUser = "/usercenter.api.usercenter.user.v1.User/UpdateCurrentUser"
const OperationUserUpdateUser = "/usercenter.api.usercenter.user.v1.User/UpdateUser"
const OperationUserUpdateUserStatus = "/usercenter.api.usercenter.user.v1.User/UpdateUserStatus"

type UserHTTPServer interface {
	// CreateUser CreateUser 创建用户信息
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	// DeleteTrashUser DeleteTrashUser 彻底删除用户信息
	DeleteTrashUser(context.Context, *DeleteTrashUserRequest) (*DeleteTrashUserReply, error)
	// DeleteUser DeleteUser 删除用户信息
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	// ExportUser ExportUser 导出用户信息
	ExportUser(context.Context, *ExportUserRequest) (*ExportUserReply, error)
	// GetCurrentUser GetUser 获取当前的用户信息
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserReply, error)
	// GetTrashUser GetTrashUser 查看指定用户信息回收站数据
	GetTrashUser(context.Context, *GetTrashUserRequest) (*GetTrashUserReply, error)
	// GetUser GetUser 获取指定的用户信息
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	// ImportUser ImportUser 导入用户信息
	ImportUser(context.Context, *ImportUserRequest) (*ImportUserReply, error)
	// ListTrashUser ListTrashUser 查看用户信息列表回收站数据
	ListTrashUser(context.Context, *ListTrashUserRequest) (*ListTrashUserReply, error)
	// ListUser ListUser 获取用户信息列表
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	// RevertTrashUser RevertTrashUser 还原用户信息
	RevertTrashUser(context.Context, *RevertTrashUserRequest) (*RevertTrashUserReply, error)
	// UpdateCurrentUser UpdateCurrentUser 更新当前的用户信息
	UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserReply, error)
	// UpdateUser UpdateUser 更新用户信息
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	// UpdateUserStatus UpdateUserStatus 更新用户信息状态
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/usercenter/client/v1/user", _User_GetCurrentUser0_HTTP_Handler(srv))
	r.PUT("/usercenter/client/v1/user", _User_UpdateCurrentUser0_HTTP_Handler(srv))
	r.GET("/usercenter/api/v1/user", _User_GetUser0_HTTP_Handler(srv))
	r.GET("/usercenter/api/v1/users", _User_ListUser0_HTTP_Handler(srv))
	r.POST("/usercenter/api/v1/user", _User_CreateUser0_HTTP_Handler(srv))
	r.POST("/usercenter/api/v1/user/import", _User_ImportUser0_HTTP_Handler(srv))
	r.POST("/usercenter/api/v1/user/export", _User_ExportUser0_HTTP_Handler(srv))
	r.PUT("/usercenter/api/v1/user", _User_UpdateUser0_HTTP_Handler(srv))
	r.PUT("/usercenter/api/v1/user/status", _User_UpdateUserStatus0_HTTP_Handler(srv))
	r.DELETE("/usercenter/api/v1/user", _User_DeleteUser0_HTTP_Handler(srv))
	r.GET("/usercenter/api/v1/user/trash", _User_GetTrashUser0_HTTP_Handler(srv))
	r.GET("/usercenter/api/v1/user/trashes", _User_ListTrashUser0_HTTP_Handler(srv))
	r.DELETE("/usercenter/api/v1/user/trash", _User_DeleteTrashUser0_HTTP_Handler(srv))
	r.PUT("/usercenter/api/v1/user/trash", _User_RevertTrashUser0_HTTP_Handler(srv))
}

func _User_GetCurrentUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetCurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrentUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateCurrentUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCurrentUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateCurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateCurrentUser(ctx, req.(*UpdateCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCurrentUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ImportUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImportUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserImportUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ImportUser(ctx, req.(*ImportUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ImportUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ExportUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExportUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserExportUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ExportUser(ctx, req.(*ExportUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExportUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUserStatus0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUserStatus)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateUserStatus(ctx, req.(*UpdateUserStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserStatusReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetTrashUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTrashUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetTrashUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTrashUser(ctx, req.(*GetTrashUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTrashUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListTrashUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTrashUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListTrashUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListTrashUser(ctx, req.(*ListTrashUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTrashUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteTrashUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTrashUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteTrashUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteTrashUser(ctx, req.(*DeleteTrashUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTrashUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_RevertTrashUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RevertTrashUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRevertTrashUser)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.RevertTrashUser(ctx, req.(*RevertTrashUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RevertTrashUserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	DeleteTrashUser(ctx context.Context, req *DeleteTrashUserRequest, opts ...http.CallOption) (rsp *DeleteTrashUserReply, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserReply, err error)
	ExportUser(ctx context.Context, req *ExportUserRequest, opts ...http.CallOption) (rsp *ExportUserReply, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *GetCurrentUserReply, err error)
	GetTrashUser(ctx context.Context, req *GetTrashUserRequest, opts ...http.CallOption) (rsp *GetTrashUserReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ImportUser(ctx context.Context, req *ImportUserRequest, opts ...http.CallOption) (rsp *ImportUserReply, err error)
	ListTrashUser(ctx context.Context, req *ListTrashUserRequest, opts ...http.CallOption) (rsp *ListTrashUserReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	RevertTrashUser(ctx context.Context, req *RevertTrashUserRequest, opts ...http.CallOption) (rsp *RevertTrashUserReply, err error)
	UpdateCurrentUser(ctx context.Context, req *UpdateCurrentUserRequest, opts ...http.CallOption) (rsp *UpdateCurrentUserReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
	UpdateUserStatus(ctx context.Context, req *UpdateUserStatusRequest, opts ...http.CallOption) (rsp *UpdateUserStatusReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/usercenter/api/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) DeleteTrashUser(ctx context.Context, in *DeleteTrashUserRequest, opts ...http.CallOption) (*DeleteTrashUserReply, error) {
	var out DeleteTrashUserReply
	pattern := "/usercenter/api/v1/user/trash"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDeleteTrashUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserReply, error) {
	var out DeleteUserReply
	pattern := "/usercenter/api/v1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ExportUser(ctx context.Context, in *ExportUserRequest, opts ...http.CallOption) (*ExportUserReply, error) {
	var out ExportUserReply
	pattern := "/usercenter/api/v1/user/export"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserExportUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*GetCurrentUserReply, error) {
	var out GetCurrentUserReply
	pattern := "/usercenter/client/v1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetTrashUser(ctx context.Context, in *GetTrashUserRequest, opts ...http.CallOption) (*GetTrashUserReply, error) {
	var out GetTrashUserReply
	pattern := "/usercenter/api/v1/user/trash"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetTrashUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/usercenter/api/v1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ImportUser(ctx context.Context, in *ImportUserRequest, opts ...http.CallOption) (*ImportUserReply, error) {
	var out ImportUserReply
	pattern := "/usercenter/api/v1/user/import"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserImportUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListTrashUser(ctx context.Context, in *ListTrashUserRequest, opts ...http.CallOption) (*ListTrashUserReply, error) {
	var out ListTrashUserReply
	pattern := "/usercenter/api/v1/user/trashes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserListTrashUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/usercenter/api/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RevertTrashUser(ctx context.Context, in *RevertTrashUserRequest, opts ...http.CallOption) (*RevertTrashUserReply, error) {
	var out RevertTrashUserReply
	pattern := "/usercenter/api/v1/user/trash"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRevertTrashUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...http.CallOption) (*UpdateCurrentUserReply, error) {
	var out UpdateCurrentUserReply
	pattern := "/usercenter/client/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/usercenter/api/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...http.CallOption) (*UpdateUserStatusReply, error) {
	var out UpdateUserStatusReply
	pattern := "/usercenter/api/v1/user/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUserStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
