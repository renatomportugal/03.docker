# GO

## Rodar

```CMD

Rodando os arquivos de forma externa, sem o docker-compose:

Copiar os arquivos para o servidor

docker build -t godocker .
docker run -p 8080:8080 godocker

Utilizando o docker-compose:
Vejam que o resultado é diferente.
Usando o docker-compose aparece "goDocker" em H1.
Usando a linha de comando, aparece apenas "Olá Mundo".

```
