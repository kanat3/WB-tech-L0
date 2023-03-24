# WB-tech-L0

В БД:
- [x] Развернуть локально postgresql
- [x] Создать бд
- [x] Настроить пользователя
- [ ] Создать таблицы для хранения полученных данных

В сервисе:
- [ ] Подключение и подписка на канал в nats-streaming
- [ ] Полученные данные писать в Postgres
- [x] Так же полученные данные сохранить in memory в сервисе (Кеш)
- [ ] В случае падения сервиса восстанавливать Кеш из Postgres
- [ ] Поднять http сервер и выдавать данные по id из кеша
- [x] Сделать простейший интерфейс отображения полученных данных

Также:
- [x] Cделать отдельный скрипт, для публикации данных в канал
- [x] Nats-streaming развернуть локально
- [ ] Покрыть сервис автотестами
- [ ] Устроить сервису стресс тест
