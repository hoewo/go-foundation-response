package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 统一响应结构
type Response struct {
	Code      int         `json:"code"`                 // 业务状态码
	Message   string      `json:"message"`              // 响应消息
	Data      interface{} `json:"data,omitempty"`       // 响应数据
	Error     *ErrorInfo  `json:"error,omitempty"`      // 错误信息
	Timestamp int64       `json:"timestamp"`            // 时间戳
	RequestID string      `json:"request_id,omitempty"` // 请求ID
}

// 错误信息结构
type ErrorInfo struct {
	Code    string      `json:"code"`              // 错误代码
	Message string      `json:"message"`           // 错误消息
	Details interface{} `json:"details,omitempty"` // 错误详情
}

// 分页响应结构
type PageResponse struct {
	List       interface{} `json:"list"`        // 数据列表
	Total      int64       `json:"total"`       // 总数
	Page       int         `json:"page"`        // 当前页
	PageSize   int         `json:"page_size"`   // 页大小
	TotalPages int         `json:"total_pages"` // 总页数
}

// 响应代码常量
const (
	CodeSuccess            = 200 // 成功
	CodeBadRequest         = 400 // 请求参数错误
	CodeUnauthorized       = 401 // 未授权
	CodeForbidden          = 403 // 禁止访问
	CodeNotFound           = 404 // 资源不存在
	CodeInternalError      = 500 // 内部错误
	CodeServiceUnavailable = 503 // 服务不可用
)

// 错误代码常量
const (
	ErrorCodeInvalidParam     = "INVALID_PARAM"
	ErrorCodeInvalidToken     = "INVALID_TOKEN"
	ErrorCodeTokenExpired     = "TOKEN_EXPIRED"
	ErrorCodeUserNotFound     = "USER_NOT_FOUND"
	ErrorCodeUserExists       = "USER_EXISTS"
	ErrorCodePasswordWrong    = "PASSWORD_WRONG"
	ErrorCodeAccountLocked    = "ACCOUNT_LOCKED"
	ErrorCodeAccountInactive  = "ACCOUNT_INACTIVE"
	ErrorCodeInsufficientPerm = "INSUFFICIENT_PERMISSION"
	ErrorCodeInternalError    = "INTERNAL_ERROR"
)

// 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// 成功响应（带自定义消息）
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// 分页成功响应
func SuccessWithPage(c *gin.Context, message string, page *PageResponse) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   message,
		Data:      page,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// 错误响应
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

// 错误响应（带详情）
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

// 快捷错误响应方法
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, CodeBadRequest, ErrorCodeInvalidParam, message)
}

func BadRequestWithDetails(c *gin.Context, message string, details interface{}) {
	ErrorWithDetails(c, http.StatusBadRequest, CodeBadRequest, ErrorCodeInvalidParam, message, details)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, CodeUnauthorized, ErrorCodeInvalidToken, message)
}

func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, CodeForbidden, ErrorCodeInsufficientPerm, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, CodeNotFound, ErrorCodeUserNotFound, message)
}

func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, CodeInternalError, ErrorCodeInternalError, message)
}

// 业务逻辑错误响应
func BusinessError(c *gin.Context, errorCode string, message string) {
	Error(c, http.StatusBadRequest, CodeBadRequest, errorCode, message)
}

// 获取请求ID（从上下文或生成新的）
func getRequestID(c *gin.Context) string {
	if requestID, exists := c.Get("request_id"); exists {
		return requestID.(string)
	}
	return c.GetHeader("X-Request-ID")
}

// 创建分页响应
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
