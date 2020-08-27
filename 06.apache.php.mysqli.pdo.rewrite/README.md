## Rodar
```
docker-compose down
docker-compose up
docker-compose up --build

```

## Instalar extensões no PHP
Abra o container. <id> é o id referente ao container<br>
```
docker exec -it <id> bash
```
Digite:<br>
```
docker-php-ext-install mysqli pdo pdo_mysql
```