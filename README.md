# Wallet App Backend on GoLang

- Разработано Веб-Приложений на Go, следуя дизайну REST API.
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
- Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций.
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>. Работа с переменными окружения.
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Регистрация и аутентификация. Работа с JWT. Middleware.
- Написание SQL запросов.
- Graceful Shutdown

### Для запуска приложения:

```
go build -o app cmd/app/main.go
```

Для postgres можно использовать Docker

```docker run -d --name ninja-db -e POSTGRES_PASSWORD=qwerty123 -v ${HOME}/pgdata/:/var/lib/postgresql/data -p 5432:5432 postgres```
