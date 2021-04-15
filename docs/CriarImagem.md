# Criar Imagem

## php:7.1.26-apache
Baixar a imagem<br>
```
docker pull php:7.1.26-apache
```

Criando uma imagem docker<br>
```
mkdir Docker
cd  Docker
code .
```
Criar o arquivo Dockerfile<br>
Criar o arquivo index.php<br>

insira no index.php:<br>
```
<?php 
    phpinfo();
?>
```

Insira no Dockerfile:<br>
```
FROM php:7.1.26-apache

COPY ./ /var/www/html

EXPOSE 80

CMD ["apache2-foreground"]
```

Criar a imagem (cria com a etiqueta latest)<br>
```
docker build -t minha_imagem .
```

Criar outra imagem com etiqueta v2<br>
```
docker build --tag minha_imagem:v2 .
```

Ver a imagem criada<br>
```
docker image ls
```
ou<br>
```
docker images
```

procurar as imagens que comecem com a palavra minha:<br>
```
docker images | grep minha
```

Rodar a imagem<br>
```
docker run -p 80:80 minha_imagem
```

## Ubuntu
Criando uma imagem docker<br>
```
mkdir imgUbuntu
cd  imgUbuntu
code .
```
Criar o arquivo Dockerfile<br>
```
FROM ubuntu
```
Criar a imagem (cria com a etiqueta latest)<br>
```
docker build -t demo_tree .
```
Rodar a imagem. Use --rm para substituir caso exista.<br>
```
docker run -it --rm demo_tree
```
Substitua o Dockerfile para:<br>
```
FROM ubuntu
```

Criar Imagem<br>
```
RUN apt-get update && apt-get install tree
```
Criar a imagem (cria com a etiqueta latest)<br>
```
docker build -t demo_tree .
```
Rodar a imagem. Use --rm para substituir caso exista<br>
```
docker run -it --rm demo_tree
```
Dentro do container execute:<br>
```
cd /root/
mkdir -p parent/child1
touch parent/child1/index.html
tree .
mkdir -p parent/child2
tree .
exit
```
Rode o camando tree e veja que foi instalado apenas no docker<br>
```
tree .
```

Abra outro container<br>
```
docker run -it --rm demo_tree
tree ./bin/
exit
```