# Микросервисная архитектура для управления событиями

## Архитектура

Проект построен на микросервисной архитектуре и состоит из следующих сервисов:

1. **Auth Service** (Порт: 8081)
  - Управление аутентификацией и авторизацией
  - Работа с пользователями
  - Управление токенами

2. **Event Service** (Порт: 8082)
  - Управление событиями
  - CRUD операции для событий
  - Поиск и фильтрация событий

3. **User Interaction Service** (Порт: 8083)
  - Управление отзывами
  - Регистрация на события
  - Взаимодействие пользователей с событиями

4. **Notification Service** (Порт: 8084)
  - Управление уведомлениями
  - Система подписок
  - Отправка уведомлений

### Базы данных

Каждый сервис имеет свою собственную базу данных PostgreSQL:

- **Auth DB** (Порт: 5436)
  - База данных: authdb
  - Пользователь: authuser
  - Пароль: authpass

- **Event DB** (Порт: 5437)
  - База данных: eventdb
  - Пользователь: eventuser
  - Пароль: eventpass

- **Interaction DB** (Порт: 5438)
  - База данных: interactiondb
  - Пользователь: interactionuser
  - Пароль: interactionpass

- **Notification DB** (Порт: 5439)
  - База данных: notificationdb
  - Пользователь: notificationuser
  - Пароль: notificationpass

### Redis

Все сервисы используют Redis для кэширования и управления сессиями:
- Порт: 6379
- Без пароля
- База данных: 0

## API Endpoints

### Auth Service (8081)

#### Регистрация
- **POST** `/auth/register`
  - Регистрация нового пользователя
  - Тело запроса: `{ "email": string, "password": string, "name": string }`

#### Вход
- **POST** `/auth/login`
  - Аутентификация пользователя
  - Тело запроса: `{ "email": string, "password": string }`

#### Обновление токена
- **POST** `/auth/refresh`
  - Обновление JWT токена
  - Тело запроса: `{ "refresh_token": string }`

### Event Service (8082)

#### События
- **POST** `/events`
  - Создание нового события
  - Тело запроса: `{ "title": string, "description": string, "date": string, "location": string }`

- **GET** `/events`
  - Получение списка событий
  - Параметры запроса: `page`, `limit`, `sort`

- **GET** `/events/:id`
  - Получение информации о конкретном событии
  - Параметры пути: `id`

- **PUT** `/events/:id`
  - Обновление события
  - Параметры пути: `id`
  - Тело запроса: `{ "title": string, "description": string, "date": string, "location": string }`

- **DELETE** `/events/:id`
  - Удаление события
  - Параметры пути: `id`

### User Interaction Service (8083)

#### Отзывы
- **POST** `/reviews`
  - Создание отзыва
  - Тело запроса: `{ "event_id": string, "rating": number, "comment": string }`

- **GET** `/reviews/event/:id`
  - Получение отзывов о событии
  - Параметры пути: `id`

- **PUT** `/reviews/:id`
  - Обновление отзыва
  - Параметры пути: `id`
  - Тело запроса: `{ "rating": number, "comment": string }`

- **DELETE** `/reviews/:id`
  - Удаление отзыва
  - Параметры пути: `id`

#### Регистрация на события
- **POST** `/registration/event/:id`
  - Регистрация на событие
  - Параметры пути: `id`

- **DELETE** `/registration/event/:id`
  - Отмена регистрации на событие
  - Параметры пути: `id`

- **GET** `/registration/event/:id`
  - Получение списка регистраций на событие
  - Параметры пути: `id`

### Notification Service (8084)

#### Уведомления
- **POST** `/notifications/send`
  - Отправка уведомления
  - Тело запроса: `{ "user_id": string, "title": string, "message": string }`

- **GET** `/notifications/user/:id`
  - Получение уведомлений пользователя
  - Параметры пути: `id`

- **PUT** `/notifications/:id/read`
  - Отметка уведомления как прочитанного
  - Параметры пути: `id`

- **DELETE** `/notifications/:id`
  - Удаление уведомления
  - Параметры пути: `id`

#### Подписки
- **POST** `/subscriptions`
  - Создание подписки
  - Тело запроса: `{ "user_id": string, "event_id": string }`

- **DELETE** `/subscriptions/:id`
  - Удаление подписки
  - Параметры пути: `id`

- **GET** `/subscriptions/user/:id`
  - Получение подписок пользователя
  - Параметры пути: `id`

## Технические решения

### Микросервисная архитектура
- Каждый сервис независим и может быть развернут отдельно
- Использование Docker для контейнеризации
- Отдельные базы данных для каждого сервиса

### Безопасность
- JWT для аутентификации
- Redis для управления сессиями
- Изоляция баз данных

### Масштабируемость
- Горизонтальное масштабирование сервисов
- Кэширование через Redis
- Асинхронная обработка уведомлений

### Мониторинг
- Логирование всех операций
- Метрики производительности
- Отслеживание ошибок

## Запуск проекта

1. Установите Docker и Docker Compose
2. Клонируйте репозиторий
3. Запустите все сервисы:
```bash
docker-compose up -d
```

## Разработка

### Требования
- Go 1.21 или выше
- Docker
- Docker Compose
- PostgreSQL
- Redis

### Локальная разработка
1. Установите зависимости:
```bash
go mod download
```

2. Запустите базы данных:
```bash
docker-compose up -d auth-db event-db interaction-db notification-db redis
```

3. Запустите сервисы:
```bash
go run auth-service/main.go
go run event-service/main.go
go run user-interaction-service/main.go
go run notification-service/main.go
```

## Тестирование

Запуск тестов:
```bash
go test ./...
```

## CI/CD

Проект использует GitHub Actions для автоматизации:
- Тестирование
- Сборка Docker образов
- Деплой

 
 
 