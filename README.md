# Testovoe Go

#### Микросервис Tasks - дает возможность запланировать задачи и отмечать их статус
#### Проект релизован с помощью Gin, BunORM, Postgres, Goose, Zerolog, Docker, Docker-compose, gRPC

## Чтобы запустить проект у себя локально, вам необохимо
Склонировать репозиторий

`https://github.com/fazletdinov/TestovoeGo.git`

Скопируйте содержимое .env.example в .env
и запустите команду

```
make run
```
Если не установлена утилита make, то необходимо запустить следующей командой

```
cd tasks
docker compose -f docker-compose.yaml up --build
```
Вышеуказанные команды запустят приложение
Далее можете посмотреть Api спецификацию (свагер) по адресу:
`http://localhost:8080/docs/index.html`

### Автор
[Idel Fazletdinov - fazletdinov](https://github.com/fazletdinov)