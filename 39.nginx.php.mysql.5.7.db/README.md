# 02.nginx.php.mysql.5.7

```CMD
Para rodar:
Vá na pasta nginx.php.mysql.5.7/docker no terminal, digite... (ou suba via FTP)
docker-compose up
```

```CMD
Para acessar:
localhost:8027

ou
<seu-ip>:8027
```

```CMD
Verificar:
Abra outro terminal
docker ps -a
docker-compose -a
docker network ls
docker volume ls
```

```CMD
Para parar:
docker stop $(docker ps -aq)
docker stop $(docker-compose ps -q)
```

```CMD
Para iniciar:
docker start $(docker-compose ps -q)
```

```CMD
Para verificar:
docker ps -a
docker-compose ps
```

```CMD
Destruir o container:
docker-compose down
docker ps -a
docker-compose ps
```

```CMD
Repare que o volume não foi destruído:
docker volume ls
```

```CMD
Testar o banco de dados:
http://localhost:8027/e.php
http://localhost:8027/f.php
```
