# README

## Rodar

```CMD
docker-compose up --build

```

## Instalar extensões no PHP manualmente

```CMD
Se precisar.... Note que ao rodar e.php e f.php o banco de dados consegue conectar.
Abra o container. <id> é o id referente ao container<br>
docker exec -it <id> /bin/bash
docker-php-ext-install mysqli pdo pdo_mysql
```
