# Week 2：mutex、竞态条件、worker pool、context

## 本周目标

- 理解共享内存并发和消息传递并发的区别
- 学会用 `sync.Mutex` 保护共享状态
- 用 race detector 观察并发错误
- 写一个最小 worker pool
- 在任务处理链路里正确传递 `context.Context`

## 本周交付

- 竞态条件 demo
- `sync.Mutex` 修复版 demo
- worker pool demo
- 任务处理服务的 worker 基础实现
- 一次 `go test -race` 使用记录

## 建议日程

### Day 1

- 写一个有共享计数器的并发 demo
- 故意制造竞态条件
- 输出：记录现象，并运行 `go test -race` 或最小测试验证

### Day 2

- 用 `sync.Mutex` 修复共享状态问题
- 理解加锁粒度和临界区的概念
- 输出：写修复版 demo，并对比差异

### Day 3

- 学 worker pool 的最小模型：任务队列、固定 worker、结果收集
- 输出：写一个最小 worker pool 示例

### Day 4

- 把 worker pool 引入主项目
- 先实现任务投递和消费，不急着做重试
- 输出：主项目能异步处理最小任务

### Day 5

- 学 `context.Context` 在并发中的作用
- 把 context 从接口入口一路传到任务执行逻辑
- 输出：能演示取消任务或超时停止执行

### Day 6

- 重构本周代码：梳理哪些地方该用 channel，哪些地方该用 mutex
- 记录 3 条实际判断原则
- 输出：一份简短设计说明

### Day 7

- 做周复盘
- 判断是否满足进入 Week 3 的条件

## 本周进入条件

满足以下条件再进入 Week 3：

- 能解释什么是竞态条件以及为什么会发生
- 能独立写出一个最小 worker pool
- 能在主项目中把 context 传递到任务处理逻辑
- 能说明某个共享状态为什么必须加锁

## 本周风险提醒

- 不要为了用 channel 而绕远路，简单共享状态直接加锁通常更清晰
- 不要把 `context.Background()` 到处乱传，请从入口传入
- worker pool 先做稳定的最小版，不要急着支持动态扩容