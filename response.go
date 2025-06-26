package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code      int         `json:"code"`                 // 业务状态码
	Message   string      `json:"message"`              // 响应消息
	Data      interface{} `json:"data,omitempty"`       // 响应数据
	Error     *ErrorInfo  `json:"error,omitempty"`      // 错误信息
	Timestamp int64       `json:"timestamp"`            // 时间戳
	RequestID string      `json:"request_id,omitempty"` // 请求ID
}

// ErrorInfo 错误信息结构
type ErrorInfo struct {
	Code    string      `json:"code"`              // 错误代码
	Message string      `json:"message"`           // 错误消息
	Details interface{} `json:"details,omitempty"` // 错误详情
}

// PageResponse 分页响应结构
type PageResponse struct {
	List       interface{} `json:"list"`        // 数据列表
	Total      int64       `json:"total"`       // 总数
	Page       int         `json:"page"`        // 当前页
	PageSize   int         `json:"page_size"`   // 页大小
	TotalPages int         `json:"total_pages"` // 总页数
}

// HTTP状态码常量
const (
	CodeSuccess            = 200 // 成功
	CodeBadRequest         = 400 // 请求参数错误
	CodeUnauthorized       = 401 // 未授权
	CodeForbidden          = 403 // 禁止访问
	CodeNotFound           = 404 // 资源不存在
	CodeInternalError      = 500 // 内部错误
	CodeServiceUnavailable = 503 // 服务不可用
)

// 业务错误代码常量
const (
	// ErrorCodeInvalidParam 无效的参数
	ErrorCodeInvalidParam = "INVALID_PARAM"
	// ErrorCodeInvalidToken 无效的令牌
	ErrorCodeInvalidToken = "INVALID_TOKEN"
	// ErrorCodeTokenExpired 令牌已过期
	ErrorCodeTokenExpired = "TOKEN_EXPIRED"
	// ErrorCodeUserNotFound 用户不存在
	ErrorCodeUserNotFound = "USER_NOT_FOUND"
	// ErrorCodeUserExists 用户已存在
	ErrorCodeUserExists = "USER_EXISTS"
	// ErrorCodePasswordWrong 密码错误
	ErrorCodePasswordWrong = "PASSWORD_WRONG"
	// ErrorCodeAccountLocked 账号已锁定
	ErrorCodeAccountLocked = "ACCOUNT_LOCKED"
	// ErrorCodeAccountInactive 账号未激活
	ErrorCodeAccountInactive = "ACCOUNT_INACTIVE"
	// ErrorCodeInsufficientPerm 权限不足
	ErrorCodeInsufficientPerm = "INSUFFICIENT_PERMISSION"
	// ErrorCodeInternalError 内部错误
	ErrorCodeInternalError = "INTERNAL_ERROR"
)

// Success 返回成功响应
// c: gin上下文
// data: 响应数据
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// SuccessWithMessage 返回带自定义消息的成功响应
// c: gin上下文
// message: 自定义消息
// data: 响应数据
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// SuccessWithPage 返回分页成功响应
// c: gin上下文
// message: 自定义消息
// page: 分页数据
func SuccessWithPage(c *gin.Context, message string, page *PageResponse) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   message,
		Data:      page,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// Error 返回错误响应
// c: gin上下文
// httpCode: HTTP状态码
// code: 业务状态码
// errorCode: 错误代码
// message: 错误消息
func Error(c *gin.Context, httpCode int, code int, errorCode string, message string) {
	c.JSON(httpCode, &Response{
		Code:    code,
		Message: "error",
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// ErrorWithDetails 返回带详情的错误响应
// c: gin上下文
// httpCode: HTTP状态码
// code: 业务状态码
// errorCode: 错误代码
// message: 错误消息
// details: 错误详情
func ErrorWithDetails(c *gin.Context, httpCode int, code int, errorCode string, message string, details interface{}) {
	c.JSON(httpCode, &Response{
		Code:    code,
		Message: "error",
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
			Details: details,
		},
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// BadRequest 返回400错误响应
// c: gin上下文
// message: 错误消息
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, CodeBadRequest, ErrorCodeInvalidParam, message)
}

// BadRequestWithDetails 返回带详情的400错误响应
// c: gin上下文
// message: 错误消息
// details: 错误详情
func BadRequestWithDetails(c *gin.Context, message string, details interface{}) {
	ErrorWithDetails(c, http.StatusBadRequest, CodeBadRequest, ErrorCodeInvalidParam, message, details)
}

// Unauthorized 返回401错误响应
// c: gin上下文
// message: 错误消息
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, CodeUnauthorized, ErrorCodeInvalidToken, message)
}

// Forbidden 返回403错误响应
// c: gin上下文
// message: 错误消息
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, CodeForbidden, ErrorCodeInsufficientPerm, message)
}

// NotFound 返回404错误响应
// c: gin上下文
// message: 错误消息
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, CodeNotFound, ErrorCodeUserNotFound, message)
}

// InternalError 返回500错误响应
// c: gin上下文
// message: 错误消息
func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, CodeInternalError, ErrorCodeInternalError, message)
}

// BusinessError 返回业务逻辑错误响应
// c: gin上下文
// errorCode: 错误代码
// message: 错误消息
func BusinessError(c *gin.Context, errorCode string, message string) {
	Error(c, http.StatusBadRequest, CodeBadRequest, errorCode, message)
}

// getRequestID 获取请求ID
// c: gin上下文
// 返回请求ID字符串
func getRequestID(c *gin.Context) string {
	if requestID, exists := c.Get("request_id"); exists {
		return requestID.(string)
	}
	return c.GetHeader("X-Request-ID")
}

// NewPageResponse 创建分页响应
// list: 数据列表
// total: 总数
// page: 当前页码
// pageSize: 每页大小
// 返回分页响应结构
func NewPageResponse(list interface{}, total int64, page, pageSize int) *PageResponse {
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	return &PageResponse{
		List:       list,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}
