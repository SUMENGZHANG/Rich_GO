# 字节跳动内部框架选择情况

## 核心结论

**字节跳动内部主要使用自研的 Hertz 框架，而不是 Gin 或 Echo。**

## 历史演进

### 早期阶段：使用 Gin
- 字节跳动在早期的 Go 服务端开发中，**主要使用 Gin 框架**
- 对 Gin 进行了封装以满足业务需求
- Gin 在当时是主流选择，生态丰富，易于上手

### 转型阶段：自研 Hertz
随着业务的快速发展，字节跳动面临以下挑战：
1. **性能需求**：峰值 QPS 超过 4000 万
2. **扩展性需求**：需要支持更多业务场景
3. **定制化需求**：需要深度定制以满足企业级需求

Gin 框架在扩展性和性能优化方面的局限性逐渐显现，因此字节跳动决定自研框架。

### 当前阶段：Hertz 成为主流
- **2020 年初**：开始基于自研网络库 Netpoll 开发 Hertz
- **现在**：Hertz 已成为字节跳动内部最大的 HTTP 框架
- **规模**：线上接入的服务数量超过 **1 万**
- **性能**：峰值 QPS 超过 **4000 万**

## Hertz 框架介绍

### 基本信息
- **全称**: CloudWeGo Hertz
- **GitHub**: https://github.com/cloudwego/hertz
- **官网**: https://www.cloudwego.io/zh/docs/hertz/
- **开源时间**: 2022 年开源
- **设计理念**: 参考了 fasthttp、Gin 和 Echo 的优势

### 核心特点

#### 1. 高性能 ⚡
- 基于自研网络库 **Netpoll**
- 零拷贝技术
- 协程池优化
- 性能优于 Gin 和 Echo

#### 2. 高易用性 🎯
- API 设计参考 Gin，学习成本低
- 丰富的中间件支持
- 完善的文档和示例

#### 3. 高扩展性 🔧
- 插件化架构
- 支持自定义扩展
- 与 CloudWeGo 生态深度集成

#### 4. 企业级特性 🏢
- 服务治理完善（限流、熔断、监控）
- 与 Kitex 无缝集成
- 支持多种协议（HTTP/1.1, HTTP/2, gRPC）

### Hertz vs Gin vs Echo

| 特性 | Hertz | Gin | Echo |
|------|-------|-----|------|
| **性能** | ⭐⭐⭐⭐⭐ 极佳 | ⭐⭐⭐⭐ 优秀 | ⭐⭐⭐⭐⭐ 极佳 |
| **易用性** | ⭐⭐⭐⭐⭐ 优秀 | ⭐⭐⭐⭐⭐ 优秀 | ⭐⭐⭐⭐ 良好 |
| **生态** | ⭐⭐⭐⭐ CloudWeGo 生态 | ⭐⭐⭐⭐⭐ 非常丰富 | ⭐⭐⭐⭐ 丰富 |
| **企业级支持** | ⭐⭐⭐⭐⭐ 完善 | ⭐⭐⭐ 基础 | ⭐⭐⭐⭐ 良好 |
| **学习成本** | ⭐⭐⭐⭐ 低（类似 Gin） | ⭐⭐⭐⭐⭐ 很低 | ⭐⭐⭐⭐ 低 |

## 字节跳动技术栈

### HTTP 框架
- **主要使用**: Hertz（自研）
- **历史使用**: Gin（早期）
- **不使用**: Echo（未大规模采用）

### RPC 框架
- **主要使用**: Kitex（自研）
- **其他**: gRPC（部分场景）

### 完整技术栈
```
字节跳动 Go 微服务技术栈
├── HTTP 框架: Hertz
├── RPC 框架: Kitex
├── 服务注册发现: 自研（与 Kitex 集成）
├── 配置中心: 自研
├── 监控追踪: 自研 + OpenTelemetry
└── 网关: 自研网关（支持 Hertz）
```

## 为什么选择自研而不是继续使用 Gin？

### 1. 性能瓶颈
- **Gin 基于 httprouter**，在高并发场景下存在性能瓶颈
- **Hertz 基于 Netpoll**，性能更优，适合超大规模场景

### 2. 扩展性限制
- Gin 的扩展性有限，难以满足字节跳动的定制化需求
- Hertz 采用插件化架构，易于扩展

### 3. 企业级需求
- 需要与内部基础设施深度集成
- 需要统一的服务治理方案
- 需要与 Kitex 无缝配合

### 4. 技术积累
- 字节跳动有强大的基础架构团队
- 自研可以更好地控制技术路线
- 可以针对业务场景深度优化

## Hertz 代码示例

### 基础示例

```go
package main

import (
    "context"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
    h := server.Default()
    
    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        c.JSON(consts.StatusOK, map[string]interface{}{
            "message": "pong",
        })
    })
    
    h.Spin()
}
```

### 路由示例

```go
package main

import (
    "context"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
    h := server.Default()
    
    // 路由组
    v1 := h.Group("/api/v1")
    {
        v1.GET("/users/:id", getUser)
        v1.POST("/users", createUser)
    }
    
    h.Spin()
}

func getUser(ctx context.Context, c *app.RequestContext) {
    id := c.Param("id")
    c.JSON(consts.StatusOK, map[string]interface{}{
        "id":   id,
        "name": "用户" + id,
    })
}

func createUser(ctx context.Context, c *app.RequestContext) {
    var user struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    
    if err := c.Bind(&user); err != nil {
        c.JSON(consts.StatusBadRequest, map[string]interface{}{
            "error": err.Error(),
        })
        return
    }
    
    c.JSON(consts.StatusCreated, user)
}
```

## 对外的技术选型建议

### 如果你不是字节跳动

**推荐选择**:
1. **Gin** - 如果追求生态丰富和快速开发
2. **Echo** - 如果追求性能和统一错误处理
3. **Hertz** - 如果追求极致性能，且愿意使用 CloudWeGo 生态

### 选择 Hertz 的场景

✅ **适合选择 Hertz**:
- 需要极致性能
- 愿意使用 CloudWeGo 生态（Kitex + Hertz）
- 需要企业级服务治理
- 团队有技术实力进行深度定制

❌ **不适合选择 Hertz**:
- 需要丰富的第三方中间件（Gin 生态更丰富）
- 团队规模小，需要快速上手
- 不需要极致性能优化

## 总结

### 字节跳动内部情况
- **早期**: 使用 Gin（封装后使用）
- **现在**: 主要使用自研的 **Hertz** 框架
- **不使用**: Echo（未大规模采用）

### 原因分析
1. **性能需求**: 超大规模场景（4000万+ QPS）
2. **扩展性需求**: 需要深度定制
3. **企业级需求**: 需要与内部基础设施集成
4. **技术积累**: 有实力自研框架

### 对外建议
- **普通项目**: 推荐 Gin 或 Echo
- **高性能项目**: 可以考虑 Hertz
- **微服务项目**: 可以考虑 CloudWeGo 生态（Kitex + Hertz）

## 参考资源

- [Hertz 官方文档](https://www.cloudwego.io/zh/docs/hertz/)
- [Hertz GitHub](https://github.com/cloudwego/hertz)
- [CloudWeGo 官网](https://www.cloudwego.io/)
- [字节跳动开源技术](https://github.com/cloudwego)

