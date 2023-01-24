# K8s

## Tutoriais

```CMD
https://kubernetes.io/pt-br/docs/tutorials/kubernetes-basics/_print/

```

## 01.DESENVOLVIMENTO

### minikube

```CMD
https://minikube.sigs.k8s.io/docs/start/

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube config set driver docker
minikube start --driver=docker


Problemas com proxy:
https://minikube.sigs.k8s.io/docs/handbook/vpn_and_proxy/


kubectl get po -A
minikube kubectl -- get po -A
alias kubectl="minikube kubectl --"
minikube dashboard

adicionar os certificados da empresa para ~/.minikube/certs
minikube delete
minikube start
minikube start --docker-env http_proxy=$http_proxy --docker-env https_proxy=$https_proxy --docker-env no_proxy=$no_proxy

[Service]
kubectl create deployment hello-minikube --image=kicbase/echo-server:1.0
kubectl expose deployment hello-minikube --type=NodePort --port=8080

kubectl get services hello-minikube
minikube service hello-minikube
kubectl port-forward service/hello-minikube 7080:8080
http://localhost:7080/

[LoadBalance]
kubectl create deployment balanced --image=kicbase/echo-server:1.0
kubectl expose deployment balanced --type=LoadBalancer --port=8080

minikube tunnel
kubectl get services balanced
<EXTERNAL-IP>:8080

[Ingress]
minikube addons enable ingress

___ingress-example.yaml________________________________________________________
kind: Pod
apiVersion: v1
metadata:
  name: foo-app
  labels:
    app: foo
spec:
  containers:
    - name: foo-app
      image: 'kicbase/echo-server:1.0'
---
kind: Service
apiVersion: v1
metadata:
  name: foo-service
spec:
  selector:
    app: foo
  ports:
    - port: 8080
---
kind: Pod
apiVersion: v1
metadata:
  name: bar-app
  labels:
    app: bar
spec:
  containers:
    - name: bar-app
      image: 'kicbase/echo-server:1.0'
---
kind: Service
apiVersion: v1
metadata:
  name: bar-service
spec:
  selector:
    app: bar
  ports:
    - port: 8080
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-ingress
spec:
  rules:
    - http:
        paths:
          - pathType: Prefix
            path: /foo
            backend:
              service:
                name: foo-service
                port:
                  number: 8080
          - pathType: Prefix
            path: /bar
            backend:
              service:
                name: bar-service
                port:
                  number: 8080
_______________________________________________________________________________

kubectl apply -f https://storage.googleapis.com/minikube-site-examples/ingress-example.yaml

Note for Docker Desktop Users:
To get ingress to work you’ll need to open a new terminal window and run minikube tunnel and in the following step use 127.0.0.1 in place of <ip_from_above>.

Now verify that the ingress works

$ curl <ip_from_above>/foo
Request served by foo-app
...

$ curl <ip_from_above>/bar
Request served by bar-app
...




minikube pause
minikube unpause
minikube stop
minikube config set memory 9001
minikube addons list
minikube start -p aged --kubernetes-version=v1.16.1
minikube delete --all


gcr.io/k8s-minikube/kicbase:v0.0.36



```

### kind

```CMD
https://kind.sigs.k8s.io/docs/user/quick-start/#installation

```

## 02.PRODUÇÃO

### kubeadm

```CMD
https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

```

### kubespray

```CMD
https://kubernetes.io/docs/setup/production-environment/tools/kubespray/

```

### container_runtimes

```CMD
https://kubernetes.io/docs/setup/production-environment/container-runtimes/


```

## Tentativas

### Primeiros_Passos

```CMD
Kind e Minikube não devem ser utilizados para produção.

https://microk8s.io/
MicroK8S - Produção, Edge Computing e IoT (Internet of things)

https://k3s.io/
k3s (Lightweight) - Produção, Edge, IoT (Internet of things), CI, ARM.

https://k0sproject.io/
k0s - Produção x86-64, ARM64, ARMv7.
```

### Primeiros_Passos_No_Windows

```CMD
Instalar o Docker Desktop e ativar o kubernetes nas opções (apenas a opção Enable kubernetes)

https://livro.descomplicandokubernetes.com.br/pt/day_one/descomplicando_kubernetes.html

kubectl get nodes
NAME             STATUS   ROLES                  AGE     VERSION
docker-desktop   Ready    control-plane,master   3m46s   v1.22.4

kubectl get pods
No resources found in default namespace.

kubectl get pods --all-namespaces
NAMESPACE     NAME                                     READY   STATUS    RESTARTS   AGE
kube-system   coredns-78fcd69978-6kzzk                 1/1     Running   0          8m52s
kube-system   coredns-78fcd69978-snhsc                 1/1     Running   0          8m52s
kube-system   etcd-docker-desktop                      1/1     Running   1          8m52s
kube-system   kube-apiserver-docker-desktop            1/1     Running   1          8m50s
kube-system   kube-controller-manager-docker-desktop   1/1     Running   1          8m54s
kube-system   kube-proxy-2qf7m                         1/1     Running   0          8m52s
kube-system   kube-scheduler-docker-desktop            1/1     Running   1          8m48s
kube-system   storage-provisioner                      1/1     Running   0          8m24s
kube-system   vpnkit-controller                        1/1     Running   0          8m24s

kubectl get all --all-namespaces
NAMESPACE     NAME                                         READY   STATUS    RESTARTS   AGE
kube-system   pod/coredns-78fcd69978-6kzzk                 1/1     Running   0          9m41s
kube-system   pod/coredns-78fcd69978-snhsc                 1/1     Running   0          9m41s
kube-system   pod/etcd-docker-desktop                      1/1     Running   1          9m41s
kube-system   pod/kube-apiserver-docker-desktop            1/1     Running   1          9m39s
kube-system   pod/kube-controller-manager-docker-desktop   1/1     Running   1          9m43s
kube-system   pod/kube-proxy-2qf7m                         1/1     Running   0          9m41s
kube-system   pod/kube-scheduler-docker-desktop            1/1     Running   1          9m37s
kube-system   pod/storage-provisioner                      1/1     Running   0          9m13s
kube-system   pod/vpnkit-controller                        1/1     Running   0          9m13s

NAMESPACE     NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
default       service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP                  9m46s
kube-system   service/kube-dns     ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   9m45s

NAMESPACE     NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
kube-system   daemonset.apps/kube-proxy   1         1         1       1            1           kubernetes.io/os=linux   9m44s

NAMESPACE     NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
kube-system   deployment.apps/coredns   2/2     2            2           9m45s

NAMESPACE     NAME                                 DESIRED   CURRENT   READY   AGE
kube-system   replicaset.apps/coredns-78fcd69978   2         2         2       9m41s

No windows não vai...
cat /var/lib/rancher/k3s/server/node-token

Então iremos instalar no linux

```

### Primeiros_Passos_No_Linux_Com_K3S

```CMD
MÁQUINA US01

Para instalar:

curl -sfL https://get.k3s.io | K3S_KUBECONFIG_MODE="644" sh -s -

kubectl version --client

kubectl get nodes

ERRO
The connection to the server 127.0.0.1:6443 was refused - did you specify the right host or port?

kubectl cluster-info

Kubernetes control plane is running at https://127.0.0.1:6443
CoreDNS is running at https://127.0.0.1:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://127.0.0.1:6443/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy

sudo -i
swapoff -a
exit
strace -eopenat kubectl version

kubectl get nodes

NAME        STATUS   ROLES                  AGE     VERSION
tcnct-166   Ready    control-plane,master   7m28s   v1.21.7+k3s1

kubectl get pods
No resources found in default namespace.

kubectl get pods --all-namespaces
NAMESPACE     NAME                                      READY   STATUS      RESTARTS   AGE
kube-system   helm-install-traefik-crd-gzvkg            0/1     Completed   0          8m22s
kube-system   helm-install-traefik-7xgx5                0/1     Completed   1          8m22s
kube-system   metrics-server-86cbb8457f-zdk5s           1/1     Running     1          8m22s
kube-system   local-path-provisioner-5ff76fc89d-g55sq   1/1     Running     1          8m22s
kube-system   svclb-traefik-qv5js                       2/2     Running     2          5m47s
kube-system   coredns-7448499f4d-j7mwx                  1/1     Running     1          8m22s
kube-system   traefik-6b84f7cbc-9dssq                   1/1     Running     1          5m48s

NAMESPACE     NAME                                          READY   STATUS      RESTARTS   AGE
kube-system   pod/helm-install-traefik-crd-gzvkg            0/1     Completed   0          9m8s
kube-system   pod/helm-install-traefik-7xgx5                0/1     Completed   1          9m8s
kube-system   pod/metrics-server-86cbb8457f-zdk5s           1/1     Running     1          9m8s
kube-system   pod/local-path-provisioner-5ff76fc89d-g55sq   1/1     Running     1          9m8s
kube-system   pod/svclb-traefik-qv5js                       2/2     Running     2          6m33s
kube-system   pod/coredns-7448499f4d-j7mwx                  1/1     Running     1          9m8s
kube-system   pod/traefik-6b84f7cbc-9dssq                   1/1     Running     1          6m34s

NAMESPACE     NAME                     TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE
default       service/kubernetes       ClusterIP      10.43.0.1       <none>          443/TCP                      9m24s
kube-system   service/kube-dns         ClusterIP      10.43.0.10      <none>          53/UDP,53/TCP,9153/TCP       9m22s
kube-system   service/metrics-server   ClusterIP      10.43.213.191   <none>          443/TCP                      9m21s
kube-system   service/traefik          LoadBalancer   10.43.20.135    192.168.32.15   80:31613/TCP,443:30371/TCP   6m34s

NAMESPACE     NAME                           DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
kube-system   daemonset.apps/svclb-traefik   1         1         1       1            1           <none>          6m34s

NAMESPACE     NAME                                     READY   UP-TO-DATE   AVAILABLE   AGE
kube-system   deployment.apps/metrics-server           1/1     1            1           9m21s
kube-system   deployment.apps/local-path-provisioner   1/1     1            1           9m22s
kube-system   deployment.apps/coredns                  1/1     1            1           9m22s
kube-system   deployment.apps/traefik                  1/1     1            1           6m34s

NAMESPACE     NAME                                                DESIRED   CURRENT   READY   AGE
kube-system   replicaset.apps/metrics-server-86cbb8457f           1         1         1       9m9s
kube-system   replicaset.apps/local-path-provisioner-5ff76fc89d   1         1         1       9m9s
kube-system   replicaset.apps/coredns-7448499f4d                  1         1         1       9m9s
kube-system   replicaset.apps/traefik-6b84f7cbc                   1         1         1       6m34s

NAMESPACE     NAME                                 COMPLETIONS   DURATION   AGE
kube-system   job.batch/helm-install-traefik-crd   1/1           2m9s       9m20s
kube-system   job.batch/helm-install-traefik       1/1           2m36s      9m20s

CONFIGURAR O AUTOCOMPLETE
sudo apt-get install bash-completion

configura o autocomplete na sua sessão atual
source <(kubectl completion bash)

Add autocomplete permanentemente ao seu shell.
echo "source <(kubectl completion bash)" >> ~/.bashrc

Crie o alias k para kubectl:
alias k=kubectl
complete -F __start_kubectl k

Pegar os nós
k get nodes

Vamos pegar o token
sudo cat /var/lib/rancher/k3s/server/node-token
K1088b89be1dbb53cbf952911613ca2b3254980a1e46b3ec383e8b35ddd15deb6a0::server:7268420907567a52f015191e6ae3dc68

ifconfig
192.168.1.116


MÁQUINA US02
ADICIONAR OUTROS NÓS

curl -sfL https://get.k3s.io | K3S_URL=https://192.168.1.116:6443 K3S_TOKEN=K1088b89be1dbb53cbf952911613ca2b3254980a1e46b3ec383e8b35ddd15deb6a0::server:7268420907567a52f015191e6ae3dc68 sh -

k get nodes

```

### Outra_Instalacao

```CMD
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

BAIXAR
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

VALIDAR
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
echo "$(<kubectl.sha256) kubectl" | sha256sum --check

INSTALAR
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

Verify kubectl configuration
kubectl cluster-info

kubectl cluster-info dump

sudo nano /etc/rancher/k3s/k3s.yaml
server: https://172.17.128.10:6443

sudo nano ~/.kube/config

```

### Mais_Uma_Tentativa_22JUN22

```CMD
https://livro.descomplicandokubernetes.com.br/pt/day_one/descomplicando_kubernetes.html

01. Instalação

curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl

Resultado:
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 43.5M  100 43.5M    0     0  2720k      0  0:00:16  0:00:16 --:--:-- 2838k

chmod +x ./kubectl

sudo mv ./kubectl /usr/local/bin/kubectl

kubectl version --client

Saída:
WARNING: This version information is deprecated and will be replaced with the output from kubectl version --short.  Use --output=yaml|json to get the full version.
Client Version: version.Info{Major:"1", Minor:"24", GitVersion:"v1.24.2", GitCommit:"f66044f4361b9f1f96f0053dd46cb7dce5e990a8", GitTreeState:"clean", BuildDate:"2022-06-15T14:22:29Z", GoVersion:"go1.18.3", Compiler:"gc", Platform:"linux/amd64"}
Kustomize Version: v4.5.4

kubectl version --client --output=yaml

clientVersion:
  buildDate: "2022-06-15T14:22:29Z"
  compiler: gc
  gitCommit: f66044f4361b9f1f96f0053dd46cb7dce5e990a8
  gitTreeState: clean
  gitVersion: v1.24.2
  goVersion: go1.18.3
  major: "1"
  minor: "24"
  platform: linux/amd64
kustomizeVersion: v4.5.4


02. kubectl: alias e autocomplete

Execute o seguinte comando para configurar o alias e autocomplete para o kubectl.

No Bash: configura o autocomplete na sua sessão atual (antes, certifique-se de ter instalado o pacote bash-completion).

source <(kubectl completion bash) 

Adicionar autocomplete permanentemente:

echo "source <(kubectl completion bash)" >> ~/.bashrc 

Então fica assim, Crie o alias k para kubectl:

alias k=kubectl

complete -F __start_kubectl k




kubectl cluster-info

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
The connection to the server localhost:8080 was refused - did you specify the right host or port?

kubectl cluster-info dump

```

### Testar_Depois

```CMD
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"

echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check

Retorno:
kubectl: OK
ou
kubectl: FAILED
sha256sum: WARNING: 1 computed checksum did NOT match

sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

chmod +x kubectl
mkdir -p ~/.local/bin
mv ./kubectl ~/.local/bin/kubectl
# and then append (or prepend) ~/.local/bin to $PATH

kubectl version --client

kubectl version --client --output=yaml

```

### Primeiros_Passos_No_Linux_28SET22

```CMD
https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
1. Ubuntu 20.04
cat /etc/os-release
Master-plane (192.168.1.106), Worker nodes (192.168.1.108, 192.168.1.110, 192.168.1.112)
TODOS:
NAME="Ubuntu"
VERSION="20.04.5 LTS (Focal Fossa)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 20.04.5 LTS"
VERSION_ID="20.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=focal
UBUNTU_CODENAME=focal

2. 2 GB or more of RAM
htop
Master-plane
192.168.1.106 - 7,7GB
Worker nodes
192.168.1.108 - 3.64GB
192.168.1.110 - 5,54GB
192.168.1.112 - 1,81GB

3. 2 CPUs or more
htop
Master-plane
192.168.1.106 - 2
Worker nodes
192.168.1.108 - 2
192.168.1.110 - 2
192.168.1.112 - 2

4. Full network connectivity between all machines in the cluster
ip link
ifconfig -a
Para instalar...
sudo apt install net-tools
Master-plane
192.168.1.106 - ens34
Worker nodes
192.168.1.108 - enp4s0
192.168.1.110 - enp2s0
192.168.1.112 - enp0s4

5. Unique hostname, MAC address, and product_uuid for every node
cat /proc/sys/kernel/hostname
ifconfig
Master-plane
192.168.1.106 - tcnct-106, 44:87:fc:b1:50:9d
Worker nodes
192.168.1.108 - tcnct-108, c8:9c:dc:01:45:2f
192.168.1.110 - tcnct-110, 60:eb:69:db:69:62
192.168.1.112 - tcnct-112, 00:03:0d:aa:ad:f5

5.
sudo cat /sys/class/dmi/id/product_uuid

Master-plane
192.168.1.106 - 00020003-0004-0005-0006-000700080009

Worker nodes
192.168.1.108 - 01010101-0101-0101-0101-010101010101
192.168.1.110 - f6a533b9-6f72-0740-89bf-60eb69db6962
192.168.1.112 - 0034e9e7-4064-0010-b785-0e34e9efe205

6. Certain ports are open on your machines
nc 127.0.0.1 6443
nc 127.0.0.1 10250

Control plane
Protocol	Direction	Port Range	Purpose	                Used By
TCP	      Inbound	  6443	      Kubernetes API server   All
TCP	      Inbound	  2379-2380	  etcd server client API	kube-apiserver, etcd
TCP	      Inbound	  10250	      Kubelet API             Self, Control plane
TCP	      Inbound	  10259	      kube-scheduler	        Self
TCP	      Inbound	  10257	      kube-controller-manager	Self

Worker node(s)
Protocol	Direction	Port Range	Purpose	            Used By
TCP	      Inbound	  10250	      Kubelet API	        Self, Control plane
TCP	      Inbound	  30000-32767	NodePort Services†	All


7. Swap disabled
sudo swapon --show
free -h
sudo swapoff -a
sudo rm /swap.img

Remove following line from /etc/fstab
sudo nano /etc/fstab
/swap.img       none    swap    sw      0       0
ou comentar...
#/swap.img       none    swap    sw      0       0

Verificar o br_filter
lsmod | grep br_netfilter

Para carregar
sudo modprobe br_netfilter
Verifique novamente...

Garanta que a configuração net.bridge.bridge-nf-call-iptables do seu sysctl está configurada com valor 1
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system



```

### Primeiros_Passos_No_Linux_28SET22_B

```CMD
https://citizix.com/how-to-set-up-kubernetes-cluster-on-ubuntu-20-04-with-kubeadm-and-cri-o/
sudo apt update && sudo apt -y upgrade
sudo apt -y install curl apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt update && sudo apt -y install vim git curl wget kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
kubectl version --client && kubeadm version
Retorno:
WARNING: This version information is deprecated and will be replaced with the output from kubectl version --short.  Use --output=yaml|json to get the full version.
Client Version: version.Info{Major:"1", Minor:"25", GitVersion:"v1.25.2", GitCommit:"5835544ca568b757a8ecae5c153f317e5736700e", GitTreeState:"clean", BuildDate:"2022-09-21T14:33:49Z", GoVersion:"go1.19.1", Compiler:"gc", Platform:"linux/amd64"}
Kustomize Version: v4.5.7
kubeadm version: &version.Info{Major:"1", Minor:"25", GitVersion:"v1.25.2", GitCommit:"5835544ca568b757a8ecae5c153f317e5736700e", GitTreeState:"clean", BuildDate:"2022-09-21T14:32:18Z", GoVersion:"go1.19.1", Compiler:"gc", Platform:"linux/amd64"}

Desabilitar o SWAP
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
sudo swapoff -a

sudo modprobe overlay
sudo modprobe br_netfilter

sudo tee /etc/sysctl.d/kubernetes.conf<<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF

sudo sysctl --system

sudo -i
OS=xUbuntu_20.04
VERSION=1.22

echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/$OS/ /" > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
echo "deb http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/$VERSION/$OS/ /" > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable:cri-o:$VERSION.list

curl -L https://download.opensuse.org/repositories/devel:kubic:libcontainers:stable:cri-o:$VERSION/$OS/Release.key | apt-key add -
curl -L https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/$OS/Release.key | apt-key add -

sudo apt update && sudo apt install cri-o cri-o-runc

apt-cache policy cri-o
Retorno:
cri-o:
  Instalado: (nenhum)
  Candidato: 1.23.3~0
  Tabela de versão:
     1.23.3~0 500
        500 http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/1.23/xUbuntu_20.04  Packages
     1.22.5~0 500
        500 http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable:/cri-o:/1.22/xUbuntu_20.04  Packages

sudo systemctl daemon-reload
sudo systemctl enable crio --now

sudo systemctl enable kubelet

sudo kubeadm config images pull

kubeadm init options that are used to bootstrap cluster.
--control-plane-endpoint :  set the shared endpoint for all control-plane nodes. Can be DNS/IP
--pod-network-cidr : Used to set a Pod network add-on CIDR
--cri-socket : Use if have more than one container runtime to set runtime socket path
--apiserver-advertise-address : Set advertise address for this particular control-plane node's API server

sudo kubeadm init --pod-network-cidr=10.10.0.0/16

sudo nano /etc/hosts

sudo kubeadm init \
  --pod-network-cidr=10.10.0.0/16 \
  --upload-certs \
  --control-plane-endpoint=citizix.k8s.local

mkdir -p $HOME/.kube
sudo cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl cluster-info

kubectl taint nodes --all node-role.kubernetes.io/master-

kubectl taint nodes --all node-role.kubernetes.io/master-node/ubuntusrv.citizix.com untainted

kubectl create -f https://projectcalico.docs.tigera.io/manifests/tigera-operator.yaml

kubectl create -f https://projectcalico.docs.tigera.io/manifests/custom-resources.yaml

watch kubectl get pods -ncalico-system

kubectl get nodes -o wide

kubeadm join 10.2.40.239:6443 --token fegkje.9uu0g8ja0kqvhll1 \
	--discovery-token-ca-cert-hash sha256:9316503c53c0fd98daca54d314c2040a5a9690358055aeb2460872f1bd28ba78


kubectl create deploy nginx --image nginx:latest
deployment.apps/nginx created

kubectl get pods




https://blog.kubesimplify.com/kubernetes-125-dockerd

Here I have 4 instances in place

controlplane	192.168.1.106
worker1	      192.168.1.108
worker2	      192.168.1.110
worker3	      192.168.1.112

Step 1 - Run this on all the machines
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt update -y
sudo apt -y install vim git curl wget kubelet=1.25.0-00 kubeadm=1.25.0-00 kubectl=1.25.0-00
sudo apt-mark hold kubelet kubeadm kubectl

Disable swap
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
sudo swapoff -a

Load the br_netfilter module and let iptables see bridged traffic
sudo modprobe overlay && sudo modprobe br_netfilter

sudo tee /etc/sysctl.d/kubernetes.conf<<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
sudo sysctl --system

Setup Dockerd
sudo apt install docker.io -y
systemctl status docker
systemctl start docker
systemctl enable docker

Now, cri-dockerd setup
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.2.5/cri-dockerd-0.2.5.amd64.tgz
tar -xvf cri-dockerd-0.2.5.amd64.tgz
cd cri-dockerd/ && mkdir -p /usr/local/bin
sudo install -o root -g root -m 0755 ./cri-dockerd /usr/local/bin/cri-dockerd

Add the files cri-docker.socker cri-docker.service
sudo tee /etc/systemd/system/cri-docker.service << EOF
[Unit]
Description=CRI Interface for Docker Application Container Engine
Documentation=https://docs.mirantis.com
After=network-online.target firewalld.service docker.service
Wants=network-online.target
Requires=cri-docker.socket
[Service]
Type=notify
ExecStart=/usr/local/bin/cri-dockerd --container-runtime-endpoint fd:// --network-plugin=
ExecReload=/bin/kill -s HUP $MAINPID
TimeoutSec=0
RestartSec=2
Restart=always
StartLimitBurst=3
StartLimitInterval=60s
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
TasksMax=infinity
Delegate=yes
KillMode=process
[Install]
WantedBy=multi-user.target
EOF

sudo tee /etc/systemd/system/cri-docker.socket << EOF
[Unit]
Description=CRI Docker Socket for the API
PartOf=cri-docker.service
[Socket]
ListenStream=%t/cri-dockerd.sock
SocketMode=0660
SocketUser=root
SocketGroup=docker
[Install]
WantedBy=sockets.target
EOF

#Daemon reload
systemctl daemon-reload && systemctl enable cri-docker.service && systemctl enable --now cri-docker.socket

# Setup required sysctl params, these persist across reboots.
cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables  = 1
net.ipv4.ip_forward                 = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF

# Apply sysctl params without reboot
sudo sysctl --system

Configuring the kubelet cgroup driver
https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/
Pull the images - Pull the images for Kubernetes 1.25 version.
sudo kubeadm config images pull --cri-socket unix:///var/run/cri-dockerd.sock --kubernetes-version v1.25.0

Step 2 - Run the cluster init command on the control plane node
Here the pod network CIDR is dependent on the CNI you will be installing later on, so in this case, I am using flannel and --apiserver-advertise-address will be the public IP for the instance (it can be private IP as well but if you want to access it from outside of the node by using Kubeconfig then you need to give the public IP).

sudo kubeadm init   --pod-network-cidr=10.244.0.0/16   --upload-certs --kubernetes-version=v1.25.0   --control-plane-endpoint=192.168.1.106   --cri-socket unix:///var/run/cri-dockerd.sock

sudo kubeadm init   --pod-network-cidr=10.244.0.0/16   --upload-certs --kubernetes-version=v1.25.0   --control-plane-endpoint=192.168.1.108   --cri-socket unix:///var/run/cri-dockerd.sock



sudo kubeadm init   --pod-network-cidr=10.244.0.0/16   --upload-certs --kubernetes-version=v1.25.0   --control-plane-endpoint=192.168.1.106   --cri-socket unix:///var/run/cri-dockerd.sock --ignore-preflight-errors=All


sudo systemctl stop kubelet



netstat -ltnp | grep -w ":10250"
kill -9 pid


sudo systemctl stop kubelet.service


...erro...
systemctl status kubelet
journalctl -xeu kubelet
...acho que não é nada, o PC tá trabalhando...

Export KUBECONFIG and install CNI Flannel
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=/etc/kubernetes/admin.conf
kubectl apply -f https://github.com/coreos/flannel/raw/master/Documentation/kube-flannel.yml

...erro...
The connection to the server localhost:8080 was refused - did you specify the right host or port?

kube-flannel.yml

---
kind: Namespace
apiVersion: v1
metadata:
  name: kube-flannel
  labels:
    pod-security.kubernetes.io/enforce: privileged
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: flannel
rules:
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - nodes/status
  verbs:
  - patch
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: flannel
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: flannel
subjects:
- kind: ServiceAccount
  name: flannel
  namespace: kube-flannel
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: flannel
  namespace: kube-flannel
---
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-flannel-cfg
  namespace: kube-flannel
  labels:
    tier: node
    app: flannel
data:
  cni-conf.json: |
    {
      "name": "cbr0",
      "cniVersion": "0.3.1",
      "plugins": [
        {
          "type": "flannel",
          "delegate": {
            "hairpinMode": true,
            "isDefaultGateway": true
          }
        },
        {
          "type": "portmap",
          "capabilities": {
            "portMappings": true
          }
        }
      ]
    }
  net-conf.json: |
    {
      "Network": "10.244.0.0/16",
      "Backend": {
        "Type": "vxlan"
      }
    }
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: kube-flannel-ds
  namespace: kube-flannel
  labels:
    tier: node
    app: flannel
spec:
  selector:
    matchLabels:
      app: flannel
  template:
    metadata:
      labels:
        tier: node
        app: flannel
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kubernetes.io/os
                operator: In
                values:
                - linux
      hostNetwork: true
      priorityClassName: system-node-critical
      tolerations:
      - operator: Exists
        effect: NoSchedule
      serviceAccountName: flannel
      initContainers:
      - name: install-cni-plugin
       #image: flannelcni/flannel-cni-plugin:v1.1.0 for ppc64le and mips64le (dockerhub limitations may apply)
        image: docker.io/rancher/mirrored-flannelcni-flannel-cni-plugin:v1.1.0
        command:
        - cp
        args:
        - -f
        - /flannel
        - /opt/cni/bin/flannel
        volumeMounts:
        - name: cni-plugin
          mountPath: /opt/cni/bin
      - name: install-cni
       #image: flannelcni/flannel:v0.19.2 for ppc64le and mips64le (dockerhub limitations may apply)
        image: docker.io/rancher/mirrored-flannelcni-flannel:v0.19.2
        command:
        - cp
        args:
        - -f
        - /etc/kube-flannel/cni-conf.json
        - /etc/cni/net.d/10-flannel.conflist
        volumeMounts:
        - name: cni
          mountPath: /etc/cni/net.d
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
      containers:
      - name: kube-flannel
       #image: flannelcni/flannel:v0.19.2 for ppc64le and mips64le (dockerhub limitations may apply)
        image: docker.io/rancher/mirrored-flannelcni-flannel:v0.19.2
        command:
        - /opt/bin/flanneld
        args:
        - --ip-masq
        - --kube-subnet-mgr
        resources:
          requests:
            cpu: "100m"
            memory: "50Mi"
          limits:
            cpu: "100m"
            memory: "50Mi"
        securityContext:
          privileged: false
          capabilities:
            add: ["NET_ADMIN", "NET_RAW"]
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: EVENT_QUEUE_DEPTH
          value: "5000"
        volumeMounts:
        - name: run
          mountPath: /run/flannel
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
        - name: xtables-lock
          mountPath: /run/xtables.lock
      volumes:
      - name: run
        hostPath:
          path: /run/flannel
      - name: cni-plugin
        hostPath:
          path: /opt/cni/bin
      - name: cni
        hostPath:
          path: /etc/cni/net.d
      - name: flannel-cfg
        configMap:
          name: kube-flannel-cfg
      - name: xtables-lock
        hostPath:
          path: /run/xtables.lock
          type: FileOrCreate


source /usr/share/bash-completion/bash_completion
echo 'source <(kubectl completion bash)' >>~/.bashrc
echo 'alias k=kubectl' >>~/.bashrc
echo 'complete -o default -F __start_kubectl k' >>~/.bashrc
exec bash

sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update && sudo apt-get install -y kubelet kubeadm kubectl && sudo apt-mark hold kubelet kubeadm kubectl

export KUBERNETES_MASTER=http://192.168.1.106:8080



nano sed -i "s/cgroupDriver: systemd/cgroupDriver: cgroupfs/g" /var/lib/kubelet/config.yaml
systemctl daemon-reload
systemctl restart kubelet

sudo kubeadm init   --pod-network-cidr=10.244.0.0/16   --upload-certs --kubernetes-version=v1.25.0   --control-plane-endpoint=192.168.1.106   --cri-socket unix:///var/run/cri-dockerd.sock --ignore-preflight-errors=All

sudo nano /var/lib/kubelet/config.yaml
...
#cgroupDriver: systemd
cgroupDriver: cgroupfs
...
Ctrl+O, Enter, Ctrl+X


nano sed -i "s/cgroupDriver: systemd/cgroupDriver: cgroupfs/g" /var/lib/kubelet/config.yaml
systemctl daemon-reload
systemctl restart kubelet


sudo systemctl stop kubelet
sudo rm -rf /etc/kubernetes/manifests/

sudo kubeadm init   --pod-network-cidr=10.244.0.0/16   --upload-certs --kubernetes-version=v1.25.0   --control-plane-endpoint=192.168.1.106   --cri-socket unix:///var/run/cri-dockerd.sock 




Step 3 - Run the join command on all the worker nodes
Remember to add the --cri-socket flag at the end

kubeadm join 192.168.1.106:6443 --token 6gh7gq.yxxvl9c0tjauu7up       --discovery-token-ca-cert-hash sha256:e3ecc16a7c7fa9ccf3c334e98bd53c18c86e9831984f1f7c8398fbd54d5e37e9  --cri-socket unix:///var/run/cri-dockerd.sock
Retorno:
[preflight] Running pre-flight checks
    [WARNING SystemVerification]: missing optional cgroups: blkio
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.

Step 4 - Nginx Test
You can copy the kubeconfig file from the controlplane node(~/.kube/config ) to local and export the KUBECONFIG variable or directly access the cluster from the controlplane node.

kubectl get nodes
NAME                       STATUS   ROLES           AGE   VERSION
kube1-25-cp-eb85-c20f1a    Ready    control-plane   13m   v1.25.0
kube1-25-w-1-b507-7ebeb0   Ready    <none>          64s   v1.25.0
kube1-25-w-2-0888-7ebeb0   Ready    <none>          54s   v1.25.0
kube1-25-w-3-142f-7ebeb0   Ready    <none>          51s   v1.25.0


The cluster is up and running with single controlplane and 3 worker nodes.

Now run nginx

kubectl run nginx --image=nginx
pod/nginx created

kubectl expose pod nginx --type=NodePort --port 80
service/nginx exposed

kubectl get pods
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          8s

kubectl get svc nginx
NAME    TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
nginx   NodePort   10.104.25.205   <none>        80:32743/TCP   8s


Access the service using Node public IP:32743(make sure your firewall rules are properly set to allow traffic to required ports)

YAY!! you have successfully setup a self managed Kubernetes cluster, version 1.25.0 and dockerd as the container runtime.
https://www.youtube.com/watch?v=V_hzP_nEOkI

https://blog.kubesimplify.com/kubernetes-crio
https://blog.kubesimplify.com/kubernetes-containerd-setup
https://gist.github.com/saiyam1814/801db1288c690a969e7174eca89c65b2
https://gist.github.com/saiyam1814/c25c100c93c8b3f38c5e3b8bc75b531b

```

### Primeiros_Passos_No_Linux_27SET22

```CMD
https://www.youtube.com/watch?v=rD6pU_b6iFg

Here I have 4 instances in place

controlplane	192.168.1.106
worker1	      192.168.1.108
worker2	      192.168.1.110
worker3	      192.168.1.112
_______________________________________________________________________________
1. Em todas as máquinas:
Desabilitar o swap
sudo swapoff -a

Alterar o FSTAB
sudo nano /etc/fstab
Comenta a linha do swap

Habilitar os módulos
sudo modprobe overlay
sudo modprobe br_netfilter

Incluir no arquivo
sudo nano /etc/modules-load.d/modules.conf

adicionar:
overlay
br_netfilter
Ctrl+O, Enter, Ctrl+X

Adicionar as configurações do sysctl
sudo nano /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
Ctrl+O, Enter, Ctrl+X
sudo sysctl --system

Instalar o containerd:
sudo apt update
sudo apt install containerd

sudo mkdir /etc/containerd
containerd config default
sudo su
containerd config default > /etc/containerd/config.tolm
systemctl restart containerd
systemctl status containerd

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -

apt-add-repository "deb http://apt.kubernetes.io/ kubernetes-xenial main"
apt update
apt install -y kubeadm kubelet kubectl

_______________________________________________________________________________
Só no control-plane:
kubeadm init --apiserver-advertise-address=192.168.1.106
... deu errado...

```

### 30SET22

```CMD
https://www.cloudsigma.com/how-to-install-and-use-kubernetes-on-ubuntu-20-04/
TODOS OS PCS
sudo su (depois testar sem)

Step 1: Install Kubernetes
sudo apt install apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" >> ~/kubernetes.list
sudo mv ~/kubernetes.list /etc/apt/sources.list.d
sudo apt update
sudo apt install kubelet
sudo apt install kubeadm
sudo apt install kubectl
sudo apt-get install -y kubernetes-cni
sudo apt-get install -y kubelet kubeadm kubectl kubernetes-cni

Step 2: Disabling Swap Memory
sudo swapoff -a
sudo nano /etc/fstab
Comentar a linha swapfile

Step 3: Setting Unique Hostnames
No master:
sudo hostnamectl set-hostname kubernetes-master

Nos workers:
sudo hostnamectl set-hostname kubernetes-worker-1
sudo hostnamectl set-hostname kubernetes-worker-2
sudo hostnamectl set-hostname kubernetes-worker-3

Step 4: Letting Iptables See Bridged Traffic
lsmod | grep br_netfilter
sudo modprobe br_netfilter
sudo sysctl net.bridge.bridge-nf-call-iptables=1

Step 5: Changing Docker Cgroup Driver
sudo mkdir /etc/docker
cat <<EOF | sudo tee /etc/docker/daemon.json
{ "exec-opts": ["native.cgroupdriver=systemd"],
"log-driver": "json-file",
"log-opts":
{ "max-size": "100m" },
"storage-driver": "overlay2"
}
EOF

Instalar docker em https://renatomportugal.github.io/03.docker/#/Instalacao?id=no-ubuntu

sudo systemctl status docker
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker

Step 6: Initializing the Kubernetes Master Node (Só no MASTER)
sudo kubeadm init --pod-network-cidr=10.244.0.0/16

ignorando os erros:
sudo kubeadm init --ignore-preflight-errors=NumCPU,Mem --pod-network-cidr=10.244.0.0/16

erro...
[ERROR CRI]: container runtime is not running

rm /etc/containerd/config.toml
systemctl restart containerd
kubeadm init

...funcionou...
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.1.106:6443 --token p1gwq1.hpe7osrsplalb8ne \
	--discovery-token-ca-cert-hash sha256:46c1f7cc1d1e6e269857ff5505604deb02465f81c50aeb69f4780ebe0f221ce9 


mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

Step 7: Deploying a Pod Network
sudo ufw allow 6443
sudo ufw allow 6443/tcp

After that, you can run the following two commands to deploy the pod network on the master node:
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/k8s-manifests/kube-flannel-rbac.yml

kubectl get pods --all-namespaces

kubectl get componentstatus
kubectl get cs

If you see 
the unhealthy status, modify the following files and delete the line at (spec->containers->command) containing this phrase - --port=0 :
sudo nano /etc/kubernetes/manifests/kube-scheduler.yaml
	
sudo nano /etc/kubernetes/manifests/kube-scheduler.yaml

Do the same for this file:
sudo nano /etc/kubernetes/manifests/kube-controller-manager.yaml
	
sudo nano /etc/kubernetes/manifests/kube-controller-manager.yaml

Finally, restart the Kubernetes service:
sudo systemctl restart kubelet.service
	
sudo systemctl restart kubelet.service

Step 8: Joining Worker Nodes to the Kubernetes Cluster
kubeadm join 192.168.1.106:6443 --token p1gwq1.hpe7osrsplalb8ne \
	--discovery-token-ca-cert-hash sha256:46c1f7cc1d1e6e269857ff5505604deb02465f81c50aeb69f4780ebe0f221ce9 

No master:
kubectl get nodes

...deu erro nos workers...


```

### 03OUT22_Lubuntu

```CMD
1. Ubuntu
cat /etc/os-release
PRETTY_NAME="Ubuntu 22.04.1 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.1 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy

2. 2 GB or more of RAM
htop
Worker-node 1.92GB

3. 2 CPUs or more
htop
4

4. Full network connectivity between all machines in the cluster
ip link
ifconfig -a

enp1s0

5. Unique hostname, MAC address, and product_uuid for every node
cat /proc/sys/kernel/hostname
ifconfig
tcnct-one

5.
sudo cat /sys/class/dmi/id/product_uuid
af1c0701-e915-f542-95bb-089e01434239

6. Certain ports are open on your machines
nc 127.0.0.1 6443
nc 127.0.0.1 10250

7. Swap disabled
sudo swapon --show
free -h
sudo swapoff -a
sudo rm /swap.img

Remove following line from /etc/fstab
sudo nano /etc/fstab
/swap.img       none    swap    sw      0       0
ou comentar...
#/swap.img       none    swap    sw      0       0

Verificar o br_filter
lsmod | grep br_netfilter

Para carregar
sudo modprobe br_netfilter
Verifique novamente...

Garanta que a configuração net.bridge.bridge-nf-call-iptables do seu sysctl está configurada com valor 1
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add
sudo apt-add-repository "deb http://apt.kubernetes.io/ kubernetes-xenial main"
sudo apt install kubeadm
kubeadm version

Implantação do Kubernetes
sudo hostnamectl set-hostname kubernetes-master

Nos workers
sudo hostnamectl set-hostname kubernetes-worker-1

kubectl get nodes

kubectl get pods


Step 1: Install Kubernetes
sudo apt install apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" >> ~/kubernetes.list
sudo mv ~/kubernetes.list /etc/apt/sources.list.d
sudo apt update
sudo apt install kubelet
sudo apt install kubeadm
sudo apt install kubectl
sudo apt-get install -y kubernetes-cni

sudo apt update -y && sudo apt-get install -y kubelet kubeadm kubectl kubernetes-cni

Step 2: Disabling Swap Memory
sudo swapoff -a
sudo nano /etc/fstab
Comentar a linha swapfile

Step 3: Setting Unique Hostnames
No master:
sudo hostnamectl set-hostname kubernetes-master

Nos workers:
sudo hostnamectl set-hostname kubernetes-worker-1
sudo hostnamectl set-hostname kubernetes-worker-2
sudo hostnamectl set-hostname kubernetes-worker-3

Step 4: Letting Iptables See Bridged Traffic
lsmod | grep br_netfilter
sudo modprobe br_netfilter
sudo sysctl net.bridge.bridge-nf-call-iptables=1

Step 5: Changing Docker Cgroup Driver
sudo mkdir /etc/docker
cat <<EOF | sudo tee /etc/docker/daemon.json
{ "exec-opts": ["native.cgroupdriver=systemd"],
"log-driver": "json-file",
"log-opts":
{ "max-size": "100m" },
"storage-driver": "overlay2"
}
EOF

Instalar docker em https://renatomportugal.github.io/03.docker/#/Instalacao?id=no-ubuntu

sudo systemctl status docker
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker

Step 6: Initializing the Kubernetes Master Node (Só no MASTER)
sudo kubeadm init --pod-network-cidr=10.244.0.0/16

ignorando os erros:
sudo kubeadm init --ignore-preflight-errors=NumCPU,Mem --pod-network-cidr=10.244.0.0/16

erro...
[ERROR CRI]: container runtime is not running

sudo rm /etc/containerd/config.toml
systemctl restart containerd
kubeadm init


```

### 04OUT22

```CMD
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

1.Download the latest release with the command:
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

...ou...
curl -LO https://dl.k8s.io/release/v1.25.0/bin/linux/amd64/kubectl

Validate the binary (optional)
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check

Install kubectl
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

If you do not have root access on the target system, you can still install kubectl to the ~/.local/bin directory:
chmod +x kubectl
mkdir -p ~/.local/bin
mv ./kubectl ~/.local/bin/kubectl
# and then append (or prepend) ~/.local/bin to $PATH

kubectl version --client
kubectl version --client --output=yaml

Verify kubectl configuration
kubectl cluster-info

kubectl cluster-info dump


source /usr/share/bash-completion/bash_completion
echo 'source <(kubectl completion bash)' >>~/.bashrc

echo 'alias k=kubectl' >>~/.bashrc
echo 'complete -o default -F __start_kubectl k' >>~/.bashrc

exec bash


Install kubectl convert plugin
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl-convert"

curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl-convert.sha256"

echo "$(cat kubectl-convert.sha256) kubectl-convert" | sha256sum --check

Install kubectl-convert
sudo install -o root -g root -m 0755 kubectl-convert /usr/local/bin/kubectl-convert
kubectl convert --help



```

```CMD
https://github.com/franciscojsc/scripts-install-kubernetes
master

sudo apt-get update

sudo swapoff -a
sudo cp /etc/fstab /etc/fstab.bkp
sudo sed -i.bak '/ swap / s/^\(.*\)$/#/g' /etc/fstab

Instalar o docker
https://renatomportugal.github.io/03.docker/#/Instalacao?id=no-ubuntu

docker ps -a

sudo su
echo '{"exec-opts": ["native.cgroupdriver=systemd"],"log-driver": "json-file","log-opts": {"max-size": "100m"},"storage-driver": "overlay2"}' > /etc/docker/daemon.json

mkdir -p /etc/systemd/system/docker.service.d

systemctl daemon-reload && systemctl restart docker

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/k8s.list

apt-get update

apt-get -y install kubectl && apt-get -y install kubeadm && apt-get -y install kubelet

kubeadm config images pull
kubeadm init

mkdir -p $HOME/.kube 
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
chown  $(id -u):$(id -g)  $HOME/.kube/config

echo 
echo "**** autocompletion kubectl ****"
echo 

echo "source <(kubectl completion bash)" >> $HOME/.bashrc

echo 
"**** pod network - weave net ****"
echo 

kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"

echo 
echo "**** install helm ****"
echo 

curl -L https://git.io/get_helm.sh | bash

helm init

kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'

echo 
echo "**** view status cluster ****"
echo

kubectl get nodes,svc,deploy,rs,rc,po -o wide

echo 
echo "**** add node worker with token ****"
echo 

kubeadm token create --print-join-command

echo 
echo "finish install"
```

```CMD
https://github.com/franciscojsc/scripts-install-kubernetes
WORKER NODE

sudo apt-get update

swapoff -a
sudo cp /etc/fstab /etc/fstab.bkp
sudo sed -i.bak '/ swap / s/^\(.*\)$/#/g' /etc/fstab

sudo su

sudo nano /etc/docker/daemon.json

{ "exec-opts": ["native.cgroupdriver=systemd"],
"log-driver": "json-file",
"log-opts":
{ "max-size": "100m" },
"storage-driver": "overlay2"
}
Ctrl+O, Enter, Ctrl+X

echo '{"exec-opts": ["native.cgroupdriver=systemd"],"log-driver": "json-file","log-opts": {"max-size": "100m"},"storage-driver": "overlay2"}' > /etc/docker/daemon.json

mkdir -p /etc/systemd/system/docker.service.d

systemctl daemon-reload
systemctl restart docker

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/k8s.list
exit

...ou...

sudo nano /etc/apt/sources.list.d/k8s.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
Ctrl+O, Enter, Ctrl+X

sudo apt-get update

sudo apt-get -y install kubectl && sudo apt-get -y install kubeadm && sudo apt-get -y install kubelet
 
sudo su
echo "source <(kubectl completion bash)" >> $HOME/.bashrc
exit
kubectl version && kubeadm version && kubelet --version

```

### 05OUT22

```CMD
Instalar o docker...
sudo apt install apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add

sudo apt-add-repository "deb http://apt.kubernetes.io/ kubernetes-xenial main"

sudo apt install kubeadm kubelet kubectl kubernetes-cni
sudo swapoff -a

sudo nano /etc/fstab.conf
comentar a linha do swap

sudo hostnamectl set-hostname kubernetes-master


E no nó do trabalhador:
sudo hostnamectl set-hostname kubernetes-worker

sudo kubeadm init

sudo nano /etc/containerd/config.toml


Instalar cri-dockerd
sudo apt install golang-go
git clone https://github.com/Mirantis/cri-dockerd.git
cd cri-dockerd
mkdir bin
VERSION=$((git describe --abbrev=0 --tags | sed -e 's/v//') || echo $(cat VERSION)-$(git log -1 --pretty='%h')) PRERELEASE=$(grep -q dev <<< "${VERSION}" && echo "pre" || echo "") REVISION=$(git log -1 --pretty='%h')

# Run these commands as root
###Install GO###
wget https://storage.googleapis.com/golang/getgo/installer_linux
chmod +x ./installer_linux
./installer_linux
source ~/.bash_profile

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


```
