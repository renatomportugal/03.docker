# Dashy

```CMD
https://hub.docker.com/r/lissy93/dashy

docker pull lissy93/dashy

docker run -p 80:80 lissy93/dashy

docker run -d \
  -p 80:80 \
  -v /root/my-local-conf.yml:/app/public/conf.yml \
  --name my-dashboard \
  --restart=always \
  lissy93/dashy:latest

```
