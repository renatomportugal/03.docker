# Dicas

Verificar o nome do container<br>

```CMD
hostname
```

Verificar o ip (notará que os IPs são diferentes)<br>

```CMD
hostname -i
```

Mostra todas as informações, sem esconder caracteres<br>

```CMD
docker ps --no-trunc
```

Mostra o help do comando container<br>

```CMD
docker container --help
```

Exclui todos os containers parados.<br>

```CMD
docker container prune
```

Para remover apenas dois containers basta colocar os dois primeiros dígitos do ID separados por espaço<br>

```CMD
docker rm 05 7a
```

Limpeza. Veja quais containers estejam dorando<br>

```CMD
docker ps
```

Pare todos os containers.<br>

```CMD
docker stop id1 id2 id3
docker ps -a
docker container prune -f
```

ACESSO ao HOST pelo Windows ou Mac<br>

```CMD
docker run -it --rm --privileged --pid=host justincormack/nsenter1
```

Limpeza da lista dos containers<br>

```CMD
docker ps
docker ps -a
docker container prune -f
docker ps -a
```

Listar imagens com a formatação desejada<br>

```CMD
docker ps -a --format "table {{.ID}}\t{{.Image}}\t{{.Names}}"
```

Copiar arquivos para dentro do container<br>

```CMD
docker cp config.php id_do_container:/var/www/html/config/config.php
```

Copiar arquivos para fora do container. Copiará para a pasta que estiver. Use pwd para saber.<br>

```CMD
docker cp id_do_container:/var/www/html/config/config.php config.php
```

## Copiar_dados

```CMD
Copiar do container para a pasta.
docker cp id_do_container:/pasta pasta

Copiar da pasta para o container.
Estando num nível acima da pasta...
ls
docker cp data 824ef00007a6:/

Copiar a pasta data do container para o diretório raiz.
docker cp af006d151ffc:/data ./


```

## Busca de Containers

```CMD
docker search redis
```

## Exportar Containers

```CMD
docker ps -a
CONTAINER ID   IMAGE                                     COMMAND                  CREATED          STATUS                      PORTS     NAMES
47ca31b69886   tensorflow/tensorflow:2.0.0-py3-jupyter   "bash -c 'source /et…"   21 minutes ago   Exited (0) 22 seconds ago             tensorflow

docker export tensorflow > tensorflow.tar
docker export --output="tensorflow.tar" tensorflow

...outra forma...
e895b5bf0e11   solr:8.11-slim                     "/bin/bash -c 'init-…"   45 minutes ago   Up 3 minutes   0.0.0.0:8983->8983/tcp                           dspacesolr
c92a286f2302   dspace/dspace:dspace-7_x-test      "/bin/bash -c 'while…"   45 minutes ago   Up 3 minutes   8000/tcp, 8009/tcp, 0.0.0.0:8080->8080/tcp       dspace
a71acf84102a   dspace/dspace-postgres-pgcrypto    "docker-entrypoint.s…"   45 minutes ago   Up 3 minutes   0.0.0.0:5432->5432/tcp                           dspacedb
f414f11a98a6   dspace/dspace-angular:dspace-7_x   "docker-entrypoint.s…"   45 minutes ago   Up 3 minutes   0.0.0.0:4000->4000/tcp, 0.0.0.0:9876->9876/tcp   dspace-angular

docker save -o C:\dspace\dspacesolr.tar solr:8.11-slim
docker save -o C:\dspace\dspace.tar dspace/dspace:dspace-7_x-test
docker save -o C:\dspace\dspacedb.tar dspace/dspace-postgres-pgcrypto
docker save -o C:\dspace\dspace-angular.tar dspace/dspace-angular:dspace-7_x

```

## Importar Containers

```CMD
docker load < tensorflow.tar
docker load --input tensorflow.tar
docker image ls

```

## Outros Comandos

```CMD
docker inspect <friendly-name|container-id>
docker logs <friendly-name|container-id>
```
