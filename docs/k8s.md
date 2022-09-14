# K8s

## Primeiros_Passos

```CMD
Kind e Minikube não devem ser utilizados para produção.

https://microk8s.io/
MicroK8S - Produção, Edge Computing e IoT (Internet of things)

https://k3s.io/
k3s (Lightweight) - Produção, Edge, IoT (Internet of things), CI, ARM.

https://k0sproject.io/
k0s - Produção x86-64, ARM64, ARMv7.
```

## Primeiros_Passos_No_Windows

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

## Primeiros_Passos_No_Linux_Com_K3S

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

## Outra_Instalacao

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

## Mais_Uma_Tentativa_22JUN22

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

## Testar_Depois

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

## Primeiros_Passos_No_Linux_13SET22

```CMD
https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
1. Ubuntu 20.04
2. 2 GB or more of RAM
3. 2 CPUs or more
4. Full network connectivity between all machines in the cluster
5. Unique hostname, MAC address, and product_uuid for every node
6. Certain ports are open on your machines
7. Swap disabled

4.
ip link
ifconfig -a

5.
sudo cat /sys/class/dmi/id/product_uuid

6.
nc 127.0.0.1 6443
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


7. Desabilitar Swap
sudo swapon --show
free -h
sudo swapoff -a
sudo rm /swap.img

Remove following line from /etc/fstab
sudo nano /etc/fstab
/swap.img       none    swap    sw      0       0
ou comentar...
#/swap.img       none    swap    sw      0       0


```

## Primeiros_Passos_No_Linux_14SET22

```CMD
https://blog.kubesimplify.com/kubernetes-125-dockerd

Here I have 4 instances in place

controlplane	74.220.19.161
worker1	      74.220.18.110
worker2	      74.220.22.48
worker3	      74.220.23.63

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
sudo modprobe overlay
sudo modprobe br_netfilter
sudo tee /etc/sysctl.d/kubernetes.conf<<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF
sysctl --system

Setup Dockerd
apt install docker.io -y
systemctl start docker
systemctl enable docker

Now, cri-dockerd setup
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.2.5/cri-dockerd-0.2.5.amd64.tgz
tar -xvf cri-dockerd-0.2.5.amd64.tgz
cd cri-dockerd/
mkdir -p /usr/local/bin
install -o root -g root -m 0755 ./cri-dockerd /usr/local/bin/cri-dockerd

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
systemctl daemon-reload
systemctl enable cri-docker.service
systemctl enable --now cri-docker.socket

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

sudo kubeadm init   --pod-network-cidr=10.244.0.0/16   --upload-certs --kubernetes-version=v1.25.0   --control-plane-endpoint=74.220.19.161   --cri-socket unix:///var/run/cri-dockerd.sock

Export KUBECONFIG and install CNI Flannel
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=/etc/kubernetes/admin.conf
kubectl apply -f https://github.com/coreos/flannel/raw/master/Documentation/kube-flannel.yml

Step 3 - Run the join command on all the worker nodes
Remember to add the --cri-socket flag at the end

kubeadm join 74.220.19.161:6443 --token 6gh7gq.yxxvl9c0tjauu7up       --discovery-token-ca-cert-hash sha256:e3ecc16a7c7fa9ccf3c334e98bd53c18c86e9831984f1f7c8398fbd54d5e37e9  --cri-socket unix:///var/run/cri-dockerd.sock
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
