# Task keeper

Сквозной пример для книги **«Вайбкодер»**: HTTP API на Go (`net/http`), in-memory storage, unit-тесты.

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

- `cmd/taskkeeper/main.go` — точка входа
- `internal/httpapi/server.go` — HTTP-сервер
- `internal/storage/memory.go` — in-memory хранилище
- `internal/task/task.go` — доменная модель
- `internal/*/*_test.go` — тесты

## Артефакты книги (`notes/`)

Шаблоны для читателей — [notes/README.md](notes/README.md):

| Файл | Описание |
|------|----------|
| [matrix-tasks.md](notes/matrix-tasks.md) | Матрица задач: когда отдавать ИИ |
| [vibe-journal.md](notes/vibe-journal.md) | Журнал вайбкодера |
| [tools.md](notes/tools.md) | Мой стек ИИ-инструментов |

```bash
git clone https://github.com/dontagr/taskkeeper.git
cd taskkeeper
go test ./...
```
