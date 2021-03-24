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

## bonita
```
https://hub.docker.com/_/bonita
_______________________________________________________________________________
__Baixar_______________________________________________________________________
docker pull bonita
_______________________________________________________________________________
__Instalar_____________________________________________________________________
docker run \
  -d \
  -p 8081:8080 \
  --network mysql \
  --name bonita \
  bonita
07
_______________________________________________________________________________
__Acessar______________________________________________________________________
http://localhost:8081/bonita/
_______________________________________________________________________________
__Configurar___________________________________________________________________
usuário: install
password: install

Settings, Language, Português do Brasil.

Crie um usuário, adicione ele aos administradores, 
Logue com o novo usuário, Aplicações, Novo aplicativo, 

Para instalar um modelo (Ainda não criamos)
Menu, BPM Services, Pause, BDM, Install Business Data Model, 

Resources, 

Documentação:
https://documentation.bonitasoft.com/bonita/7.9/
_______________________________________________________________________________
MySQL
There are known issues with the management of XA transactions by MySQL engine and driver: see MySQL bugs 17343 and 12161 for more details. Thus, using MySQL database in a production environment is not recommended.

Increase the packet size which is set by default to 1M:

$ mkdir -p custom_mysql
$ echo "[mysqld]" > custom_mysql/bonita.cnf
$ echo "max_allowed_packet=16M" >> custom_mysql/bonita.cnf
Mount that directory location as /etc/mysql/conf.d inside the MySQL container:

$ docker run --name mydbmysql -v "$PWD"/custom_mysql/:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=mysecretpassword -d mysql:8.0
See the official MySQL documentation for more details.

Start your application container to link it to the MySQL container:

$ docker run --name bonita_mysql --link mydbmysql:mysql -d -p 8080:8080 bonita

Modify default credentials
docker run --name=bonita -e "TENANT_LOGIN=tech_user" -e "TENANT_PASSWORD=secret" -e "PLATFORM_LOGIN=pfadmin" -e "PLATFORM_PASSWORD=pfsecret" -d -p 8080:8080 bonita

Now you can access the Bonita Portal on localhost:8080/bonita and login using: tech_user / secret
```

## busybox
```
docker run busybox

docker run -it busybox

digite ls e verá que a lista de pastas é bem menor.
liste ls /bin e verá que também tem o comando echo, que poderá ser usado
echo "Olá mundo"

teste o comando uptime

crie uma pasta
mkdir my-folder

entre na pasta e crie um arquivo de texto
mkdir my-folder
cd my-folder
touch file.txt
ls
```

## centos8
```
docker pull centos:centos8
docker run --name centos08.server -ti -d --privileged=true centos:8 "/sbin/init"
cat /etc/centos-release
    CentOS Linux release 8.2.2004 (Core)

dnf update
dnf check-update
dnf clean all
dnf install nano vim wget curl net-tools lsof bash-completion

__Instalar Docker

dnf remove docker \
    docker-client \
    docker-client-latest \
    docker-common \
    docker-latest \
    docker-latest-logrotate \
    docker-logrotate \
    docker-engine

dnf -y install epel-release
dnf repolist
dnf info epel-release
dnf install 'dnf-command(config-manager)'
dnf config-manager --set-enabled PowerTools
dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
dnf list docker-ce
dnf list containerd.io
dnf update --skip-broken
dnf update --nobest
dnf update
dnf repolist
dnf --disablerepo="*" --enablerepo="epel" list available | wc -l
dnf install docker-ce docker-ce-cli containerd.io --nobest
docker -v
    Docker version 19.03.12, build 48a66213fe

docker ps
Deu o erro:
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?






dnf repolist -v
usermod -aG docker root
id root

systemctl enable docker
systemctl enable docker.service
systemctl start docker
systemctl status docker

/var/run/docker.sock

systemctl unmask docker
systemctl unmask docker.service
systemctl unmask docker.socket

systemctl enable --now docker

dnf module enable container-tools && \
dnf install container-selinux && \
dnf module disable container-tools && \
dnf install docker-ce

dnf update
yum install -y yum-utils


dnf install docker-ce --nobest
dnf list docker-ce --showduplicates | sort -r

dnf install docker-ce-19.03.12 docker-ce-cli-19.03.12 containerd.io --skip-broken

dnf -y upgrade

curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh



dnf install https://download.docker.com/linux/centos/7/x86_64/stable/Packages/containerd.io-1.2.6-3.3.el7.x86_64.rpm


yum install docker-ce



wget https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/master/files/docker/systemctl.py -O /usr/local/bin/systemctl


docker run hello-world

docker run -i -t centos:8 /bin/bash
yum update
cat /etc/issue
ps -ef
exit
docker start 42
docker ps
docker attach 42
exit
docker diff 42
docker exec -it 42 bash
yum update


ifconfig
ip a
ping -c2 google.com

ethtool eth0
mii-tool eth0

netstat -tulpn

ss -tulpn
lsof -i4 -6

useradd tcnct
usermod -aG wheel tcnct

su - tcnct
exit
dnf update

netstat -tulpn

__Instalar Docker

yum remove docker \
    docker-client \
    docker-client-latest \
    docker-common \
    docker-latest \
    docker-latest-logrotate \
    docker-logrotate \
    docker-engine

yum -y install epel-release
yum repolist
dnf info epel-release

dnf install 'dnf-command(config-manager)'
dnf config-manager --set-enabled PowerTools

dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo

dnf list docker-ce
dnf list containerd.io
yum update --skip-broken
yum update --nobest

dnf update
dnf repolist
dnf --disablerepo="*" --enablerepo="epel" list available | wc -l

dnf install docker-ce --nobest
yum install docker-ce docker-ce-cli containerd.io --nobest

yum list docker-ce --showduplicates | sort -r

yum install docker-ce-19.03.12 docker-ce-cli-19.03.12 containerd.io --skip-broken

yum -y upgrade

curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

yum install -y https://download.docker.com/linux/centos/7/x86_64/stable/Packages/containerd.io-1.2.6-3.3.el7.x86_64.rpm

dnf install https://download.docker.com/linux/centos/7/x86_64/stable/Packages/containerd.io-1.2.6-3.3.el7.x86_64.rpm

yum install docker-ce

docker -v

dnf repolist -v

usermod -aG docker root
id root

wget https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/master/files/docker/systemctl.py -O /usr/local/bin/systemctl

systemctl enable docker

systemctl enable docker.service

systemctl enable --now docker

dnf module enable container-tools && \
dnf install container-selinux && \
dnf module disable container-tools && \
dnf install docker-ce





/var/run/docker.sock

systemctl unmask docker
systemctl unmask docker.service
systemctl unmask docker.socket

dnf update

yum install -y yum-utils



systemctl enable --now docker
systemctl status docker
systemctl disable firewalld

docker run hello-world

Instalar docker-compose
dnf install curl -y

curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

chmod +x /usr/local/bin/docker-compose

docker-compose -v

```

## centos8
```
https://developer.valvesoftware.com/wiki/Counter-Strike:_Global_Offensive_Dedicated_Servers

docker run -d --net=host --name=csgo-dedicated -e SRCDS_TOKEN={YOURTOKEN} cm2network/csgo

Tentar por aqui
https://linuxgsm.com/lgsm/csgoserver/

https://hub.docker.com/_/ubuntu
docker pull ubuntu

docker run \
  -it \
  -p 27015:27015 \
  -p 27016:27016 \
  -p 27021:27021 \
  -e LANG=C.UTF-8 \
  --name ubuntu \
  ubuntu

e8f  
  

  
  

docker network ls
docker update \
  --network bridge.66.6 \
  --ip 66.6.0.17 \
  ubuntu



exit
docker ps -a | grep ubuntu
docker exec -it ubuntu bash

dpkg --add-architecture i386; apt update; apt install mailutils postfix curl wget file tar bzip2 gzip unzip bsdmainutils python util-linux ca-certificates binutils bc jq tmux netcat lib32gcc1 lib32stdc++6 steamcmd
adduser csgoserver
su - csgoserver
wget -O linuxgsm.sh https://linuxgsm.sh && chmod +x linuxgsm.sh && bash linuxgsm.sh csgoserver
./csgoserver install

./csgoserver

./csgoserver status
./csgoserver start
./csgoserver stop
./csgoserver restart
./csgoserver console


./csgoserver update
./csgoserver force-update
./csgoserver validate
./csgoserver details
./csgoserver debug
/home/csgoserver/logs
./csgoserver monitor

Para reconhecer o ifconfig
apt-get update
apt-get install pciutils

Veja qual a placa de rede
lspci

apt-get install net-tools
apt-get install nano

apt-get update
apt-get install iputils-ping

nano /etc/resolv.conf
nameserver 192.168.1.1

apt-get install systemd
Ou para reinstalar
apt-get install --reinstall systemd




__outro________________________________________________
  docker run \
  -it \
  -p 27015:27015 \
  -p 27016:27016 \
  -p 27021:27021 \
  -e LANG=C.UTF-8 \
  --name ubuntu2 \
  --network bridge.66.6 \
  --ip 66.6.0.17 \
  ubuntu

c6d

Para reconhecer o lspci
apt-get update
apt-get install pciutils
lspci

Para reconhecer o ifconfig
apt-get update
apt-get install net-tools
ifconfig

Instalar o nano
apt-get install nano

Instalar o ping
apt-get update
apt-get install iputils-ping

https://hub.docker.com/r/dorowu/ubuntu-desktop-lxde-vnc

sudo dpkg --add-architecture i386; sudo apt update; sudo apt install mailutils postfix curl wget file tar bzip2 gzip unzip bsdmainutils python util-linux ca-certificates binutils bc jq tmux netcat lib32gcc1 lib32stdc++6 steamcmd

```

## centos8
```
CentOS (Funcionou mexendo no firewall)
sudo yum install httpd -y
sudo systemctl enable httpd && sudo systemctl start httpd

yum install firewalld
systemctl start firewalld
systemctl status firewalld -l

httpd -M | grep proxy
_______________________
No Ubuntu (funcionou)

sudo su
apt-get install apache2
a2enmod proxy
a2enmod proxy_wstunnel
a2enmod proxy_http
a2enmod ssl
exit
________________
docker run -d -p 80:80 --name nc2 nextcloud
docker run -t -d -p 9980:9980 -e "extra_params=--o:ssl.enable=false" --name co2 collabora/code

docker run -d -p 8000:80 nextcloud
docker run -t -d -p 9980:9980 -e "extra_params=--o:ssl.enable=false" collabora/code

no nextcloud:
docker exec -it nc2 bash
apt-get update
apt install nano
nano /var/www/html/config/config.php

array (
    0 => 'localhost',
    1 => '192.168.1.115',
  ),
  


https://github.com/owncloud/richdocuments
https://www.collaboraoffice.com/code/docker/
https://www.collaboraoffice.com/code/quick-tryout-nextcloud-docker/

docker pull collabora/code

docker network ls

docker network create \
  -d \
  bridge \
  --subnet 66.6.0.0/16 \
  --gateway 66.6.0.1 \
  bridge.66.6

docker network ls

docker network inspect bridge.66.6

docker volume ls

docker volume create mysql

docker volume ls

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

docker run \
    -d \
    -p 8080:80 \
    -e PMA_HOST=mysql.5.7 \
    --network bridge.66.6 \
    --name phpmyadmin \
    --ip 66.6.0.15 \
    --restart=always \
phpmyadmin/phpmyadmin

docker volume ls

docker volume create nextcloud

docker volume ls

docker run \
  -d \
  -p 80:80 \
  -v nextcloud:/var/www/html \
  --network bridge.66.6 \
  --name nextcloud \
  --ip 66.6.0.20 \
  --restart=always \
nextcloud

Instale o nextcloud e adicione o app Collabora.

docker run \
  -t \
  -d \
  -p 9980:9980 \
  -e "extra_params=--o:ssl.enable=false" \
  --network bridge.66.6 \
  --name collabora\
  --ip 66.6.0.21 \
  --restart=always \
collabora/code

Set up the Collabora Online server in Nextcloud Settings – Collabora Online to http://localhost:9980






docker run -t -d -p 127.0.0.1:9980:9980 -e "domain=<your-dot-escaped-domain>" \
        -e "username=admin" -e "password=S3cRet" --restart always collabora/code


docker run -t -d \
-p 9980:9980 \
-e "username=admin" \
-e "password=admin" \
--network mysql \
--name collabora \
--cap-add MKNOD \
collabora/code

docker exec -it collabora /bin/bash -c "apt-get -y update && apt-get -y install xmlstarlet && xmlstarlet ed --inplace -u \"/config/ssl/enable\" -v false /etc/loolwsd/loolwsd.xml && xmlstarlet ed --inplace -u \"/config/ssl/termination\" -v false /etc/loolwsd/loolwsd.xml"

Reading package lists... Done
E: List directory /var/lib/apt/lists/partial is missing. - Acquire (13: Permission denied)



_____________

CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
2f7b4b5a0da4        collabora/code      "/bin/sh -c 'bash st…"   2 hours ago         Up 16 minutes       0.0.0.0:9980->9980/tcp   awesome_tesla
02774f11162d        nextcloud           "/entrypoint.sh apac…"   3 hours ago         Up 16 minutes       0.0.0.0:8000->80/tcp     pensive_noyce

docker network inspect bridge
[
 {
        "Name": "bridge",
        "Id": "c8245f92e7d0dc249addc703a3864a8c11b3e989c54bf322b6757abb89668de2",
        "Created": "2020-07-13T21:24:58.538497159Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "02774f11162dc357518554d267e101925908998c3de5ffc499ed7d29c5bb4946": {
                "Name": "pensive_noyce",
                "EndpointID": "739267a52dcd686d6f01a1fdd2aaf866c62725a2de3177ec88516a094e8cce9f",
                "MacAddress": "02:42:ac:11:00:03",
                "IPv4Address": "172.17.0.3/16",
                "IPv6Address": ""
            },
            "2f7b4b5a0da4ea9fa0533a171a544744eeb3ecd501b708957f63129967061e49": {
                "Name": "awesome_tesla",
                "EndpointID": "06904d6bd5eb9893d34470d78d3dde8c12d093ea236395f25690eeb92dd60f0e",
                "MacAddress": "02:42:ac:11:00:02",
                "IPv4Address": "172.17.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {}
    }
]

docker inspect awesome_tesla
[
    {
        "Id": "2f7b4b5a0da4ea9fa0533a171a544744eeb3ecd501b708957f63129967061e49",
        "Created": "2020-07-13T19:20:37.96400956Z",
        "Path": "/bin/sh",
        "Args": [
            "-c",
            "bash start-libreoffice.sh"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 1788,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2020-07-13T21:25:42.513979741Z",
            "FinishedAt": "2020-07-13T21:24:05.372773702Z"
        },
        "Image": "sha256:8ae6850294e5b69618d7295303169804790caf076bfb1b0c82c1c9cbbbb668f5",
        "ResolvConfPath": "/var/lib/docker/containers/2f7b4b5a0da4ea9fa0533a171a544744eeb3ecd501b708957f63129967061e49/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/2f7b4b5a0da4ea9fa0533a171a544744eeb3ecd501b708957f63129967061e49/hostname",
        "HostsPath": "/var/lib/docker/containers/2f7b4b5a0da4ea9fa0533a171a544744eeb3ecd501b708957f63129967061e49/hosts",
        "LogPath": "/var/lib/docker/containers/2f7b4b5a0da4ea9fa0533a171a544744eeb3ecd501b708957f63129967061e49/2f7b4b5a0da4ea9fa0533a171a544744eeb3ecd501b708957f63129967061e49-json.log",
        "Name": "/awesome_tesla",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {
                "9980/tcp": [
                    {
                        "HostIp": "",
                        "HostPort": "9980"
                    }
                ]
            },
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "Capabilities": null,
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/079933b3b759ef95f6244e7574135f0cd9f9465d9c0bf38c9a167f972a8b4562-init/diff:/var/lib/docker/overlay2/a14a22f4bd0275f120963002c792c6d9e11e206ef9c81eaf856042a4a4161039/diff:/var/lib/docker/overlay2/8458c11e3fc7f4dfe8e04c8390ffdf2af0abfba141fc378ac003b50d4b545945/diff:/var/lib/docker/overlay2/6d50f87d15168277f84710ef9718017621844cdb6a019fd1728b1cd1dc9e2868/diff:/var/lib/docker/overlay2/61f79fccf85086cceebaf6562233dc17950716641b0bbcd416b0836b50fe9f77/diff:/var/lib/docker/overlay2/9e766d953ed49eb9e1d8e38dfadfb1b5ce30d6d03b84592ee313de49b21e0252/diff:/var/lib/docker/overlay2/c74742803e733fe06f22bcadfe73b949a1d505c75ed99cc6bda810e3cd6762f6/diff:/var/lib/docker/overlay2/aa5e127fa42b6d7fbcf1310ba2e1017052fd3ca824ef9db8db049f37d1fb1487/diff",
                "MergedDir": "/var/lib/docker/overlay2/079933b3b759ef95f6244e7574135f0cd9f9465d9c0bf38c9a167f972a8b4562/merged",
                "UpperDir": "/var/lib/docker/overlay2/079933b3b759ef95f6244e7574135f0cd9f9465d9c0bf38c9a167f972a8b4562/diff",
                "WorkDir": "/var/lib/docker/overlay2/079933b3b759ef95f6244e7574135f0cd9f9465d9c0bf38c9a167f972a8b4562/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "2f7b4b5a0da4",
            "Domainname": "",
            "User": "101",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "9980/tcp": {}
            },
            "Tty": true,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "extra_params=--o:ssl.enable=false",
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "domain=localhost",
                "LC_CTYPE=C.UTF-8"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "bash start-libreoffice.sh"
            ],
            "Image": "collabora/code",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "55494357cbe337204b24014cff6294050369386a14c25c0bb755c82943cab239",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "9980/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "9980"
                    }
                ]
            },
            "SandboxKey": "/var/run/docker/netns/55494357cbe3",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "06904d6bd5eb9893d34470d78d3dde8c12d093ea236395f25690eeb92dd60f0e",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "c8245f92e7d0dc249addc703a3864a8c11b3e989c54bf322b6757abb89668de2",
                    "EndpointID": "06904d6bd5eb9893d34470d78d3dde8c12d093ea236395f25690eeb92dd60f0e",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
        }
    }
]

docker inspect pensive_noyce
[
    {
        "Id": "02774f11162dc357518554d267e101925908998c3de5ffc499ed7d29c5bb4946",
        "Created": "2020-07-13T19:05:01.588289607Z",
        "Path": "/entrypoint.sh",
        "Args": [
            "apache2-foreground"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 1916,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2020-07-13T21:25:46.335580177Z",
            "FinishedAt": "2020-07-13T21:24:03.900585528Z"
        },
        "Image": "sha256:327476ebe3280c7b570d8463edd136956eab120959976b643cb7dbfaa73f98c1",
        "ResolvConfPath": "/var/lib/docker/containers/02774f11162dc357518554d267e101925908998c3de5ffc499ed7d29c5bb4946/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/02774f11162dc357518554d267e101925908998c3de5ffc499ed7d29c5bb4946/hostname",
        "HostsPath": "/var/lib/docker/containers/02774f11162dc357518554d267e101925908998c3de5ffc499ed7d29c5bb4946/hosts",
        "LogPath": "/var/lib/docker/containers/02774f11162dc357518554d267e101925908998c3de5ffc499ed7d29c5bb4946/02774f11162dc357518554d267e101925908998c3de5ffc499ed7d29c5bb4946-json.log",
        "Name": "/pensive_noyce",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {
                "80/tcp": [
                    {
                        "HostIp": "",
                        "HostPort": "8000"
                    }
                ]
            },
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "Capabilities": null,
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/d6145a14fa0d39b23671e537aa98c2fc9dc353b64f797ae0c00dfe188f3a539a-init/diff:/var/lib/docker/overlay2/aa2f3916e4b2c690e90c824d4e6a7708072da7fc7b3a4e5f220e9c3c656961d0/diff:/var/lib/docker/overlay2/ee43d8b1c0c7cb9f58892f2129faae2f6604d7ca95e06024def080595db99412/diff:/var/lib/docker/overlay2/580356a45f3528b4ceaaf758bfbccb978b0fd9fc8aea0e2909c33333e9950143/diff:/var/lib/docker/overlay2/b86af3b98a349d60a396042d569ba5ec910caa0ba36b2d87ed07cecdd71312af/diff:/var/lib/docker/overlay2/64c2a8a0513328d5e83f7f423587579d92de279ad1e9ce68f4eba95bc3c1a4bb/diff:/var/lib/docker/overlay2/9fac805e7be1bc41d389b9465320330f4748fb5951f1e7c5efe2c1e4f85963b8/diff:/var/lib/docker/overlay2/f14dee23b2a760b7a39fedb2a219bc966c870ae8c3d0aebcf3ada241a8c2e7e6/diff:/var/lib/docker/overlay2/5e6a99a8f75ce1bc4ae3ee7f380561bc2c48d974c0eac08052d6a48885c90533/diff:/var/lib/docker/overlay2/ced94cf9034dd0d5f3f9cea0848e3667598a2a579ca4df4970d1ffbc2b3c16f1/diff:/var/lib/docker/overlay2/f2df497ad8819ac545cf73aefe1c4c714779394a8e03d9192d64e8204d8f4efa/diff:/var/lib/docker/overlay2/ce5a50c70bad47078f6767ebcb7695abab093df7e2a18b862a8128e1f4691b3c/diff:/var/lib/docker/overlay2/37157d8f1f44d24e0e987167e8526c0a302dcd63e1a16831d082b918e78a78fc/diff:/var/lib/docker/overlay2/f6a8074f627c13e3d9e32b7720660bc034fb67e0d34fdc8464f23511e2cf12d6/diff:/var/lib/docker/overlay2/3e5b01340f81762a662c06277d768da4f5e88b433e090b87ccd94457862cebb7/diff:/var/lib/docker/overlay2/8d87dee9fb831bea0bd54b14dee2fadf706c9b3393e59e42382b5dbbdbc58ef7/diff:/var/lib/docker/overlay2/b152a482f45b595eebbd3a25c9f564d511ed10e3ebab41ecd656cd488616db1c/diff:/var/lib/docker/overlay2/c67cd8345b64fcaaba9a5c0e3a36cae3521d4cbaa90ec4c3da6dbf074e550dcf/diff:/var/lib/docker/overlay2/0b6370368ee0d935bd0ad2b19f85849eb1f4485b165ff149d542bd7ef14a4190/diff:/var/lib/docker/overlay2/1d5ec75ee1c91fc0b88bee00525df9a49d5d80b8a1ef495393ec56494250147d/diff:/var/lib/docker/overlay2/bb9570cd1dfd2ed0538249ffe01bc521c9428332c2bccf61230fc10a6282b487/diff",
                "MergedDir": "/var/lib/docker/overlay2/d6145a14fa0d39b23671e537aa98c2fc9dc353b64f797ae0c00dfe188f3a539a/merged",
                "UpperDir": "/var/lib/docker/overlay2/d6145a14fa0d39b23671e537aa98c2fc9dc353b64f797ae0c00dfe188f3a539a/diff",
                "WorkDir": "/var/lib/docker/overlay2/d6145a14fa0d39b23671e537aa98c2fc9dc353b64f797ae0c00dfe188f3a539a/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [
            {
                "Type": "volume",
                "Name": "b6b0818295646780e8e69a681faa890ce6ba7e2736462cf21d1f5f686710d047",
                "Source": "/var/lib/docker/volumes/b6b0818295646780e8e69a681faa890ce6ba7e2736462cf21d1f5f686710d047/_data",
                "Destination": "/var/www/html",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            }
        ],
        "Config": {
            "Hostname": "02774f11162d",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "80/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "PHPIZE_DEPS=autoconf \t\tdpkg-dev \t\tfile \t\tg++ \t\tgcc \t\tlibc-dev \t\tmake \t\tpkg-config \t\tre2c",
                "PHP_INI_DIR=/usr/local/etc/php",
                "APACHE_CONFDIR=/etc/apache2",
                "APACHE_ENVVARS=/etc/apache2/envvars",
                "PHP_EXTRA_BUILD_DEPS=apache2-dev",
                "PHP_EXTRA_CONFIGURE_ARGS=--with-apxs2 --disable-cgi",
                "PHP_CFLAGS=-fstack-protector-strong -fpic -fpie -O2 -D_LARGEFILE_SOURCE -D_FILE_OFFSET_BITS=64",
                "PHP_CPPFLAGS=-fstack-protector-strong -fpic -fpie -O2 -D_LARGEFILE_SOURCE -D_FILE_OFFSET_BITS=64",
                "PHP_LDFLAGS=-Wl,-O1 -pie",
                "GPG_KEYS=42670A7FE4D0441C8E4632349E4FDC074A4EF02D 5A52880781F755608BF815FC910DEB46F53EA312",
                "PHP_VERSION=7.4.8",
                "PHP_URL=https://www.php.net/distributions/php-7.4.8.tar.xz",
                "PHP_ASC_URL=https://www.php.net/distributions/php-7.4.8.tar.xz.asc",
                "PHP_SHA256=642843890b732e8af01cb661e823ae01472af1402f211c83009c9b3abd073245",
                "PHP_MD5=",
                "NEXTCLOUD_VERSION=19.0.0"
            ],
            "Cmd": [
                "apache2-foreground"
            ],
            "Image": "nextcloud",
            "Volumes": {
                "/var/www/html": {}
            },
            "WorkingDir": "/var/www/html",
            "Entrypoint": [
                "/entrypoint.sh"
            ],
            "OnBuild": null,
            "Labels": {},
            "StopSignal": "SIGWINCH"
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "660f3e0c505c5823cdc0a525d16af4e740195be1726614323f5cccfdef40fd78",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "80/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "8000"
                    }
                ]
            },
            "SandboxKey": "/var/run/docker/netns/660f3e0c505c",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "739267a52dcd686d6f01a1fdd2aaf866c62725a2de3177ec88516a094e8cce9f",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.3",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:03",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "c8245f92e7d0dc249addc703a3864a8c11b3e989c54bf322b6757abb89668de2",
                    "EndpointID": "739267a52dcd686d6f01a1fdd2aaf866c62725a2de3177ec88516a094e8cce9f",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.3",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:03",
                    "DriverOpts": null
                }
            }
        }
    }
]


docker exec -it collabora bash

su root -s /bin/sh

Espere um temo e reinicie.
docker restart collabora

Access Collabora Admin at http://[your-host-public-ip]:9980/loleaflet/dist/admin/admin.html e.g. 172.16.12.95, 
and set in Settings -> Admin -> Additional -> Collabora Online server -> http://[your-host-public-ip]:9980
```

## docker
```
docker pull docker

docker network create \
  -d \
  bridge \
  --subnet 66.6.0.0/16 \
  --gateway 66.6.0.1 \
  bridge.66.6

docker run --privileged --name some-docker -d \
    --network bridge.66.6 --network-alias docker \
    -e DOCKER_TLS_CERTDIR=/certs \
    -v some-docker-certs-ca:/certs/ca \
    -v some-docker-certs-client:/certs/client \
    docker:dind


docker run --rm --network  bridge.66.6 \
    -e DOCKER_TLS_CERTDIR=/certs \
    -v some-docker-certs-client:/certs/client:ro \
    docker:latest version

docker run -it --rm --network  bridge.66.6 \
    -e DOCKER_TLS_CERTDIR=/certs \
    -v some-docker-certs-client:/certs/client:ro \
    docker:latest sh

docker run --rm --network  bridge.66.6 \
    -e DOCKER_TLS_CERTDIR=/certs \
    -v some-docker-certs-client:/certs/client:ro \
    docker:latest info

docker run --rm -v /var/run/docker.sock:/var/run/docker.sock docker:latest version

docker run --privileged --name some-docker -d \
    --network  bridge.66.6 --network-alias docker \
    -e DOCKER_TLS_CERTDIR=/certs \
    -v some-docker-certs-ca:/certs/ca \
    -v some-docker-certs-client:/certs/client \
    docker:dind --storage-driver overlay2

Where to store data
docker run --privileged --name some-docker -v /my/own/var-lib-docker:/var/lib/docker -d docker:dind
```

## drupal
```
https://hub.docker.com/_/drupal

docker pull drupal

docker run \
  -d \
  -p 8200:80 \
  --name drupal \
  --network bridge.66.6 \
  --ip 66.6.0.200 \
  --restart=always \
drupal

__Configuração (8.9.1)
Escolha a língua: Português, Brasil, botão Save and continue
Selecione o perfil de instalação, escolheremos DEMO para uma melhor primeiro contato, botão Salvar e continuar.
MySQL, nome drupal, usuário root, senha my-password
Configurações avançadas, servidor mysql.5.7, porta 3306, botão Salvar e continuar.
```

## elasticsearch
```
__Elasticsearch
dentro da pasta containers
cd /home/usuario/Desktop/containers
mkdir elasticsearch
cd elasticsearch
touch elasticsearch-commands.txt
touch commands.txt

tentar instalar sem java da próxima vez
docker pull openjdk

docker network create elasticsearch
docker network ls

docker run \
    --network elasticsearch \
    --name elasticsearch \
    -e "discovery.type=single-node" \
    -p 9200:9200 \
    -e xpack.ml.enabled=false \
    elasticsearch:7.7.1
    
docker ps

5c

docker run \
    -it \
    --network elasticsearch \
    --name curl \
     appropriate/curl sh

bb


__reiniciar os containers parados
docker start 5c


docker start -i bb

no container do curl
ping elasticsearch

pare o ping
Ctrl+C

curl -XPUT http://elasticsearch:9200/my-index
curl -XGET http://elasticsearch:9200/_cat/indices?v

curl -XPOST http://elasticsearch:9200/my-index/cities/1 \
  -H 'Content-Type: application/json' \
  -d '{"city":"New York"}'

curl -XPOST http://elasticsearch:9200/my-index/cities/2 \
  -H 'Content-Type: application/json' \
  -d '{"city":"Paris"}'

curl -XPOST http://elasticsearch:9200/my-index/cities/3 \
  -H 'Content-Type: application/json' \
  -d '{"city":"London"}'
  
curl -XGET http://elasticsearch:9200/my-index/_mapping?pretty
curl -XGET http://elasticsearch:9200/my-index/cities/1?pretty
curl -XGET http://elasticsearch:9200/my-index/cities/2?pretty

curl -XGET http://elasticsearch:9200/my-index/_search?q=city:new
```

## friendica
```
https://hub.docker.com/_/friendica

Baixar
docker pull friendica

_______________________________________________________________________________
__INSATALANDO COM MYSQL________________________________________________________
Instale o MySQL e o PHPmyAdmin, veja em no arquivo deste mesmo diretório mysql.txt

docker run \
  -d \
  -p 80:80 \
  --network mysql \
  --name friendica \
  friendica
  
  56 (deu errado, parece que é certificado)
  93
  
  mysql.5.7.wp
  
  o banco de dados friendica deve ser criado
  
Por algum motivo, depois de configurar, não permitiu executar o site sem o https, mesmo eliminando.


Friendica Settings

FRIENDICA_URL The Friendica URL.
FRIENDICA_TZ The default localization of the Friendica server.
FRIENDICA_LANG The default language of the Friendica server.
FRIENDICA_SITENAME The Sitename of the Friendica server.
FRIENDICA_NO_VALIDATION If set to true, the URL and E-Mail validation will be disabled.
FRIENDICA_DATA If set to true, the fileystem will be used instead of the DB backend.
FRIENDICA_DATA_DIR The data directory of the Friendica server (Default: /var/www/data).
Friendica Logging

FRIENDICA_DEBUGGING If set to true, the logging of Friendica is enabled.
FRIENDICA_LOGFILE (optional) The path to the logfile (Default: /var/www/friendica.log).
FRIENDICA_LOGLEVEL (optional) The loglevel to log (Default: notice).
Database (required at installation)

MYSQL_USER Username for the database user using mysql / mariadb.
MYSQL_PASSWORD Password for the database user using mysql / mariadb.
MYSQL_DATABASE Name of the database using mysql / mariadb.
MYSQL_HOST Hostname of the database server using mysql / mariadb.
MYSQL_PORT Port of the database server using mysql / mariadb (Default: 3306)
Lock Driver (Redis)

REDIS_HOST The hostname of the redis instance (in case of locking).
REDIS_PORT (optional) The port of the redis instance (in case of locking).
REDIS_PW (optional) The password for the redis instance (in case of locking).
REDIS_DB (optional) The database instance of the redis instance (in case of locking).

Administrator account
Because Friendica links the administrator account to a specific mail address, you have to set a valid address for MAILNAME.

Mail settings
The binary ssmtp is used for the mail() support of Friendica.

You have to set the --hostname/-h parameter correctly to use the right domainname for the mail() command.

You have to set a valid SMTP-MTA for the SMTP environment variable to enable mail support in Friendica. A valid SMTP-MTA would be, for example, mx.example.org.

The following environment variables are possible for the SMTP examples.

SMTP Address of the SMTP Mail-Gateway. (required)
SMTP_DOMAIN The sender domain. (required - e.g. friendica.local)
SMTP_FROM Sender user-part of the address. (Default: no-reply - e.g. no-reply@friendica.local)
SMTP_TLS Use TLS for connecting the SMTP Mail-Gateway. (Default: empty)
SMTP_STARTTLS Use STARTTLS for connecting the SMTP Mail-Gateway. (Default: empty)
SMTP_AUTH_USER Username for the SMTP Mail-Gateway. (Default: empty)
SMTP_AUTH_PASS Password for the SMTP Mail-Gateway. (Default: empty)
SMTP_AUTH_METHOD Authentication method for the SMTP Mail-Gateway. (Default: empty/plain text)
Database settings
You have to add the Friendica container to the same network as the running database container, e. g. --network some-network, and then use mysql as the database host on setup.

Persistent data
The Friendica installation and all data beyond what lives in the database (file uploads, etc) is stored in the unnamed docker volume volume /var/www/html. The docker daemon will store that data within the docker directory /var/lib/docker/volumes/.... That means your data is saved even if the container crashes, is stopped or deleted. To make your data persistent to upgrading and get access for backups is using named docker volume or mount a host folder. To achieve this you need one volume for your database container and Friendica.

Friendica:

/var/www/html/ folder where all Friendica data lives
$ docker run -d \
  -v friendica-vol-1:/var/www/html \
  --network some-network
  friendica
Database:

/var/lib/mysql MySQL / MariaDB Data
$ docker run -d \
  -v mysql-vol-1:/var/lib/mysql \
  --network some-network
  mariadb
Automatic installation
The Friendica image supports auto configuration via environment variables. You can preconfigure everything that is asked on the install page on first run. To enable the automatic installation, you have to the following environment variables:

FRIENDICA_ADMIN_MAIL E-Mail address of the administrator.
MYSQL_USER Username for the database user using mysql / mariadb.
MYSQL_PASSWORD Password for the database user using mysql / mariadb.
MYSQL_DATABASE Name of the database using mysql / mariadb.
MYSQL_HOST Hostname of the database server using mysql / mariadb.
```

## ghost
```
docker pull ghost

https://hub.docker.com/_/ghost

docker run \
  -d \
  -p 3001:2368 \
  -e url=http://localhost:3001 \
  --network mysql \
  --name ghost \
  ghost

Não funcionou

docker run -d --name some-ghost -e url=http://localhost:3001 -p 3001:2368 ghost
Parece que esse de cima funcionou mas não sei usar.

If all goes well, you'll be able to access your new site on http://localhost:3001 and http://localhost:3001/ghost to access Ghost Admin (or http://host-ip:3001 and http://host-ip:3001/ghost, respectively).
```

## grafana
```
docker pull grafana/grafana

https://hub.docker.com/r/grafana/grafana/

docker run -d --name=grafana -p 3000:3000 grafana/grafana
```

## grav
```
https://getgrav.org/
https://github.com/getgrav/docker-grav

docker build -t grav:latest

docker run -p 8000:80 grav:latest

docker run -d -p 8000:80 --restart always -v grav_data:/var/www/html grav:latest
```

## httpd
```
docker pull httpd

__APACHE
docker run -p 8080:80 -d httpd
a5

acesse localhost:8080 e veja o resultado.
docker stop a5
docker rm a5
```

## influx
```
cd ~
mkdir influxdb
cd influxdb
git clone https://github.com/EspenAlbert/tick-udemy
cd tick-udemy/00_config/influx-only/
docker-compose docker-compose.yml

Abrir outro terminal
docker ps
a349f8a03b7c        influxdb:1.5.1      "/entrypoint.sh infl…"   11 minutes ago      Up 10 minutes       0.0.0.0:8086->8086/tcp   influx-only_influxdb_1

Instale o Postman
https://github.com/renatomportugal/docker/blob/master/17.InstallCentOS.md#postman
```

## jekyll
```
docker pull jekyll/jekyll

https://hub.docker.com/r/jekyll/jekyll/
```

## jenkins
```
docker pull jenkins/jenkins

docker run \
  -d \
  -p 9090:8080 \
  --name jenkins \
  jenkins/jenkins

1bb

docker exec -it 1bb bash

cat /var/jenkins_home/secrets/initialAdminPassword

Clique em continue
Select Plugins to install, None, Install
Crie o usuario admin, senha admin, Nome completo Admin, Save and continue, Save and Finish, Start using Jenkins
```

## joomla
```
https://hub.docker.com/_/joomla

docker pull joomla

_______________________________________________________________________________
__INSATALANDO COM MYSQL________________________________________________________
Instale o MySQL e o PHPmyAdmin, veja em no arquivo deste mesmo diretório mysql.txt

Não é necessário criar a base de dados, mas é necessário excluir caso exista
docker run \
  -d \
  -p 80:80 \
  --network mysql \
  nextcloud
  
docker run \
  -d \
  -p 80:80 \
  -e JOOMLA_DB_HOST=mysql.5.7.wp \
  -e JOOMLA_DB_USER=root \
  -e JOOMLA_DB_PASSWORD=my-password \
  -e JOOMLA_DB_NAME=joomla \
  --name joomla22JUN \
  --network mysql \
  joomla
  
20
_______________________________________________________________________________
docker run --name some-joomla --link some-mysql:mysql -d joomla

-e JOOMLA_DB_HOST=... (defaults to the IP and port of the linked mysql container)
-e JOOMLA_DB_USER=... (defaults to "root")
-e JOOMLA_DB_PASSWORD=... (defaults to the value of the MYSQL_ROOT_PASSWORD environment variable from the linked mysql container)
-e JOOMLA_DB_NAME=... (defaults to "joomla")
If the JOOMLA_DB_NAME specified does not already exist on the given MySQL server, it will be created automatically upon startup of the joomla container, provided that the JOOMLA_DB_USER specified has the necessary permissions to create it.

If you'd like to be able to access the instance from the host without the container's IP, standard port mappings can be used:

docker run --name some-joomla --link some-mysql:mysql -p 8080:80 -d joomla

If you'd like to use an external database instead of a linked mysql container, specify the hostname and port with JOOMLA_DB_HOST along with the password in JOOMLA_DB_PASSWORD and the username in JOOMLA_DB_USER (if it is something other than root):
```

## known
```

https://hub.docker.com/_/known

docker run \
  -d \
  --link some-mysql:db \
  known
  
docker run \
  -d \
  -p 80:9000 \
  -e KNOWN_DB_HOST=mysql.5.7.wp:3306 \
  -e KNOWN_DB_USER=rooot \
  -e KNOWN_DB_PASSWORD=my-password \
  --link mysql.5.7.wp:known \
  --network mysql \
  --name some-known \
  known

não funcionou

Now you can get access to fpm running on port 9000 inside the container. If you want to access it from the Internets, we recommend using a reverse proxy in front. You can find more information on that on the docker-compose section.

The following environment variables are also honored for configuring your Known instance:

-e KNOWN_DB_HOST=... (defaults to the IP and port of the linked mysql container)
-e KNOWN_DB_USER=... (defaults to "root")
-e KNOWN_DB_PASSWORD=... (defaults to the value of the MYSQL_ROOT_PASSWORD environment variable from the linked mysql container)
-e KNOWN_DB_NAME=... (defaults to "known")
-e MAIL_HOST=...
-e MAIL_PORT=...
-e MAIL_SECURE=... ("starttls" for instance)
-e MAIL_USER=...
-e MAIL_PASS=...
If the KNOWN_DB_NAME specified does not already exist on the given MySQL server, it will be created automatically upon startup of the known container, provided that the KNOWN_DB_USER specified has the necessary permissions to create it.

If you'd like to use an external database instead of a linked mysql container, specify the hostname and port with KNOWN_DB_HOST along with the password in KNOWN_DB_PASSWORD and the username in KNOWN_DB_USER (if it is something other than root):

docker run --name some-known -e KNOWN_DB_HOST=10.1.2.3:3306 \
    -e KNOWN_DB_USER=... -e KNOWN_DB_PASSWORD=... -d known 
```

## lightstreamer
```
https://hub.docker.com/_/lightstreamer
docker pull lightstreamer

docker run --name ls-server -d -p 80:8080 lightstreamer

4da
```

## mariadb
```
https://hub.docker.com/_/mariadb

docker pull mariadb

https://hub.docker.com/_/mariadb

docker run --name some-mariadb -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mariadb:tag
docker run -it --network some-network --rm mariadb mysql -hsome-mariadb -uexample-user -p
docker run -it --rm mariadb mysql -hsome.mysql.host -usome-mysql-user -p
docker exec -it some-mariadb bash
docker logs some-mariadb

Using a custom MySQL configuration file
The startup configuration is specified in the file /etc/mysql/my.cnf, and that file in turn includes any files found in the /etc/mysql/conf.d directory that end with .cnf. Settings in files in this directory will augment and/or override settings in /etc/mysql/my.cnf. If you want to use a customized MySQL configuration, you can create your alternative configuration file in a directory on the host machine and then mount that directory location as /etc/mysql/conf.d inside the mariadb container.

If /my/custom/config-file.cnf is the path and name of your custom configuration file, you can start your mariadb container like this (note that only the directory path of the custom config file is used in this command):

$ docker run --name some-mariadb -v /my/custom:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mariadb:tag
This will start a new container some-mariadb where the MariaDB instance uses the combined startup settings from /etc/mysql/my.cnf and /etc/mysql/conf.d/config-file.cnf, with settings from the latter taking precedence.

Configuration without a cnf file
Many configuration options can be passed as flags to mysqld. This will give you the flexibility to customize the container without needing a cnf file. For example, if you want to change the default encoding and collation for all tables to use UTF-8 (utf8mb4) just run the following:

docker run --name some-mariadb -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mariadb:tag --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

docker run -it --rm mariadb:tag --verbose --help
```

## maven
```
docker pull maven
docker pull maven:3.6-alpine
```

## mediawiki
```
https://hub.docker.com/_/mediawiki

docker pull mediawiki

docker run --name some-mediawiki -d mediawiki
docker run --name some-mediawiki -p 8080:80 -d mediawiki
docker run --name some-mediawiki --link some-mysql:mysql -d mediawiki

docker run \
  -d \
  -p 80:80 \
  --network mysql \
  --name mediawiki \
  --link mysql.5.7.wp:3306 \
  mediawiki
  
  49
  
  LRTZ@4aCg8MNLVB
  Após concluir a instalação baixará automaticamente o LocalSettings.php
  Copie para a pasta que contém o index.php
  docker cp LocalSettings.php 49:/var/www/html/LocalSettings.php
  
  Abra outra janela e veja se realmente copiou
  docker exec -it 49 bash
  ls
  
  ## The protocol and server name to use in fully-qualified URLs
  $wgServer = "http://tecnocrata.ddns.net";
```

## mongodb
```
__Baixar
docker pull mongo
_______________________________________________________________
__Rodar
Sem rede
docker run -d -p 27017:27017 --name mongo mongo

Com rede

docker run \
    -d \
    -p 27017:27017 \
    --network mysql \
    --name mongo \
    mongo
d31

docker run -d -p 27017:27017 --network mysql --name mongo mongo

_______________________________________________________________
__Executar
docker exec -it d31 bash
mongo
show dbs
show collections
use admin
db.stats()
exit

netstat -nltp

docker cp d31:/etc/mongod.conf.orig mongod.conf.orig
_______________________________________________________________




_______________________________________________________________
docker run mongo

abra outro terminal
docker ps
veja qual o id
docker exec -it 64 bash
ls
mostrar os processos
ps -e

abrir outro processo
abra outra janela de terminal
docker exec -it 64 sh
ps -e

_______________________________________________________________
__Inspecionando o Docker
docker ps
docker inspect 64

na tela de terminal do 64, digite
ls /usr/local/bin

more /usr/local/bin/docker-entrypoint.sh
ls /usr/bin

_______________________________________________________________
__Criando um Mongo Database
docker ps
docker exec -it 64 mongo
db.version()
show dbs
use test
db
db.animals.insert({"animal": "cat"})
db.animals.insert({"animal": "dog"})
db.animals.insert({"animal": "monkey"})
db.animals.find()
exit
docker ps
docker stop 64
docker ps

vamos criar mais um container
docker run mongo

conectar com o mongo
docker exec -it 57 mongo
show dbs
exit
docker stop 57

agora vamos abrir o container 64
docker ps -a
docker start 64
docker ps
docker exec -it 64 mongo
show dbs
use test
db.animals.find()

_______________________________________________________________
__Persistent database
exit
docker ps -a
docker container prune -f
docker stop 64
docker start 64
docker ps
docker exec -it 64 bash
ls
ls data
ls data/db

exit
docker stop 64
docker ps
docker rm 64
docker ps -a

vamos criar uma pasta mongo dentro da pasta container
cd /home/usuario/Desktop/containers
mkdir mongo
cd mongo
mkdir db
ls db


vamos rodar uma instância do mongo em background (-d)
docker run -d -v $PWD/db:/data/db mongo
docker ps

agora verá  que a pasta contém os arquivos de banco de dados
ls db

docker ps
docker exec -it 57 bash
mongo
show dbs

vamos criar um db
use mydb
db.posts.insert({"post": "Hey there."})
db.posts.insert({"post": "How are you?"})
db.posts.insert({"post": "Doing well."})

vejamos o que tem
db.posts.find()
exit
exit
docker ps
ls db
ls -la db

cat db/index-9-2421615610053107849.wt

agora vamos remover o container e ver se os dados irão persistir
docker ps
docker stop 57
docker rm 57
docker ps
docker ps -a

não há mais containers, iremos criar um
docker run -d -v $PWD/db:/data/db mongo

docker ps
docker exec -it 70 mongo

veja que mydb está no novo container
show dbs

use mydb
db.posts.find()
exit
docker ps
docker stop 70
docker rm 70
_______________________________________________________________
__Outro caso de persistência de dados, com mapeamento de porta externa.

create volume create mongo

docker run \
    -d \
    -p 27017:27017 \
    -v mongo:/data/db
    --network mysql \
    --name mongo \
    mongo

_______________________________________________________________
```

## mysql
```
docker pull mysql:5.7
docker pull phpmyadmin/phpmyadmin

__REDE
docker network ls
docker network create mysql
docker network inspect mysql

__Rodar o MySQL
docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=my-password \
    -e MYSQL_USER=root \
    -e MYSQL_PASSWORD=my-password \
    --network mysql \
    --name mysql.5.7 \
    --restart=always \
    mysql:5.7

416

__Rodar o PHPmyadmin
docker run \
    -d \
    -p 8080:80 \
    -e PMA_HOST=mysql.5.7 \
    --network mysql \
    --name phpmyadmin \
    --restart=always \
    phpmyadmin/phpmyadmin

c81

Para impedir o acesso externo deve alterar uma configuração, primeiro copie o arquivo para a máquina
docker cp 88:/etc/mysql/mysql.conf.d/mysqld.cnf mysqld.cnf

Comentar (colocando a cerquilha da frente)
#bind-address	= 127.0.0.1

Copiar de volta
docker cp mysqld.cnf 88:/etc/mysql/mysql.conf.d/mysqld.cnf

Reinicia o docker
docker restart 88

docker ps

acesse localhos:8080
usuário root senha my-password

_____________________________________________________________________________________
Por padáo a pasta que guarda os dados é /var/lib/mysql
docker exec -it <id> bash
cd /var/lib/mysql

Para fazer backup dos dados abra outra janela de terminal e poie a pasta
cd ~
mkdir volumemysql
cd volumemysql
docker cp <id>:/var/lib/mysql .

Para terstar se o backup deu certo crie outro container mysql e outro phpmyadmin (use as portas 3307 e 8081)
docker run \
    -d \
    -p 3307:3306 \
    -e MYSQL_ROOT_PASSWORD=my-password \
    -e MYSQL_USER=root \
    -e MYSQL_PASSWORD=my-password \
    --network mysql \
    --name mysql.5.7.3307 \
    mysql:5.7

226

docker run \
    -d \
    -p 8081:80 \
    -e PMA_HOST=mysql.5.7.3307 \
    --network mysql \
    --name phpmyadmin.8081 \
    phpmyadmin/phpmyadmin

28d

Acesse o localhost:8081 para ver se tudo está certo. Agora iremos passar o conteúdo da pasta volumemysql/mysql para a pasta de dados do mysql.5.7.3307.
cd ~/volumesql/mysql
docker cp . 226:/var/lib/mysql

Agora acesse a pasta mysql do 226 para mudar o dono:
docker exec 226 -t bash
cd /var/lib
ls -la
chown -R mysql mysql
cd mysql
ls -la

Reinicie o container do mysql
docker restart 226

Acesse o localhost:8081 e verá que está funcionando. Se não reiniciar dará um erro ao acessar as tabelas.
Foi utilizado o mysql da mesma imagem. Para outros parâmetros faça outro estudo.
_____________________________________________________________________________________
Persistência de dados (na pasta usuário)
Pare o docker e renomeie para criarmos outro com o mesmo nome.
Isso evita mudar o host do phpmyadmin.
docker ps
docker stop 416
docker rename 416 mysql.5.7.old

docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=my-password \
    -e MYSQL_USER=root \
    -e MYSQL_PASSWORD=my-password \
    -v ~/volumesdocker/mysql.5.7:/var/lib/mysql \
    --network mysql \
    --name mysql.5.7 \
    mysql:5.7

948
Tentei copiar a pasta, está travada.
docker stop 948
docker rm 948

Para manusear a pasta é necessário mudar o dono.
sudo chown -R root mysql.5.7

Fazendo um backup
mkdir mysql.5.7.bck
sudo cp -R mysql.5.7/ mysql.5.7.bck/

Veja que a pasta mysql.5.7 foi copiada para a pasta mysql.5.7.bck.Faremos...
cd mysql.5.7.bck
cd mysql.5.7
sudo cp -r * ..
cd ..
sudo rm -rf mysql.5.7

Agora iremos testar a criação de um container com volume externo de uma pasta existente.
docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=my-password \
    -e MYSQL_USER=root \
    -e MYSQL_PASSWORD=my-password \
    -v ~/volumemysql/mysql:/var/lib/mysql \
    --network mysql \
    --name mysql.5.7 \
    mysql:5.7
    
7a8

O Container foi criado com os dados que já estavam. Funcionou.
_____________________________________________________________________________________

vai aparecer uma mensagem de erro para especificar algumas variáveis.
MYSQL_ROOT_PASSWORD
MYSQL_ALLOW_EMPTY_PASSWORD
MYSQL_RANDOM_ROOT_PASSWORD


docker ps
docker --help
docker run -e MYSQL_ROOT_PASSWORD=my-password mysql

em outra janela de terminal
docker ps

verifique se a variável foi atribuída
docker exec 09 env
```

## nextcloud
```
Baixar
docker pull nextcloud

Para instalação do ambiente siga os passos em:
https://github.com/renatomportugal/docker/blob/master/02.Producao.md#servidor-padr%C3%A3o

__Configuração
Coloque IP fixo no computador para que consiga usar o collabora.
No meu pc é 192.168.1.115:8000
Escolha o nome de usuário e senha do administrador.
Em Armazenamento & banco de dados:
Pasta de dados: /var/www/html/data
Configurar o Banco de dados: MySQL/MariaDB
usuário de banco de dados: root
senha do banco de dados: my-password
nome do banco de dados: nextcloud
host do banco de dados: mysql.5.7
Desmarque a opção Instalar aplicativos recomendados
Aperte o botão Concluir configuração

_______________________________________________________________________________
APLICATIVOS (que estou usando)

__Collabora
https://github.com/renatomportugal/docker/blob/master/ImagensOficiais/collabora-code.txt

__Adicionar
Na engrenagem do lado direito, Aplicativos,
Menu do lado esquerdo (ícone de 3 barras horizontais sobrepostas), Aplicativos em destaque, 
ache o Collabora Online (3.7.3), botão Baixar e ativar.
Na engrenagem do lado direito, Configurações,
No lado esquedo, Collabora Online, 
Selecione "Usar seu próprio servidor"
Selecione "Desativar a verificação do certificado (inseguro)"
Colque o endereço do serviço do Collabora online (sem a barra no final)
http://192.168.1.115:9980
Botão Salvar... (se der erro adicione as portas 8000 e 9980 no firewall)
Para testar abra um arquivo docx.

__Music
Na engrenagem do lado direito, Aplicativos,
Buscar, music, 

__Draw.io

__Mind Map

__Quick Notes

__Duplicate Finder

__Quota Warning

__Metadata
_______________________________________________________________________________
APLICATIVOS REPROVADOS
_Bookmarks
Pontos bons: Etiquetas e Anotações. Preview para fotos e videos.
Pontos Ruins: Não há como exportar nem importar.


_______________________________________________________________________________
APLICATIVOS (para testar)
Accessibility
Activity
Appointments
Auditing / Logging
_Automated PDF conversion

Brute-force settings
Calendar
Comments
_Contacts
Contacts Interaction
_Cookbook
_Cospend
_Custom Properties
_Deck
_Deleted files
Federation
_File access control
File sharing
First run wizard
_Flow upload
Forms
_Group folders
Log Reader
_Mail
_Maps
Monitoring
Nextcloud announcements
_Notes
Notifications
Password policy
PDF viewer
Photos
Privacy
Recommendations
_Retantion
Right click
Share by mail
Support
_Talk
Tasks
Text
Theming
Update notification
Usage survey
_User retention
Versions
Video player
Workflow external scripts
Default encryption module
External storage support
LDAP user and group backend

__APARENCIA
Verificar como instalar a extensão Imagemagick do PHP, para transformar a imagem enviada em ícone.
Configurações, Menu do lado esquerdo, Personalização, 

Apague os arquivos do diretório:
docker exec -it id bash
cd /var/www/html/core/skeleton
rm -rf *

Configurar Usuários
__Adicionar grupo
Na conta de administrador, Menu esquerdo, Adicionar Grupo, "Biblioteca", confirmar com a senha.
Coloque o administrador no grupo criado, e selecione para ser administrador do novo grupo. Usuários, editar, adicione o grupo, Concluído.

Desabilitar App
Contacts Interaction
Federation
Share by mail
Comments
First Run Wizard

Baixar e Ativar
File access control

_______________________________________________________________________________
Limitar a quantidade de espaço para cada usuário em 100MB.
_______________________________________________________________________________
Com a conta de administrador, clique no seu usuário, lado superior direito, Usuários, 
Clique na engrenagem no lado inferior esquerdo, Quota padrão 1 B.
Selecione as opções:
Exibir o último login.

Para o administrador dê a cota ilimitada para postar os arquivos.

__Configurar Compartilhamento__________________________________________________
Com a conta de administrador, clique no seu usuário, lado superior direito, Configurações
No lado esquerdo, Admin, Compartilhamento
Permitir que aplicativos usem a API de Compartilhamento
Permitir o compartilhamento com grupos
Excluir grupos de compartilhamento
 Biblioteca

Compartilhe a pasta com o grupo Biblioteca, retirando todas as opções.

docker exec -it id bash
apt-get update
apt-get install nano
cd config
nano config.php

Adicione no fim do arquivo: 
'default_language' => 'pt_BR',
'force_language' => 'pt_BR',
'default_locale' => 'pt_BR',


__Criar Usuários
Na conta de administrador, Usuários. Menu esquerdo, Novo usuário: usuário: usuario01, Nome: Usuário 01, senha usuario01, grupo Biblioteca, Cota padrão, ok.









__Configuração via occ
su www-data -s /bin/bash
php occ config:list
php occ config:app:set theming slogan --value "Seu local de estudos"



__Testando... não faça ainda
Instalar o app face recognition
1. pdlib
Como verificar?

https://github.com/goodspb/pdlib

docker exec -it 125 bash
apt-get update
apt-get install libbz2-dev
apt-get update
apt-get install libx11-dev
apt-get update
apt-get install libopenblas-dev liblapack-dev

__Install Dlib as shared library
cd /
mkdir mylib
cd mylib
apt-get install git
git clone https://github.com/davisking/dlib.git
cd dlib/dlib
mkdir build
cd build
apt-get install cmake
cmake -DBUILD_SHARED_LIBS=ON ..
make
make install

__Installation
cd /
mkdir mylib2
cd mylib2
git clone https://github.com/goodspb/pdlib.git
cd pdlib
phpize

-irá retornar:
Configuring for:
PHP Api Version:         20190902
Zend Module Api No:      20190902
Zend Extension Api No:   320190902

./configure --enable-debug
make

-deu uns warnings, mas instalou.

make install

__Configure PHP installation
nano youpath/php.ini

apt-get install nano
nano /usr/local/etc/php/php.ini

__Append the content below into php.ini
[pdlib]
extension="pdlib.so"

2. bz2

__Test
apt-get install php-bz2

erro: Package 'php-bz2' has no installation candidate

apt-get install software-properties-common
apt-get update
apt-get upgrade
sudo apt-get install libbz2-dev

/usr/local/bin/php-config --with-bz2

add-apt-repository ppa:ondrej/php
add-apt-repository ppa:ondrej/apache2

apt install software-properties-common
add-apt-repository ppa:deadsnakes/ppa

apt install python3.7
python3.7 --version

php -v

apt-cache search php7.3-*

apt-get install php5-dev php-pear

php --ini | grep "Loaded Configuration File"

nano /usr/local/etc/php/php.ini

apt install php7.3-bz2

apt install apt-transport-https lsb-release
apt-get install wget
wget -O /etc/apt/trusted.gpg.d/php.gpg https://packages.sury.org/php/apt.gpg
sh -c 'echo "deb https://packages.sury.org/php/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/php.list' # Add Ondrej's repo to sources list. sudo apt update


make test

_______________________________________________________________________________
__CONFIGURAÇÕES________________________________________________________________
Autorizar domínio
docker ps -a
2d
docker stop 2d
Abra outra tela de terminal com Ctrl+Shift+T
Como o container não tem editor de texto, copiaremos o arquivo para fora, editaremos, depois devolveremos.
copie a pasta do container para o seu diretório atual
pode usar os dois primeiros dígitos
docker cp 2d:/var/www/html/config/config.php config.php

verifique se realmente copiou
ls -la

Abra-os para edição, destrave o terminal com &
gedit config.php &

Faça as alterações conforme a seguir, não substitua tudo, apenas complete com sua necessidade.
No meu caso adicionei o site "seuSite.com"

'trusted_domains' =>
array(
    0 => 'localhost',
    1 => 'seuSite.com'
)

Salve com Ctrl+S

Cole novamente dentro do container (não se esqueça que cd é o id do container) .
docker cp config.php 2d:/var/www/html/config/config.php

Mude o dono imediatamente (se mexer no site dará erro)
docker start 2d
docker exec -it 2d bash
cd config
ls -la
chown www-data config.php
ls -la

Reinicie o container...
docker restart 2d
________________________________________________
FAZENDO O BAKUP DOS DADOS (MySQL)
Faça o backup do banco de dados com o phpMyAdmin

docker stop 48
cd ~
mkdir ncMySQL
cd ncMySQL
docker stop 48
docker cp 48:/var/www/html ./

Compactar o arquivo
tar -czvf nc_html_mysql_20jun20.tar.gz html

Também deixe junto o arquivo sql
_______________________________________________________________________________
__A aplicação parou de funcionar.
Fiz o backup da pasta html e do banco.
Reinstalei normalmente, funcionou.
docker stop c5
docker cp ./ c5:/var/www/

Abra outra tela de terminal
docker start c5
docker exec -it c5 bash
cd /var/www/
mudar o dono de tudo
chown -R www-data html
ls -la
docker stop c5
Agora restore a base de dados no phpMyAdmin
docker start c5

Deu erro, agora estou tentando retornar a base de dados.
Também deu erro. Vou fazer novamente, mas copiando apenas a pasta de dados.
Deu erro de novo, e parece que não vai funcionar.

__Vou tentar restaurar a base de dados primeiro
Mexendo só na base não deu erro, mas precisa ver a restauração dos APPS.
Agora restaurar só a pasta data.
docker exec -it 2d bash
cd data
ls
docker stop 2d

Abra outro terminal
docker cp 2d:/var/www/html/data ./

agora acesse a pasta data do bacjup a restaurar
cd data

vamos copiar a pasta data para o servidor
docker cp ./ 2d:/var/www/html/data

Abra outro terminal
docker exec -it 2d bash
chown -R www-data data

Agora a pasta dos aplicativos
mv appdata_ochsuvqvqiwm/ appdata_oclhjk56fpwo/

Ainda está funcionando, precisamos fazer mais testes.
_______________________________________________________________________________
__TABELAS MYSQL________________________________________________________________
oc_group_user - usuários admin

_______________________________________________________________________________
__RECUPERAR ADMIN______________________________________________________________
Logue com uma conta normar, substitua em oc_group_user,  altere a senha do admin padrão,
volte a tabela com o nome padrão e teste

_______________________________________________________________________________
__INSTALANDO COM SQLite________________________________________________________
Rodar
docker run -d -p 80:80 nextcloud

Acesse localhost:80
Escolha um usuário e senha para a conta de administrador
Selecione o banco de dados, no nosso caso deixaremos o SQLite, com a pasta padrão /var/www/html/data.
Deixe selecionado para instalar os aplicativos recomendados (Calendar, Contacts, Talk, Mail e Collaborative Editing)
Clique no botão concluir

_______________________________________________________________________________
FAZENDO O BAKUP DOS DADOS (SQLite)
docker stop 59 1e
cd ~
mkdir nc20
cd nc20
docker cp 59:/var/www/html ./

Compactar o arquivo
tar -czvf html20jun20.tar.gz html

RECUPERAR OS DADOS
cd nc20

Deve ter arquivos compactados
ls

Descompactar o arquivo
Estando no diretório onde a pasta esteja
tar -vzxf html20jun20.tar.gz

Substitua a pasta html no destino (o sistema não perguntará, fará direto)
docker cp ./ 1e:/var/www/

Abra outra tela de terminal
docker exec -it 1e bash
cd /var/www/
mudar o dono de tudo
chown -R www-data html

_______________________________________________________________________________

https://hub.docker.com/_/nextcloud
A safe home for all your data. Access & share your files, calendars, contacts, mail & more from any device, on your terms.

Using an external database
By default this container uses SQLite for data storage, but the Nextcloud setup wizard (appears on first run) allows connecting to an existing MySQL/MariaDB or PostgreSQL database. You can also link a database container, e. g. --link my-mysql:mysql, and then use mysql as the database host on setup. More info is in the docker-compose section.

Persistent data
The Nextcloud installation and all data beyond what lives in the database (file uploads, etc) is stored in the unnamed docker volume volume /var/www/html. The docker daemon will store that data within the docker directory /var/lib/docker/volumes/.... That means your data is saved even if the container crashes, is stopped or deleted.

A named Docker volume or a mounted host directory should be used for upgrades and backups. To achieve this you need one volume for your database container and one for Nextcloud.
/var/www/html/ folder where all Nextcloud data lives

$ docker run -d \
-v nextcloud:/var/www/html \
nextcloud

Database:

/var/lib/mysql MySQL / MariaDB Data

/var/lib/postgresql/data PostgreSQL Data

$ docker run -d \
-v db:/var/lib/mysql \
mariadb

If you want to get fine grained access to your individual files, you can mount additional volumes for data, config, your theme and custom apps. The data, config are stored in respective subfolders inside /var/www/html/. The apps are split into core apps (which are shipped with Nextcloud and you don't need to take care of) and a custom_apps folder. If you use a custom theme it would go into the themes subfolder.

Overview of the folders that can be mounted as volumes:

/var/www/html Main folder, needed for updating
/var/www/html/custom_apps installed / modified apps
/var/www/html/config local configuration
/var/www/html/data the actual data of your Nextcloud
/var/www/html/themes/<YOU_CUSTOM_THEME> theming/branding
If you want to use named volumes for all of these it would look like this

$ docker run -d \
    -v nextcloud:/var/www/html \
    -v apps:/var/www/html/custom_apps \
    -v config:/var/www/html/config \
    -v data:/var/www/html/data \
    -v theme:/var/www/html/themes/<YOUR_CUSTOM_THEME> \
    nextcloud
```

## nginx
```
docker pull nginx
docker images
docker run nginx

vamos rodar na porta 80
docker run -p 8080:80 nginx

testar para iniciar com o Sistema
docker run -d -p 8080:80 --restart=always nginx

abra o browser e acesse o nginx digitando localhost:8080

Criando uma página e mostrando no container
abra o terminal, vá até o Desktop
cd 'Área de Trabalho'/
mkdir containers
cd containers
mkdir nginx

vamos baixar o visual studio code, abra o browser, baixe e instale...
abra o visual studio code, crie um index.html dentro da pasta nginx.
liste os arquivos da pasta (ls) e veja se realmente foi criado, e mostre o conteúdo
more index.html

agora vamos expor a porta 80 do docker através da porta 8081 do localhost, com volume definido
o caminho deve ser absoluto, e para descobrir digite pwd na pasta
/home/user/Área de Trabalho/containers/nginx

caminho absoluto
docker run -p 8081:80 -v /home/user/'Área de Trabalho'/containers/nginx:/usr/share/nginx/html nginx

tente
docker run -p 8081:80 -v $PWD:/usr/share/nginx/html nginx

caminho relativo (deu erro por causa do nome composto Area de Trabalho)
Basta renomear a pasta para Desktop
mv 'Área de Trabalho' Desktop

tente novamente
docker run -p 8081:80 -v $PWD:/usr/share/nginx/html nginx

veja se o container está rodando
docker ps

acesse o nginx digitando localhost:8081 e verá a página que foi criada

__________________________
https://www.katacoda.com/courses/docker/create-nginx-static-web-server

nginx
 |_Dockerfile
 |_index.html
 
__Dockerfile
FROM nginx:alpine
COPY . /usr/share/nginx/html

__index.html
<h1>Hello World</h1>

Abra o terminal na pasta nginx e digite o comando
docker build -t webserver-image:v1 .
docker images
docker run -d -p 80:80 webserver-image:v1
curl docker

https://www.katacoda.com/courses/docker/2

nginx
 |_Dockerfile
 |_index.html

__Dockerfile
FROM nginx:1.11-alpine
COPY index.html /usr/share/nginx/html/index.html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

__index.html
<h1>Hello World</h1>

Abra o terminal na pasta nginx e digite o comando
docker build -t my-nginx-image:latest .
docker images
docker run -d -p 80:80 my-nginx-image:latest
curl -i http://docker
docker ps
```

## nodejs
```
no terminal podemos ver que não há o nodejs instalado (no Sistema Operacional)
node -v
bem como o npm
npm -v

Então instalaremos via docker
verificar se tem containers rodando
docker ps

verificar a lista de containers
docker ps -a

apagaremos a lista (-f forçar)
docker container prune -f

verificar
docker ps -a

baixar a imagem
docker pull node

verificar se a imagem existe
docker images

docker run node
docker ps
docker ps -a

rodar com interação
docker run -it node

chame ajuda
.help

console.log('Hello from Node.js container')
10+5
const a = 10;
a

sair
.exit

dentro do Visual Studio code, na pasta containers, crie uma pasta node e inclua o script hello.js
console.log('This application was executed by Node.js container');

Salve com control+s
acesse a pasta node
cd /home/usuario/Desktop/containers/node
ls

vamos criar a imagem do docker
docker run -v $PWD:/app -w /app node node hello.js
docker ps

MÓDULO EXPRESS
https://www.npmjs.com/package/express

Criar uma pasta dentro da pasta node, nomear como express.
criar index.js dentro da pasta express
copiar o texto para o arquivo

const express = require('express')
const app = express()
app.get('/', function (req, res) {
  res.send('This Express app was executed by Node.js container inside of the Docker')
})
app.listen(3000)

cd express
ls
docker run -v $PWD:/app -w /app node node index.js

dará um erro porque o módulo express não foi achado
para resolver deveremos instalar o modulo no docker
docker run -v $PWD:/app -w /app node npm install

deu outro erro, ENOENT: no such file or directory
o projeto precisa ser iniciado
docker run -v $PWD:/app -w /app node npm init

o comando anterior necessita de interação, vamos acrescentar o parâmetro -it
vamos novamente...
docker run -v $PWD:/app -w /app -it node npm init

selecione as opções padrões, ele mostrará para conferência, enter para yes
verifique se foi criado o arquivo
ls

vamos instalar o expres (atualmente express@4.17.1)
docker run -v $PWD:/app -w /app -it node npm install express

veja que foi criada a pasta node_modules
ls

docker run -v $PWD:/app -w /app -it node node index.js

verifica se está rodando
docker ps

acessando http://localhost:3000/ pelo browser dá o erro ERR_CONNECTION_REFUSED
docker stop 57

vamos expor a porta
docker run -v $PWD:/app -w /app -it -p 3000:3000 node node index.js

teste no browser e verá que está funcionando

__SIGINT
o código a seguir identifica o fim do processo.

const express = require('express')
const process = require('process')
const app = express()
process.on('SIGINT', () =>{
    console.log('Application is being interrupted...')
    process.exit(0)
})
process.on('SIGTERM', () =>{
    console.log('Application is being terminated...')
    process.exit(0)
})
app.get('/', function (req, res) {
  res.send('This Express app was executed by Node.js container inside of the Docker')
})
app.listen(3000)

docker ps
docker stop e4
docker ps

docker run -v $PWD:/app -w /app -it -p 3000:3000 node node index.js

abra outro terminal...
docker ps
docker stop 98

volte no terminal que rodou a aplicação e veja a mensagem
Application is being terminated...

__Aplicação Node
cd /home/usuario/Desktop/containers/node/express
cd ..
ls
mkdir files
cd files
touch index.js

const fs = require('fs')
const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
})
readline.question('Enter filename: ', filename => {
    readline.question('Enter some text: ', text=> {
        fs.writeFile(`${filename}.txt`, text, err =>{
            if (err) throw err
            console.log('File was created')
            readline.close()
        })
    })
})

docker run -v $PWD:/app -w /app -it node node index.js

digita
myfile1
Hi

ls
more file1

mostrará o conteúdo (Hi)

_____
__Testar
https://www.katacoda.com/courses/docker/3
```

## nuxeo
```
https://hub.docker.com/_/nuxeo

docker pull nuxeo

docker run --name nuxeo -p 80:8080 -d nuxeo

44c


docker run \
  --name mynuxeo \
  --rm \
  -ti \
  -p 80:8080 \
  -e NUXEO_PACKAGES="\
    nuxeo-web-ui \
    nuxeo-dam \
    nuxeo-drive \
    nuxeo-showcase-content \
    nuxeo-template-rendering \
    nuxeo-template-rendering-samples \
    nuxeo-spreadsheet" \
  nuxeo

cdc

Senha não abre
Administrator/Administrator







docker run --name mynuxeo --rm -ti -p 8080:8080 -e NUXEO_PACKAGES="nuxeo-web-ui nuxeo-dam nuxeo-drive nuxeo-showcase-content nuxeo-template-rendering nuxeo-template-rendering-samples nuxeo-spreadsheet" nuxeo

NUXEO_DB_TYPE
This defines the database type to use. By default it sets an H2 embedded database that is suitable for test purpose only. When specifying a DB type, other variable mays help :

NUXEO_DB_HOST : If NUXEO_DB_TYPE is defined, this variable is mandatory and has to point to the DB server host.
NUXEO_DB_NAME : name of the database to use (nuxeo by default)
NUXEO_DB_USER : user to connect to the database (nuxeo by default)
NUXEO_DB_PASSWORD : the password to connect to the database (nuxeo by default)

NUXEO_ES_HOSTS
This variables allows to setup an external Elasticsearch cluster. Use a comma separated list of Elasticsearch hosts with the 9300 port. Additional environment vars may be setup like :

NUXEO_ES_CLUSTER_NAME : name of the Elasticsearch cluster to join
NUXEO_ES_INDEX_NAME: name of the index (nuxeo by default)
NUXEO_ES_REPLICAS : number or replicas (1 by default). If not 0, it means that your ES cluster must have enough node to fullfill the replicas setup.
NUXEO_ES_SHARDS : number or shards (5 by default).

NUXEO_ES_HOSTS=es1:9300,es2:9300
NUXEO_ES_CLUSTER_NAME=dockerCluster
NUXEO_ES_INDEX_NAME=nuxeo1
NUXEO_ES_REPLICAS=0
NUXEO_ES_SHARDS=5
```

## odoo
```
https://hub.docker.com/_/odoo

docker pull odoo

testar com postgres:10
```

## onlyoffice community Server 5.5
```
Fazer em 12JUL20
cd oo
sudo bash opensource-install.sh -ims false


https://sempreupdate.com.br/como-instalar-o-onlyoffice-no-linux-ubuntu-debian-opensuse/

mkdir oo
cd oo
wget http://download.onlyoffice.com/install/opensource-install.sh

Tente instalar com o Docker porque se a arquitetura de tudo não for x64 não funciona. No Docker funciona bem.
sudo bash opensource-install.sh -ims false

Se der o erro: Minimal requirements are not met: need at least 5500 MB of RAM
Para instalar com memória menor que 5500 MB mude no arquivo abaixo e rode normalmente

nano opensource-install.sh
Dentro do if [ "$DOCKER" == "true" ]; then
comente as linhas colocando o # na frente
sudo curl -s -O http://download.onlyoffice.com/install/install.sh
sudo rm install.sh
Vai ficar assim:
#sudo curl -s -O http://download.onlyoffice.com/install/install.sh
#sudo rm install.sh
Ctrl+O, Enter e Ctrl+X pra sair

Baixar o arquivo install manualmente
wget  http://download.onlyoffice.com/install/install.sh

nano install.sh
Mude MEMORY_REQUIREMENTS=5500;
Para MEMORY_REQUIREMENTS=2800;
Caso tenha 3GB de RAM, ajuste de acordo com sua necessidade.
Ctrl+O, Enter e Ctrl+X pra sair

Para rodar local
bash opensource-install.sh -ims false

Ele baixará as imagens e ajustará conforme a necessidade.

Ao terminar acesse localhost

__Configuração
```

## onlyoffice documentserver
```
https://hub.docker.com/r/onlyoffice/documentserver
docker pull onlyoffice/documentserver:5.5

docker run \
  -i \
  -t \
  -d \
  -p 8005:80 \
  --network mysql \
  onlyoffice/documentserver
  
Acesse http://localhost:8005/ e verá que aparecerá uma página, significa que está rodando.

https://www.youtube.com/watch?v=GAd-x_sP6ng
Configurações, Admin, Adicional, Endereço do servidor: http://localhost:8005, 
Chave Secreta: mysecret
Endereço do servidor para pedidos internos: http://localhost:8005
Antes de clicar no botão Salvar, abra o terminal

cd ~
docker cp onlyoffice:/etc/onlyofficec/documentserver/default.json default.json

Ache servives.CoAuthoring.secret.inbox e altere para mysecret o browser, inbox, outbox e session.
Habilitar a validação de token
token.enable.browser, request (inbox e outbox) = true
Salve o arquivo e cole de volta
docker cp default.json onlyoffice:/etc/onlyofficec/documentserver/default.json

Precisamos mudar o dono do arquivo
docker exec -it onlyoffice bash
cd /etc/onlyofficec/documentserver
ls -la
chown root:root default.json
supervisorctl restart all




docker ps

nano



Após instalar ele não está reconhecendo o adicionar arquivos.
Abra o prompt de comando do container do owncloud.
docker exec -it bc2 bash
pwd
/var/www/html
cd apps/
```

## openjdk
```
docker pull openjdk
docker pull openjdk:8-alpine
```

## openproject-community
```
admin/admin
https://hub.docker.com/r/openproject/community

https://docs.openproject.org/installation-and-operations/installation/docker/

docker pull openproject/community
mkdir openproject
cd openproject
git clone --depth=1 --branch=stable/10 https://github.com/opf/openproject

Instalação do Docker Compose: https://github.com/renatomportugal/docker/blob/master/08.Docker-compose.md

docker-compose up -d
docker-compose down

_________________________________________________________________________________________________________

docker run -it -p 8080:80 -e SECRET_KEY_BASE=secret openproject/community:10

docker run -d -p 8080:80 -e SECRET_KEY_BASE=secret openproject/community:10

sudo mkdir -p /var/lib/openproject/{pgdata,assets}
docker run -d -p 8080:80 --name openproject -e SECRET_KEY_BASE=secret \
  -v /var/lib/openproject/pgdata:/var/openproject/pgdata \
  -v /var/lib/openproject/assets:/var/openproject/assets \
  openproject/community:10
  
  Note: Make sure to replace secret with a random string. One way to generate one is to run head /dev/urandom | tr -dc A-Za-z0-9 | head -c 32 ; echo '' if you are on Linux.
  Since we named the container, you can now stop it by running:
  docker stop openproject
  docker start openproject
  ```

## ownCloud
```
Baixar
docker pull owncloud

Primeiro iremos instalar o onlyoffice, vá no arquivo https://github.com/renatomportugal/docker/blob/master/ImagensOficiais/onlyoffice-comunityserver.txt

docker network ls

veja que já existe a rede onlyoffice, e que há 3 containers.
docker inspect onlyoffice

docker run \
  -d \
  -p 8000:80 \
  --network onlyoffice \
  --name owncloud \
  owncloud

__Configuração
Escolha um usuário e senha
Pasta de dados padrão: /var/www/html/data

Escolha MySQL...
Para saber o IP do host
docker network inspect onlyoffice
veja o nome "onlyoffice-mysql-server"

usuario root
senha my-secret-pw
nome do banco owncloud
host 172.18.0.2

docker cp 8c57e21fdfae:/var/www/html/config/config.php config.php

incluir:
'onlyoffice' => array (
    'verify_peer_off' => true
)

docker cp config.php 8c57e21fdfae:/var/www/html/config/config.php
docker exec -it 8c57e21fdfae bash
cd config
ls -la
chown www-data config.php


I did not find the real solution, so I decided to simply remove the healthcheck in the NextCloud app.
So I commented in

docker exec -it 8c57e21fdfae bash
cd /var/www/html/apps/onlyoffice/controller/
ls -la

docker cp 8c57e21fdfae:/var/www/html/apps/onlyoffice/controller/settingscontroller.php settingscontroller.php

these lines

249 /*$healthcheckResponse = $documentService->HealthcheckReques 249 t();
250 * &if (!$healthcheckResponse) {
251 *
252 }*/

Additionally I had to adjust
/etc/nginx/includes/onlyoffice-communityserver-proxy-to-documentserver.conf
and replaced
proxy_pass {{DOCUMENT_SERVER_HOST_ADDR}};
by
proxy_pass http://127.0.0.1:80;

Now it seems to work. At least until the next update of the OnlyOffice-App in Nextcloud.

docker cp settingscontroller.php 8c57e21fdfae:/var/www/html/apps/onlyoffice/controller/settingscontroller.php
chown www-data settingscontroller.php
ls -la





docker exec -it 6583aace0bf9 bash
nano /etc/onlyoffice/documentserver/default.json
de "rejectUnauthorized": true para "rejectUnauthorized": false
Ctrl + O, Enter, Ctrl + X
supervisorctl restart all

_______________________________________________________________________________
__Nova instalação MySQL________________________________________________________
Depois de rodar o container, acesse o local onde instalou, aqui estou no localhost.
Criar conta de administrador, coloque usuário e senha.
Expanda a opção Armazenamento & Banco de dados, selecione a pasta de dados padrão (/var/www/html/data)
Selecione MySQL/MariaDB, coloque as informações do banco:
usuário, senha, nome do banco e host(nome ou IP). O banco não precisa existir, mas se tiver deverá apagar.
A aplicação cria a tabela a partir do nome colocado.
Clique em concluir a instalação.
Após a instalação coloque o usuário e senha de administrador.

__Configurando_________________________________________________________________
Defina o ESCOPO do seu projeto, ou as REGRAS de utilização. Explicarei as minhas e poderá copiar.
_______________________________________________________________________________
R01 (Regra 01) - Nossos usuários terão de fazer diagramas.
_______________________________________________________________________________
Vamos instalar a aplicação Drawio.
Vá em menu Arquivos, Market, procure por Drawio, abra e instale.
DESENVOLVEDOR Tom Needham, VERSÃO 0.0.8, RELEASE 13 de Ago de 2018, DATE General Public License, LICENÇA Version 2
Para testar vá nos arquivos e adicione um diagrama. Se abrir a tela de edição é porque funcionou.
_______________________________________________________________________________
R02 - Compartilhamento de pasta do administrador para grupo específico
_______________________________________________________________________________
__Criar um grupo_______________________________________________________________
Com a conta de administrador, clique no seu usuário, lado superior direito, Usuários, 
Adicione um grupo chamado Biblioteca
_______________________________________________________________________________
__Configurar Compartilhamento__________________________________________________
Com a conta de administrador, clique no seu usuário, lado superior direito, Configurações
No lado esquerdo, Admin, Compartilhamento
Permitir que aplicativos usem a API de Compartilhamento
Permitir o compartilhamento com grupos
Excluir grupos de compartilhamento
 biblioteca
Restringir a enumeração aos membros do grupo
Campo extra para exibir em resultados de preenchimento automático
 ID do usuário
Group Sharing Blacklist
 Exclude groups from receiving shares
  biblioteca
_______________________________________________________________________________
R03 - Limitar a quantidade de espaço para cada usuário em 1GB.
_______________________________________________________________________________
Com a conta de administrador, clique no seu usuário, lado superior direito, Usuários, 
Clique na engrenagem no lado inferior esquerdo, Quota padrão 1BG.
Selecione as opções:
Mostrar o último acesso
Mostrar o endereço de email

Para o administrador dê a cota de 10GB para postar os arquivos.
_______________________________________________________________________________
R04 - Leitura de PDF
_______________________________________________________________________________
Ao clicar num PDF o sistema baixa, mas queremos ler. Vamos instalar o leitor.
Vá em Menu, Market, ache o PDF Viewer.
DESENVOLVEDOR: ownCloud, 	VERSÃO: 0.11.0, RELEASE DATE: 11 de Abr de 2019,LICENÇA: GNU Affero General Public License
_______________________________________________________________________________
R05 - Gallery
_______________________________________________________________________________
Mostrar as imagens em formato de galeria.
Vá em Menu, Market, ache o Gallery
DESENVOLVEDOR: ownCloud, VERSÃO: 16.1.1, RELEASE DATE: 10 de Dez de 2018,LICENÇA: GNU Affero General Public License
_______________________________________________________________________________
R06 - Brute-Force Protection
_______________________________________________________________________________
Impedir que tentem logar com brute force.

Vá em Menu, Market, ache o Brute-Force Protection
DESENVOLVEDOR: Semih Serhat Karakaya, VERSÃO: 1.0.1,	RELEASE DATE: 6 de Dez de 2018,	LICENÇA: General Public License, Version 2
Funcionamento: depois de várias tentativas ele trava, e devemos esperar. Mesmo digitando a correta tem de esperar.
_______________________________________________________________________________
Update no Raspberry Pi
https://ssi.le-piolot.fr/updating-an-owncloud-instance-running-in-docker/
_______________________________________________________________________________
https://doc.owncloud.org/server/10.1/admin_manual/configuration/server/occ_command.html#command-line-upgrade
NÃO ATUALIZAR SENÃO PERDE OS PROGRAMAS

__DESATIVAR O AVISO SOBRE AS ATUALIZAÇÕES
Com a conta de administrador, clique no seu usuário, lado superior direito, Configurações
Lado esquerdo inferior, Admin, Geral
No item atualizador retire o grupo admin de Notificar membros dos seguintes grupos sobre atualizações disponíveis.
Lado esquerdo inferior, Admin, Configurações, Update notification, Desabilitar.

UPGRADE
docker exec -it 895 bash
su www-data -s /bin/sh

php occ maintenance:mode --on

php occ upgrade --help
php occ upgrade

Atualizando para a versaõ 10.1.1
Deu defeito...
php occ app:disable files_opds
php occ app:disable files_reader
php occ app:disable groupalert
php occ app:disable polls
php occ app:disable files_paperhive

php occ upgrade
Atualizando para a versão 10.3.2.2
php occ app:disable drawio
php occ app:disable wallpaper

Atualizando para a versão 10.4.1.3


php occ maintenance:mode --off

Para habilitar
php occ app:enable drawio (Não dá para habilitar porque funciona no máximo com a versão 10.1)
php occ app:enable wallpaper (Não dá para habilitar porque funciona no máximo com a versão 10.1)


php occ config:list brute_force_protection
php occ config:app:set brute_force_protection <Key> --value=<Value> --update-only
brute_force_protection_fail_tolerance, padrao 3
brute_force_protection_time_threshold, padrao 60
brute_force_protection_ban_period, padrao 300

Anti-Virus - https://marketplace.owncloud.com/apps/files_antivirus
php occ config:list files_antivirus
php occ config:app:set files_antivirus <Key> --value=<Value> --update-only
Key av_mode
Default 'executable'
Possible Values 'executable' 'daemon' 'socket'


_______________________________________________________________________________
https://software.opensuse.org/download/package?project=isv:ownCloud:desktop&package=owncloud-client
owncloud-client

Para CentOS 7.6 execute o seguinte como root:
cd /etc/yum.repos.d/
wget https://download.opensuse.org/repositories/isv:ownCloud:desktop/CentOS_7.6/isv:ownCloud:desktop.repo
yum install owncloud-client

_______________________________________________________________________________
__INSTALANDO COM MYSQL_________________________________________________________
Instale o MySQL e o PHPmyAdmin, veja em no arquivo deste mesmo diretório mysql.txt

Vamos rodar outra instância do nextcloud, na mesma rede do mysql
docker run \
  -d \
  -p 80:80 \
  --network mysql \
  owncloud

89

Não é necessário criar a base de dados, mas é necessário excluir caso exista

Crie um usuário e uma senha de administrador
Em Armazenamento & banco de dados...
Pasta de dados /var/www/html/data

Selecione o banco de dados: MySQL
usuario (do mySQL): root
senha (do mySQL): my-password
nome (do banco): owncloud
nome do host e porta: mysql.5.7.wp

Evite trabalhar com IP de host. Trabalhe sempre com o nome (neste caso é mysql.5.7.wp)
_______________________________________________________________________________
__CONFIGURAÇÕES________________________________________________________________
Autorizar domínio
docker ps -a
89
docker stop 89
Abra outra tela de terminal com Ctrl+Shift+T
Como o container não tem editor de texto, copiaremos o arquivo para fora, editaremos, depois devolveremos.
copie a pasta do container para o seu diretório atual
pode usar os dois primeiros dígitos
docker cp 89:/var/www/html/config/config.php config.php

verifique se realmente copiou
ls -la

Abra-os para edição, destrave o terminal com &
gedit config.php &

Faça as alterações conforme a seguir, não substitua tudo, apenas complete com sua necessidade.
No meu caso adicionei o site "seuSite.com"

'trusted_domains' =>
array(
    0 => 'localhost',
    1 => 'seuSite.com'
)

Salve com Ctrl+S

Cole novamente dentro do container (não se esqueça que cd é o id do container) .
docker cp config.php 89:/var/www/html/config/config.php

Mude o dono imediatamente (se mexer no site dará erro)
docker start 89
docker exec -it 89 bash
cd config
ls -la
chown www-data config.php
ls -la

Reinicie o container...
docker restart 89

_______________________________________________________________________________
Instalar (SQLite)
Crie uma conta de administrador
Apenas para teste escolha a base SQLite, na pasta padrão /var/www/html/data
Clique em concluir
___________________________________________________________________________________
APLICAÇÕES
Custom Groups
ok Gallery
ok PDF Viewer
ok Drawio
Carnet ()
Contacts
Activity
Polls
Tasks
Wallpaper
OK Brute-Force Protection

___________________________________________________________________________________
https://hub.docker.com/_/owncloud
owncloud.org

docker run -d -p 80:80 owncloud:8.1

For a MySQL database you can link an database container, e.g. --link my-mysql:mysql, and then use mysql as the database host on setup.

ownCloud is a self-hosted file sync and share server.

ownCloud is a self-hosted file sync and share server. It provides access to your data through a web interface, sync clients or WebDAV while providing a platform to view, sync and share across devices easily—all under your control. ownCloud’s open architecture is extensible via a simple but powerful API for applications and plugins and it works with any storage.

Persistent data
All data beyond what lives in the database (file uploads, etc) is stored within the default volume /var/www/html. With this volume, ownCloud will only be updated when the file version.php is not present.

-v /<mydatalocation>:/var/www/html
For fine grained data persistence, you can use 3 volumes, as shown below.

-v /<mydatalocation>/apps:/var/www/html/apps installed / modified apps
-v /<mydatalocation>/config:/var/www/html/config local configuration
-v /<mydatalocation>/data:/var/www/html/data the actual data of your ownCloud

docker exec -u www-data some-owncloud php occ status

Example stack.yml for owncloud:

# ownCloud with MariaDB/MySQL
#
# Access via "http://localhost:8080" (or "http://$(docker-machine ip):8080" if using docker-machine)
#
# During initial ownCloud setup, select "Storage & database" --> "Configure the database" --> "MySQL/MariaDB"
# Database user: root
# Database password: example
# Database name: pick any name
# Database host: replace "localhost" with "mysql"

version: '3.1'

services:

  owncloud:
    image: owncloud
    restart: always
    ports:
      - 8080:80

  mysql:
    image: mariadb
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: example

Run docker stack deploy -c stack.yml owncloud (or docker-compose -f stack.yml up), wait for it to initialize completely, and visit http://swarm-ip:8080/, http://localhost:8080/, or http://host-ip:8080 (as appropriate).
```

## php
```
docker pull php:7.1.26-apache

Instale o MySQL e o PHPmyAdmin, veja no arquivo https://github.com/renatomportugal/docker/blob/master/02.Producao.md
mkdir app01
cd app01

docker run \
  -d \
  -p 80:80 \
  -v "$PWD":/var/www/html \
  --name app01 \
  --network bridge.66.6 \
  --ip 66.6.0.25 \
  --restart=always \
php:7.1.26-apache

docker run \
  -d \
  -p 80:80 \
  -v "$PWD":/var/www/html \
  --name app01 \
php:7.1.26-apache

docker run \
  -d \
  -p 80:80 \
  -v ./app01:/var/www/html \
  --name app01 \
php:7.1.26-apache

crie um arquivo de index:
code index.php
Insira no arquivo:
<?php phpinfo(); ?>

Pronto, acesse o localhost, e todas as mudanças automaticamente refletirão.
```

## phpmyadmin
```
__PHPMYDMIN
docker pull phpmyadmin/phpmyadmin
docker images
docker run phpmyadmin/phpmyadmin

abra outra janela de terminal
docker ps

docker stop 05
docker ps

volte na janela do phpmyadmin, que deve estar parado
docker run -p 8080:80 phpmyadmin/phpmyadmin


rode o mysql
docker run -e MYSQL_ROOT_PASSWORD=my-password mysql
docker ps
docker exec -it 2e sh
hostname -i
exit

copia o endereço 172.17.0.2
docker run -p 8080:80 -e PMA_HOST=172.17.0.2 phpmyadmin/phpmyadmin

docker exec 3e env

criar pasta mysql dentro da pasta containers
mkdir mysql
cd mysql
touch commands.txt
ls

coloque em commands.txt
docker stop 2e
docker stop 3e
docker container prune -f
docker ps
docker ps -a
```

## plone
```
https://hub.docker.com/_/plone
docker pull plone

docker run -p 80:8080 plone

Funcionou
```

## portainer
```
docker pull portainer/portainer

https://hub.docker.com/r/portainer/portainer

docker volume create portainer_data
docker run -d -p 8000:8000 -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer

curl -L https://downloads.portainer.io/portainer-agent-stack.yml -o portainer-agent-stack.yml
docker stack deploy --compose-file=portainer-agent-stack.yml portainer

docker service create --name portainer_agent --network portainer_agent_network --publish mode=host,target=9001,published=9001 -e AGENT_CLUSTER_ADDR=tasks.portainer_agent --mode global --mount type=bind,src=//var/run/docker.sock,dst=/var/run/docker.sock --mount type=bind,src=//var/lib/docker/volumes,dst=/var/lib/docker/volumes –-mount type=bind,src=/,dst=/host portainer/agent

No swarm
docker run -d -p 9001:9001 --name portainer_agent --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v /var/lib/docker/volumes:/var/lib/docker/volumes portainer/agent
```

## postgres
```
https://hub.docker.com/_/postgres

docker pull postgres:10

Instale o MySQL e o PHPmyAdmin, veja em no arquivo deste mesmo diretório mysql.txt

docker run \
  -p 5432:5432 \
  -d \
  -e "POSTGRES_PASSWORD=12345678" \
  -e "POSTGRES_USER=docker_user" \
  -e "POSTGRES_DB=docker_db" \
  --name postgres \
  --network mysql \
  postgres:10

ab9

docker exec -it ab9 bash
psql -d docker_db -U docker_user
```

## prometheus
```
docker pull prom/prometheus

https://hub.docker.com/r/prom/prometheus/
```

## pythom
```
docker pull python

rodar o python
docker run python

ver quais containers estão rodando
docker ps

ver quais containers estão ativos
docker ps -a

vamos rodar com o terminal interativo
docker run -it python

help do python
help()

mostra os módulos do pythom
modules

mais informações do zlib
zlib

para sair envie o comando q
q

para sair do python
exit()

dentro do Visual Studio code, na pasta containers, crie uma pasta python e inclua o script hello-world.py
print ('hello from the Python container')

Salve com control+s
acesse a pasta python
cd /home/usuario/Desktop/containers/python
ls

vamos criar a imagem do docker
docker run -it -v $PWD:/app python python3 /app/hello-world.py

vamos informar o working directory
docker run -it -v $PWD:/app -w /app python python3 hello-world.py

Crindo a aplicação calendário
crie o script calendar-app.py e insira:
import calendar
print ('Welcome to the Calendar application!')
year = int(input('Please enter any year:'))
month = int(input('Please enter any month number:'))
print(calendar.month(year, month))
print('Have a nice day!')

vamos criar o container
docker run -it -v $PWD:/app -w /app python python3 calendar-app.py
```

## rancher-nginx
```
https://hub.docker.com/r/rancher/nginx

docker pull rancher/nginx:1.17.4-alpine

docker run \
    -d \
    -p 8081:8080 \
    --network mysql \
    --name rancherAlpine \
    rancher/nginx:1.17.4-alpine

Não funcionou
```

## rancher-rancher
```
https://rancher.com/docs/rancher/v2.x/en/installation/other-installation-methods/single-node-docker/

1.Baixar
docker pull rancher/rancher

2.Criar Volume
Verifique se náo existe o nome
docker volume ls
Crie um volume
docker volume create rancher
Verifique se foi criado
docker volume ls

3.Ligar na rede desejada
Verifique quais redes existem
docker network ls
Verifique quais containers tëm na rede
Para criar a rede vá em:
https://github.com/renatomportugal/docker/blob/master/05.RedeDeContainers.md#organizando-os-containers
docker network inspect 
Utilize a rede escolhida

docker run \
  -d \
  -p 80:80 \
  -p 443:443 \
  -v rancher:/var/lib/rancher \
  --name rancher \
  --network bridge.66.6 \
  --ip 66.6.0.180 \
  --restart=always \
  rancher/rancher:latest

__ADICIONAR CLUSTER
Acesse o localhost ou endereço que foi instalado.
Vá em menu Clusters, botão Add Cluster, From existing nodes (Custom), Nomeie como cluster01, botão next.
Aguarde pelo provisionamento.

__
Edite o cluster01, em Customize Node Run Command, Selecione:
etcd
Control Plane

Copie a linha de comando, parecendo com:
sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run rancher/rancher-agent:v2.4.5 --server https://192.168.1.106 --token h86gqvnq5wb4fw8hpnhzqqnbmprthrkf2k5xf5gd8x45bdfqdpnbrc --ca-checksum 70db14697fb1a4ae6d9d11da1c9568b31c3792306e7957576fb3f6afb83154a4 --etcd --controlplane --worker

Cole no terminal e execute, ele baixará o que for necessário para rodar.
Aguarde na tela de monitoramento que ficará pronta.




__ADICIONAR OS CONTAINERS
Vamos implantar (Deploy, adicionar) um container pelo próprio rancher.
Vá em Menu Cluster, cluster01, Default (active), botão deploy.
Nome: php
Docker Image: php:7.1.26-apache
Add Port
Port Mapping:
Port Name: php-8001
Publish the container port: 80
Protocol: TCP
As a: HostPort
On listening port: 8001
Clicar em Launch

_Testar depois com volume
Clica em Volumes, Add Volume (Bind-mount)
Volume Name: php-vol
Mount Point: /var/www/html
Clicar em Add Mount



FUNCIONOU, ver se dá para mudar de porta.
__________________________________________________  
__Configurações
Dark mode
Clique na imagem, lado superior direito, Preferences, Them, Dark

Adicionar Cluste
Clique em Menu Cluster, botão Add Cluster no lado superior direito, From existing mode (Custom)
Nomeie (cluster01), Next
Em Node options selecione etcd, Control Plane e Worker. Copie o código...

sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run rancher/rancher-agent:v2.4.5 --server https://192.168.1.106 --token v52ttj9shzwhzjs7t6vqgvl6h5mnlwsrcgsk5fk92qsz4jgfs898mg --ca-checksum 24393920386be9fd8d36ea0f52cf939e534be152b74c81574bc0a8c7c020c2ce --etcd --controlplane --worker

clique em Done, Vá no terminal, cole e execute o código.

Volte ao terminal e execute o comando top para verificar o processamento.

Vamos implantar (Deploy, adicionar) um container pelo próprio rancher.
Vá em Default, active, botão deploy.
Nome: nginx-8001
Docker Image: nginx
Add Port
Port Mapping:
Port Name: nginx.8001
Publish the container port: 80
Protocol: TCP
As a: HostPort
On listening port: 8001
Clica em Volumes, Add Volume (Bind-mount)
Volume Name: nginx.vol
Clicar em Launch



__________________________________________________  
 


docker run -d --restart=unless-stopped \
  -p 80:80 -p 443:443 \
  rancher/rancher:latest
  
  docker run -d --restart=unless-stopped \
  -p 80:80 -p 443:443 \
  -v /<CERT_DIRECTORY>/<FULL_CHAIN.pem>:/etc/rancher/ssl/cert.pem \
  -v /<CERT_DIRECTORY>/<PRIVATE_KEY.pem>:/etc/rancher/ssl/key.pem \
  -v /<CERT_DIRECTORY>/<CA_CERTS.pem>:/etc/rancher/ssl/cacerts.pem \
  rancher/rancher:latest
```

## rancher-server
```
https://hub.docker.com/r/rancher/server/
https://rancher.com/docs/rancher/v1.6/en/

docker pull rancher/server
docker pull rancher/server:stable
rancher/agent:v1.2.11

Instale o MySQL e o PHPmyAdmin, veja em no arquivo deste mesmo diretório mysql.txt

docker run \
  -d \
  -p 8080:8080 \
  --network mysql \
  --name rancher \
  rancher/server
  
79c  
docker start -i 79c




docker run \
  -d \
  -p 8080:8080 \
  --restart=unless-stopped \
  rancher/server
```

## redis
```
__REDIS
docker run redis

em outra janela de terminal (Ctrl+Shift+T)
docker images
docker ps
76

docker exec -it 76 redis-cli
ping
info
SET key1 "Hey there"
GET key1

em outra janela de terminal (Ctrl+Shift+T)
docker ps
docker stop 76
docker rm 76
docker ps -a

dentro da pasta containers
cd /home/usuario/Desktop/containers
mkdir redis
cd redis
touch commands.txt
code . &


docker network create redis
77

docker run \
    --name redis \
    --network redis \
    -d redis

eb

docker run \
    --name redis-commander \
    --network redis \
    -p 8081:8081 \
    -e REDIS_HOST=redis \
    -d rediscommander/redis-commander

25

docker ps
docker stop 25 eb
docker container prune -f
docker ps -a
________________________________
https://www.katacoda.com/courses/docker/deploying-first-container

docker run -d --name redisHostPort -p 6379:6379 redis:latest
docker run -d --name redisDynamic -p 6379 redis:latest
docker port redisDynamic 6379
docker ps

PERSISTENT DATA
docker run -d --name redisMapped -v /opt/docker/data/redis:/data redis
```

## redmine
```
https://hub.docker.com/_/redmine

docker pull redmine

Crie o banco de dados redmine

docker run \
  -d \
  -p 80:3000 \
  -e REDMINE_DB_MYSQL=mysql.5.7.wp \
  -e REDMINE_DB_USERNAME=root \
  -e REDMINE_DB_PASSWORD=my-password \
  --name redmine \
  --network mysql \
  redmine

d43

senha
admin
admin

Crie conta para o admin aprovar.
Funcionou.
```

## registry
```
https://hub.docker.com/_/registry

docker pull registry

docker run -d -p 5000:5000 --restart always --name registry registry:2

Now, use it from within Docker:

$ docker pull ubuntu
$ docker tag ubuntu localhost:5000/ubuntu
$ docker push localhost:5000/ubuntu
```

## sakai
```
https://github.com/hypothesis/sakai-docker

Não funcionou o comando abaixo
git submodules update --init

docker-compose up

Acessar  http://localhost:8080/portal

Installing the Hypothesis LTI tool
To test an LTI tool, you will need to:

Create a new "project site".
Click "Worksite Setup" in the left navbar
Select "Build your own site" => "project site"
Click "Continue".
Enter a title for the site and click "Continue".
On the "Project Site Tools" page, check the "External Tool" option and click Continue.
Click "Continue" on the Project Site Access page
Click "Create Site" on the confirmation page.
Select the new site in the list of sites. It should be displayed as active in the tabs at the top.
Click the "External Tool" link in the left navbar and then click "Edit" to configure it.
Register a new LMS application instance at https://hypothes.is/welcome and enter the launch URL, and consumer key and secret in Sakai.




OUTRO (deu erro)
https://wiki.centos.org/HowTos/Virtualization/VirtualBox
https://github.com/sakaicontrib/docker-sakai
export VIRTUALBOX_CPU_COUNT="2"
export VIRTUALBOX_MEMORY_SIZE="2048"

Copia as proximas linhas e execute:
curl -L https://github.com/docker/machine/releases/download/v0.16.1/docker-machine-`uname -s`-`uname -m` >/tmp/docker-machine &&
chmod +x /tmp/docker-machine &&
sudo cp /tmp/docker-machine /usr/local/bin/docker-machine
    
sudo yum install kernel-devel kernel-devel-$(uname -r) kernel-headers kernel-headers-$(uname -r) make patch gcc
sudo wget https://download.virtualbox.org/virtualbox/rpm/el/virtualbox.repo -P /etc/yum.repos.d
sudo yum install VirtualBox-5.2
systemctl status vboxdrv
wget https://download.virtualbox.org/virtualbox/5.2.20/Oracle_VM_VirtualBox_Extension_Pack-5.2.20.vbox-extpack
sudo VBoxManage extpack install  Oracle_VM_VirtualBox_Extension_Pack-5.2.20.vbox-extpack

docker-machine create -d virtualbox dev
eval $(docker-machine env dev)
```

## teamspeak
```
https://hub.docker.com/_/teamspeak

docker pull teamspeak

docker run -e TS3SERVER_LICENSE=view teamspeak

docker run -p 9987:9987/udp -p 10011:10011 -p 30033:30033 -e TS3SERVER_LICENSE=accept teamspeak

docker exec -it some-teamspeak sh
```

## ubuntu
```
docker run ubuntu ps
docker run -it ubuntu bash


FROM ubuntu
RUN apt-get update && apt-get install tree
```

## wordpress
```
Instale o MySQL e o PHPmyAdmin, veja em no arquivo deste mesmo diretório mysql.txt

docker pull wordpress:5.4

docker run \
    --network mysql \
    --name wordpress \
    -p 80:80 \
    -d wordpress:5.4
  
__Configuração
Abra o browser e acesse
localhost

prosseguindo a configuração, escolha língua (Português do Brasil), continue
Nome do banco de dados: wordpress
Nome de usuário: root
Senha: my-password
Servidor do banco de dados: 172.18.0.5
Prefixo da tabela: wp_





__Estabelecendo comunicação entre dois containers
docker stop aa
docker container prune -f
docker ps -a
docker run -it busybox

abra outra aba
docker run -it busybox

para saber qual id
hostname

para saber qual ip
hotname -i

faremos um ping da máquina ip final 2 para a máquina ip final 3
ping 172.17.0.3
aperte Ctrl+C para parar

faremos a mesma coisa com a outra, do final 3 para o final 2
ping 172.17.0.2
aperte Ctrl+C para parar

em outra tela de terminal (externo)
docker inspect 57

abra outra tela de terminal
docker inspect ef

veja que os gateway são os mesmos, porque estão no mesmo host
sobre o IPPrefixLen 16, os IP podem ser acessados:
172.17.5.6
172.17.100.7
172.17.0.5
172.17.0.10

__Environment variables
dentro do primeiro container
env

saída:
HOSTNAME=572e43ad5588
SHLVL=1
HOME=/root
TERM=xterm
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
PWD=/

em outra janela de terminal
docker run -it ubuntu
env

na primeira tela do terminal
hostname

em outra tela que não sejam as duas primeiras
docker exec 57 env

vamos sair de todos containers que estão rodando
exit

docker ps
docker ps -a
docker container prune -f
docker ps -a
```

## wso2

### Instalação Separada

```
Testado em 05MAR21.
For integrator:

Sem Iteração (considere fazer com interação se for novato ou se não conhecer a máquina que está trabalhando). Mesmo assim será necessário abrir outra janela quando terminar.

docker run -p 8280:8280 -p 8243:8243 -p 9443:9443 wso2/wso2ei-integrator & docker run -p 9713:9713 -p 9643:9643 -p 9613:9613 -p 7713:7713 -p 7613:7613  wso2/wso2ei-analytics-dashboard & docker run -p 9091:9091 -p 9444:9444 -p 9712:9712 -p 9612:9612 -p 7712:7712 -p 7612:7612 -p 7070:7070 -p 7443:7443 wso2/wso2ei-analytics-worker & docker run -p 9445:9445 wso2/wso2ei-business-process & docker run -p 5675:5675 -p 9446:9446 -p 8675:8675 -p 1886:1886 -p 8836:8836 -p 7614:7614 wso2/wso2ei-broker &

Com Interação (Será necessário uma janela para cada instalação, bom para ver os erros, principalmente se for a primeira vez)
docker run -it \
-p 8280:8280 \
-p 8243:8243 \
-p 9443:9443 \
wso2/wso2ei-integrator

For analytics dashboard:

docker run -it \
-p 9713:9713 \
-p 9643:9643 \
-p 9613:9613 \
-p 7713:7713 \
-p 7613:7613  \
wso2/wso2ei-analytics-dashboard

For analytics worker:

docker run -it \
-p 9091:9091 \
-p 9444:9444 \
-p 9712:9712 \
-p 9612:9612 \
-p 7712:7712 \
-p 7612:7612 \
-p 7070:7070 \
-p 7443:7443 \
wso2/wso2ei-analytics-worker

For business process:

docker run -it \
-p 9445:9445 \
wso2/wso2ei-business-process

For broker:

docker run -it \
-p 5675:5675 \
-p 9446:9446 \
-p 8675:8675 \
-p 1886:1886 \
-p 8836:8836 \
-p 7614:7614 \
wso2/wso2ei-broker

Once the container is started, access the management consoles of each profile via the following URLs on your favorite web browser using the credentials
username: admin and password: admin.

For integrator:

https://localhost:9443/carbon
admin, admin

For analytics dashboard:

https://localhost:9643/portal
admin, admin

For business process:

https://localhost:9445/carbon
admin, admin

For broker:

https://localhost:9446/carbon
admin, admin
```
### Com Docker Composer

```
________________________________________________________________________
__Instalação sem alterar nada
Fazer o downloado do código (https://github.com/wso2/docker-apim)
Passe via FTP para a máquina remota.
Descompacte a pasta em home/SeuUsuario/github
Acesse a pasta apim-with-analytics, veja seu caminho.
cd github/docker-apim-master/docker-compose/apim-with-analytics/
```
#### CentOS:
```
Altere nos arquivos Dockerfile:
docker-apim-3.2.x\docker-compose\apim-with-analytics\dockerfiles\apim\
FROM wso2/wso2am:3.2.0-centos

docker-apim-3.2.x\docker-compose\apim-with-analytics\dockerfiles\apim-analytics-dashboard
FROM wso2/wso2am-analytics-dashboard:3.2.0-centos

docker-apim-3.2.x\docker-compose\apim-with-analytics\dockerfiles\apim-analytics-worker
FROM wso2/wso2am-analytics-worker:3.2.0-centos
```
#### Ubuntu
```
Altere nos arquivos Dockerfile:
docker-apim-3.2.x\docker-compose\apim-with-analytics\dockerfiles\apim\
FROM wso2/wso2am:3.2.0

docker-apim-3.2.x\docker-compose\apim-with-analytics\dockerfiles\apim-analytics-dashboard
FROM wso2/wso2am-analytics-dashboard:3.2.0

docker-apim-3.2.x\docker-compose\apim-with-analytics\dockerfiles\apim-analytics-worker
FROM wso2/wso2am-analytics-worker:3.2.0
```
Compilar<br>
```
docker-compose up --build
```
Acessar<br>
```
Access the WSO2 API Manager web UIs using the below URLs via a web browser.
https://localhost:9443/publisher
https://localhost:9443/devportal
https://localhost:9443/admin
https://localhost:9443/carbon
Login to the web UIs using following credentials.

Username: admin
Password: admin
Please note that API Gateway will be available on following ports.

https://localhost:8243
https://localhost:8280
Access the WSO2 API Manager Analytics web UIs using the below URL via a web browser.

https://localhost:9643/analytics-dashboard
Login to the web UIs using following credentials.
Username: admin
Password: admin
```
### Instalação Personalizada
```
Fazer o downloado do código (https://github.com/wso2/docker-apim)
Descompacte para docker-apim-master
Com o WinSCP, passe via FTP para a máquina remota (na pasta home do usuário).
Acesse com Putty ou via SSH

```
#### Ubuntu:
```
1º Passo - Alterar os Dockfiles para baixar a versão CENTOS dos containers
1. Altere nos arquivos Dockerfile:

cd ~
sudo nano docker-apim-master/docker-compose/apim-with-analytics/dockerfiles/apim/Dockerfile
FROM wso2/wso2am:3.2.0
Ctrl+O, Enter, Ctrl+X

sudo nano docker-apim-master/docker-compose/apim-with-analytics/dockerfiles/apim-analytics-dashboard/Dockerfile
FROM wso2/wso2am-analytics-dashboard:3.2.0

sudo nano docker-apim-master/docker-compose/apim-with-analytics/dockerfiles/apim-analytics-worker/Dockerfile
FROM wso2/wso2am-analytics-worker:3.2.0

________________________________________________________________________
__2º Passo - Personalizando os CONF
__PASTA apim
Alterar o conf do APIMANAGER
sudo nano docker-apim-master/docker-compose/apim-with-analytics/conf/apim/repository/conf/deployment.toml

Alterar todos os localhost para url do servidor
Para procurar, pressione Ctrl+W, digite localhost, Enter
Troque todas as ocorrências para seuDominio.com

Se tiver IP fixo, altere...

pode alterar também a linha 3
de
node_ip = "127.0.0.1"
para 
node_ip = "SeuIPdoDominio"

alterar a senha de admin para adminPower
de
[super_admin]
username = "admin"
password = "admin"
create_admin_account = true

para
[super_admin]
username = "admin"
password = "adminPower"
create_admin_account = true

Ctrl+O, Enter, Ctrl+X

__ALTERAR O ARQUIVO DE CONFIGURAÇÃO DO ANALYTICS DASHBOARD
__Pasta apim-analytics-dashboard
sudo nano docker-apim-master/docker-compose/apim-with-analytics/conf/apim-analytics-dashboard/conf/dashboard/deployment.yaml

Alterar todos os localhost para url do servidor
Para procurar, pressione Ctrl+W, digite localhost, Enter
Troque todas as ocorrências para seuDominio.com

ALTERAR TODOS OS CAMPOS DE PASSWORD admin para adminPower
Ctrl+-, 340
de
linha 340 password: admin
linha 386 adminPassword: admin
linha 391 kmPassword: admin
para
linha 340 password: adminPower
linha 386 adminPassword: adminPower
linha 391 kmPassword: adminPower

Ctrl+O, Enter, Ctrl+X

__ALTERAR O ARQUIVO DE CONFIGURAÇÃO DO ANALYTICS WORKER
sudo nano docker-apim-master/docker-compose/apim-with-analytics/conf/apim-analytics-worker/conf/worker/deployment.yaml

alterar todos os localhost para url do servidor
seuDominio.com

SÓ TERA UM LOCALHOST PARA ALTERAR (linha 440)
grpc://localhost:9806/org.wso2.analytics.mgw.grpc.service.AnalyticsSendService/sendAnalytics
para
grpc://seuDominio.com:9806/org.wso2.analytics.mgw.grpc.service.AnalyticsSendService/sendAnalytics


Alterar a senha (está codificada em BASE64)
https://www.base64encode.org/


Ctrl+-, 503
senha admin = YWRtaW4=
senha adminPower = YWRtaW5Qb3dlcg==

Vai ficar assim:

de
# Authentication configuration

auth.configs:
  type: 'local'        # Type of the IdP client used
  userManager:
    adminRole: admin   # Admin role which is granted all permissions
    userStore:         # User store
      users:
        -
          user:
            username: admin
            password: YWRtaW4=
            roles: 1
            
para
# Authentication configuration

auth.configs:
  type: 'local'        # Type of the IdP client used
  userManager:
    adminRole: admin   # Admin role which is granted all permissions
    userStore:         # User store
      users:
        -
          user:
            username: admin
            password: YWRtaW5Qb3dlcg==
            roles: 1

Ctrl+O, Enter, Ctrl+X

cd docker-apim-master/docker-compose/apim-with-analytics/

docker-compose up --build
```
#### Se o PC for fraco
```
Comente as linhas no docker-compose:
47 a 49
    #depends_on:
      #mysql:
        #condition: service_healthy

59 a 63
    #depends_on:
      #mysql:
        #condition: service_healthy
      #am-analytics-worker:
        #condition: service_healthy

79 a 83
    #depends_on:
      #mysql:
        #condition: service_healthy
      #api-manager:
        #condition: service_healthy

Com o Build, serão criados 4 Containers: Api Manager, Analytics, Worker e MySQL.
```
#### Ajustar a memória
```
Mudar a memória. Entre no Container do Api Manager.
Aperte Ctrl+C para parar de rodar o container.

Suponhamos que só existam os 4 containers que criamos...
docker ps -a
Vai aparecer todos os containers que foram parados...
f7336cbab40b   wso2_am-analytics-worker      "/home/wso2carbon/do…"            wso2_am-analytics-worker_1
ae410d5f8e61   wso2_am-analytics-dashboard   "/home/wso2carbon/do…"            wso2_am-analytics-dashboard_1
f3652a678f82   wso2_api-manager              "/home/wso2carbon/do…"            wso2_api-manager_1
e7b64f9cb3f7   mysql:5.7.31                  "docker-entrypoint.s…"            wso2_mysql_1

Reinicie o container para editar
docker start f3652a678f82

docker exec -it IdDoContainerApiManager bash

Entrar como Root
docker exec -it -u 0 f3652a678f82 bash
apt-get update
apt-get install nano

Edite o arquivo
nano wso2am-3.2.0/bin/wso2server.sh

Encontrar essa parte e incrementar os valores deixando como está abaixo: 

Procurar por $JVM_MEM_OPTS com Ctrl+W ou Ctrl+-, 294


if [ -z "$JVM_MEM_OPTS" ]; then 
   java_version=$("$JAVACMD" -version 2>&1 | awk -F '"' '/version/ {print $2}') 
   JVM_MEM_OPTS="-Xms512m -Xmx1024m" 
   if [ "$java_version" \< "1.8" ]; then 
      JVM_MEM_OPTS="$JVM_MEM_OPTS -XX:MaxPermSize=512m" 

Ctrl+O, Enter, Ctrl+X

Sair do container
exit

Parar o container e subir todos.
docker stop $(docker ps -q)
docker start $(docker ps -aq)


Access the WSO2 API Manager web UIs using the below URLs via a web browser.

https://localhost:9443/publisher
https://localhost:9443/devportal
https://localhost:9443/admin
https://localhost:9443/carbon

Login to the web UIs using following credentials.
Username: admin
Password: adminPower
Please note that API Gateway will be available on following ports.

https://localhost:8243
https://localhost:8280
Access the WSO2 API Manager Analytics web UIs using the below URL via a web browser.

https://localhost:9643/analytics-dashboard
Login to the web UIs using following credentials.

Username: admin
Password: adminPower


ERROS:
https://apim.docs.wso2.com/en/latest/troubleshooting/troubleshooting-invalid-callback-error/

```
#### CentOS:
```
1º Passo - Alterar os Dockfiles para baixar a versão CENTOS dos containers
1. Altere nos arquivos Dockerfile:

cd ~
sudo nano docker-apim-master/docker-compose/apim-with-analytics/dockerfiles/apim/Dockerfile
FROM wso2/wso2am:3.2.0-centos
Ctrl+O, Enter, Ctrl+X

sudo nano docker-apim-master/docker-compose/apim-with-analytics/dockerfiles/apim-analytics-dashboard/Dockerfile
FROM wso2/wso2am-analytics-dashboard:3.2.0-centos

sudo nano docker-apim-master/docker-compose/apim-with-analytics/dockerfiles/apim-analytics-worker/Dockerfile
FROM wso2/wso2am-analytics-worker:3.2.0-centos

________________________________________________________________________
__2º Passo - Personalizando os CONF
__PASTA apim
Alterar o conf do APIMANAGER em docker-apim-3.2.x\docker-compose\apim-with-analytics\conf\apim\repository\conf
cd docker-apim-master/docker-compose/apim-with-analytics/conf/apim/repository/conf

editar o arquivo deployment.toml
nano deployment.toml

alterar todos os localhost para url do servidor
seuDominio.com

pode alterar também a linha 3
de
node_ip = "127.0.0.1"
para 
node_ip = "SeuIPdoDominio"

alterar a senha de admin para adminPower
de
[super_admin]
username = "admin"
password = "admin"
create_admin_account = true

para
[super_admin]
username = "admin"
password = "adminPower"
create_admin_account = true

__ALTERAR O ARQUIVO DE CONFIGURAÇÃO DO ANALYTICS DASHBOARD
__Pasta apim-analytics-dashboard
cd docker-apim-master/docker-compose/apim-with-analytics/conf/apim-analytics-dashboard/conf/dashboard
docker-apim-3.2.x\docker-compose\apim-with-analytics\conf\apim-analytics-dashboard\conf\dashboard

editar o arquivo deployment.yaml
nano deployment.yaml

alterar todos os localhost para url do servidor 
seuDominio.com

ALTERAR TODOS OS CAMPOS DE PASSWORD admin para adminPower
de
linha 340 password: admin
linha 386 adminPassword: admin
linha 391 kmPassword: admin
para
linha 340 password: adminPower
linha 386 adminPassword: adminPower
linha 391 kmPassword: adminPower

__ALTERAR O ARQUIVO DE CONFIGURAÇÃO DO ANALYTICS WORKER
__Pasta apim-analytics-worker
docker-apim-3.2.x\docker-compose\apim-with-analytics\conf\apim-analytics-worker\conf\worker
editar o arquivo deployment.yaml

alterar todos os localhost para url do servidor
seuDominio.com

SÓ TERA UM LOCALHOST PARA ALTERAR 
grpc://localhost:9806/org.wso2.analytics.mgw.grpc.service.AnalyticsSendService/sendAnalytics
para
grpc://seuDominio.com:9806/org.wso2.analytics.mgw.grpc.service.AnalyticsSendService/sendAnalytics

Alterar a senha na seção. A senha está codificada em BASE64
https://www.base64encode.org/

senha admin = YWRtaW4=
senha adminPower = YWRtaW5Qb3dlcg==

de
# Authentication configuration

auth.configs:
  type: 'local'        # Type of the IdP client used
  userManager:
    adminRole: admin   # Admin role which is granted all permissions
    userStore:         # User store
      users:
        -
          user:
            username: admin
            password: YWRtaW4=
            roles: 1
            
para
# Authentication configuration

auth.configs:
  type: 'local'        # Type of the IdP client used
  userManager:
    adminRole: admin   # Admin role which is granted all permissions
    userStore:         # User store
      users:
        -
          user:
            username: admin
            password: YWRtaW5Qb3dlcg==
            roles: 1


cd docker-apim/docker-compose/apim-with-analytics




docker-compose up --build

Access the WSO2 API Manager web UIs using the below URLs via a web browser.

https://localhost:9443/publisher
https://localhost:9443/devportal
https://localhost:9443/admin
https://localhost:9443/carbon
Login to the web UIs using following credentials.

Username: admin
Password: admin
Please note that API Gateway will be available on following ports.

https://localhost:8243
https://localhost:8280
Access the WSO2 API Manager Analytics web UIs using the below URL via a web browser.

https://localhost:9643/analytics-dashboard
Login to the web UIs using following credentials.

Username: admin
Password: admin


ERROS:
https://apim.docs.wso2.com/en/latest/troubleshooting/troubleshooting-invalid-callback-error/


______________________________________________________________________________________________

https://wso2.com/api-management/install/docker/get-started/

docker run -it -p 9443:9443 -p 8243:8243 -p 8280:8280 wso2/wso2am:3.2.0
   
   Once the container is started, access the following URLs on your favorite web browser using the credentials username: admin and password: admin.

Publisher - https://localhost:9443/publisher
Store - https://localhost:9443/devportal
Admin console - https://localhost:9443/admin
Carbon console - https://localhost:9443/carbon

Please note that the WSO2 API Manager Gateway will be available on the following ports.

https://localhost:8243
http://localhost:8280

https://apim.docs.wso2.com/en/latest/getting-started/quick-start-guide/
```

## xampp
```
docker pull xampp/app:7.1
```

## xwiki
```
https://hub.docker.com/_/xwiki

docker pull xwiki

docker run \
  -d \
  -p 80:80 \
  --network mysql \
  xwiki

5d
```

