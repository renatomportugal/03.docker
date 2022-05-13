# Producao

## Tabela_de_IP

Aplicação | PORTA | IP          | NOME
--------- | ---- | ----------- | ----
Rede      | ---- | bridge.66.6 | ----
Subnet    | ---- | 66.6.0.0/16 | ----
Gateway   | ---- | 66.6.0.1    | ----
phpmyadmin| 8080 | 66.6.0.14   | ----
MySQL     | 3306 | 66.6.0.5    | ----
colabora  | 9980 | 66.6.0.15   | ----
nc-mib    | 777  | 66.6.0.20   | ----

## Rede_666

Verifica se ter a rede bridge.66.6

```CMD
docker network ls
```

Se não tiver, crie

```CMD
docker network create \
  -d \
  bridge \
  --subnet 66.6.0.0/16 \
  --gateway 66.6.0.1 \
  bridge.66.6
```

Verifique quais IPs estão vagos (Todos que não aparecem estão vagos)

```CMD
docker network inspect bridge.66.6
```

Estou vendo uma forma de melhorar a apresentação dos dados:

```CMD
docker network inspect bridge.66.6 -f {{.Containers}}
```

## Volumes

Verifique quais volumes existem:

```CMD
docker volume ls
```

Crie o volume para a aplicação (caso saiba qual pasta deseja persistir os dados)

```CMD
docker volume create mysql.5.7
docker volume create nextcloud
```

Verifique se foi criado:

```CMD
docker volume ls
```

## Iniciar o Container com o Docker

Rode a imagem que deseja que seja iniciada com o sistema

Liste a imagem para ver o id

```CMD
docker ps
```
Vamos suport que o id inicie com 98...<br>
Pare a imagem.<br>
```
docker stop 98
```
Vamos listar todas as imagens.<br>
```
docker ps -a
```
Vamos dar o comando para que a imagem inicie junto com o docker.<br>
```
docker update --restart=always 98
```

Comando para reverter.<br>
```
docker update --restart=no 98
```

Para verificar:<br>
```
docker inspect -f "{{ .HostConfig.RestartPolicy }}" <container_id>
docker inspect -f "{{ .RestartCount }}" <container_id>
docker inspect -f "{{ .State.StartedAt }}" <container_id>
```

Verificar todos os containers:

```CMD
docker inspect -f "{{ .HostConfig.RestartPolicy }}" $(docker ps -q)
```

Iniciar todos os Containers com o Docker:

```CMD
docker update --restart=always $(docker ps -q)
```

Lista de todos os containers com id, status de restart e quantidadede (de restart) (que estão rodando)<br>
```
docker inspect -f "{{.Id}} {{.HostConfig.RestartPolicy.Name}} {{.RestartCount}}" $(docker container ls -q)
```

Lista de todos os containers com id, status de restart e quantidadede (de restart) (incluindo os que não estão rodando)<br>
docker inspect -f "{{.Id}} {{.HostConfig.RestartPolicy.Name}} {{.RestartCount}}" $(docker container ps -a -q)

Explicando:
Este comando lista todos os ids dos containers que estão rodando. Para incluir todos adicione -a.

```CMD
$(docker container ls -a -q)
```

## Entrar no container

```CMD
docker ps -a
Padrão
docker exec -it id_do_container bash

Forma Completa
docker exec -it id_do_container /bin/bash

Para containers de outros tipos
docker exec -it id_do_container /bin/sh
```

## Remover

### Remover todos os containers parados

```CMD
CUIDADO!!!!!!!!
docker rm -v $(docker ps -a -q -f status=exited)
```

## Parar o container

```CMD
docker stop <id>
```

### Parar todos os containers

```CMD
docker stop $(docker ps -q)
```
