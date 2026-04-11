# Week 1：goroutine、channel、select、waitgroup

## 本周目标

- 建立 Go 并发编程的最小心智模型
- 理解 goroutine 和线程的区别，不把它神化
- 学会用 channel 做简单通信
- 学会用 `select` 处理多路等待和超时
- 学会用 `sync.WaitGroup` 等待并发任务收敛

## 本周交付

- goroutine demo
- buffered 和 unbuffered channel demo
- `select` 超时 demo
- `WaitGroup` 协作 demo
- 记录 3 个最常见并发误区

## 建议日程

### Day 1

- 学 goroutine 的启动方式和调度直觉
- 对比 Java 线程池心智和 Go goroutine 心智
- 输出：写 1 个最小 goroutine demo，观察输出顺序
- 参考产物：`demos/month-03/goroutine-basic-demo`

### Day 2

- 学 unbuffered channel 的发送接收语义
- 理解为什么没有接收方时发送会阻塞
- 输出：写 2 个 demo，分别验证同步通信和阻塞现象
- 参考产物：`demos/month-03/channel-unbuffered-sync-demo`
- 参考产物：`demos/month-03/channel-unbuffered-blocking-demo`

### Day 3

- 学 buffered channel 的容量、阻塞边界和典型用途
- 输出：写 1 个 buffered channel demo，记录满和空时的行为
- 参考产物：`demos/month-03/channel-buffered-demo`

### Day 4

- 学 `select` 的基本写法
- 用 `time.After` 做超时控制
- 输出：写 1 个超时等待 demo

### Day 5

- 学 `sync.WaitGroup` 的使用方式
- 理解它适合做等待，不负责传值
- 输出：写 1 个并发执行后统一收尾的 demo

### Day 6

- 回到主项目，设计任务处理服务的数据流：任务提交、入队、消费、结果回写
- 先不要写复杂业务，先画清楚任务生命周期
- 输出：一张简化流程图或一份文字版流程说明

### Day 7

- 做周复盘
- 判断是否满足进入 Week 2 的条件

## 本周进入条件

满足以下条件再进入 Week 2：

- 能解释 goroutine、channel、waitgroup 各自解决的问题
- 能独立写出一个 `select + time.After` 的超时 demo
- 能说清 buffered channel 和 unbuffered channel 的区别
- 已经设计出任务处理服务的最小流转过程

## 本周风险提醒

- 不要一开始就看复杂并发模式，先把阻塞和通信语义吃透
- 不要把 channel 当成万能方案，先理解它的代价和适用场景
- 看懂不等于会写，必须亲手写出最小 demo