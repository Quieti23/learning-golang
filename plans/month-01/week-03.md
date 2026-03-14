# Week 3：HTTP 服务与 Todo API 骨架

## 本周目标

- 进入 Go 后端开发主场景
- 能用标准库写最小 HTTP 服务
- 开始搭建 Todo API 骨架

## 本周交付

- 一个最小 HTTP server demo
- Todo API 初始目录结构
- 至少 2 个基础接口可用

## 建议日程

### Day 1

- 学 `net/http` 的最小 server 写法
- 输出：一个 `/ping` 接口返回 `pong`

### Day 2

- 学处理 query、path、method、status code
- 输出：补一个简单查询接口

### Day 3

- 学 request body 解析和 JSON response
- 输出：补一个 POST 接口

### Day 4

- 设计 Todo API 的最小结构：handler、service、store、model
- 输出：完成项目骨架

### Day 5

- 实现创建任务和查询列表
- 输出：两个接口跑通

### Day 6

- 实现查询单个任务或更新状态
- 输出：补 1 到 2 个接口

### Day 7

- 周复盘
- 判断 Week 4 需要补哪些短板

## 本周进入条件

- 能独立写一个最小 HTTP 服务
- 知道 handler 和业务逻辑如何分开
- 能处理 JSON 请求和状态码返回
- Todo API 已经不是空壳，而是有实际接口可调用