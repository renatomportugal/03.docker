# Redis

```CMD
start a redis instance
docker run --name some-redis -d redis

start with persistent storage
$ docker run --name some-redis -d redis redis-server --save 60 1 --loglevel warning

```
