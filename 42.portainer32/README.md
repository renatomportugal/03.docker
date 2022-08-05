# docker32bit

```CMD

sudo apt install -y docker.io
sudo usermod -aG docker tcnct
sudo systemctl status docker
sudo reboot now
docker version
docker pull alpine
docker pull nginx
docker run --name nginx -d -p 80:80 nginx

https://hub.docker.com/r/thibaudlabat/portainer-ce-32
Subir a pasta via FTP
cd 42.portainer32
cd docker
docker image build . -t portainer-ce-32

Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?









docker run -d --restart always \
    -p 9000:9000 -p 8000:8000 \
    -v "/var/run/docker.sock:/var/run/docker.sock" \
    portainer-ce-32

DEV
docker run -it --rm \
    -p 9000:9000 -p 8000:8000 \
    -v "/var/run/docker.sock:/var/run/docker.sock" \
    portainer-ce-32 /bin/sh

sudo apt-get update

sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common -y

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo apt-key fingerprint 0EBFCD88

sudo add-apt-repository \
   "deb [arch=386] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

sudo nano /etc/apt/sources.list


cat /etc/os-release
lsb_release -a
lsb_release -cs


Desinstalar

Uninstall old versions:

sudo apt-get remove docker docker-engine docker.io containerd runc

Uninstall the Docker Engine - Community package:

sudo apt-get purge docker-ce
sudo rm -rf /var/lib/docker




```
