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