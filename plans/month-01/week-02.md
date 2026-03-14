# Week 2：常用数据结构、错误处理、JSON

## 本周目标

- 掌握后端常用的 Go 基础能力
- 建立对 slice、map、pointer 的直觉
- 能完成 JSON 编解码
- 用一个小练习把数据处理串起来

## 本周交付

- 至少 3 个关于 slice、map、pointer 的 demo
- 至少 1 个 JSON 读写示例
- 至少 1 个小练习，例如命令行 Todo 或数据转换器

## 建议日程

### Day 1

- 学 slice 的长度、容量、append、拷贝行为
- 输出：写 demo 验证容量变化和共享底层数组

### Day 2

- 学 map 的基本操作和零值语义
- 输出：写 demo 验证 key 不存在时的行为

### Day 3

- 学 pointer 和值传递
- 输出：写 demo 验证传值与传指针修改差异

### Day 4

- 学 defer、错误包装、基础资源释放
- 输出：写一个带文件读取或资源关闭的例子

### Day 5

- 学 struct tag 与 JSON 编解码
- 输出：写 Task 的 JSON encode/decode 示例

### Day 6

- 把前面内容串成一个小练习
- 输出：完成一个可运行的小程序

### Day 7

- 周复盘
- 判断是否进入 Week 3

## 本周进入条件

- 能说清 slice 和数组的区别
- 能解释 map 的常见使用方式
- 能判断什么时候该用指针
- 能处理 JSON 请求或文件中的数据结构