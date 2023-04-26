# Instalação do Docker

https://docs.docker.com/engine/install/ubuntu/<br>

## Verificar se já está instalado <!-- {docsify-ignore} -->

```CMD
docker -v
```

## No CentOS 7

### Instalação do Sistema Operacional <!-- {docsify-ignore} -->

```CMD
Baixar a ISO

http://isoredirect.centos.org/centos/7/isos/x86_64/
Exclua as partições do Disco
Selecione Servidor com GUI
Bibliotecas de Compatibilidade
Ferramentas de Administração de Sistema
https://docs.docker.com/engine/install/centos/
```

### Passo 00 <!-- {docsify-ignore} -->

```CMD
Configuração do proxy para empresas que usam. Residencial e conexões comuns nem precisa.
sudo nano /etc/yum.conf
proxy=http://proxy.example.com:3128 
proxy_username=yum-user 
proxy_password=qwerty

Ctrl+O, Enter
Ctrl+x

```

### Passo 01 <!-- {docsify-ignore} -->
Setup do repositório<br>
```
sudo yum install -y yum-utils

sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```

### Passo 02 <!-- {docsify-ignore} -->
Instalar Docker Engine<br>
```
sudo yum install docker-ce docker-ce-cli containerd.io
```
Veja se a fingerprint está correta e aceite<br>
```
060A 61C5 1B55 8A7F 742B 77AA C52F EB6B 621E 9F35,
```
### Passo 03 <!-- {docsify-ignore} -->
Inicie o Docker<br>
```
sudo systemctl start docker
```
### Passo 04 <!-- {docsify-ignore} -->
Inclua o usuário atual no grupo Docker<br>
```
whoami
sudo usermod -aG docker usuario
```
### Passo 05 <!-- {docsify-ignore} -->
Reiniciar o sistema.<br>
```
reboot
```
### Passo 06 <!-- {docsify-ignore} -->
Fazer o Docker iniciar coma máquina:<br>
Verificar se o serviço docker está iniciado:<br>
```
sudo systemctl status docker
```

Iniciar o serviço:<br>
```
sudo systemctl start docker
sudo systemctl status docker
docker ps
```
Habilitar o docker com o início da máquina:<br>
```
sudo systemctl enable docker
sudo systemctl status docker
```

## Passo07 <!-- {docsify-ignore} -->
Instalar o docker-compose<br>
### Instalação Normal <!-- {docsify-ignore} --> 

https://docs.docker.com/compose/install/<br>

Run this command to download the current stable release of Docker Compose:<br>

```CMD
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

Apply executable permissions to the binary:<br>

```CMD
sudo chmod +x /usr/local/bin/docker-compose
```

Testar<br>

```CMD
docker-compose --version
```

Install command completion<br>

```CMD
cd /etc/bash_completion.d/
sudo curl -L https://raw.githubusercontent.com/docker/compose/1.26.0/contrib/completion/bash/docker-compose -o /etc/bash_completion.d/docker-compose
```

#### Possíveis erros <!-- {docsify-ignore} -->
Note: If the command docker-compose fails after installation, check your path. You can also create a symbolic link to /usr/bin or any other directory in your path.<br>
```
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

### Instalação alternativa <!-- {docsify-ignore} -->
```
https://github.com/docker/compose/releases/

Estando na pasta, copie para /usr/local/bin/docker-compose
sudo cp ./docker-compose-Linux-x86_64 /usr/local/bin/docker-compose
```
Apply executable permissions to the binary:<br>
```
sudo chmod +x /usr/local/bin/docker-compose
```
Testar<br>
```
docker-compose --version
```

Install command completion<br>
```
Acesse https://raw.githubusercontent.com/docker/compose/1.27.3/contrib/completion/bash/docker-compose e salve como docker-compose-completion em qualquer pasta.
sudo cp docker-compose-completion /etc/bash_completion.d/docker-compose
```


## Passo 08 <!-- {docsify-ignore} -->
Configurar o proxy (se não usar não faça)<br>
```
sudo nano /etc/sysconfig/docker
```

Insira no arquivo:<br>
```
HTTP_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA
HTTPS_PROXY=https://USUARIO:SENHA@proxy.DOMINIO:PORTA
NO_PROXY=localhost, 127.0.0.1, localaddress,.localdomain.com
```
Para salvar e sair:<br>
```
Ctrl+O, Enter, Ctrl+X
```
Criar a pasta docker.service.d e o arquivo http-proxy.conf<br>
```
cd /etc/systemd/system
sudo mkdir docker.service.d
cd docker.service.d
sudo nano http-proxy.conf
```
Adicione as linhas no arquivo:<br>
```
[Service]
Environment="HTTP_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA"
Environment="HTTPS_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA"
Environment="NO_PROXY=localhost, 127.0.0.1, 10.*"
```
Para salvar e sair:<br>
```
Ctrl+O, Enter, Ctrl+X
```
Reiniciar os serviços:<br>
```
sudo systemctl daemon-reload
sudo systemctl restart docker
```
Para verificar se está tudo em ordem:<br>
```
systemctl show --property=Environment docker
```
Para testar, deverá ter sucesso ao baixar:<br>
```
docker pull alpine
```


## No CentOS 8
Estou como usuário root, para usuários comuns utilize o sudo na frente dos comandos que não funcionarem.<br>
Verificar versão:<br>
```
cat /etc/centos-release
```
## Passo 01 <!-- {docsify-ignore} -->
Preparar o proxy.<br>
```
nano /etc/dnf/dnf.conf
```
Veja o arquivo e ajuste:<br>
```
[main]
gpgcheck=1
installonly_limit=3
clean_requirements_on_remove=True
best=True
skip_if_unavailable=False
proxy=http://proxy.empresa.com:porta
proxy_username=user
proxy_password=pass
proxy_auth_method=basic
sslverify=false
```
Para gravar Ctrl+O, Enter, e para sair Ctrl+X<br>

## Passo 02 <!-- {docsify-ignore} -->
Update do sistema.<br>
```
dnf update
dnf check-update
dnf clean all
```

## Passo 03 <!-- {docsify-ignore} -->
Remover as versões anteriores.<br>
```
dnf remove docker \
    docker-client \
    docker-client-latest \
    docker-common \
    docker-latest \
    docker-latest-logrotate \
    docker-logrotate \
    docker-engine
```
## Passo 04 <!-- {docsify-ignore} -->
Adicionar o repositório:<br>
```
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
dnf repolist -v
```
## Passo 05 <!-- {docsify-ignore} -->
Instalar o Docker:<br>
```
dnf install docker-ce docker-ce-cli containerd.io --nobest
```

## Passo 06 <!-- {docsify-ignore} -->
Verificar:<br>
```
docker -v
```

## Passo 07 <!-- {docsify-ignore} -->
Colocar o usuário no grupo Docker:<br>
```
whoami
usermod -aG docker USER
```
Verificar:<br>
```
id USER
```

## Passo 08 <!-- {docsify-ignore} -->
Iniciar com a máquina:<br>
```
systemctl status docker
systemctl enable docker
systemctl enable docker.service
systemctl start docker
systemctl status docker
reboot
```
## Passo 09 <!-- {docsify-ignore} -->
Configurar o proxy (se não usar não faça)<br>
```
nano /etc/sysconfig/docker
```

Insira no arquivo:<br>
```
HTTP_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA
HTTPS_PROXY=https://USUARIO:SENHA@proxy.DOMINIO:PORTA
NO_PROXY=localhost, 127.0.0.1, localaddress,.localdomain.com
```
Para salvar e sair:<br>
```
Ctrl+O, Enter, Ctrl+X
```
Criar a pasta docker.service.d e o arquivo http-proxy.conf<br>
```
cd /etc/systemd/system
mkdir docker.service.d
cd docker.service.d
nano http-proxy.conf
```
Adicione as linhas no arquivo:<br>
```
[Service]
Environment="HTTP_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA"
Environment="HTTPS_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA"
Environment="NO_PROXY=localhost, 127.0.0.1, 10.*"
```
Para salvar e sair:<br>
```
Ctrl+O, Enter, Ctrl+X
```
Reiniciar os serviços:<br>
```
systemctl daemon-reload
systemctl restart docker
```
Para verificar se está tudo em ordem:<br>
```
systemctl show --property=Environment docker
```
Para testar, deverá ter sucesso ao baixar:<br>
```
docker pull alpine
```

## Passo 10 <!-- {docsify-ignore} -->
Instalar o docker-compose<br>
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.27.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```
### Método alternativo <!-- {docsify-ignore} -->
```
wget https://github.com/docker/compose/releases/download/1.27.1/docker-compose-Linux-x86_64
```
Se não conseguir, baixe pelo navegador e suba via ftp para qualquer pasta do servidor:<br>
```
https://github.com/docker/compose/releases/download/1.27.1/docker-compose-Linux-x86_64
```
Renomeie o pacote:<br>
```
mv docker-compose-Linux-x86_64 docker-compose
```
Copiar para a pasta correta e fazer executável:<br>
```
sudo mv docker-compose /usr/local/bin/
sudo chmod +x /usr/local/bin/docker-compose
```
Para testar:<br>
```
docker-compose --help
```
## Passo 11 <!-- {docsify-ignore} -->
Desabilitar firewall (caso esteja ativo)<br>
Este passo é para quem não precisa. Assim que possível pesquisarei sobre o assunto.<br>
```
service firewalld stop
systemctl disable firewalld
```
## Passo 12 <!-- {docsify-ignore} -->
Reiniciar serviços de Daemon e docker.<br>
```
systemctl daemon-reload
systemctl restart docker
```

## No Ubuntu

https://docs.docker.com/engine/install/ubuntu/<br>

### Passo U00 <!-- {docsify-ignore} -->

```CMD
Se precisar configure o proxy

sudo nano /etc/environment

Adicione no fim do arquivo:
http_proxy=http://USUARIO:SENHA@proxy.DOMINIO:PORTA
https_proxy=http://USUARIO:SENHA@proxy.DOMINIO:PORTA
ftp_proxy=http://USUARIO:SENHA@proxy.DOMINIO:PORTA
no_proxy=10.*

Ctrl+O, Enter, Ctrl+X

```

###  Passo U01 <!-- {docsify-ignore} -->

```CMD
Atualizar o sistema
sudo apt-get update

Desinstalar
Anote a quantidade de espaço que foi usado
df -h
No meu caso é:
/dev/mapper/ubuntu--vg-ubuntu--lv   63%

Desisntalando o Docker-compose:
docker-compose --version
sudo rm -rf /usr/local/bin/docker-compose
docker-compose --version
Retorno:
-bash: /usr/local/bin/docker-compose: No such file or directory

Remover o docker:
docker --version
sudo apt-get remove docker docker-engine docker.io containerd runc
sudo apt-get purge docker-ce docker-ce-cli containerd.io docker-compose-plugin
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd

df -h
Retorno:
/dev/mapper/ubuntu--vg-ubuntu--lv   36G   19G   16G  56% /

```
  
### Passo U02 <!-- {docsify-ignore} -->

Copiar e colar todas essas linhas (-y aceita tudo por padrão)<br>
```
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common -y
```

### Passo U03 <!-- {docsify-ignore} -->

Adicionar a chave oficial do Docker<br>
```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

#### Método alternativo <!-- {docsify-ignore} -->
Baixar o arquivo (https://download.docker.com/linux/ubuntu/gpg) no Browser ou em outro computador.<br>
Instalar o programa WinSCP (https://winscp.net/eng/download.php)<br>
Conectar via FTP no PC. Transferir o arquivo para a pasta do usuario que está usando.<br>
```
cd ~
sudo apt-key add gpg
```

```
gpg --list-keys
gpg --import gpg
gpg --list-keys

Aparecerá
------------------------------
pub   rsa4096 2017-02-22 [SCEA]
      9DC858229FC7DD38854AE2D88D81803C0EBFCD88
uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
```

### Passo 04 <!-- {docsify-ignore} -->
Verificar a impressão digital<br>
```
sudo apt-key fingerprint 0EBFCD88
```
Tela de retorno:<br>
```
pub   rsa4096 2017-02-22 [SCEA]
      9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
```

### Passo 05 <!-- {docsify-ignore} -->
Setup do repositório<br>
```
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

...erro...
Traceback (most recent call last):
  File "/usr/bin/add-apt-repository", line 363, in <module>
    addaptrepo = AddAptRepository()
  File "/usr/bin/add-apt-repository", line 41, in __init__
    self.distro.get_sources(self.sourceslist)
  File "/usr/lib/python3/dist-packages/aptsources/distro.py", line 91, in get_sources
    raise NoDistroTemplateException(
aptsources.distro.NoDistroTemplateException: Error: could not find a distribution template for Ubuntu/jammy

Adicione manualmente (sem add-apt-repository):
echo "deb [arch=amd64] https://download.docker.com/linux/ubuntu groovy stable" |sudo tee /etc/apt/sources.list.d/docker-ce.list
...resolveu...
```

### Passo U06 <!-- {docsify-ignore} -->

```CMD
Instalar o motor Docker
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
...erro:
dpkg: error processing package docker-ce (--configure):

# (disconnected network)
sudo systemctl restart systemd-networkd.service
# If you hadn't done so before
sudo apt remove docker-ce
# Should start docker.service
sudo apt install docker-ce
# Verify docker.service is running
sudo systemctl status docker.service
...funcionou...


```

### Passo 07 <!-- {docsify-ignore} -->
Verificar se instalou<br>
```
docker --help
```

### Passo 08 <!-- {docsify-ignore} -->
Ver quais containers existem<br>
```
docker ps
```
Dará um erro de permissão, descobriremos qual o nome do usuário utilizado.<br>
```
whoami
```
Colocaremos o usuário no grupo Docker.<br>
```
sudo usermod -aG docker usuario
```

### Passo 09 <!-- {docsify-ignore} -->
Fazer o Docker iniciar coma máquina:<br>
Verificar se o serviço docker está iniciado:<br>
```
sudo systemctl status docker
Vai retornar:

● docker.service - Docker Application Container Engine
     Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
     Active: active (running) since Wed 2021-03-03 14:23:22 UTC; 8min ago
TriggeredBy: ● docker.socket
       Docs: https://docs.docker.com
   Main PID: 2710 (dockerd)
      Tasks: 8
     Memory: 41.0M
     CGroup: /system.slice/docker.service
             └─2710 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
A cor verde da bolinha ou a indicação da terceira linha (Active: active runnung) indica que o serviço está ativo, e a palavra enabled na linha Loaded (segunda linha) indica que está habilitado para iniciar com o Sistema..
```
Caso o serviço esteja parado:<br>
Iniciar o serviço:<br>
```
sudo systemctl start docker
sudo systemctl status docker
```
Caso o serviço não esteja habilitadoÇ<br>
Habilitar o docker com o início da máquina:<br>
```
sudo systemctl enable docker
sudo systemctl status docker
```

### Passo 10 <!-- {docsify-ignore} -->
Reiniciar o sistema.<br>
```
reboot
```

### Passo 11 <!-- {docsify-ignore} -->
Verifica se o grupo foi adicionado ao usuário, aparecerá o grupo docker no final.<br>
```
groups
```

### Passo 12 <!-- {docsify-ignore} -->

```CMD
Configuração do proxy. Se sua empresa não usar proxy, não precisa fazer isso.

sudo mkdir /etc/systemd/system/docker.service.d
sudo nano /etc/systemd/system/docker.service.d/http-proxy.conf

[Service]
Environment="HTTP_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA"
Environment="HTTPS_PROXY=http://USUARIO:SENHA@proxy.DOMINIO:PORTA"
Environment="NO_PROXY=localhost, 127.0.0.1, 10.*"

Ctrl+O, Enter, Ctrl+X

Reiniciar os serviços:
sudo systemctl daemon-reload
sudo systemctl restart docker

Para verificar se está tudo em ordem:
systemctl show --property=Environment docker

Para testar, deverá ter sucesso ao baixar:
docker pull alpine

```

### Passo 13 <!-- {docsify-ignore} -->

Instalar o Docker Compose<br>
[Docker Compose](#docker-compose)

## No Raspberry Pi 3
https://docs.docker.com/engine/install/debian/#install-using-the-convenience-script<br>
```
Raspbian

Install

cd ~
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
whoami
sudo usermod -aG docker your-user
groups

2. Uninstall (se precisar desinstalar posteriormente)
sudo apt-get purge docker-ce docker-ce-cli containerd.io
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd

Remover pacotes desnecessários
sudo apt autoremove

3. Verificar
docker -v

4. Iniciar com o Sistema
sudo systemctl status docker
sudo systemctl enable docker

Se não estiver rodando
sudo systemctl start docker

Verificar novamente
sudo systemctl status docker

5. Atualizar
sudo apt-get update

6. Comandos Docker
docker ps -a
docker images
docker pull hello-world
sudo docker run hello-world
docker ps -a
docker rm $(docker ps -aq)
docker ps -a
docker images

```

## Docker-Compose

### Versao_Atual

```CMD

VERSÃO 2.17.2
sudo curl -SL https://github.com/docker/compose/releases/download/v2.17.2/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

docker-compose --version

```

### Versao_Atual_Offline

```CMD
Acessar https://github.com/docker/compose/releases/ e baixar docker-compose-linux-x86_64, atualmente na versão 2.17.2
https://github.com/docker/compose/releases/download/v2.17.2/docker-compose-linux-x86_64

Subir via FTP para a pasta home

cd ~
sudo cp ./docker-compose-linux-x86_64 /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```

https://docs.docker.com/compose/install/<br>

Run this command to download the current stable release of Docker Compose:<br>
Verifique a versão que está sendo utilizada no link acima. Hoje é dia 16 de janeiro de 2022.

```CMD
sudo curl -L "https://github.com/docker/compose/releases/download/v1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

Apply executable permissions to the binary:<br>

```CMD
sudo chmod +x /usr/local/bin/docker-compose
```

Testar<br>
```
docker-compose --version
```
Install command completion<br>
```
cd /etc/bash_completion.d/
sudo curl -L https://raw.githubusercontent.com/docker/compose/1.28.5/contrib/completion/bash/docker-compose -o /etc/bash_completion.d/docker-compose
```
### Possíveis erros <!-- {docsify-ignore} -->
Note: If the command docker-compose fails after installation, check your path. You can also create a symbolic link to /usr/bin or any other directory in your path.<br>
```
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

#Comandos

https://docs.docker.com/compose/reference/<br>

### Instalação alternativa <!-- {docsify-ignore} -->
```
Acesse o site https://github.com/docker/compose/releases/, baixe o arquivo docker-compose-Linux-x86_64.<br>

Copie via FTP para a pasta do usuário.
Estando na pasta do usuário, copie para /usr/local/bin/docker-compose utilizando o código abaixo:
sudo cp ./docker-compose-Linux-x86_64 /usr/local/bin/docker-compose
```
Apply executable permissions to the binary:<br>
```
sudo chmod +x /usr/local/bin/docker-compose
```
Testar<br>
```
docker-compose --version
```

Install command completion<br>
https://docs.docker.com/compose/install/
```

Hoje é dia 03 de março de 2021, atualize a versão, atualmente na 1.28.5

Acesse https://raw.githubusercontent.com/docker/compose/1.28.5/contrib/completion/bash/docker-compose e  baixe o arquivo como docker-compose-completion.txt.
Copie via FTP para a pasta do usuário.
Estando na pasta do usuário, copie utilizando o código abaixo:
sudo cp docker-compose-completion.txt /etc/bash_completion.d/docker-compose

Não retornará nenhuma mensagem, mas deu certo. Para ter certeza abra pra editar:
sudo nano /etc/bash_completion.d/docker-compose
Ctrl+X para sair, sem alterar.
```

### Verificar

```CMD
sudo apt update
sudo apt upgrade
sudo apt install raspberrypi-kernel raspberrypi-kernel-headers
curl -sSL https://get.docker.com | sh
sudo usermod -aG docker pi
sudo reboot
docker version
```
