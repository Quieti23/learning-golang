# Week 3：中间件、Context、错误码、参数校验

## 本周目标

- 理解并实现 HTTP 中间件模式
- 掌握 `context.Context` 在请求链路中的传递和使用
- 建立统一的错误码返回体系
- 实现基本的请求参数校验

## 本周交付

- 至少 2 个中间件（日志中间件、panic recovery 中间件）
- context 超时控制 demo
- 统一错误返回格式
- 请求参数校验示例

## 建议日程

### Day 1

- 学 Go HTTP 中间件的实现原理：函数包装函数
- 实现一个日志中间件：记录请求方法、路径、耗时
- 输出：所有请求自动输出日志

### Day 2

- 实现一个 panic recovery 中间件：捕获 panic 并返回 500
- 学习中间件的组合方式
- 输出：服务不会因为单个请求 panic 而崩溃

### Day 3

- 学 `context.Context` 的基本用法
- 学 `context.WithTimeout`、`context.WithCancel`
- 输出：写一个 demo 演示超时取消和 context 值传递

### Day 4

- 在 Blog API 中传递 context：从 handler 到 service 到 repository
- 数据库操作使用 `QueryContext`、`ExecContext`
- 输出：Blog API 的数据库操作支持 context 超时控制

### Day 5

- 设计统一错误返回结构：`{"code": "NOT_FOUND", "message": "post not found"}`
- 区分业务错误和系统错误
- 实现错误码常量和错误响应工具函数
- 输出：所有接口返回统一格式

### Day 6

- 实现请求参数校验：检查必填字段、长度限制、格式要求
- 校验失败返回 400 和具体错误信息
- 输出：创建文章接口有完整的参数校验

### Day 7

- 周复盘
- 判断 Week 4 需要补哪些短板

## 本周进入条件

满足以下条件再进入 Week 4：

- 能解释中间件的实现原理并自己写一个
- 能说清 context 在请求链路中的作用
- Blog API 有统一的错误返回格式
- 关键接口有基本参数校验

## 本周风险提醒

- 中间件不要写太多，2 到 3 个就够
- context 不要滥用 `context.WithValue`，只传真正需要的请求级信息
- 错误码设计保持简单，不要一开始就搞复杂的错误分类体系
- 参数校验先手写，不要一开始引入校验框架
