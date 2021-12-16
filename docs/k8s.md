# K8s

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

## Primeiros_Passos_No_Linux

```CMD
Para instalar:

curl -sfL https://get.k3s.io | sh -s - --write-kubeconfig-mode 644
ou
curl -sfL https://get.k3s.io | K3S_KUBECONFIG_MODE="644" sh -s -

PULEI ESSAS PARTES POR CAUSA DO COMANDO DE CIMA
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl

kubectl version --client

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
K10c79bf801e14c84cbab06422b2d5454de09afad32b743d232169dae4df5ff079f::server:a1899415ac5e4c60829ed45481658c79

ifconfig
172.17.128.10

ADICIONAR OUTROS NÓS


curl -sfL https://get.k3s.io | K3S_URL=https://172.17.128.10:6443 K3S_TOKEN=K10c79bf801e14c84cbab06422b2d5454de09afad32b743d232169dae4df5ff079f::server:a1899415ac5e4c60829ed45481658c79 sh -

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
