# Week 1：分层设计、配置管理、结构化日志

## 本周目标

- 理解后端服务分层的基本原则
- 用 interface 实现层间解耦
- 引入配置文件管理服务参数
- 接入结构化日志替代 fmt 输出
- 搭建 Blog API 项目骨架

## 本周交付

- Blog API 项目初始目录和分层结构
- 配置管理 demo
- 结构化日志 demo
- 至少一个可调用的接口

## 建议日程

### Day 1

- 回顾 Todo API 的结构，识别哪些地方不够清晰
- 学习 handler → service → repository 分层模式
- 输出：画一张简单依赖图，理解每层职责

### Day 2

- 初始化 Blog API 项目，建立分层目录
- 定义核心 model：`Post`（标题、内容、作者、创建时间、更新时间）
- 输出：`projects/month-02-blog-api/` 基本骨架建好

### Day 3

- 学习用 interface 定义 repository 和 service 的契约
- 先用内存实现 repository
- 输出：service 通过 interface 调用 repository，handler 调用 service

### Day 4

- 学配置管理：用环境变量或配置文件读取端口、数据库地址等
- 推荐尝试 `github.com/spf13/viper` 或简单的 JSON/YAML 配置
- 输出：写一个配置加载 demo，Blog API 的端口从配置读取

### Day 5

- 学结构化日志：推荐 `log/slog`（Go 1.21+ 标准库）或 `go.uber.org/zap`
- 输出：写一个日志 demo，Blog API 中替换所有 fmt 输出为结构化日志

### Day 6

- 串联前面内容：Blog API 具备分层结构、配置读取、结构化日志
- 实现创建文章和查询文章列表接口
- 输出：至少 2 个接口可调用

### Day 7

- 做周复盘
- 判断是否满足进入 Week 2 的条件

## 本周进入条件

满足以下条件再进入 Week 2：

- 能解释 handler、service、repository 各自的职责
- 项目配置不再硬编码
- 日志输出是结构化的，不是 fmt.Println
- Blog API 骨架可运行，至少有 1 个接口

## 本周风险提醒

- 不要在分层上过度设计，先做最小可用的分层
- 配置管理不需要一开始就支持所有来源，先支持一种即可
- 日志库选择不纠结，标准库 `slog` 够用就先用
