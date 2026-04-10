# channel unbuffered blocking demo

## 这个 demo 验证什么

- 如果 unbuffered channel 没有接收方，发送会阻塞在 `ch <- value`
- 发送 goroutine 不会继续往下执行，直到某个接收方真的来取值
- 如果接收方永远不来，程序最终会陷入 deadlock

## 怎么运行

```bash
go run .
```

## 你应该观察什么

- `sender: trying to send 42` 出现后，`sender: send finished...` 不会立刻出现
- 中间会先看到 `main: sleep 100ms, no receiver yet` 和 `main: ready to receive`
- 只有 `main` 执行 `<-ch` 之后，发送方才继续打印，说明之前一直卡在发送动作上