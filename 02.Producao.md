# Servidor Padrão
## Rede
Verifica se ter a rede bridge.66.6<br>
```
docker network ls
```
Se não tiver, crie<br>
```
docker network create \
  -d \
  bridge \
  --subnet 66.6.0.0/16 \
  --gateway 66.6.0.1 \
  bridge.66.6
```
Verifique quais IPs estão vagos<br>
```
docker network inspect bridge.66.6
```
Estou vendo uma forma de melhorar a apresentação dos dados:<br>
```
docker network inspect bridge.66.6 -f {{.Containers}}
```

## Volume
Verifique quais volumes existem:<br>
```
docker volume ls
```
Crie o volume para a aplicação (caso saiba qual pasta deseja persistir os dados)<br>
```
docker volume create mysql
docker volume create nextcloud
```
Verifique se foi criado<br>
```
docker volume ls
```
## Banco de dados
### MySQL
```
docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=my-password \
    -e MYSQL_USER=root \
    -e MYSQL_PASSWORD=my-password \
    -v mysql:/var/lib/mysql \
    --network bridge.66.6 \
    --name mysql.5.7 \
    --ip 66.6.0.5 \
    --restart=always \
    mysql:5.7
```
### phpmyadmin
```
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

### nextcloud
```
docker run \
  -d \
  -p 8000:80 \
  -v nextcloud:/var/www/html \
  --network bridge.66.6 \
  --name nextcloud \
  --ip 66.6.0.20 \
  --restart=always \
nextcloud
```
### nextcloud
```
777, nextcloud-mib, nc-mib, ip 66.6.0.20
8000, nextcloud-estudo, nc-estudo, ip 66.6.0.21
8001, nextcloud-media, nc-media, ip 66.6.0.22
8002, nextcloud-sp, nc-sp, ip 66.6.0.23
8003, nextcloud-code, nc-code, ip 66.6.0.25
8004, nextcloud-ha, nc-ha, ip 66.6.0.26
```
### collabora
```
sudo su
apt-get install apache2
a2enmod proxy
a2enmod proxy_wstunnel
a2enmod proxy_http
a2enmod ssl
systemctl restart apache2
exit

docker run \
  -t \
  -d \
  -p 9980:9980 \
  -e "extra_params=--o:ssl.enable=false" \
  --network bridge.66.6 \
  --name collabora\
  --ip 66.6.0.24 \
  --restart=always \
collabora/code

```

# Iniciar o Container com o Docker
Rode a imagem que deseja que seja iniciada com o sistema<br>

Liste a imagem para ver o id<br>
```
docker ps
```
Vamos suport que o id inicie com 98...<br>
Pare a imagem.<br>
```
docker stop 98
```
Vamos listar todas as imagens.<br>
```
docker ps -a
```
Vamos dar o comando para que a imagem inicie junto com o docker.<br>
```
docker update --restart=always 98
```

Comando para reverter.<br>
```
docker update --restart=no 98
```

Para verificar:<br>
```
docker inspect -f "{{ .HostConfig.RestartPolicy }}" <container_id>
docker inspect -f "{{ .RestartCount }}" <container_id>
docker inspect -f "{{ .State.StartedAt }}" <container_id>
```

Lista de todos os containers com id, status de restart e quantidadede (de restart) (que estão rodando)<br>
```
docker inspect -f "{{.Id}} {{.HostConfig.RestartPolicy.Name}} {{.RestartCount}}" $(docker container ls -q)
```

Lista de todos os containers com id, status de restart e quantidadede (de restart) (incluindo os que não estão rodando)<br>
docker inspect -f "{{.Id}} {{.HostConfig.RestartPolicy.Name}} {{.RestartCount}}" $(docker container ps -a -q)


Explicando:<br>
Este comando lista todos os ids dos containers que estão rodando. Para incluir todos adicione -a.<br>
```
$(docker container ls -a -q)
```
# Tranformando seu Ubuntu num servidor com DNS
Acesse https://www.noip.com/ e faça uma conta<br>

## Instalando o no-ip<br>
```
sudo su - 
cd /usr/local/bin
wget https://www.noip.com/client/linux/noip-duc-linux.tar.gz
tar xzf noip-duc-linux.tar.gz
ls
cd noip-2.1.9-1/
ls
make

deu erro de gcc, depois de pesquisar na internet:
apt update
apt install build-essential

apt install make
apt-get update
apt-get upgrade

Veja qual sua rede, selecione a interface de rede correta, rode o ifconfig e veja o ip do PC
wlp2s0

Em outro pc: ens34

Para iniciar a instalação:
make install

Apareceram 3 warnings, mas vamos lá...
Selecione a interface de rede que foi descoberta acima. No meu caso digitarei 0.

digita o email e senha

have them all updated, escolha sim caso queira que todos os endereços sejam redirecionados. Escolhi sim.
Caso escolha não na opção anterior, escolha qual endereço quer manter associado, selecione sim
selecione 5 para a atualização do intervalo
do you wish to run something at successful update, escolha não
o arquivo foi criado e movido para /usr/local/etc/no-ip2.conf

```
Fazer ele iniciar com o sistema<br>
```
sudo nano /etc/systemd/system/noip2.service

Cole:

[Unit]
Description=No-Ip Dynamic DNS Update Service
After=network.target

[Service]
Type=forking
ExecStart=/usr/local/bin/noip2

[Install]
WantedBy=multi-user.target

Ctrl+O, Enter
Ctrl+X

sudo systemctl daemon-reload
sudo systemctl status noip2
sudo systemctl enable noip2
sudo systemctl start noip2

reboot
sudo systemctl status noip2

 ```

# Formatando a saída
https://docs.docker.com/config/formatting/<br>
Use aspas simples externamente, internamente use aspas duplas como no exemplo a seguir:<br>
Este exemplo separa o id quando encontra a letra "e".<br>
Entrada c81e8baea8<br>
Saída c81 8ba a8<br>
```
docker inspect -f '{{split .Id "e"}} {{.HostConfig.RestartPolicy.Name}} {{.RestartCount}}' $(docker container ls -q)
```
Tamanho do campo<br>
```
docker inspect -f '{{len .Id}} {{.HostConfig.RestartPolicy.Name}} {{.RestartCount}}' $(docker container ls -q)
```

# Rancher
80,443   - Rancher - Server<br>
8001     - Rancher - php:7.1.26-apache<br>

## Containers
20,21,47400-47470 - bogem/ftp
3306 - mysql:5.7
8080 - phpmyadmin/phpmyadmin
8200 - 
