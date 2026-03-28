# Лабораторная работа №10: Веб-разработка: FastAPI (Python) vs Gin (Go)

**Студент:** Ражина Маргарита Александровна
**Группа:** 220032-11
**Вариант:** 8

## Выполненные задания

### Средняя сложность:
1. **Задание 2:** Добавить middleware для логирования в Go.
Инструкция по запуску
Подготовка зависимостей:
cd task1_2
go mod tidy
Запуск сервера:
go run main.go
Проверка:
Сервер по умолчанию доступен по адресу http://localhost:8080. Проверить эндпоинты можно через браузер или curl:
Статус: curl http://localhost:8080/status
Инфо: curl http://localhost:8080/info
Данные (POST): curl -X POST http://localhost:8080/data
Запуск тестов:
go test -v


2. **Задание 4:** Создать FastAPI-сервис, который вызывает Go-сервис через HTTP.
Инструкция по запуску
Подготовка зависимостей:
cd task4
pip install fastapi uvicorn httpx
Запуск сервера:
Запустите FastAPI на стандартном порту 8000 (убедитесь, что Go-сервис уже запущен на порту 8080):
Терминал 1 (Go):
cd lab10/task1_2
go run main.go
Терминал 2 (Python):
cd lab10/task2_4
uvicorn main:app --port 8000
Проверка:
Получение статуса: curl http://localhost:8000/fetch-status
Получение инфо: curl http://localhost:8000/fetch-info
Отправка данных (POST):
Bash
curl -X POST http://localhost:8000/send-data \
     -H "Content-Type: application/json" \
     -d '{"test": "hello"}'
Интерактивная документация:
Откройте в браузере: http://localhost:8000/docs
Запуск тестов:
В первом терминале 
cd lab10/task1_2
go run main.go
Во втором терминале
cd lab10/task2_4
pytest


3. **Задание 8:** Добавить Swagger-документацию для FastAPI и OpenAPI для Gin.
Инструкция по запуску (Go / Gin)
Подготовка зависимостей:
Перейдите в папку с проектом и обновите модули для поддержки Swagger:
cd task3_8
go install github.com/swaggo/swag/cmd/swag@latest
swag init
go mod tidy
Запуск сервера:
go run main.go
Проверка документации:
Swagger UI: http://localhost:8080/swagger/index.html
OpenAPI JSON: http://localhost:8080/swagger/doc.json
Инструкция по запуску (Python / FastAPI)
Подготовка зависимостей:
Убедитесь, что установлены необходимые библиотеки (FastAPI генерирует документацию автоматически):
cd task3_8
pip install fastapi uvicorn httpx
Запуск сервера:
Запустите шлюз на порту 8000:
uvicorn main:app --port 8000 --reload
Проверка документации:
Swagger UI (основной): http://localhost:8000/docs
ReDoc (альтернативный): http://localhost:8000/redoc  
Запуск тестов:
Тестирование Go (Gin) Documentation
cd task3_8
go get github.com/stretchr/testify/assert
go mod tidy
go test -v .
Тестирование Python (FastAPI) Documentation
В первом терминале:
cd task3_8
go run main.go
Во втором терминале:
cd task3_8
uvicorn main:app --port 8000
В третьем терминале:
cd task3_8
pip install pytest-asyncio
pytest test_docs.py




### Повышенная сложность:
1. **Задание 2:** Создать API-шлюз на Go, который маршрутизирует запросы к разным микросервисам (Python, Go).
   - Реализована маршрутизация входящих запросов между микросервисами на Python и Go.
2. **Задание 4:** Использовать WebSocket: реализовать чат на Go и подключиться к нему из Python.
   - Сервер на Go обрабатывает соединения, Python-скрипт подключается как клиент.

