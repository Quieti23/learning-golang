# month-02-blog-api

第 2 个月 Blog API 项目骨架。

## 当前结构

- `main.go`：装配依赖并启动 HTTP 服务
- `config/`：读取项目配置
- `model/`：定义 `Post`
- `repository/`：定义 repository 契约和内存实现
- `service/`：处理业务规则
- `handler/`：处理 HTTP 请求和响应

## 当前日志

- 使用标准库 `log/slog`
- 启动日志和关键请求日志使用结构化字段输出

## 当前配置

- `config.json`：当前用来配置服务端口

## 当前接口

- `GET /ping`
- `GET /posts`
- `POST /posts`

## 运行方式

```powershell
go run .
```

默认会读取项目根目录下的 `config.json`：

```json
{
	"server_port": "8081"
}
```
