# O que é Swarm?
.<br>

# Criar swarm
Com o comando a seguir ele inicia como manager.<br>
```
docker swarm init --advertise-addr <ip-interno>
```
## Verifica se foi criado
```
docker node ls
```
## Adicionar worker
Digitar o comando no principal para pegar o token.<br>
```
docker swarm join-token worker
```

## Adicionar o manager
Copie e cole no outro nó<br>
```
docker swarm join-token manager
```
## Iniciar nova rede
```
docker network create -d overlay interna
docker network ls
```
