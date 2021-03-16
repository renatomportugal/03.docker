# MongoDB

## Instalar no Docker
[Docker](Imagens.md#mongodb)

## ExecutarNoDocker
```
docker exec -it <id do Container> bash
```

## Primeiros Comandos
```
Conexão Implícita
mongo

Ver as Coleções
show dbs


Trabalhar com uma versão específica
use test

Verificar se selecionou corretamente
db

Inserir
db.animals.insert({"animal": "cat"})
db.animals.insert({"animal": "dog"})
db.animals.insert({"animal": "monkey"})

Ler Coleção
db.animals.find()

```