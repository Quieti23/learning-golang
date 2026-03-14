# Week 2：数据库接入、SQL 操作、事务

## 本周目标

- 接入真实数据库，脱离纯内存存储
- 掌握 Go 中 `database/sql` 的基本用法
- 理解连接池配置和事务操作
- 在 Blog API 中完成数据库 CRUD

## 本周交付

- 数据库连接和基础 SQL demo
- Blog API 的 repository 层从内存切换到数据库
- 至少完成文章的增删改查数据库操作

## 建议日程

### Day 1

- 选定数据库：推荐 PostgreSQL 或 SQLite（本地学习用 SQLite 更轻量）
- 学 `database/sql` 的基本连接方式
- 输出：写一个连接数据库并执行简单查询的 demo

### Day 2

- 学建表和 migration：手动写 SQL 建表或用简单 migration 工具
- 设计 `posts` 表结构
- 输出：数据库中有可用的表

### Day 3

- 学 `database/sql` 的 Query、QueryRow、Exec 用法
- 实现 Blog repository 的 Create 和 List 方法
- 输出：创建文章和查询列表走数据库

### Day 4

- 实现 GetByID、Update、Delete 方法
- 学习处理 `sql.ErrNoRows` 等常见错误
- 输出：CRUD 全部走数据库

### Day 5

- 学连接池配置：`SetMaxOpenConns`、`SetMaxIdleConns`、`SetConnMaxLifetime`
- 学事务基础：`db.BeginTx`、`tx.Commit`、`tx.Rollback`
- 输出：写一个事务操作 demo，理解事务失败时的回滚行为

### Day 6

- 整合：Blog API 完全使用数据库存储
- 验证所有接口正常工作
- 处理数据库错误到 HTTP 响应的映射
- 输出：Blog API 的数据持久化可用

### Day 7

- 周复盘
- 判断是否进入 Week 3

## 本周进入条件

满足以下条件再进入 Week 3：

- Blog API 的数据存储已经切换到数据库
- 能解释 `database/sql` 中连接池的作用
- 能写基本事务操作
- 能正确处理数据库层面的错误并返回合理响应

## 本周风险提醒

- 不要一开始就追 ORM，先用 `database/sql` 理解底层原理
- SQLite 适合本地学习，生产项目一般用 PostgreSQL 或 MySQL
- 数据库错误处理要认真对待，不要直接把内部错误暴露给客户端
- 如果装数据库环境卡住，先用 SQLite 推进，不要在环境上花太多时间
