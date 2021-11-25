# ReadMe

```CMD
https://www.youtube.com/watch?v=_2GRXgYswhI


Subir a pasta 29.KongKongaPortainer para o servidor, via WinSCP
cd 29.KongKongaPortainer
```

## Rede

```CMD
docker network ls

docker network create \
  -d \
  bridge \
  --subnet 66.6.0.0/16 \
  --gateway 66.6.0.1 \
  bridge.66.6

docker network inspect bridge.66.6
docker network inspect bridge.66.6 -f {{.Containers}}
```

## Build_godocker

```CMD
cd services/service
docker build -t godocker .
docker image ls
```

## Services

```CMD
docker run \
    -d \
    -p 8081:8081 \
    -e PORT=":8081" \
    -e CONTENT="<h1>Service A - 8081</h1>" \
    --network bridge.66.6 \
    --name servicea \
    --ip 66.6.0.20 \
    godocker


docker run \
    -d \
    -p 8082:8082 \
    -e PORT=":8082" \
    -e CONTENT="<h1>Service B - 8082</h1>" \
    --network bridge.66.6 \
    --name serviceb \
    --ip 66.6.0.21 \
    godocker

docker run \
    -d \
    -p 8083:8083 \
    -e PORT=":8083" \
    -e CONTENT="<h1>Service C - 8083</h1>" \
    --network bridge.66.6 \
    --name servicec \
    --ip 66.6.0.22 \
    godocker

```

## Volume

```CMD
docker volume ls
docker volume create postgres_9.6
docker volume create kong
docker volume create konga
docker volume ls
```

## Kong_DB

```CMD
docker run  \
    -d \
    -p 5432:5432 \
    -e "POSTGRES_USER=kong" \
    -e "POSTGRES_DB=kong" \
    -e "POSTGRES_PASSWORD=kong" \
    -v postgres_9.6:/var/lib/postgresql/data \
    --network bridge.66.6 \
    --name kong-database \
    --ip 66.6.0.23 \
    postgres:9.6

docker run --rm \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_USER=kong" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    --network bridge.66.6 \
    --ip 66.6.0.24 \
    kong/kong-gateway kong migrations bootstrap

docker run -d \
    --link kong-database:kong-database \
    -v "kong:/usr/local/kong/declarative" \
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
    --network bridge.66.6 \
    --name kong \
    --ip 66.6.0.25 \
    kong/kong-gateway

docker run  \
    -d \
    -e TOKEN_SECRET=ahjkscoenfudkjsiqoijoisoisjhysdr \
    -e "DB_ADAPTER=postgres" \
    -e "DB_HOST=kong-database" \
    -e "DB_PORT=5432" \
    -e "DB_USER=kong" \
    -e "DB_PASSWORD=kong" \
    -e "DB_DATABASE=postgres" \
    -e "NODE_ENV=development" \
    -p 1337:1337 \
    --network bridge.66.6 \
    --name konga \
    --ip 66.6.0.26 \
    pantsel/konga

```

## ACESSAR

```CMD
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

## Add_User

```CMD
Preencher username, Email, Password, Confirm password
Botão Create Admin
```

## Add_Connections

```CMD
No dashboard:
Botão New Connections
Criar conexão Default
Name: Kong, 
Kong Admin URL: localhost:8001
Botão Create Connection
Ativar a Conexão
```

## Add_Services

```CMD
Menu Services, Botão Add new service
Name: servicea
Protocol: http
Host: servicea
Port: 8081
Path: /
Botão Submit service
```

## Cadastrar_Rotas

```CMD
Entrar no servicea, Routes, Add Route,
Name: a
Paths: /a (apertar Enter)
Botão Submit Route

Acessar http://10.39.52.113:8000/a
```

## Cadastrar_Servicos_B_C

```CMD
Add new service, Name: serviceb, Protocol: http, Host: serviceb, Port: 8082, Path: /, Botão Submit service
Add new service, Name: servicec, Protocol: http, Host: servicec, Port: 8083, Path: /, Botão Submit service
Services, serviceb, Routes, Add Route, Name: b, Paths: /b (apertar Enter), Botão Submit Route
Services, servicec, Routes, Add Route, Name: c, Paths: /c (apertar Enter), Botão Submit Route
```

## tc

```CMD
Services, Name: tc, Url: https://tecnocrata.org/, Submit Service
Services, tc, Routes, Add Route, Name: tc, Paths: /tc (apertar Enter), Botão Submit Route
```
