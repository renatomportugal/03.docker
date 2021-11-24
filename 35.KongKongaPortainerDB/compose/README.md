# Build_godocker

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

```


# Kong in Docker Compose

This is the official Docker Compose template for [Kong][kong-site-url].

# What is Kong?

You can find the official Docker distribution for Kong at [https://hub.docker.com/_/kong](https://hub.docker.com/_/kong).

# How to use this template

This Docker Compose template provisions a Kong container with a Postgres database, plus a nginx load-balancer. After running the template, the `nginx-lb` load-balancer will be the entrypoint to Kong.

To run this template execute:

```shell
$ docker-compose up
```

To scale Kong (ie, to three instances) execute:

```shell
$ docker-compose scale kong=3
```

Kong will be available through the `nginx-lb` instance on port `8000`, and `8001`. You can customize the template with your own environment variables or datastore configuration.

Kong's documentation can be found at [https://docs.konghq.com/][kong-docs-url].

## Issues

If you have any problems with or questions about this image, please contact us through a [GitHub issue][github-new-issue].

## Contributing

You are invited to contribute new features, fixes, or updates, large or small; we are always thrilled to receive pull requests, and do our best to process them as fast as we can.

Before you start to code, we recommend discussing your plans through a [GitHub issue][github-new-issue], especially for more ambitious contributions. This gives other contributors a chance to point you in the right direction, give you feedback on your design, and help you find out if someone else is working on the same thing.

[kong-site-url]: https://konghq.com/
[kong-docs-url]: https://docs.konghq.com/
[github-new-issue]: https://github.com/Kong/docker-kong/issues/new



