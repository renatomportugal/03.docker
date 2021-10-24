# NodeJS

## Imagem

```CMD
docker pull node
```

## Rodar_Imagem

```CMD
docker run -d -p 8000:8000 node

No cmd da máquina
curl --request POST \
  --url http://localhost:8000/test \
  --header 'content-type: application/json' \
  --data '{"msg": "testing"}'
```
