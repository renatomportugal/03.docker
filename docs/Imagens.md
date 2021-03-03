## Adminer
```
https://hub.docker.com/_/adminer

docker pull 

docker run \
  -d \
  -p 8080:8080 \
  --network mysql \
  --link mysql.5.7.wp \
  adminer
  
  469
  
  Substitui o phpMyAdmin
  
  docker run --link some_database:db -p 8080:8080 adminer
```

## Alpine
```
docker pull alpine

cria o container e vai para o shell dele
docker run -it alpine
ls

ver a versáo
cat /etc/os-release
ls /bin

verá que no alpine existem links para a busybox container
ls -la bin
```

## Backdrop
```
https://hub.docker.com/_/backdrop

docker pull backdrop

docker run \
  -d \
  -p 80:80 \
  -e BACKDROP_DB_HOST=mysql.5.7.wp \
  -e BACKDROP_DB_PORT=3306 \
  -e BACKDROP_DB_USER=root \
  -e BACKDROP_DB_PASSWORD=my-password \
  -e BACKDROP_DB_NAME=backdrop \
  --name backdrop \
  --network mysql \
  backdrop

e82

Deu erro

docker run --name some-backdrop --link some-mysql:mysql -d backdrop

The following environment variables are also honored for configuring your Backdrop CMS instance:

-e BACKDROP_DB_HOST=... (defaults to the IP and port of the linked mysql container)
-e BACKDROP_DB_USER=... (defaults to "root")
-e BACKDROP_DB_PASSWORD=... (defaults to the value of the MYSQL_ROOT_PASSWORD environment variable from the linked mysql container)
-e BACKDROP_DB_NAME=... (defaults to "backdrop")
-e BACKDROP_DB_PORT=... (defaults to 3306)
-e BACKDROP_DB_DRIVER=... (defaults to "mysql")
The BACKDROP_DB_NAME must already exist on the given MySQL server. Check out the official mysql image for more info on spinning up a DB.

If you'd like to be able to access the instance from the host without the container's IP, standard port mappings can be used:

$ docker run --name some-backdrop --link some-mysql:mysql -p 8080:80 -d backdrop
Then, access it via http://localhost:8080 or http://host-ip:8080 in a browser.

If you'd like to use an external database instead of a linked mysql container, specify the hostname and port with BACKDROP_DB_HOST/BACKDROP_DB_PORT along with the password in BACKDROP_DB_PASSWORD and the username in BACKDROP_DB_USER (if it is something other than root):

docker run --name some-backdrop \
  -e BACKDROP_DB_HOST=10.1.2.3 \
  -e BACKDROP_DB_PORT=10432 \
  -e BACKDROP_DB_USER=... \
  -e BACKDROP_DB_PASSWORD=... \
  -d backdrop
```

## bitnami-docker-phabricator
```
https://github.com/bitnami/bitnami-docker-phabricator
https://bitnami.com/stack/phabricator/containers
```

## bitnami-edx
```
sudo apt-get update -y
sudo apt-get upgrade -y
sudo reboot

sudo apt-get install make

Instalar o Docker:
https://github.com/renatomportugal/docker/blob/master/01.Install.md

Instalar o Docker-Compose:
https://github.com/renatomportugal/docker/blob/master/08.Docker-compose.md

Instalar o Python:
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository ppa:deadsnakes/ppa
sudo apt install python3.7

https://python.org.br/instalacao-linux/
Instalar o Python no Centos 7:
which python
which python3
sudo yum install python3
sudo yum -y install python3-pip

docker info | grep -i 'storage driver'
Saída:
 Storage Driver: overlay2
WARNING: No swap limit support

__Bitnami
sudo su
mkdir bitnami-edx
cd bitnami-edx
wget https://bitnami.com/redirect/to/1084950/bitnami-edx-ironwood.2-9-linux-x64-installer.run
chmod +x ./bitnami-edx-ironwood.2-9-linux-x64-installer.run
./bitnami-edx-ironwood.2-9-linux-x64-installer.run

Demo course for Open edX [Y/n] :y
Is the selection above correct? [Y/n]: Y
Select a folder [/opt/edx-ironwood.2-9]:
Your real name [User Name]:
Email Address [user@example.com]:
Login [user]:
Password :
Please confirm your password :
Hostname [127.0.1.1]:192.168.1.106
Do you want to configure mail support? [y/N]:N
(Desmarque a opção para colocar na nuvem)
Do you want to continue? [Y/n]:Y

Installing
 0% ______________ 50% ______________ 100%
 #############
 
 Setup has finished installing Open edX powered by Bitnami on your computer.
 
 __Para rodar...
 sudo su
 cd /opt/edx-ironwood.2-9
 ./bitnami... manager
 
 Acesse o endereço no browser...
 

Info: In order to use CodeJail, additional configuration is required. Find more
information at:
https://docs.bitnami.com/installer/apps/edx/configuration/sandbox-codejail/
Press [Enter] to continue:

Info: To access the Open edX powered by Bitnami, go to
http://127.0.0.1:80 from your browser.
Press [Enter] to continue:

Funcionou!!!
 
__Modo Nativo (Não testado)
https://openedx.atlassian.net/wiki/spaces/OpenOPS/pages/146440579/Native+Open+edX+platform+Ubuntu+16.04+64+bit+Installation
```

## bitnami-laravel
```
https://laravel.com/docs/7.x
https://hub.docker.com/r/bitnami/laravel

docker pull bitnami/laravel

mkdir ~/myapp && cd ~/myapp
curl -LO https://raw.githubusercontent.com/bitnami/bitnami-docker-laravel/master/docker-compose.yml
docker-compose up

acesse localhost:30000
```

## bitnami-moodle
```
docker pull bitnami/moodle
https://hub.docker.com/r/bitnami/moodle/

Se der falha na instalação
docker volume ls
docker volume rm bitnami-moodle_mariadb_data
docker volume rm bitnami-moodle_moodle_data
docker volume ls
______________________________
__Usando Composer
Instale o docker compose na sua máquina, veja o arquivo docker/08.Docker-compose.md, acesso em:
https://github.com/renatomportugal/docker/blob/master/08.Docker-compose.md

mkdir bitnami-moodle
cd bitnami-moodle
curl -sSL https://raw.githubusercontent.com/bitnami/bitnami-docker-moodle/master/docker-compose.yml > docker-compose.yml
docker-compose up -d

Run the application manually
If you want to run the application manually instead of using docker-compose, these are the basic steps you need to run:

Create a new network for the application and the database:

$ docker network create moodle-tier
Create a volume for MariaDB persistence and create a MariaDB container

$ docker volume create --name mariadb_data
$ docker run -d --name mariadb \
 -e ALLOW_EMPTY_PASSWORD=yes \
 -e MARIADB_USER=bn_moodle \
 -e MARIADB_DATABASE=bitnami_moodle \
 --net moodle-tier \
 --volume mariadb_data:/bitnami \
 bitnami/mariadb:latest
Note: You need to give the container a name in order to Moodle to resolve the host

Create volumes for Moodle persistence and launch the container

$ docker volume create --name moodle_data
$ docker run -d --name moodle -p 80:80 -p 443:443 \
 -e ALLOW_EMPTY_PASSWORD=yes \
 -e MOODLE_DATABASE_USER=bn_moodle \
 -e MOODLE_DATABASE_NAME=bitnami_moodle \
 --net moodle-tier \
 --volume moodle_data:/bitnami \
 bitnami/moodle:latest
 
 Persisting your application
If you remove the container all your data and configurations will be lost, and the next time you run the image the database will be reinitialized. To avoid this loss of data, you should mount a volume that will persist even after the container is removed.

For persistence you should mount a volume at the /bitnami path. Additionally you should mount a volume for persistence of the MariaDB data.

The above examples define docker volumes namely mariadb_data and moodle_data. The Moodle application state will persist as long as these volumes are not removed.

To avoid inadvertent removal of these volumes you can mount host directories as data volumes. Alternatively you can make use of volume plugins to host the volume data.

Mount persistent folders in the host using docker-compose
This requires a minor change to the docker-compose.yml file present in this repository:

services:
  mariadb:
  ...
    volumes:
      - '/path/to/mariadb-persistence:/bitnami'
  ...
  moodle:
  ...
    depends_on:
      - mariadb
  ...
Mount persistent folders manually
In this case you need to specify the directories to mount on the run command. The process is the same than the one previously shown:

If you haven't done this before, create a new network for the application and the database:

$ docker network create moodle-tier
Start a MariaDB database in the previous network:

$ docker run -d --name mariadb \
 -e ALLOW_EMPTY_PASSWORD=yes \
 -e MARIADB_USER=bn_moodle \
 -e MARIADB_DATABASE=bitnami_moodle \
 -v /path/to/mariadb-persistence:/bitnami \
 --net moodle-tier \
 bitnami/mariadb:latest
Note: You need to give the container a name in order to Moodle to resolve the host

Run the Moodle container:

$ docker run -d -p 80:80 -p 443:443 --name moodle \
 -e ALLOW_EMPTY_PASSWORD=yes \
 -e MOODLE_DATABASE_USER=bn_moodle \
 -e MOODLE_DATABASE_NAME=bitnami_moodle \
 --net moodle-tier \
 --volume /path/to/moodle-persistence:/bitnami \
 bitnami/moodle:latest
```

## bogem-ftp
```
https://hub.docker.com/r/bogem/ftp

docker volume create bogem

docker run \
  -d \
  -v bogem:/home/vsftpd \
  -p 20:20 \
  -p 21:21 \
  -p 47400-47470:47400-47470 \
  -e FTP_USER=ekode \
  -e FTP_PASS=ekode123 \
  -e PASV_ADDRESS=192.168.1.106 \
  --name bogem-ftp \
  --network bridge.66.6 \
  --restart=always \
bogem/ftp
```