# README

## Rodar

```CMD
docker-compose down
docker-compose up
docker-compose up --build

```

## Insalar extensões no PHP

Abra o container. <id> é o id referente ao container<br>

```CMD
docker exec -it <id> bash
```

Digite:<br>

```CMD
docker-php-ext-install mysqli pdo pdo_mysql
```
