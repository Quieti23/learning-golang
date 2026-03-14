# 第 1 个月计划

第 1 个月的目标不是学完整个 Go，而是建立 Go 后端的最小工作能力。

## 本月目标

本月结束时，你需要达到以下状态：

- 能独立写出一个基础 Go HTTP 服务
- 能理解 Go 中 struct、method、interface、slice、map、pointer、error 的常见使用方式
- 能使用标准库处理 JSON 请求和响应
- 能完成一个内存版 Todo API
- 能写基础 table-driven tests
- 能说清楚 Java 与 Go 在错误处理、抽象方式、工程组织上的关键差异

## 本月学习主线

主线只有一条：围绕 Todo API 项目学习 Go 基础。

不是先学完语法再做项目，而是边学边做：

- 语法点用 demo 验证
- 常用能力用 exercise 固化
- 项目中真正落地
- 每周复盘判断是否进入下一周

## 周计划概览

### Week 1：环境、基础语法、开发方式入门

目标：建立最基本的 Go 手感。

交付：

- 完成开发环境和 `go mod` 初始化
- 写出基础语法 demo
- 理解 package、module、struct、method、error 的基本写法

### Week 2：数据结构、函数、错误处理、JSON

目标：开始写更接近后端的代码。

交付：

- 完成 slice、map、pointer、interface、defer 的练习
- 能读写 JSON
- 能写一个简单命令行 Todo 或小型数据处理示例

### Week 3：HTTP 服务与路由处理

目标：进入后端主场景。

交付：

- 写出最小 HTTP server
- 处理 query、path、body、status code
- 开始搭 Todo API 骨架

### Week 4：完善 Todo API + 基础测试 + 月复盘

目标：形成第一个可展示成果。

交付：

- 完成 Todo API 基础功能
- 补核心测试
- 做一次项目复盘和月复盘

## 每周完成定义

每周都必须满足：

- 有代码产出
- 有最小可运行示例
- 有至少一次复盘
- 有下周进入条件判断

## 每日执行模板

每天 2 小时建议这样分配：

- 20 分钟：看当天任务和资料
- 70 分钟：编码和实验
- 20 分钟：整理结论和问题
- 10 分钟：安排明天起点

## 本月验收标准

月末时，检查自己是否满足以下条件：

- 不看资料也能写出一个最小 HTTP 服务
- 能解释为什么 Go 更偏好显式错误处理
- 能说出至少 5 个 Java 转 Go 常见坑
- 能独立完成 Todo API 的主要接口
- 能写 3 到 5 个基础测试用例

## 文件导航

- `plans/month-01/week-01.md`
- `plans/month-01/week-02.md`
- `plans/month-01/week-03.md`
- `plans/month-01/week-04.md`
- `plans/month-01/checklist.md`

## 开始方式

如果今天是第 1 天，直接从 Week 1 的 Day 1 开始，不要继续扩展计划。先执行，再根据真实进度调整。