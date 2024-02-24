// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotFoundError.String() && e.Code == 200
}

// 为某个枚举单独设置错误码
func NotFoundErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotFoundError.String(), "数据不存在:"+fmt.Sprintf(format, args...))
}

// 为某个枚举单独设置错误码
func NotFoundError() *errors.Error {
	return errors.New(200, ErrorReason_NotFoundError.String(), "数据不存在")
}

func IsDatabaseError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DatabaseError.String() && e.Code == 200
}

func DatabaseErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DatabaseError.String(), "数据库错误:"+fmt.Sprintf(format, args...))
}

func DatabaseError() *errors.Error {
	return errors.New(200, ErrorReason_DatabaseError.String(), "数据库错误")
}

func IsMetadataError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_MetadataError.String() && e.Code == 200
}

func MetadataErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_MetadataError.String(), "元数据异常:"+fmt.Sprintf(format, args...))
}

func MetadataError() *errors.Error {
	return errors.New(200, ErrorReason_MetadataError.String(), "元数据异常")
}

func IsTransformError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TransformError.String() && e.Code == 200
}

func TransformErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_TransformError.String(), "数据转换失败:"+fmt.Sprintf(format, args...))
}

func TransformError() *errors.Error {
	return errors.New(200, ErrorReason_TransformError.String(), "数据转换失败")
}

func IsGenCaptchaError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GenCaptchaError.String() && e.Code == 200
}

func GenCaptchaErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_GenCaptchaError.String(), "验证码生成失败:"+fmt.Sprintf(format, args...))
}

func GenCaptchaError() *errors.Error {
	return errors.New(200, ErrorReason_GenCaptchaError.String(), "验证码生成失败")
}

func IsVerifyCaptchaError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_VerifyCaptchaError.String() && e.Code == 200
}

func VerifyCaptchaErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_VerifyCaptchaError.String(), "验证码验证失败:"+fmt.Sprintf(format, args...))
}

func VerifyCaptchaError() *errors.Error {
	return errors.New(200, ErrorReason_VerifyCaptchaError.String(), "验证码验证失败")
}

func IsNotRecordError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotRecordError.String() && e.Code == 200
}

func NotRecordErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotRecordError.String(), "暂无数据记录:"+fmt.Sprintf(format, args...))
}

func NotRecordError() *errors.Error {
	return errors.New(200, ErrorReason_NotRecordError.String(), "暂无数据记录")
}

func IsRsaDecodeError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RsaDecodeError.String() && e.Code == 200
}

func RsaDecodeErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_RsaDecodeError.String(), "rsa解密失败:"+fmt.Sprintf(format, args...))
}

func RsaDecodeError() *errors.Error {
	return errors.New(200, ErrorReason_RsaDecodeError.String(), "rsa解密失败")
}

func IsPasswordFormatError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PasswordFormatError.String() && e.Code == 200
}

func PasswordFormatErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_PasswordFormatError.String(), "密码格式错误:"+fmt.Sprintf(format, args...))
}

func PasswordFormatError() *errors.Error {
	return errors.New(200, ErrorReason_PasswordFormatError.String(), "密码格式错误")
}

func IsPasswordExpireError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PasswordExpireError.String() && e.Code == 200
}

func PasswordExpireErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_PasswordExpireError.String(), "密码已过期:"+fmt.Sprintf(format, args...))
}

func PasswordExpireError() *errors.Error {
	return errors.New(200, ErrorReason_PasswordExpireError.String(), "密码已过期")
}

func IsUsernameOrPasswordError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UsernameOrPasswordError.String() && e.Code == 200
}

func UsernameOrPasswordErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UsernameOrPasswordError.String(), "账户或密码错误:"+fmt.Sprintf(format, args...))
}

func UsernameOrPasswordError() *errors.Error {
	return errors.New(200, ErrorReason_UsernameOrPasswordError.String(), "账户或密码错误")
}

func IsUserDisableError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UserDisableError.String() && e.Code == 200
}

func UserDisableErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UserDisableError.String(), "用户已被禁用:"+fmt.Sprintf(format, args...))
}

func UserDisableError() *errors.Error {
	return errors.New(200, ErrorReason_UserDisableError.String(), "用户已被禁用")
}

func IsGenTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GenTokenError.String() && e.Code == 200
}

func GenTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_GenTokenError.String(), "token生成失败:"+fmt.Sprintf(format, args...))
}

func GenTokenError() *errors.Error {
	return errors.New(200, ErrorReason_GenTokenError.String(), "token生成失败")
}

func IsParseTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ParseTokenError.String() && e.Code == 200
}

func ParseTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_ParseTokenError.String(), "token解析失败:"+fmt.Sprintf(format, args...))
}

func ParseTokenError() *errors.Error {
	return errors.New(200, ErrorReason_ParseTokenError.String(), "token解析失败")
}

func IsRefreshTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RefreshTokenError.String() && e.Code == 401
}

func RefreshTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败:"+fmt.Sprintf(format, args...))
}

func RefreshTokenError() *errors.Error {
	return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败")
}

func IsEmptyTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_EmptyTokenError.String() && e.Code == 200
}

func EmptyTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_EmptyTokenError.String(), "token不能为空:"+fmt.Sprintf(format, args...))
}

func EmptyTokenError() *errors.Error {
	return errors.New(200, ErrorReason_EmptyTokenError.String(), "token不能为空")
}

func IsNotUserAppError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotUserAppError.String() && e.Code == 200
}

func NotUserAppErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotUserAppError.String(), "用户暂无应用使用权限:"+fmt.Sprintf(format, args...))
}

func NotUserAppError() *errors.Error {
	return errors.New(200, ErrorReason_NotUserAppError.String(), "用户暂无应用使用权限")
}

func IsNotUserChannelError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotUserChannelError.String() && e.Code == 200
}

func NotUserChannelErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotUserChannelError.String(), "用户暂无渠道使用权限:"+fmt.Sprintf(format, args...))
}

func NotUserChannelError() *errors.Error {
	return errors.New(200, ErrorReason_NotUserChannelError.String(), "用户暂无渠道使用权限")
}

func IsNotAppError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotAppError.String() && e.Code == 200
}

func NotAppErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotAppError.String(), "错误的应用标识:"+fmt.Sprintf(format, args...))
}

func NotAppError() *errors.Error {
	return errors.New(200, ErrorReason_NotAppError.String(), "错误的应用标识")
}

func IsNotAppChannelError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotAppChannelError.String() && e.Code == 200
}

func NotAppChannelErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotAppChannelError.String(), "应用为开通此渠道:"+fmt.Sprintf(format, args...))
}

func NotAppChannelError() *errors.Error {
	return errors.New(200, ErrorReason_NotAppChannelError.String(), "应用为开通此渠道")
}

func IsNotExistEmailError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotExistEmailError.String() && e.Code == 200
}

func NotExistEmailErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotExistEmailError.String(), "不存在此邮箱:"+fmt.Sprintf(format, args...))
}

func NotExistEmailError() *errors.Error {
	return errors.New(200, ErrorReason_NotExistEmailError.String(), "不存在此邮箱")
}

func IsAlreadyExistEmailError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AlreadyExistEmailError.String() && e.Code == 200
}

func AlreadyExistEmailErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_AlreadyExistEmailError.String(), "邮箱已存在:"+fmt.Sprintf(format, args...))
}

func AlreadyExistEmailError() *errors.Error {
	return errors.New(200, ErrorReason_AlreadyExistEmailError.String(), "邮箱已存在")
}

func IsAlreadyExistPhoneError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AlreadyExistPhoneError.String() && e.Code == 200
}

func AlreadyExistPhoneErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_AlreadyExistPhoneError.String(), "手机已存在:"+fmt.Sprintf(format, args...))
}

func AlreadyExistPhoneError() *errors.Error {
	return errors.New(200, ErrorReason_AlreadyExistPhoneError.String(), "手机已存在")
}

func IsAlreadyExistUsernameError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AlreadyExistUsernameError.String() && e.Code == 200
}

func AlreadyExistUsernameErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_AlreadyExistUsernameError.String(), "账号已存在:"+fmt.Sprintf(format, args...))
}

func AlreadyExistUsernameError() *errors.Error {
	return errors.New(200, ErrorReason_AlreadyExistUsernameError.String(), "账号已存在")
}

func IsRegisterError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RegisterError.String() && e.Code == 200
}

func RegisterErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_RegisterError.String(), "用户注册失败:"+fmt.Sprintf(format, args...))
}

func RegisterError() *errors.Error {
	return errors.New(200, ErrorReason_RegisterError.String(), "用户注册失败")
}

func IsForbiddenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ForbiddenError.String() && e.Code == 403
}

func ForbiddenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(403, ErrorReason_ForbiddenError.String(), "无应用权限:"+fmt.Sprintf(format, args...))
}

func ForbiddenError() *errors.Error {
	return errors.New(403, ErrorReason_ForbiddenError.String(), "无应用权限")
}

func IsGetAuthInfoError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GetAuthInfoError.String() && e.Code == 200
}

func GetAuthInfoErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_GetAuthInfoError.String(), "获取授权信息失败:"+fmt.Sprintf(format, args...))
}

func GetAuthInfoError() *errors.Error {
	return errors.New(200, ErrorReason_GetAuthInfoError.String(), "获取授权信息失败")
}

func IsUnBindError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UnBindError.String() && e.Code == 401
}

func UnBindErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(401, ErrorReason_UnBindError.String(), "用户未绑定:"+fmt.Sprintf(format, args...))
}

func UnBindError() *errors.Error {
	return errors.New(401, ErrorReason_UnBindError.String(), "用户未绑定")
}

func IsNotUserError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotUserError.String() && e.Code == 200
}

func NotUserErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotUserError.String(), "用户不存在:"+fmt.Sprintf(format, args...))
}

func NotUserError() *errors.Error {
	return errors.New(200, ErrorReason_NotUserError.String(), "用户不存在")
}