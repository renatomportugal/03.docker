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

## ExecutarNaMaquina

```CMD
Depois de executar a instalação do express...
npm install
npm start

Abra o browser em http://localhost:3000/, vai aparecer a mensagem Hello World!
```

## ExecutarNoDocker

```CMD
Antes de copiar, não se esqueça dos comandos abaixo
npm install express
npm install
npm start

Agora sim, copiar a pasta toda (inclusive a node_modules) para a máquina que tem o docker

Compilar
docker build -t nodedocker .
docker run -p 3000:3000 -d nodedocker
docker-compose up --build

```
