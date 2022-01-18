# 41.nginx.php.mysql.5.7

```CMD
Para rodar:
Vá na pasta 41.nginx.php.mysql.5.7/docker no terminal, digite... (ou suba via FTP)
docker-compose up
```

```CMD
Para acessar:
localhost:80

ou
<seu-ip>:80
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
http://localhost/e.php
http://localhost/f.php
```
