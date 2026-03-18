# month-01-todo-api

最小 Todo API，使用 Go 标准库 HTTP、内存存储和简单分层结构实现。

## 当前接口

- `GET /ping`
- `GET /tasks`
- `POST /tasks`
- `GET /tasks/{id}`
- `PUT /tasks/{id}`
- `DELETE /tasks/{id}`

## 运行方式

```powershell
go run .
```

## 测试方式

```powershell
go test ./...
```

## 复盘

- 项目复盘见 `PROJECT-REVIEW.md`
