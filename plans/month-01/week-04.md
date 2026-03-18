# Week 4：完善 Todo API、测试、月度复盘

## 本周目标

- 完成 Todo API 的核心闭环
- 建立基础测试意识
- 做第 1 个月阶段复盘

## 本周交付

- Todo API 基础 CRUD 完成
- 关键逻辑测试和 handler 测试若干
- 一次项目复盘
- 一次月复盘

## 建议日程

### Day 1

- 补全创建、查询、更新、删除接口
- 输出：CRUD 主流程跑通

### Day 2

- 整理错误返回结构和参数校验
- 输出：接口行为更稳定

### Day 3

- 学 table-driven tests
- 输出：为纯逻辑或 service 层补测试

### Day 4

- 为 handler 补 1 到 2 个关键路径测试
- 输出：掌握基础 HTTP 测试写法

### Day 5

- 重构明显别扭的结构，但不做大改
- 输出：代码更清晰

### Day 6

- 做项目复盘
- 输出：记录设计取舍、问题和后续优化点

### Day 7

- 做月复盘
- 判断是否进入第 2 个月

## 本周完成标准

- Todo API 能跑通基础 CRUD
- 至少有 3 到 5 个有效测试用例
- 你能讲清项目的结构和下一步改进方向

## 月末判断问题

- 如果面试官问你怎么用 Go 写一个简单 HTTP 服务，你能否说清并写出骨架
- 如果让你解释 Go 错误处理和 Java 异常差异，你能否举例说明
- 如果让你展示一个练习项目，你是否有东西可以打开就讲

第一问，你已经基本可以回答“怎么用 Go 写一个简单 HTTP 服务”，而且也能写出骨架。最小骨架你已经实际做过，在 demos/month-01/day11/http-ping-demo/main.go 和 projects/month-01-todo-api/main.go。你现在至少应该能说清这条主线：用标准库 net/http，先定义 main，再注册路由和 handler，最后用 http.ListenAndServe 启动服务。handler 通过 http.ResponseWriter 写响应，通过 *http.Request 读请求。如果面试官让你现场写，一个最小 /ping 服务你应该已经能写出来。

第二问，你也已经可以解释 Go 错误处理和 Java 异常差异，而且这是你这个阶段的一个优势点。你的回答不用太大，抓住 4 个点就够了：Java 常见是 try/catch 和异常链，Go 更常见是函数显式返回 error；Java 容易把错误处理藏在调用链里，Go 会在调用处显式判断 if err != nil；Go 代码会更啰嗦，但错误路径更直接；Go 不鼓励把常规业务问题都走异常机制。你前面做过的错误处理例子就在 demos/month-01/day03/error.go 和 projects/month-01-todo-api/handler/task_handler.go，完全可以拿“空标题返回业务错误”“任务不存在返回明确错误”来举例。

第三问，你现在已经有一个可以打开就讲的练习项目，就是 projects/month-01-todo-api。而且这个项目不只是 demo，已经具备了可讲述的结构：入口在 projects/month-01-todo-api/main.go，HTTP 层在 projects/month-01-todo-api/handler/task_handler.go，业务层在 projects/month-01-todo-api/service/task_service.go，存储层在 projects/month-01-todo-api/store/task_store.go。如果让你讲项目，你可以按“做了什么接口”“为什么这样分层”“怎么做错误返回”“怎么做测试”这条顺序讲，复盘材料也已经在 projects/month-01-todo-api/PROJECT-REVIEW.md。