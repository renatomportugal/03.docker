# Dicas
# Reiniciar Automaticamente
Fazer com que todos os containers reiniciem automaticamente<br>
```
docker inspect -f "{{ .HostConfig.RestartPolicy }}" $(docker ps -q)
```
Iniciar todos os Containers com o Docker<br>
```
docker update --restart=always $(docker ps -q)
```

# Acessar o container
```
docker ps -a
docker exec -it id_do_container bash
```

# Remover
## Remover todos os containers parados
```
CUIDADO!!!!!!!!
docker rm -v $(docker ps -a -q -f status=exited)
```

# Parar o container
```
docker stop <id>
```
## Parar todos os containers
```
docker stop $(docker ps -q)
```

# Pesquisas
## Mostrar todos os containers parados
```
docker ps -a -f status=exited
```

## Lista de Retorno
### Ordenar pelas portas
```
docker ps -a --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}" | (read -r; printf "%s\n" "$REPLY"; sort -k 3 )
```
### Ordenar pelos nomes
```
docker ps -a --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}" | (read -r; printf "%s\n" "$REPLY"; sort -k 2 )
```
### Ordenar pelos ids
```
docker ps -a --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}" | (read -r; printf "%s\n" "$REPLY"; sort -k 1 )
```

## Procurar as imagens que comecem com a palavra minha
```
docker images | grep minha
```
## Procurar containers que tenham no nome a palavra tutorial
docker ps -a | grep tutorial

## Ver o log do container (colocar pelo menos os três primeiros dígitos do id)<br>
```
docker logs 921
```

## Deletar imagem<br>
```
docker images
docker rmi nginx:1.16.0-alpine-perl
```

Ou usando o id (verifique se não existe mais de uma imagem com o mesmo início, -f é para forçar)<br>
```
docker rmi -f a2
```

Verificar o nome do container<br>
```
hostname
```

Verificar o ip (notará que os IPs são diferentes)<br>
```
hostname -i
```

Mostra todas as informações, sem esconder caracteres<br>
```
docker ps --no-trunc
```


Mostra o help do comando container<br>
```
docker container --help
```

Exclui todos os containers parados.<br>
```
docker container prune
```

Para remover apenas dois containers basta colocar os dois primeiros dígitos do ID separados por espaço<br>
```
docker rm 05 7a
```

Limpeza. Veja quais containers estejam dorando<br>
```
docker ps
```

Pare todos os containers.<br>
```
docker stop id1 id2 id3
docker ps -a
docker container prune -f
```

ACESSO ao HOST pelo Windows ou Mac<br>
```
docker run -it --rm --privileged --pid=host justincormack/nsenter1
```

Limpeza da lista dos containers<br>
```
docker ps
docker ps -a
docker container prune -f
docker ps -a
```

Listar imagens com a formatação desejada<br>
```
docker ps -a --format "table {{.ID}}\t{{.Image}}\t{{.Names}}"
```

Copiar arquivos para dentro do container<br>
```
docker cp config.php id_do_container:/var/www/html/config/config.php
```

Copiar arquivos para fora do container. Copiará para a pasta que estiver. Use pwd para saber.<br>
```
docker cp id_do_container:/var/www/html/config/config.php config.php
```
# Busca de Containers
```
docker search redis
```
# Outros Comandos
```
docker inspect <friendly-name|container-id>
docker logs <friendly-name|container-id>
```
