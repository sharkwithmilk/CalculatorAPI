# CalculatorAPI

## Общее описание

- CalculatorAPI — это API для вычисления арифметических выражений.
- Поддерживает базовые математические операции: сложение, вычитание, умножение, деление.
- Работает со скобками и числами с плавающей точкой.
## Требования
- [Go](https://go.dev/)
- [Git](https://git-scm.com/)
## Запуск

### Шаги для запуска
- Склонируйте репозиторий:
   ```bash
    git clone https://github.com/sharkwithmilk/CalculatorAPI
  
- Перейдите в корневую папку проекта:
    ```bash
    cd calculator_api

- Установите зависимости:
    ```bash
    go mod tidy
  
- Запустите сервер:
    ```bash
    go run ./cmd/main.go

- По умолчанию сервис доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate)
   
> Примечание: Если порт 8080 уже занят, вы можете указать другой порт, задав переменную PORT:
> PORT=8081 go run ./cmd/main.go
> Примечание: Используйте GitBash(cmd не принимает значение порта) а иначе копируйте без порта: go run ./cmd/main.go        

## Как использовать

### Эндпоинт POST /api/v1/calculate

#### Описание

- Эндпоинт принимает JSON с математическим выражением в поле expression.
- Возвращает результат вычисления или ошибку.

#### Примеры

1. Успешный запрос:
   ```bash
      curl --location 'http://localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
       "expression": "2+2*2"
   }'
   
Ответ:
      {
       "result": :"6.000000"
   }
   

2. Невалидное выражение:
    ```bash
   curl --location 'http://localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
       "expression": "2++2"
   }'
   
Ответ:
      {
       "error": "Expression is not valid"
   }
   

3. Попытка деления на 0:
   ```bash
   curl --location 'http://localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
     "expression": "10/0"
   }'

Ответ:
      {
       "error": "Division by zero"
   }
   

## Особенности

- Поддерживаемые операции: +, -, *, /.
- Обработка скобок для учета приоритета операций.
- Поддержка чисел с плавающей точкой.
- Коды статуса HTTP для ошибок (напр., 422 для невалидного выражения).

## Ошибки

- 200 OK: Успешный запрос. Результат вычисления передается в поле result.
- 422 Unprocessable Entity: Ошибка в выражении. Возможные причины:
  - Неверный формат математического выражения.
  - Попытка деления на 0.
- 500 Internal Server Error: Внутренняя ошибка сервера. Например, из-за некорректного JSON в теле запроса.

## Тестирование

- Для запуска автотестов:
  Переходите в папку application
  ```bash
     cd application
- или в папку pkg/Calculator
   ```bash
     cd pkg/Calculator
- Запускаете тест
  ```bash
    go test -v
  
- Ручное тестирование:
  Используйте curl или Postman для проверки работы API.
