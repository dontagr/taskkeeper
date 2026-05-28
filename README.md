# Task keeper

Сквозной пример для книги «Вайбкодер»: HTTP API на Go (`net/http`), in-memory storage.

**Спека:** [../book/specs/0001-task-keeper.md](../book/specs/0001-task-keeper.md)

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

- `internal/task` — модель и валидация title
- `internal/storage` — in-memory store
- `internal/httpapi` — HTTP handlers
- `cmd/taskkeeper` — точка входа
