# channel unbuffered sync demo

## 这个 demo 验证什么

- unbuffered channel 没有缓冲区，发送和接收必须当场配对
- 接收方已经在等待时，发送动作才能立刻完成
- 这是一种同步通信，不只是“把值塞进去以后马上走人”

## 怎么运行

```bash
go run .
```

## 你应该观察什么

- `receiver: waiting for value` 先打印，说明接收方已经阻塞等待
- `sender: sending value` 之后，接收和发送几乎配对发生
- `sender: send completed after receiver was ready` 只有在接收动作发生后才会出现
- `sender: send completed...` 有可能先于 `receiver: got value...` 打印，因为 channel 只同步收发动作，不保证后续打印的调度顺序