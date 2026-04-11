# channel buffered demo

## 这个 demo 验证什么

- buffered channel 有容量，缓冲区没满时发送可以先完成，不必立刻等接收方
- 缓冲区满了以后，再发送就会阻塞，直到有人接收腾出空位
- 缓冲区空了以后，再接收也会阻塞，直到有人发送新值进来

## 怎么运行

```bash
go run .
```

## 你应该观察什么

- 前两次发送 `task-1` 和 `task-2` 会直接完成，因为容量是 2
- `sender: trying to send task-3 into full buffer` 出现后，发送方会卡住，直到主流程先做一次接收
- 第一次接收后看到的 `len` 仍然可能是 2，因为刚腾出的空位会立刻被被阻塞的 `task-3` 发送补上
- 清空缓冲区后，`main: trying to receive from empty buffer` 会先出现，然后要等 `late sender` 发来 `task-4` 才能继续

## 典型用途

- 做短时间削峰，让生产方和消费方不必每次都严格同步
- 做有限排队，但它不是无限队列，容量打满后一样会形成背压
- 当你需要明确控制吞吐和阻塞边界时，比随手开 goroutine 更容易建立心智模型