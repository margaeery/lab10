# Лабораторная работа №10: Веб-разработка: FastAPI (Python) vs Gin (Go)

**Студент:** Ражина Маргарита Александровна
**Группа:** 220032-11
**Вариант:** 8

## Выполненные задания

### Средняя сложность:
1. **Задание 2:** Добавить middleware для логирования в Go.
Инструкция по запуску
Подготовка зависимостей:
```bash
cd task1_2
go mod tidy
```
Запуск сервера:
```bash
go run main.go
```
Проверка:
Сервер по умолчанию доступен по адресу http://localhost:8080. Проверить эндпоинты можно через браузер или curl:
Статус: curl http://localhost:8080/status
Инфо: curl http://localhost:8080/info
Данные (POST): curl -X POST http://localhost:8080/data
Запуск тестов:
```bash
go test -v
```


2. **Задание 4:** Создать FastAPI-сервис, который вызывает Go-сервис через HTTP.
Инструкция по запуску
Подготовка зависимостей:
```bash
cd task2_4
pip install fastapi uvicorn httpx
```
Запуск сервера:
Запустите FastAPI на стандартном порту 8000 (убедитесь, что Go-сервис уже запущен на порту 8080):
Терминал 1 (Go):
```bash
cd lab10/task1_2
go run main.go
```
Терминал 2 (Python):
```bash
cd lab10/task2_4
uvicorn main:app --port 8000
```
Проверка:
Получение статуса: curl http://localhost:8000/fetch-status
Получение инфо: curl http://localhost:8000/fetch-info
Отправка данных (POST):
curl -X POST http://localhost:8000/send-data \
     -H "Content-Type: application/json" \
     -d '{"test": "hello"}'
Интерактивная документация:
Откройте в браузере: http://localhost:8000/docs
Запуск тестов:
В первом терминале 
```bash
cd lab10/task1_2
go run main.go
```
Во втором терминале
```bash
cd lab10/task2_4
pytest
```


3. **Задание 8:** Добавить Swagger-документацию для FastAPI и OpenAPI для Gin.
Инструкция по запуску (Go / Gin)
Подготовка зависимостей:
Перейдите в папку с проектом и обновите модули для поддержки Swagger:
```bash
cd task3_8
go install github.com/swaggo/swag/cmd/swag@latest
swag init
go mod tidy
```
Запуск сервера:
```bash
go run main.go
```
Проверка документации:
Swagger UI: http://localhost:8080/swagger/index.html
OpenAPI JSON: http://localhost:8080/swagger/doc.json
Инструкция по запуску (Python / FastAPI)
Подготовка зависимостей:
Убедитесь, что установлены необходимые библиотеки (FastAPI генерирует документацию автоматически):
```bash
cd task3_8
pip install fastapi uvicorn httpx
```
Запуск сервера:
Запустите шлюз на порту 8000:
```bash
uvicorn main:app --port 8000 --reload
```
Проверка документации:
Swagger UI (основной): http://localhost:8000/docs
ReDoc (альтернативный): http://localhost:8000/redoc  
Запуск тестов:
Тестирование Go (Gin) Documentation
```bash
cd task3_8
go get github.com/stretchr/testify/assert
go mod tidy
go test -v .
```
Тестирование Python (FastAPI) Documentation
В первом терминале:
```bash
cd task3_8
go run main.go
```
Во втором терминале:
```bash
cd task3_8
uvicorn main:app --port 8000
```
В третьем терминале:
```bash
cd task3_8
pip install pytest-asyncio
pytest test_docs.py
```



### Повышенная сложность:
1. **Задание 2:** Создать API-шлюз на Go, который маршрутизирует запросы к разным микросервисам (Python, Go).
Инструкция по запуску
Подготовка среды
```bash
cd task4_2
go mod tidy
pip install fastapi uvicorn pydantic
```
Запуск компонентов
Для работы системы нужно открыть три терминала и запустить сервисы по отдельности:
Терминал 1: API-шлюз 
```bash
go run gateway/gateway.go
```
Терминал 2: Go-сервис (товары)
```bash
go run service_go/main.go
```
Терминал 3: Python-сервис (аналитика)
```bash
cd service_python
uvicorn analytics:app --port 8001
```
Проверка эндпоинтов (через Шлюз :8000)
Сервис товаров (Go)
Список товаров (GET):curl http://localhost:8000/api/v1/items
Создание товара (POST):Bashcurl -X POST http://localhost:8000/api/v1/items -H "Content-Type: application/json" -d "{\"id\":\"3\", \"name\":\"Клавиатура\", \"price\":2500}"
(После этого снова проверь GET, чтобы увидеть новый товар в списке).
Сервис аналитики (Python)
Статус (GET):curl http://localhost:8000/analytics/health
Проверка данных (POST):Bashcurl -X POST http://localhost:8000/analytics/check -H "Content-Type: application/json" -d "{\"users\":100, \"revenue\":50000.0}"
Ссылки на документацию (Swagger)
Go Service Swagger UI http://localhost:8002/swagger/index.html
Python ServiceFastAPI Docs http://localhost:8001/docs

Запуск тестов:
Подготовка и запуск сервисов
Перед запуском тестов необходимо поднять всю инфраструктуру в разных терминалах(см. Инструкцию по запуску)
Тестирование Go (Проверка Go-сервиса напрямую и через Шлюз):
В четвертом терминале:
```bash
cd task4_2/tests
go test -v go_service_test.go
```
Тестирование Python (Проверка Python-сервиса напрямую и через Шлюз):
В том же четвертом терминале:
```bash
cd task4_2/tests
pip install pytest requests
pytest test_python_service.py
```


2. **Задание 4:** Использовать WebSocket: реализовать чат на Go и подключиться к нему из Python.
Инструкция по запуску
Подготовка среды
```bash
go mod tidy
pip install websockets 
```
Запуск компонентов
Терминал 1: Сервер чата (Go) 
```bash
go run server.go
``` 
Терминал 2: Консольный клиент (Python) 
```bash
python client.py 
```
Терминал 3: Веб-интерфейс (Браузер) 
Найдите в папке проекта файл index.html.
Откройте его двойным кликом в любом современном браузере.

Запуск тестов:
Подготовка и запуск сервисов
Перед выполнением тестов необходимо поднять серверную часть в отдельном терминале (см. Инструкцию по запуску)
```bash
go run server.go
```
Тестирование Go:
В новом терминале:
```bash
go test -v .
```
Тестирование Python:
В том же терминале:
```bash
pip install pytest pytest-asyncio websockets
pytest test_client.py
```