# Raspberry

## Comandos
Verificar arquitetura do Raspberry<br>
```
cat /proc/cpuinfo
```
Memória disponível<br>
```
free -h
```
Ver processamento<br>
```
htop
```

## Imagens
https://hub.docker.com/u/arm32v7<br>
```
Por problemas de arquitetura nem todas as imagens funcionarão.
```
### Hello-world
```
docker run hello-world
```
### Postgres
```
docker pull postgres:9.6.19
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres:9.6.19
```
### Testar
```
https://hub.docker.com/r/arm32v7/rabbitmq
docker run -d --hostname my-rabbit --name some-rabbit arm32v7/rabbitmq:3

https://hub.docker.com/r/arm32v7/alpine
docker pull arm32v7/alpine
FROM arm32v7/alpine:3.7
RUN apk add --no-cache mysql-client
ENTRYPOINT ["mysql"]

https://hub.docker.com/r/arm32v7/redis
docker run --name some-redis -d arm32v7/redis

https://hub.docker.com/r/arm32v7/nginx
docker run --name some-nginx -v /some/content:/usr/share/nginx/html:ro -d arm32v7/nginx

https://hub.docker.com/r/arm32v7/phpmyadmin
docker run --name myadmin -d -e PMA_HOST=dbhost -p 8080:80 arm32v7/phpmyadmin

https://hub.docker.com/r/arm32v7/influxdb
docker run -p 8086:8086 \
      -v influxdb:/var/lib/influxdb \
      arm32v7/influxdb:1.8

https://hub.docker.com/r/arm32v7/wordpress
docker run --name some-wordpress --network some-network -d arm32v7/wordpress

https://hub.docker.com/r/arm32v7/telegraf
docker run -d --name influxdb -p 8086:8086 influxdb

https://hub.docker.com/r/arm32v7/httpd
docker run -dit --name my-apache-app -p 8080:80 -v "$PWD":/usr/local/apache2/htdocs/ arm32v7/httpd:2.4

https://hub.docker.com/r/arm32v7/nats-streaming
docker run -p 4223:4223 -p 8223:8223 arm32v7/nats-streaming -p 4223 -m 8223

https://hub.docker.com/r/arm32v7/adminer
docker run --link some_database:db -p 8080:8080 arm32v7/adminer

https://hub.docker.com/r/arm32v7/owncloud
docker pull arm32v7/owncloud
docker run -d -p 80:80 arm32v7/owncloud:8.1
docker run -d -p 80:80 arm32v7/owncloud

https://hub.docker.com/r/arm32v7/piwik
docker run --name some-piwik --link some-mysql:db -d arm32v7/piwik

```

### arm32v7/php:7.4-cli
https://hub.docker.com/r/arm32v7/php<br>
```
docker pull arm32v7/php:7.4-cli

```
