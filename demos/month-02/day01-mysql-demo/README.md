# day01-mysql-demo

最小 MySQL `database/sql` 连接示例。

## 目标

- 用 `database/sql` 连接 MySQL
- 用 `Ping()` 验证连接
- 执行一个简单查询

## 使用方式

先设置环境变量：

```powershell
$env:MYSQL_DSN = "root:password@tcp(127.0.0.1:3306)/mysql?parseTime=true"
```

然后运行：

```powershell
go run .
```

## 说明

- 这里默认连接 MySQL 自带的 `mysql` 库，仅用于做最小连接演示。
- `SELECT VERSION()` 用来确认真的连到了 MySQL。
- `SELECT 1` 用来演示最小查询闭环。
