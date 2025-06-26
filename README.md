# Response Package

基于 Gin 框架的统一响应处理包，提供标准化的 HTTP 响应格式、错误处理和分页响应支持。

## 功能特性

- 统一的响应格式
- 标准化的错误处理
- 分页响应支持
- 请求ID跟踪
- 内置常用响应方法
- 业务错误码管理

## 快速开始

### 1. 安装

```bash
go get -u github.com/hoewo/go-foundation-response
```

### 2. 基本使用

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/hoewo/go-foundation-response"
)

func main() {
    r := gin.Default()
    
    r.GET("/users/:id", func(c *gin.Context) {
        user := map[string]interface{}{
            "id": "123",
            "name": "John Doe",
            "email": "john@example.com",
        }
        response.Success(c, user)
    })
    
    r.Run(":8080")
}
```

## 响应格式

### 成功响应
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "id": "123",
        "name": "John Doe",
        "email": "john@example.com"
    },
    "timestamp": 1679384523,
    "request_id": "req_abc123"
}
```

### 错误响应
```json
{
    "code": 400,
    "message": "error",
    "error": {
        "code": "INVALID_PARAM",
        "message": "用户名不能为空",
        "details": {
            "field": "username",
            "reason": "required"
        }
    },
    "timestamp": 1679384523,
    "request_id": "req_abc123"
}
```

### 分页响应
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "list": [...],
        "total": 100,
        "page": 1,
        "page_size": 10,
        "total_pages": 10
    },
    "timestamp": 1679384523,
    "request_id": "req_abc123"
}
```

## API说明

### 成功响应方法

```go
// 基础成功响应
response.Success(c *gin.Context, data interface{})

// 带消息的成功响应
response.SuccessWithMessage(c *gin.Context, message string, data interface{})

// 分页成功响应
response.SuccessWithPage(c *gin.Context, message string, page *PageResponse)
```

### 错误响应方法

```go
// 通用错误响应
response.Error(c *gin.Context, httpCode int, code int, errorCode string, message string)

// 带详情的错误响应
response.ErrorWithDetails(c *gin.Context, httpCode int, code int, errorCode string, message string, details interface{})

// 快捷错误响应
response.BadRequest(c *gin.Context, message string)
response.Unauthorized(c *gin.Context, message string)
response.Forbidden(c *gin.Context, message string)
response.NotFound(c *gin.Context, message string)
response.InternalError(c *gin.Context, message string)

// 业务错误响应
response.BusinessError(c *gin.Context, errorCode string, message string)
```

### 分页响应创建

```go
response.NewPageResponse(list interface{}, total int64, page, pageSize int)
```

## 状态码说明

### HTTP状态码
```go
CodeSuccess            = 200 // 成功
CodeBadRequest         = 400 // 请求参数错误
CodeUnauthorized       = 401 // 未授权
CodeForbidden          = 403 // 禁止访问
CodeNotFound          = 404 // 资源不存在
CodeInternalError     = 500 // 内部错误
CodeServiceUnavailable = 503 // 服务不可用
```

### 错误代码
```go
ErrorCodeInvalidParam     = "INVALID_PARAM"      // 无效参数
ErrorCodeInvalidToken     = "INVALID_TOKEN"      // 无效令牌
ErrorCodeTokenExpired     = "TOKEN_EXPIRED"      // 令牌过期
ErrorCodeUserNotFound     = "USER_NOT_FOUND"     // 用户不存在
ErrorCodeUserExists       = "USER_EXISTS"        // 用户已存在
ErrorCodePasswordWrong    = "PASSWORD_WRONG"     // 密码错误
ErrorCodeAccountLocked    = "ACCOUNT_LOCKED"     // 账号锁定
ErrorCodeAccountInactive  = "ACCOUNT_INACTIVE"   // 账号未激活
ErrorCodeInsufficientPerm = "INSUFFICIENT_PERMISSION" // 权限不足
ErrorCodeInternalError    = "INTERNAL_ERROR"     // 内部错误
```

## 使用示例

### 1. 基本成功响应
```go
func GetUser(c *gin.Context) {
    user := GetUserFromDB()
    response.Success(c, user)
}
```

### 2. 带消息的成功响应
```go
func CreateUser(c *gin.Context) {
    user := CreateUserInDB()
    response.SuccessWithMessage(c, "用户创建成功", user)
}
```

### 3. 分页响应
```go
func ListUsers(c *gin.Context) {
    page := 1
    pageSize := 10
    users, total := GetUsersFromDB(page, pageSize)
    
    pageResponse := response.NewPageResponse(users, total, page, pageSize)
    response.SuccessWithPage(c, "获取用户列表成功", pageResponse)
}
```

### 4. 错误响应
```go
func UpdateUser(c *gin.Context) {
    if err := ValidateUser(); err != nil {
        response.BadRequest(c, "用户参数验证失败")
        return
    }
    
    if err := CheckPermission(); err != nil {
        response.Forbidden(c, "没有更新权限")
        return
    }
    
    if err := UpdateUserInDB(); err != nil {
        response.BusinessError(c, "UPDATE_FAILED", "更新用户信息失败")
        return
    }
    
    response.Success(c, nil)
}
```

### 5. 带详情的错误响应
```go
func ValidateUserInput(c *gin.Context) {
    errors := ValidateInput()
    if len(errors) > 0 {
        response.BadRequestWithDetails(c, "输入验证失败", errors)
        return
    }
}
```

## 最佳实践

1. **统一错误处理**：
   - 使用预定义的错误代码
   - 提供清晰的错误消息
   - 在需要时添加错误详情

2. **请求ID跟踪**：
   - 在中间件中生成请求ID
   - 通过上下文传递请求ID
   - 用于日志关联和问题排查

3. **分页处理**：
   - 统一使用PageResponse结构
   - 保持一致的分页参数命名
   - 合理设置默认分页大小

4. **响应格式**：
   - 保持响应结构的一致性
   - 使用适当的HTTP状态码
   - 提供有意义的响应消息 