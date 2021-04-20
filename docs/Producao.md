# Producao

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
Verifique quais volumes existem:<br>
```
docker volume ls
```
Crie o volume para a aplicação (caso saiba qual pasta deseja persistir os dados)<br>
```
docker volume create mysql
docker volume create nextcloud
```
Verifique se foi criado<br>
```
docker volume ls
```

## Iniciar o Container com o Docker
Rode a imagem que deseja que seja iniciada com o sistema<br>

Liste a imagem para ver o id<br>
```
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
docker exec -it id_do_container bash
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
