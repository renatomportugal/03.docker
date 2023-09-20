# Redes

## Criando_a_rede

```CMD
docker network ls
Retorno:
NETWORK ID     NAME      DRIVER    SCOPE
ead90b558756   bridge    bridge    local
54e492915993   host      host      local
c3a5cd6eb112   none      null      local

docker inspect bridge
Retorno:
[
  {
    "Name": "bridge",
    "Id": "ead90b558756d69de6aed48cad1b22a903e47ab87683cd83cfe150c99390c235",
    "Created": "2023-09-19T09:32:46.325354575-03:00",
    "Scope": "local",
    "Driver": "bridge",
    "EnableIPv6": false,
    "IPAM": {
      "Driver": "default",
      "Options": null,
      "Config": [
        {
          "Subnet": "172.17.0.0/16",
          "Gateway": "172.17.0.1"
        }
      ]
    },
    "Internal": false,
    "Attachable": false,
    "Ingress": false,
    "ConfigFrom": {
      "Network": ""
    },
    "ConfigOnly": false,
    "Containers": {},
    "Options": {
      "com.docker.network.bridge.default_bridge": "true",
      "com.docker.network.bridge.enable_icc": "true",
      "com.docker.network.bridge.enable_ip_masquerade": "true",
      "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
      "com.docker.network.bridge.name": "docker0",
      "com.docker.network.driver.mtu": "1500"
    },
    "Labels": {}
  }
]


Criando...
docker network create \
  -d \
  bridge \
  --subnet 172.18.0.0/16 \
  --gateway 172.18.0.1 \
  bridge.172.18

docker network ls
Retorno:
NETWORK ID     NAME            DRIVER    SCOPE
ead90b558756   bridge          bridge    local
f135a3f6ad37   bridge.172.18   bridge    local
54e492915993   host            host      local
c3a5cd6eb112   none            null      local
```

## Criando_um_container

```CMD
docker run -it -d --name alpine-172-17 alpine
docker ps -a
CONTAINER ID   IMAGE     COMMAND     CREATED          STATUS          PORTS     NAMES
9922f7a5d01f   alpine    "/bin/sh"   10 minutes ago   Up 10 minutes             alpine-172-17

docker inspect alpine-172-17
...veja o IP...
 "IPAddress": "172.17.0.2",

docker exec -it alpine-172-17 sh
ping 172.17.0.1
Retorno:
PING 172.17.0.1 (172.17.0.1): 56 data bytes
64 bytes from 172.17.0.1: seq=0 ttl=64 time=0.125 ms
64 bytes from 172.17.0.1: seq=1 ttl=64 time=0.053 ms
64 bytes from 172.17.0.1: seq=2 ttl=64 time=0.066 ms
```

## Criando_um_container_no_172_18

```CMD
docker run -it -d --name alpine-172-18 --network bridge.172.18 alpine

docker ps -a
CONTAINER ID   IMAGE     COMMAND     CREATED          STATUS          PORTS     NAMES
a7d7b8dfc4c7   alpine    "/bin/sh"   7 seconds ago    Up 5 seconds              alpine-172-18
9922f7a5d01f   alpine    "/bin/sh"   24 minutes ago   Up 24 minutes             alpine-172-17

Saber o IP do container alpine-172-18
docker inspect alpine-172-18
...veja o IP...
"IPAddress": "172.18.0.2",

Agora iremos entrar no alpine do barramento 17 e fazer um ping no do barramento 18.

Na primeira janela:
docker exec -it alpine-172-17 sh
ping 172.18.0.2
Não respondeu...

Na segunda janela:
docker exec -it alpine-172-18 sh
ping 172.17.0.2
Não respondeu...

```

## Adicionar_o_Container_em_outro_barramento

```CMD

Na verdade a conexão é a inclusão do container desejado na rede que o container alvo está. Consequentemente ele consegue acessar todos os outros containers.
Em uma terceira janela... Deixe as outras duas abertas para ver a hora que começa a se comunicar...
docker network connect bridge.172.18 alpine-172-17
...imediatamente o alpine-172-17 consegue alcançar a rede bridge.172.18...
Ainda na terceira janela, dê um docker inspect alpine-172-17 para ver que este container tem dois IP, cada um com MAC Address diferente.
..."IPAddress": "172.17.0.2",
..."IPAddress": "172.18.0.3",

Para retirar:
docker network disconnect bridge.172.18 alpine-172-17
...o ping pára automaticamente...

```

## Outro teste

```CMD
Crie a primeira rede

docker network create \
  -d \
  bridge \
  --subnet 172.30.0.0/16 \
  --gateway 172.30.0.1 \
  bridge.172.30

docker network create \
  -d \
  bridge \
  --subnet 172.31.0.0/16 \
  --gateway 172.31.0.1 \
  bridge.172.31

Crie um container na primeira rede

docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=my-password \
    -e MYSQL_PASSWORD=my-password \
    --network bridge.172.30 \
    --name mysql.5.7 \
    mysql:5.7

Crie um container na segunda rede
docker run \
    -d \
    -p 8080:80 \
    -e PMA_HOST=mysql.5.7 \
    --network bridge.172.31 \
    --name phpmyadmin \
    phpmyadmin/phpmyadmin

```

## Conectar os dois containers

```CMD
Acesse no browser. http://localhost:8080
Ou http://seu-ip:8080
Não foi possível conectar.

Na verdade a conexão é a inclusão do container desejado na rede que o container alvo está. Consequentemente ele consegue acessar todos os outros containers.<br>
docker network connect bridge.172.30 phpmyadmin

Agora conectou...
```

## Desconectando um container de uma rede

```CMD
docker network disconnect bridge.172.30 phpmyadmin
```

## IP estático

### Criar a rede

```CMD
docker network create \
  -d \
  bridge \
  --subnet 172.40.0.0/16 \
  --gateway 172.40.0.1 \
  net.172.40
```

### Criar um container dentro da rede com o IP

```CMD
docker run \
  -d \
  -p 8080:80 \
  --name mginx \
  --network bridge.172.40 \
  --ip 172.40.0.100 \
  nginx
```

### Criar um container isolado

```CMD
docker run \
  -d \
  --name mginx.isolated \
  --network none \
  nginx
```

## Organizando os Containers

### Fazendo a rede

```CMD
docker network create \
  -d \
  bridge \
  --subnet 66.6.0.0/16 \
  --gateway 66.6.0.1 \
  bridge.66.6
```

### Criando os Containers de banco de dados

#### MySQL

```CMD
Criar o volume<br>
docker volume create mysql.5.7.var.lib.mysql
```

```CMD
docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=my-password \
    -e MYSQL_USER=root \
    -e MYSQL_PASSWORD=my-password \
    -v mysql.5.7.var.lib.mysql:/var/lib/mysql \
    --network bridge.66.6 \
    --name mysql.5.7 \
    --ip 66.6.0.5 \
    --restart=always \
    mysql:5.7
```

### Aplicações

#### 8080 - PhpMyAdmin

```CMD
docker run \
    -d \
    -p 8080:80 \
    -e PMA_HOST=mysql.5.7 \
    --network bridge.66.6 \
    --name phpmyadmin \
    --ip 66.6.0.15 \
    --restart=always \
    phpmyadmin/phpmyadmin
```

#### 8000 - Owncloud

```CMD
Criar o volume
docker volume create owncloud.var.www.html.data
```

```CMD
docker run \
  -d \
  -p 8000:80 \
  -v owncloud.var.www.html.data:/var/www/html/data \
  --network bridge.66.6 \
  --name owncloud \
  --ip 66.6.0.100 \
  --restart=always \
owncloud
```

```CMD
docker run -it busybox
```

```CMD
Abra outra janela
docker run -it busybox
```

```CMD
Verifique os números IP dos containers em cada janela
hostname -i
hostname
```

```CMD
exit

172.17.0.2
135774e941a8

172.17.0.3
135774e941a8o

Para que um container reconheça outro pelo Aliases
hostname é interno, name é externo

Abra o terminal id 13
```

```CMD
docker run -it --name busybox1 -h busybox-one busybox
hostname
hostname -i
```

```CMD
abra o terminal id 06
exit
```

```CMD
docker run -it --name busybox2 -h busybox-two busybox
hostname
hostname -i
```

```CMD
docker network ls
```

```CMD
Copie o id da bridge
0e195bb2ae21
docker network inspect bridge
docker network create custom
```

```CMD
8d
docker network inspect custom
```

saia dos dois terminais

```CMD
exit
```

```CMD
e elimine todas as instâncias
docker container prune -f
docker ps -a
```

abra dois terminais

```CMD
docker run -it --network custom busybox
hostname -i
hostname
```

```CMD
172.18.0.2
317653a37985

172.18.0.3
36b6a47edc2e

Agora faremos um ping do IP final 2 para o final 3, por IP e por hostname
ping 172.18.0.3
ping 36b6a47edc2e
```

Conclusão: Containers na rede custom se comunicam com hostnames.
e não conseguem se comunicar com os nomes dinâmicos de hosts.

Using custom persistent names for connectivity in the custom network
saia dos dois terminais

```CMD
exit
```

```CMD
docker run -it --network custom --name busybox1 busybox
docker run -it --network custom --name busybox2 busybox
docker ps
```

na janela do busybox1

```CMD
ping busybox2
ping 0e22d7f144a0
```

```CMD
hostname -i
hostname
```

```CMD
172.18.0.2
5a91c70cbed6

172.18.0.3
0e22d7f144a0
docker network inspect custom
docker inspect busybox1
docker inspect busybox2
```

```CMD
na janela do busybox1
exit
docker ps
docker ps -a
```

```CMD
docker start -i busybox1
ping busybox2
```

```CMD
na janela do busybox2
ping busybox1
```

```CMD
saia dos dois terminais
exit
```
