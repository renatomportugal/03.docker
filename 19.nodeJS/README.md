# NodeJS

## Executar

```CMD
docker-compose up --build
```

## Imagem

```CMD
docker pull node
```

## Estrutura

```CMD
baixar o nodeJS na máquina
https://nodejs.org/en/download/

Abrir o cmd
node -v

Criar os arquivos index.js e Dockerfile

Na raiz da pasta do projeto
npm install express
```

## Executar

```CMD
Copiar a pasta para a máquina que tem o docker

Compilar
docker build -t nodedocker .
docker run -p 3000:3000 -d nodedocker
docker-compose up --build

```
