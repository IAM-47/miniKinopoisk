# miniKinopoisk

REST API для каталога фильмов с актёрами и финансовой статистикой. Написан на Go без фреймворков — только стандартная библиотека, pgx и JWT.

## Стек

- **Go 1.23** — net/http (стандартный роутер)
- **PostgreSQL 16** — pgx/v5 (без ORM)
- **JWT** — аутентификация / авторизация
- **Docker / docker-compose** — развёртывание

## Быстрый старт

```bash
git clone https://github.com/IAM-47/miniKinopoisk.git
cd miniKinopoisk
docker-compose up --build
```

Сервер поднимется на `http://localhost:8080`. База данных создаётся и инициализируется автоматически.

> Для запуска без Docker: создайте БД вручную, выполните `migrations/001_init.sql`, обновите `configs/config.yaml` и запустите `go run ./cmd/server/main.go`.

## Переменные окружения

| Переменная   | Описание             | По умолчанию                   |
|--------------|----------------------|--------------------------------|
| `JWT_SECRET` | Секрет для JWT-токенов | `miniKinopoisk-secret-key`   |

## Роли

| Роль    | Возможности                              |
|---------|------------------------------------------|
| `user`  | Просмотр фильмов, актёров, бюджетов     |
| `admin` | Полный CRUD по всем сущностям            |

Роль назначается в таблице `users` вручную. По умолчанию новые пользователи получают роль `user`.

---

## API

### Системное

| Метод | URL | Описание |
|-------|-----|----------|
| GET | `/` | Health-check, версия |

### Пользователи

| Метод | URL | Описание |
|-------|-----|----------|
| POST | `/register` | Регистрация |
| POST | `/login` | Вход, возвращает JWT-токен |

**Тело запроса (register / login):**
```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

**Ответ login:**
```json
{ "token": "<jwt>" }
```

Все защищённые запросы требуют заголовка:
```
Authorization: Bearer <token>
```

---

### Фильмы

| Метод | URL | Доступ | Описание |
|-------|-----|--------|----------|
| GET | `/movies` | публичный | Список всех фильмов |
| GET | `/movies/{id}` | публичный | Фильм по ID |
| POST | `/movies` | admin | Создать фильм |
| PUT | `/movies/{id}` | admin | Обновить фильм |
| DELETE | `/movies/{id}` | admin | Удалить фильм |

**Тело (POST / PUT):**
```json
{
  "title": "Inception",
  "producer": "Emma Thomas",
  "director": "Christopher Nolan",
  "release_year": 2010
}
```

---

### Актёры

| Метод | URL | Доступ | Описание |
|-------|-----|--------|----------|
| GET | `/movies/{id}/actors` | публичный | Актёры фильма |
| POST | `/actors` | admin | Создать актёра |
| PUT | `/actors/{id}` | admin | Обновить актёра |
| DELETE | `/actors/{id}` | admin | Удалить актёра |
| POST | `/movies/{id}/actors` | admin | Привязать актёра к фильму |

**Тело (POST /actors):**
```json
{
  "first_name": "Leonardo",
  "last_name": "DiCaprio",
  "birth_date": "1974-11-11",
  "salary": 20000000
}
```

**Тело (POST /movies/{id}/actors):**
```json
{ "actor_id": 1 }
```

---

### Бюджет и сборы

| Метод | URL | Доступ | Описание |
|-------|-----|--------|----------|
| GET | `/movies/{id}/budget` | публичный | Бюджет фильма |
| POST | `/movies/{id}/budget` | admin | Добавить бюджет |
| PUT | `/movies/{id}/budget` | admin | Обновить бюджет |
| DELETE | `/budget/{id}` | admin | Удалить запись |

**Тело (POST / PUT):**
```json
{
  "total_budget": 160000000,
  "fees_in_prod_country": 292576195,
  "fees_in_other": 543000000
}
```

---

## Схема базы данных

```
users
  id, email (unique), password_hash, role, created_at

movies
  id, title, producer, director, release_year

actors
  id, first_name, last_name, birth_date, salary

movie_actor  (связь многие-ко-многим)
  id_movie → movies.id
  id_actor → actors.id

budget_and_fees  (один-к-одному с movies)
  id, id_movie → movies.id
  total_budget, fees_in_prod_country, fees_in_other
```

Полный DDL: [`migrations/001_init.sql`](migrations/001_init.sql)
