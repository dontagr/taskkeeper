# Task keeper

Сквозной пример для книги «Вайбкодер»: HTTP API на Go (`net/http`), in-memory storage.

**Спека:** [specs/0001-task-keeper.md](specs/0001-task-keeper.md)

## Запуск

```bash
go run ./cmd/taskkeeper
```

## Примеры

```bash
curl -s -X POST localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk"}'

curl -s localhost:8080/tasks/t_000001
curl -s -i localhost:8080/tasks/t_999999
```

## Тесты

```bash
go test ./...
```

## Структура

- Точка входа приложения в cmd/taskkeeper/main.go
- HTTP-сервер в internal/httpapi/server.go
- In-memory хранилище в internal/storage/memory.go
- Доменная модель задачи в internal/task/task.go
- Набор автотестов для ключевых модулей:
  - internal/httpapi/server_test.go
  - internal/storage/memory_test.go
  - internal/task/task_test.go
