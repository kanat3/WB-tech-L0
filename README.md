```
sudo docker-compose up --build -d
go build
./WB-tech-L0
```
##### Подключитья к базе данных:
###### Пароль: root
```
sudo docker ps -a
sudo docker exec -it <container_id> /bin/bash
psql -h localhost -p 5432 -U admin -W
```
##### Доступен по адресу: *127.0.0.1:8090*
##### Ответы на запросы доступны на: *127.0.0.1:8090/bye_page*