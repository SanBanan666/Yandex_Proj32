# Микросервисная система управления событиями

Система состоит из четырех микросервисов, каждый из которых отвечает за свою часть функциональности. Проект использует Go, PostgreSQL, Redis и Docker.

## Архитектура

### Микросервисы

1. **Auth Service** (порт 8081)
   - Аутентификация и авторизация пользователей
   - Управление пользователями
   - JWT токены

2. **Event Service** (порт 8082)
   - Управление событиями
   - CRUD операции для событий
   - Поиск и фильтрация событий

3. **User Interaction Service** (порт 8083)
   - Управление отзывами
   - Регистрация на события
   - Взаимодействие пользователей с событиями

4. **Notification Service** (порт 8084)
   - Отправка уведомлений
   - Управление подписками
   - Оповещения о событиях

### Базы данных

Каждый микросервис имеет свою собственную базу данных PostgreSQL:
- Auth DB (порт 5436)
- Event DB (порт 5437)
- Interaction DB (порт 5438)
- Notification DB (порт 5439)

### Redis

Используется для кэширования и управления сессиями (порт 6379)

## Требования

- Docker
- Docker Compose
- Go 1.24.2 или выше
- PostgreSQL 15
- Redis 7

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd <repository-name>
```

2. Запустите все сервисы:
```bash
docker-compose up -d
```

3. Проверьте статус сервисов:
```bash
docker-compose ps
```

## API Документация

### Auth Service (http://localhost:8081)

#### Регистрация пользователя
```http
POST /auth/register
Content-Type: application/json

{
    "username": "string",
    "password": "string",
    "email": "string"
}
```

#### Вход в систему
```http
POST /auth/login
Content-Type: application/json

{
    "username": "string",
    "password": "string"
}
```

### Event Service (http://localhost:8082)

#### Создание события
```http
POST /events
Content-Type: application/json

{
    "title": "string",
    "description": "string",
    "date": "string (ISO 8601)",
    "location": "string",
    "creator_id": "integer",
    "capacity": "integer"
}
```

#### Получение списка событий
```http
GET /events
```

#### Получение события по ID
```http
GET /events/{id}
```

#### Обновление события
```http
PUT /events/{id}
Content-Type: application/json

{
    "title": "string",
    "description": "string",
    "date": "string (ISO 8601)",
    "location": "string",
    "capacity": "integer"
}
```

#### Удаление события
```http
DELETE /events/{id}
```

### User Interaction Service (http://localhost:8083)

#### Создание отзыва
```http
POST /reviews
Content-Type: application/json

{
    "event_id": "integer",
    "user_id": "integer",
    "rating": "integer",
    "comment": "string"
}
```

#### Регистрация на событие
```http
POST /registration/event/{id}
Content-Type: application/json

{
    "user_id": "integer"
}
```

#### Получение отзывов о событии
```http
GET /reviews/event/{id}
```

### Notification Service (http://localhost:8084)

#### Отправка уведомления
```http
POST /notifications/send
Content-Type: application/json

{
    "user_id": "integer",
    "type": "string",
    "message": "string"
}
```

#### Получение уведомлений пользователя
```http
GET /notifications/user/{id}
```

## Разработка

### Структура проекта

```
.
├── auth-service/
│   ├── internal/
│   │   ├── database/
│   │   ├── models/
│   │   └── redis/
│   ├── main.go
│   └── Dockerfile
├── event-service/
│   ├── internal/
│   │   ├── database/
│   │   ├── models/
│   │   └── redis/
│   ├── main.go
│   └── Dockerfile
├── user-interaction-service/
│   ├── internal/
│   │   ├── database/
│   │   ├── models/
│   │   └── redis/
│   ├── main.go
│   └── Dockerfile
├── notification-service/
│   ├── internal/
│   │   ├── database/
│   │   ├── models/
│   │   └── redis/
│   ├── main.go
│   └── Dockerfile
├── docker-compose.yml
└── README.md
```

### Локальная разработка

1. Установите зависимости:
```bash
go mod download
```

2. Запустите тесты:
```bash
go test ./...
```

3. Запустите сервис локально:
```bash
go run main.go
```

## CI/CD

Проект использует GitHub Actions для автоматизации процессов CI/CD:
- Автоматические тесты при пуше и создании PR
- Сборка и публикация Docker образов
- Автоматическое развертывание

## Мониторинг

Для мониторинга состояния сервисов используйте:
```bash
docker-compose ps
docker-compose logs -f
```

## Остановка сервисов

```bash
docker-compose down
```

## Лицензия

MIT 