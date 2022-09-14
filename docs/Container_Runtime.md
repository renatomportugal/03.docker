# Container_Runtime

```CMD
https://kubernetes.io/docs/setup/production-environment/container-runtimes/
```

## containerd

```CMD
Desinstalar
sudo rm -rf /usr/local/sbin/runc
sudo rm -rf /usr/local/lib/systemd/system/

...deu errado...
https://github.com/containerd/containerd/blob/main/docs/getting-started.md
https://github.com/containerd/containerd/releases
13SET22
cd ~
wget https://github.com/containerd/containerd/releases/download/v1.6.8/containerd-1.6.8-linux-arm64.tar.gz
sudo tar Cxzvf /usr/local containerd-1.6.8-linux-arm64.tar.gz


If you intend to start containerd via systemd, you should also download the containerd.service unit file
wget https://github.com/containerd/containerd/blob/main/containerd.service
sudo mkdir /usr/local/lib/systemd/
sudo mkdir /usr/local/lib/systemd/system/
sudo mv containerd.service /usr/local/lib/systemd/system/containerd.service

systemctl daemon-reload
systemctl enable --now containerd
Neste momento deu erro....

...continuando...
Step 2: Installing runc
https://github.com/opencontainers/runc/releases
wget https://github.com/opencontainers/runc/releases/download/v1.1.4/runc.amd64
sudo install -m 755 runc.amd64 /usr/local/sbin/runc

Installing CNI plugins
https://github.com/containernetworking/plugins/releases
wget https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-amd64-v1.1.1.tgz
sudo mkdir -p /opt/cni/bin
sudo tar Cxzvf /opt/cni/bin cni-plugins-linux-amd64-v1.1.1.tgz

Configuração
https://github.com/containerd/containerd/blob/main/docs/man/containerd-config.toml.5.md

sudo mkdir /etc/containerd/
sudo nano /etc/containerd/config.toml

version = 2

root = "/var/lib/containerd"
state = "/run/containerd"
oom_score = 0
imports = ["/etc/containerd/runtime_*.toml", "./debug.toml"]

[grpc]
  address = "/run/containerd/containerd.sock"
  uid = 0
  gid = 0

[debug]
  address = "/run/containerd/debug.sock"
  uid = 0
  gid = 0
  level = "info"

[metrics]
  address = ""
  grpc_histogram = false

[cgroup]
  path = ""

[plugins]
  [plugins."io.containerd.monitor.v1.cgroups"]
    no_prometheus = false
  [plugins."io.containerd.service.v1.diff-service"]
    default = ["walking"]
  [plugins."io.containerd.gc.v1.scheduler"]
    pause_threshold = 0.02
    deletion_threshold = 0
    mutation_threshold = 100
    schedule_delay = 0
    startup_delay = "100ms"
  [plugins."io.containerd.runtime.v2.task"]
    platforms = ["linux/amd64"]
    sched_core = true
  [plugins."io.containerd.service.v1.tasks-service"]
    blockio_config_file = ""
    rdt_config_file = ""

Ctrl+O, Enter, Ctrl+X

sudo systemctl restart containerd
...erro:
Failed to restart containerd.service: Unit containerd.service has a bad unit file setting.

journalctl -xe

ctr --help
...erro:
-bash: /usr/local/bin/ctr: cannot execute binary file: Exec format error


ctr images pull docker.io/library/redis:alpine
ctr run docker.io/library/redis:alpine redis
```

## CRI-O

```CMD
https://github.com/cri-o/cri-o

```

## Docker_Engine

```CMD
1. Instalar o docker e o docker-compose
https://renatomportugal.github.io/03.docker/#/Instalacao?id=no-ubuntu
Ver arquivo de configuração
sudo nano /etc/containerd/config.toml
Ctrl+X

2. Instalar o cri-dockerd
https://github.com/Mirantis/cri-dockerd

git clone https://github.com/Mirantis/cri-dockerd.git

# Run these commands as root
###Install GO###
wget https://storage.googleapis.com/golang/getgo/installer_linux
sudo chmod +x ./installer_linux
sudo ./installer_linux
source ~/.bash_profile

sudo passwd root
source /root/.bash_profile

cd cri-dockerd
mkdir bin
go build -o bin/cri-dockerd
mkdir -p /usr/local/bin
install -o root -g root -m 0755 bin/cri-dockerd /usr/local/bin/cri-dockerd
cp -a packaging/systemd/* /etc/systemd/system
sed -i -e 's,/usr/bin/cri-dockerd,/usr/local/bin/cri-dockerd,' /etc/systemd/system/cri-docker.service
systemctl daemon-reload
systemctl enable cri-docker.service
systemctl enable --now cri-docker.socket


...outro...
https://www.mirantis.com/blog/how-to-install-cri-dockerd-and-migrate-nodes-from-dockershim
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.2.0/cri-dockerd-v0.2.0-linux-amd64.tar.gz
tar xvf cri-dockerd-v0.2.0-linux-amd64.tar.gz
sudo mv ./cri-dockerd /usr/local/bin/

wget https://raw.githubusercontent.com/Mirantis/cri-dockerd/master/packaging/systemd/cri-docker.service
wget https://raw.githubusercontent.com/Mirantis/cri-dockerd/master/packaging/systemd/cri-docker.socket
sudo mv cri-docker.socket cri-docker.service /etc/systemd/system/
sudo sed -i -e 's,/usr/bin/cri-dockerd,/usr/local/bin/cri-dockerd,' /etc/systemd/system/cri-docker.service

systemctl daemon-reload
systemctl enable cri-docker.service
systemctl enable --now cri-docker.socket

systemctl status cri-docker.socket


Configure nodes to use cri-dockerd
Here, we’ll assume we’ve used kubeadm to configure our node. Use your text editor of choice to open the node’s kubeadm-flags.env file—I’m using nano in the example below. 

mkdir /var/lib/kubelet/
nano /var/lib/kubelet/kubeadm-flags.env

Cole:
unix:///var/run/cri-dockerd.sock
Ctrl+O, Enter, Ctrl+X

kubectl cordon













sudo apt update && sudo apt install golang
git clone https://github.com/Mirantis/cri-dockerd.git

cd cri-dockerd
sudo mkdir bin
sudo -E go build -o bin/cri-dockerd



...erro...
build github.com/Mirantis/cri-dockerd: cannot load net/netip: malformed module path "net/netip": missing dot in first path element

sudo mkdir -p /usr/local/bin
install -o root -g root -m 0755 bin/cri-dockerd /usr/local/bin/cri-dockerd
sudo cp -a packaging/systemd/* /etc/systemd/system
sed -i -e 's,/usr/bin/cri-dockerd,/usr/local/bin/cri-dockerd,' /etc/systemd/system/cri-docker.service
sudo systemctl daemon-reload
sudo systemctl enable cri-docker.service
sudo systemctl enable --now cri-docker.socket

```

## Mirantis_Container_Runtime

```CMD

```
