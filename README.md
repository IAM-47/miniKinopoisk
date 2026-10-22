# miniKinopoisk

[![CI](https://github.com/IAM-47/miniKinopoisk/actions/workflows/ci.yml/badge.svg)](https://github.com/IAM-47/miniKinopoisk/actions/workflows/ci.yml)

REST API для каталога фильмов с актёрами и финансовой статистикой. Написан на Go без фреймворков — только стандартная библиотека, pgx и JWT.

## Стек

- **Go 1.25** — net/http (стандартный роутер)
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
| DELETE | `/users/{email}` | Удалить пользователя (admin) |

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
| GET | `/movies` | публичный | Список фильмов (`?limit=10&offset=0`) |
| GET | `/movies/{id}` | публичный | Фильм по ID |
| POST | `/movies` | admin | Создать фильм |
| PUT | `/movies/{id}` | admin | Обновить фильм |
| DELETE | `/movies/{id}` | admin | Удалить фильм |
| GET | `/actors/{id}/movies` | публичный | Фильмы с участием актёра |

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
| GET | `/actors/{id}` | публичный | Актёр по ID |
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

## Документация API (OpenAPI / Swagger)

Спецификация OpenAPI 3.0 лежит в [`docs/openapi.yaml`](docs/openapi.yaml): все 20 эндпоинтов, схемы запросов и ответов, коды ошибок, JWT-авторизация.

Посмотреть в браузере через Swagger UI:

```bash
docker run --rm -p 8081:8080 -e SWAGGER_JSON=/spec/openapi.yaml \
  -v "$(pwd)/docs:/spec" swaggerapi/swagger-ui
```

После запуска откройте `http://localhost:8081`. Кнопка **Authorize** принимает токен из `POST /login`, после этого можно вызывать admin-методы прямо из интерфейса. Либо вставьте содержимое файла на [editor.swagger.org](https://editor.swagger.org).

## Тесты и CI

Тесты не требуют базы данных: хендлеры зависят от небольших интерфейсов хранилищ (`internal/handlers/stores.go`), в тестах вместо PostgreSQL используются подменные реализации.

```bash
go test ./...                 # все тесты
go test -race -cover ./...    # с детектором гонок и покрытием
```

Что покрыто:

- **auth** — выдача JWT, срок жизни 24 часа, отказ при неверном секрете
- **middleware** — проверка токена (нет заголовка, неверная схема, мусор, просрочен, чужая подпись, `alg=none`), цепочка `AuthMiddleware → AdminOnly`, CORS и preflight
- **utils** — bcrypt-хеширование паролей (соль, неверный пароль, битый хеш)
- **handlers** — все обработчики: валидация входных данных, коды ответов 2xx/4xx/5xx, пагинация, конфликт при дубле email, одинаковая ошибка для неверного email и пароля

GitHub Actions ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)) на каждый push в `main` и каждый pull request запускает: проверку `gofmt`, `go vet`, тесты с `-race` и покрытием, сборку Docker-образа и валидацию `docs/openapi.yaml`.

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
