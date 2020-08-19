Para rodar:<br>
```
Vá na pasta nginx.php.mysql.5.7/docker no terminal, digite...
docker-compose up
```
Para acessar:<br>
```
localhos:8027
```
Verificar:<br>
```
Abra outro terminal
docker ps -a
docker-compose -a
docker network ls
docker volume ls
```
Para parar:<br>
```
docker stop $(docker ps -aq)
docker stop $(docker-compose ps -q)
```
Para iniciar:<br>
```
docker start $(docker-compose ps -q)
```
Para verificar:<br>
```
docker ps -a
docker-compose ps
```
Destruir o container:<br>
```
docker-compose down
docker ps -a
docker-compose ps
```
Repare que o volume não foi destruído:<br>
```
docker volume ls
```
Testar o banco de dados:<br>
```
http://localhost:8027/e.php
http://localhost:8027/f.php
```
