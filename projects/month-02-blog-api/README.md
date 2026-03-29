# month-02-blog-api

第 2 个月 Blog API 项目骨架。

## 当前结构

- `main.go`：装配依赖并启动 HTTP 服务
- `config/`：读取项目配置
- `cmd/migrate/`：最小 migration 执行器
- `model/`：定义 `Post`
- `migrations/`：建表 SQL
- `repository/`：定义 repository 契约，以及内存/MySQL 两种实现
- `service/`：处理业务规则
- `handler/`：处理 HTTP 请求和响应

## 当前日志

- 使用标准库 `log/slog`
- 启动日志和关键请求日志使用结构化字段输出
- 所有 HTTP 请求都会经过日志中间件，自动记录请求方法、路径、状态码、耗时

## 当前配置

- `config.json`：当前用来配置服务端口和 MySQL DSN
- 也包含连接池配置：`max_open_conns`、`max_idle_conns`、`conn_max_lifetime_minutes`
- `request_timeout_ms`：控制单个 HTTP 请求传入数据库操作的超时时间

## 当前接口

- `GET /ping`
- `GET /posts`
- `POST /posts`
- `GET /posts/{id}`
- `PUT /posts/{id}`
- `DELETE /posts/{id}`

## 当前数据库实现

- Blog API 当前已经切到 MySQL repository
- handler 会把请求 `context` 传到 service 和 repository
- `Create`、`List`、`GetByID`、`Update`、`Delete` 都通过 `QueryContext`、`QueryRowContext`、`ExecContext` 访问 `posts` 表
- `sql.ErrNoRows` 会在 repository 层转换成 `post not found`
- 超时会返回 `504 request timeout`
- handler 层会把常见错误映射成统一结构：`{"code":"...","message":"..."}`
- `main.go` 会应用基础连接池配置

## 当前错误返回

- 统一错误结构：`{"code":"NOT_FOUND","message":"post not found"}`
- 业务错误：`INVALID_REQUEST`、`NOT_FOUND`
- 系统错误：`REQUEST_TIMEOUT`、`REQUEST_CANCELED`、`INTERNAL_ERROR`
- 非法方法：`METHOD_NOT_ALLOWED`

## 当前数据库迁移

- `migrations/001_create_posts_table.up.sql`：创建 `posts` 表
- `migrations/001_create_posts_table.down.sql`：删除 `posts` 表

执行方式：

```powershell
$env:MYSQL_DSN = "root:password@tcp(127.0.0.1:3306)/your_database?parseTime=true"
go run ./cmd/migrate up
```

回滚方式：

```powershell
go run ./cmd/migrate down
```

## 运行方式

```powershell
go run .
```

启动后可以先验证：

```powershell
Invoke-RestMethod http://localhost:8081/ping
```

默认会读取项目根目录下的 `config.json`：

```json
{
	"server_port": "8081",
	"mysql_dsn": "root:password@tcp(127.0.0.1:3306)/your_database?parseTime=true",
	"max_open_conns": 10,
	"max_idle_conns": 5,
	"conn_max_lifetime_minutes": 5,
	"request_timeout_ms": 3000
}
```

## 事务 demo

可以运行最小事务示例：

```powershell
go run ./cmd/txdemo
```

这个 demo 会：

- 打开数据库连接并应用连接池配置
- 开启事务
- 向 `posts` 表插入一条临时数据
- 人为触发回滚
- 再检查回滚后数据是否未落库

## Day 6 手动验收

1. 启动服务：

```powershell
go run .
```

2. 新建文章：

```powershell
Invoke-RestMethod -Method Post http://localhost:8081/posts \
	-ContentType 'application/json' \
	-Body '{"title":"first post","content":"hello mysql","author":"eson"}'
```

3. 查询列表：

```powershell
Invoke-RestMethod http://localhost:8081/posts
```

4. 查询单篇文章：

```powershell
Invoke-RestMethod http://localhost:8081/posts/1
```

5. 更新文章：

```powershell
Invoke-RestMethod -Method Put http://localhost:8081/posts/1 \
	-ContentType 'application/json' \
	-Body '{"title":"updated post","content":"updated content","author":"eson"}'
```

6. 删除文章：

```powershell
Invoke-WebRequest -Method Delete http://localhost:8081/posts/1
```

7. 验证错误映射：

```powershell
Invoke-WebRequest http://localhost:8081/posts/999
Invoke-WebRequest -Method Post http://localhost:8081/posts \
	-ContentType 'application/json' \
	-Body '{}'
Invoke-WebRequest -Method Post http://localhost:8081/posts \
	-ContentType 'application/json' \
	-Body '{"title":"x","content":"y","author":"z","extra":"bad"}'
```

预期结果：

- 不存在的文章返回 `404`
- 缺少必填字段返回 `400`
- 多余字段或非法 JSON 返回 `400`
- 数据库调用超过 `request_timeout_ms` 时返回 `504`
- 正常 CRUD 返回 `200 / 201 / 204`

错误响应示例：

```json
{"code":"NOT_FOUND","message":"post not found"}
```
