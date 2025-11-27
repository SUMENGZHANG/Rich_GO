# 企业级 Go 项目模块化设计规范

## 一、概述

本文档总结了大型互联网公司（如字节跳动、阿里、腾讯等）在 Go 项目中的模块化设计规范和最佳实践。

## 二、核心设计原则

### 2.1 单一职责原则（SRP）
- 每个模块只负责一个业务领域
- 每个函数只做一件事
- 避免"上帝类"和"万能函数"

### 2.2 高内聚、低耦合
- **高内聚**：模块内部组件紧密相关
- **低耦合**：模块之间依赖最小化
- 通过接口定义依赖关系

### 2.3 依赖倒置原则（DIP）
- 高层模块不依赖低层模块，都依赖抽象
- 通过接口定义依赖，而非具体实现
- 便于测试和替换实现

### 2.4 开闭原则（OCP）
- 对扩展开放，对修改关闭
- 通过组合而非继承实现扩展
- 使用策略模式、工厂模式等

## 三、标准项目结构

### 3.1 字节跳动/CloudWeGo 推荐结构

```
project/
├── cmd/                      # 应用程序入口
│   ├── server/              # HTTP 服务器入口
│   │   └── main.go
│   └── migrate/             # 数据库迁移工具入口
│       └── main.go
│
├── internal/                # 私有应用代码（外部不可导入）
│   ├── handler/             # HTTP 处理器层
│   │   ├── user.go         # 用户相关处理器
│   │   ├── coupon.go       # 优惠券相关处理器
│   │   └── health.go       # 健康检查处理器
│   │
│   ├── service/             # 业务逻辑层
│   │   ├── user.go         # 用户业务逻辑
│   │   └── coupon.go       # 优惠券业务逻辑
│   │
│   ├── repository/          # 数据访问层
│   │   ├── user.go         # 用户数据访问
│   │   └── coupon.go       # 优惠券数据访问
│   │
│   ├── model/               # 数据模型
│   │   ├── user.go
│   │   └── coupon.go
│   │
│   ├── middleware/          # 中间件
│   │   ├── auth.go         # 认证中间件
│   │   ├── logger.go       # 日志中间件
│   │   └── recovery.go      # 恢复中间件
│   │
│   ├── router/              # 路由配置
│   │   └── router.go       # 路由注册
│   │
│   ├── server/              # 服务器封装
│   │   └── http_server.go  # HTTP 服务器
│   │
│   └── config/              # 配置管理
│       └── config.go
│
├── pkg/                     # 可被外部使用的库代码
│   ├── errors/              # 错误定义
│   ├── validator/           # 验证器
│   └── utils/               # 工具函数
│
├── api/                     # API 定义（可选）
│   └── proto/               # Protobuf 定义
│
├── configs/                 # 配置文件
│   └── config.yaml
│
├── scripts/                 # 脚本文件
│   ├── build.sh
│   └── deploy.sh
│
├── docs/                    # 文档
│   └── README.md
│
├── test/                    # 测试文件
│   ├── integration/        # 集成测试
│   └── unit/               # 单元测试
│
├── go.mod
├── go.sum
└── Makefile
```

### 3.2 分层架构详解

#### Handler 层（处理器层）
**职责**：
- HTTP 请求解析和响应
- 参数验证
- 调用 Service 层
- 错误处理和响应格式化

**示例**：
```go
// internal/handler/user.go
package handler

import (
    "net/http"
    "rich_go/internal/service"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
    users, err := h.userService.ListUsers(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var req struct {
        Name  string `json:"name" binding:"required"`
        Email string `json:"email" binding:"required,email"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user, err := h.userService.CreateUser(c.Request.Context(), req.Name, req.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"user": user})
}
```

#### Service 层（业务逻辑层）
**职责**：
- 业务逻辑实现
- 事务管理
- 调用 Repository 层
- 数据转换和组装

**示例**：
```go
// internal/service/user.go
package service

import (
    "context"
    "rich_go/internal/model"
    "rich_go/internal/repository"
)

type UserService struct {
    userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
    return s.userRepo.FindAll(ctx)
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (*model.User, error) {
    // 业务逻辑验证
    if name == "" {
        return nil, errors.New("用户名不能为空")
    }
    
    user := &model.User{
        Name:  name,
        Email: email,
    }
    
    return s.userRepo.Create(ctx, user)
}
```

#### Repository 层（数据访问层）
**职责**：
- 数据库操作
- 数据持久化
- 查询优化
- 缓存管理

**示例**：
```go
// internal/repository/user.go
package repository

import (
    "context"
    "rich_go/internal/model"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
    var users []*model.User
    err := r.db.WithContext(ctx).Find(&users).Error
    return users, err
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
    err := r.db.WithContext(ctx).Create(user).Error
    return user, err
}
```

## 四、模块化设计模式

### 4.1 按业务模块组织

```
internal/
├── user/                    # 用户模块
│   ├── handler.go          # 用户处理器
│   ├── service.go          # 用户服务
│   ├── repository.go       # 用户仓储
│   └── model.go            # 用户模型
│
├── coupon/                  # 优惠券模块
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   └── model.go
│
└── order/                   # 订单模块
    ├── handler.go
    ├── service.go
    ├── repository.go
    └── model.go
```

**优点**：
- 模块边界清晰
- 便于团队分工
- 易于独立测试和部署

### 4.2 按技术层次组织（当前项目采用）

```
internal/
├── handler/                 # 所有处理器
│   ├── user.go
│   ├── coupon.go
│   └── order.go
│
├── service/                 # 所有服务
│   ├── user.go
│   ├── coupon.go
│   └── order.go
│
└── repository/              # 所有仓储
    ├── user.go
    ├── coupon.go
    └── order.go
```

**优点**：
- 层次清晰
- 便于统一管理
- 适合中小型项目

### 4.3 混合模式（推荐）

```
internal/
├── handler/                 # 处理器层
│   ├── user.go
│   ├── coupon.go
│   └── health.go
│
├── service/                 # 服务层
│   ├── user.go
│   └── coupon.go
│
├── repository/              # 仓储层
│   ├── user.go
│   └── coupon.go
│
└── module/                  # 复杂业务模块（可选）
    └── payment/             # 支付模块（包含子模块）
        ├── handler/
        ├── service/
        └── repository/
```

## 五、依赖注入模式

### 5.1 构造函数注入（推荐）

```go
// internal/server/http_server.go
package server

import (
    "rich_go/internal/handler"
    "rich_go/internal/service"
    "rich_go/internal/repository"
    "github.com/gin-gonic/gin"
)

type HTTPServer struct {
    router *gin.Engine
    
    // 依赖注入
    userHandler   *handler.UserHandler
    couponHandler *handler.CouponHandler
}

func NewHTTPServer(
    userRepo *repository.UserRepository,
    couponRepo *repository.CouponRepository,
) *HTTPServer {
    router := gin.Default()
    
    // 创建服务层
    userService := service.NewUserService(userRepo)
    couponService := service.NewCouponService(couponRepo)
    
    // 创建处理器层
    userHandler := handler.NewUserHandler(userService)
    couponHandler := handler.NewCouponHandler(couponService)
    
    // 注册路由
    setupRoutes(router, userHandler, couponHandler)
    
    return &HTTPServer{
        router:        router,
        userHandler:   userHandler,
        couponHandler: couponHandler,
    }
}
```

### 5.2 接口定义依赖

```go
// internal/service/user_service.go
package service

type UserService interface {
    ListUsers(ctx context.Context) ([]*model.User, error)
    CreateUser(ctx context.Context, name, email string) (*model.User, error)
}

// internal/repository/user_repository.go
package repository

type UserRepository interface {
    FindAll(ctx context.Context) ([]*model.User, error)
    Create(ctx context.Context, user *model.User) (*model.User, error)
}
```

## 六、路由组织方式

### 6.1 集中式路由（适合小型项目）

```go
// internal/router/router.go
package router

import (
    "rich_go/internal/handler"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(
    router *gin.Engine,
    userHandler *handler.UserHandler,
    couponHandler *handler.CouponHandler,
) {
    // 健康检查
    router.GET("/health", handler.HealthCheck)
    
    // API v1
    v1 := router.Group("/api/v1")
    {
        // 用户路由
        users := v1.Group("/users")
        {
            users.GET("", userHandler.ListUsers)
            users.GET("/:id", userHandler.GetUser)
            users.POST("", userHandler.CreateUser)
            users.PUT("/:id", userHandler.UpdateUser)
            users.DELETE("/:id", userHandler.DeleteUser)
        }
        
        // 优惠券路由
        coupons := v1.Group("/coupons")
        {
            coupons.GET("", couponHandler.ListCoupons)
            coupons.GET("/:id", couponHandler.GetCoupon)
            coupons.POST("", couponHandler.CreateCoupon)
            coupons.PUT("/:id", couponHandler.UpdateCoupon)
            coupons.DELETE("/:id", couponHandler.DeleteCoupon)
        }
    }
}
```

### 6.2 模块化路由（推荐，适合大型项目）

```go
// internal/router/router.go
package router

import (
    "rich_go/internal/handler"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(
    router *gin.Engine,
    userHandler *handler.UserHandler,
    couponHandler *handler.CouponHandler,
) {
    router.GET("/health", handler.HealthCheck)
    
    v1 := router.Group("/api/v1")
    {
        SetupUserRoutes(v1, userHandler)
        SetupCouponRoutes(v1, couponHandler)
    }
}

// internal/router/user_routes.go
package router

import (
    "rich_go/internal/handler"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(v1 *gin.RouterGroup, handler *handler.UserHandler) {
    users := v1.Group("/users")
    {
        users.GET("", handler.ListUsers)
        users.GET("/:id", handler.GetUser)
        users.POST("", handler.CreateUser)
        users.PUT("/:id", handler.UpdateUser)
        users.DELETE("/:id", handler.DeleteUser)
    }
}

// internal/router/coupon_routes.go
package router

import (
    "rich_go/internal/handler"
    "github.com/gin-gonic/gin"
)

func SetupCouponRoutes(v1 *gin.RouterGroup, handler *handler.CouponHandler) {
    coupons := v1.Group("/coupons")
    {
        coupons.GET("", handler.ListCoupons)
        coupons.GET("/:id", handler.GetCoupon)
        coupons.POST("", handler.CreateCoupon)
        coupons.PUT("/:id", handler.UpdateCoupon)
        coupons.DELETE("/:id", handler.DeleteCoupon)
    }
}
```

## 七、字节跳动/CloudWeGo 实践

### 7.1 Hertz 框架项目结构

基于 CloudWeGo Hertz 的推荐结构：

```
project/
├── biz/                     # 业务逻辑层（类似 service）
│   ├── user/
│   │   ├── handler.go      # HTTP 处理器
│   │   ├── service.go      # 业务逻辑
│   │   └── model.go        # 数据模型
│   └── coupon/
│       ├── handler.go
│       ├── service.go
│       └── model.go
│
├── dal/                     # 数据访问层（Data Access Layer）
│   ├── init.go             # 数据库初始化
│   └── user.go             # 用户数据访问
│
├── pkg/                     # 公共库
│   ├── errno/              # 错误码定义
│   └── middleware/         # 中间件
│
├── router/                  # 路由注册
│   └── register.go
│
└── main.go                  # 入口文件
```

### 7.2 字节跳动内部规范

1. **模块命名规范**
   - Handler: `{module}_handler.go`
   - Service: `{module}_service.go`
   - Repository: `{module}_repository.go`
   - Model: `{module}.go`

2. **接口定义规范**
   - Service 层必须定义接口
   - Repository 层必须定义接口
   - 便于测试和替换实现

3. **错误处理规范**
   - 统一错误码定义（`pkg/errno/`）
   - 错误信息国际化支持
   - 错误日志记录

4. **配置管理规范**
   - 配置文件统一管理（`configs/`）
   - 支持多环境配置（dev/test/prod）
   - 配置热更新支持

## 八、最佳实践总结

### 8.1 模块划分原则

1. **按业务领域划分**
   - 用户模块、订单模块、支付模块等
   - 每个模块独立，可单独测试和部署

2. **按技术层次划分**
   - Handler → Service → Repository
   - 清晰的职责边界

3. **避免循环依赖**
   - 使用接口解耦
   - 依赖方向：Handler → Service → Repository

### 8.2 代码组织建议

1. **小型项目（< 10 个模块）**
   ```
   internal/
   ├── handler/
   ├── service/
   └── repository/
   ```

2. **中型项目（10-50 个模块）**
   ```
   internal/
   ├── handler/
   ├── service/
   ├── repository/
   └── router/        # 路由单独管理
   ```

3. **大型项目（> 50 个模块）**
   ```
   internal/
   ├── module/        # 按业务模块组织
   │   ├── user/
   │   └── coupon/
   ├── shared/        # 共享组件
   └── router/
   ```

### 8.3 命名规范

1. **文件命名**
   - Handler: `user_handler.go` 或 `user.go`
   - Service: `user_service.go` 或 `user.go`
   - Repository: `user_repository.go` 或 `user.go`

2. **结构体命名**
   - Handler: `UserHandler`
   - Service: `UserService`
   - Repository: `UserRepository`

3. **函数命名**
   - HTTP 方法 + 资源名：`ListUsers`, `CreateUser`
   - 业务方法：`ValidateUser`, `CalculateDiscount`

## 九、对比总结

| 特性 | 集中式（当前） | 模块化（推荐） |
|------|---------------|---------------|
| **代码组织** | 所有 Handler 在一个目录 | 按模块或层次分离 |
| **可维护性** | ⭐⭐⭐ 中等 | ⭐⭐⭐⭐⭐ 优秀 |
| **可扩展性** | ⭐⭐⭐ 中等 | ⭐⭐⭐⭐⭐ 优秀 |
| **团队协作** | ⭐⭐ 一般 | ⭐⭐⭐⭐⭐ 优秀 |
| **适用场景** | 小型项目 | 中大型项目 |
| **学习成本** | ⭐⭐⭐⭐⭐ 低 | ⭐⭐⭐⭐ 中等 |

## 十、迁移建议

如果要从集中式迁移到模块化：

1. **第一步**：创建 `handlers/` 目录，按模块分离处理器
2. **第二步**：创建 `service/` 和 `repository/` 层
3. **第三步**：创建 `router/` 目录，模块化路由注册
4. **第四步**：引入依赖注入，解耦各层

## 参考资源

- [Go Standard Project Layout](https://github.com/golang-standards/project-layout)
- [CloudWeGo Hertz](https://www.cloudwego.io/zh/docs/hertz/)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html)

