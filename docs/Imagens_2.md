# Imagens

## Procurar as imagens que comecem com a palavra minha

```CMD
docker images | grep minha
```

## Procurar containers que tenham no nome a palavra tutorial

```CMD
docker ps -a | grep tutorial
```

## Ver o log do container (colocar pelo menos os três primeiros dígitos do id)<br>

```CMD
docker logs 921
```

## Criar_Imagem

### php:7.1.26-apache

Baixar a imagem

```CMD
docker pull php:7.1.26-apache
```

Criando uma imagem docker<br>

```CMD
mkdir Docker
cd  Docker
code .
```

Criar o arquivo Dockerfile<br>
Criar o arquivo index.php<br>

insira no index.php:<br>

```CMD
<?php 
    phpinfo();
?>
```

Insira no Dockerfile:<br>

```CMD
FROM php:7.1.26-apache

COPY ./ /var/www/html

EXPOSE 80

CMD ["apache2-foreground"]
```

Criar a imagem (cria com a etiqueta latest)<br>

```CMD
docker build -t minha_imagem .
```

Criar outra imagem com etiqueta v2<br>

```CMD
docker build --tag minha_imagem:v2 .
```

Ver a imagem criada<br>

```CMD
docker image ls
```

ou<br>

```CMD
docker images
```

procurar as imagens que comecem com a palavra minha:<br>

```CMD
docker images | grep minha
```

Rodar a imagem<br>

```CMD
docker run -p 80:80 minha_imagem
```

### Ubuntu
Criando uma imagem docker<br>

```CMD
mkdir imgUbuntu
cd  imgUbuntu
code .
```

Criar o arquivo Dockerfile<br>

```CMD
FROM ubuntu
```

Criar a imagem (cria com a etiqueta latest)<br>

```CMD
docker build -t demo_tree .
```

Rodar a imagem. Use --rm para substituir caso exista.<br>

```CMD
docker run -it --rm demo_tree
```

Substitua o Dockerfile para:<br>

```CMD
FROM ubuntu
```

Criar Imagem<br>

```CMD
RUN apt-get update && apt-get install tree
```

Criar a imagem (cria com a etiqueta latest)<br>

```CMD
docker build -t demo_tree .
```

Rodar a imagem. Use --rm para substituir caso exista<br>

```CMD
docker run -it --rm demo_tree
```

Dentro do container execute:<br>

```CMD
cd /root/
mkdir -p parent/child1
touch parent/child1/index.html
tree .
mkdir -p parent/child2
tree .
exit
```

Rode o camando tree e veja que foi instalado apenas no docker<br>

```CMD
tree .
```

Abra outro container<br>

```CMD
docker run -it --rm demo_tree
tree ./bin/
exit
```

## Deletar imagem<br>

```CMD
docker images
docker rmi nginx:1.16.0-alpine-perl
```

Ou usando o id (verifique se não existe mais de uma imagem com o mesmo início, -f é para forçar)<br>

```CMD
docker rmi -f a2
```