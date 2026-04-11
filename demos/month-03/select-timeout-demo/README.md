# select timeout demo

## 这个 demo 验证什么

- `select` 可以同时等待多个通信分支，谁先就绪就执行谁
- `time.After(duration)` 会返回一个只读 channel，到时间后会送出一个信号
- 把 `time.After` 放进 `select`，就能给等待操作加一个超时出口

## 怎么运行

```bash
go run .
```

## 你应该观察什么

- `fast job` 的结果在 100ms 左右就准备好了，早于 300ms 超时，所以会走结果分支
- `slow job` 需要 400ms，但超时只给 200ms，所以会先走超时分支
- `select` 不是轮询语法糖，而是“等待多个事件，谁先准备好就处理谁”

## 一个实用心智

- channel 分支表示“正常结果到了”
- `time.After` 分支表示“我最多只等这么久”
- 这和 Java 里 future/get timeout 的心智接近，但 Go 通常直接把等待和超时写在并发控制流里