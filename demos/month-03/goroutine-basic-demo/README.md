# goroutine basic demo

## 这个 demo 验证什么

- `go f()` 只是把函数放到新的 goroutine 里调度，不代表它会立刻跑完
- 主 goroutine 如果先结束，进程就会退出，其他 goroutine 不会被自动等完
- 要观察稳定结果，必须显式等待，例如用 `sync.WaitGroup`

## 怎么运行

```bash
go run .
```

## 你应该观察什么

- `case 1` 里，`main` 和 `worker` 的打印先后不是你手写顺序的简单复制，而是调度结果
- `case 1` 里如果去掉那 10ms 的 `Sleep`，有时你甚至看不到 `worker` 的输出，因为主 goroutine 太快结束了
- `case 2` 里由于 `WaitGroup` 明确等待，`worker: finished` 一定出现在 `main: worker joined` 前面

## Java 线程池心智 vs Go goroutine 心智

- Java 线程池更像“先有一组工作线程，再把任务提交进去执行”
- Go goroutine 更像“先写并发任务，再让 runtime 决定如何把大量 goroutine 映射到少量线程上跑”
- 线程池主要解决线程复用和资源控制；goroutine 主要提供更轻量的并发组织方式
- 在 Go 里，`go` 关键字不是线程池提交动作，也不自带结果收集，等待和通信要靠 `WaitGroup`、channel、context 等机制补齐