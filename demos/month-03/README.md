# 第 3 个月 demos

这个目录用于放第 3 个月的最小并发实验代码。

建议按主题拆小目录，例如：

- `goroutine-basic-demo/`
- `channel-buffered-demo/`
- `select-timeout-demo/`
- `waitgroup-demo/`
- `mutex-race-demo/`
- `worker-pool-demo/`
- `context-timeout-demo/`
- `rate-limit-demo/`

要求：

- 每个 demo 只验证一个核心点
- 每个 demo 都能单独运行
- 代码旁边写下你观察到的行为，不只保留代码