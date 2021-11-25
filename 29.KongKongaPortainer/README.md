# ReadMe

## Build_godocker

```CMD
Na pasta services\service execute
docker build -t godocker .
```

## ACESSAR

```CMD
Portainer
localhost:9000

Konga
localhost:1337

Kong
localhost:8000

Api Kong
localhost:8001

Service A
localhost:8081

Service B
localhost:8082

Service C
localhost:8083
```

## User

```CMD
Preencher username, Email, Password, Confirm password
Botão Create Admin
```

## Connections

```CMD
No dashboard:
Botão New Connections
Criar conexão Default
Name: Kong, 
Kong Admin URL: localhost:8001
Botão Create Connection
```

## Services

```CMD
Menu Services, Botão Add new service
Name: servicea
Protocol: http
Host: servicea
Port: 8081
Path: /
Botão Submit service

Erro: Não dá pra criar serviço sem criar uma base de dados.
Criou usando com o DB, REFAZER.

```

## Kong_DB

```CMD
docker run -d --name kong-database \
    -p 5432:5432 \
    -e "POSTGRES_USER=kong" \
    -e "POSTGRES_DB=kong" \
    -e "POSTGRES_PASSWORD=kong" \
    postgres:9.6

docker run --rm \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_USER=kong" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    kong/kong-gateway kong migrations bootstrap

docker run -d --name kong \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
    -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
    -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
    -p 8000:8000 \
    -p 8443:8443 \
    -p 8001:8001 \
    -p 8444:8444 \
    kong/kong-gateway

```
