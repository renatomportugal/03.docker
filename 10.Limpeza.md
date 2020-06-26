# Limpeza do sistema

# Dangling images

Dangling images são as imagens que estão penduradas, ou seja, imagens órfãs.
## Procurar por dangling images
```
docker images -f dangling=true
```
Mostrar por id, lista em linha<br>
```
echo $(docker images -f dangling=true -q)
```
## Removendo imagens
Passe a busca como parâmetro para remover todas as imagens ao mesmo tempo.
```
docker rmi $(docker images -f dangling=true -q)
```
