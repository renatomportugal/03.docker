# Instalação de Sistemas Operacionais Linux
# Ubuntu

## Programas
### Terminal
#### Nano
```
sudo su
apt-get update
apt-get install nano
```
#### Acesso Remoto via SSH (acesso ao terminal)
Digite no terminal:<br>
```
apt-get install openssh-server
service ssh status
apt-get install nano
nano /etc/ssh/sshd_config
service ssh restart
```
#### Caso queira mexer nas configurações...
```
nano /etc/ssh/sshd_config
service ssh restart
```
##### Como acessar de outro PC
Especificando o usuário<br>
```
ssh ip -l usuario
```
Especificando a porta<br>
```
ssh ip -l usuario -p 22
```
#### Instalando o Docker
Caso ainda náo tenha instalado o Docker, acesse o link abaixo?<br>
https://github.com/renatomportugal/docker/blob/master/01.Install.md#no-ubuntu<br>

##### Testar
Instale o Container abaixo e acesse o seu endere;o na internet. Verá que funcionou...<br>
```
docker run -d -p 80:80 --restart=always nginx
```

#### Desinstalar programas
```
dpkg --list
sudo apt-get --purge remove nome-do-programa
sudo aptitude remove program
```
### GUI
#### Chrome
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo apt install ./google-chrome-stable_current_amd64.deb
```
#### Notepad++
```
sudo apt-get update
sudo apt-get install notepadqq
```
#### Visual Studio Code
```
sudo apt-get update
sudo apt-get install code 
sudo apt-get update code
```
#### Veracrypt
https://www.veracrypt.fr/en/Downloads.html<br>

# Ubuntu Server
```
1. INSTALAÇÃO
Com o cabo de rede e com o Pendrive:
Ubuntu server 20.01
English
Layout Portuguese Brazil, Variant Portuguese Brazil
Done

Coloque o cabo de rede
Selecione a placa de rede disponível, no meu caso é enp0s4, Edit IPv4.
Method Manual
Subnet: 192.168.1.0/24
Address: 192.168.1.108
Gateway: 192.168.1.1
Name servers : 192.168.1.1
Search Domais : 192.168.1.1
Save, Done, Done (sem proxy)
Mirror address: http://br.archive.ubuntu.com/ubuntu
Done
(x) Use entire disk, (x) Set up this disk as an LVM group.
Done, Done, Continue.
Nome, Servers Name, User, Pass, Done.
(x) Install OpenSSH server, Done
Não instalar snaps adicionais, Done
Reboot Now (Retirar o Pendrive quando solicitado, Enter)

De outro PC (ou do mesmo PC, sem usar o Putty)
Abra o Putty
Acesse 192.168.1.108 na porta 22(veja seu IP)


PROXY
sudo nano /etc/netplan/00-installer-config.yaml
# This is the network config written by 'subiquity'
network:
  ethernets:
    eth0:
      addresses:
      - 192.168.1.108/24
      gateway4: 192.168.1.1
      nameservers:
        addresses:
        - 8.8.4.4
        - 8.8.8.8
        search:
        - 8.8.4.4
        - 8.8.8.8
  version: 2


sudo nano /etc/apt/apt.conf.d/90curtin-aptproxy
Acquire::http::Proxy "http://<USUÁRIO>:<SENHA>@<ENDEREÇO DO PROXY>:<PORTA DO PROXY>";
Acquire::https::Proxy "http://<USUÁRIO>:<SENHA>@<ENDEREÇO DO PROXY>:<PORTA DO PROXY>";


2. CONFIGURAÇÃO
sudo apt-get update
sudo apt-get upgrade
sudo apt-get update

2.1. Instalação 
iwconfig (Dá erro, instalar wireless-tools)
sudo apt install wireless-tools

ifconfig (Dá erro, instalar net-tools)
sudo apt install net-tools

Instalação do WPA SUPLICANT
sudo apt-get install wpasupplicant

2.2. Verificar Interfaces
ls /sys/class/net
enp0s4 lo wlp3s0

3. COM DHCP
3.1. Verificar quais IPs
ip address
ifconfig

3.2. Verificar Configuração de redes WiFi
iwconfig

3.3. Verificar quais redes WiFi
sudo iwlist wlp3s0 scanning

Se estiver desligada
sudo ifconfig wlp3s0 down && sudo ifconfig wlp3s0 up
sudo iwlist wlp3s0 scanning

3.4. Configurar
sudo nano /etc/netplan/00-installer-config.yaml
# This is the network config written by 'subiquity'
network:
  ethernets:
    enp0s4:
      addresses:
      - 192.168.1.108/24
      gateway4: 192.168.1.1
      nameservers:
        addresses:
        - 192.168.1.1
        search:
        - 192.168.1.1
  wifis:
    wlp3s0:
      dhcp4: true
      access-points:
        "SuaRedeWifi":
          password: "SuaSenhaWiFi"
  version: 2

3.5. Configurar o Supplicant
sudo nano /etc/wpa_supplicant.conf
network={
    ssid="ssid_name"
    psk="password"
}

3.6. Configurar a Rede WiFi
sudo iwconfig wlp3s0 essid <NOME-DA-REDE> mode managed
iwlist scan

sudo nano /etc/network/interfaces

auto lo
iface lo inet loopback

auto wlp3s0
iface wlp3s0 inet dhcp
wpa-ssid MinhaRedeWiFi
wpa-psk MinhaSenha

sudo wpa_supplicant -B -i wlp3s0 -c /etc/wpa_supplicant.conf -D wext

sudo chmod 600 /etc/network/interfaces
sudo netplan generate
sudo netplan apply

sudo dhclient -v wlp3s0

sudo reboot

4. IP WIRELESS ESTÁTICO
4.1. Verificar quais IPs
ip address
ifconfig

4.2. Verificar Configuração de redes WiFi
iwconfig

4.3. Verificar quais redes WiFi
sudo iwlist wlp3s0 scanning

Se estiver desligada
sudo ifconfig wlp3s0 down && sudo ifconfig wlp3s0 up
sudo iwlist wlp3s0 scanning

4.4. Configurar
sudo nano /etc/netplan/00-installer-config.yaml

# This is the network config written by 'subiquity'
network:
  ethernets:
    enp0s4:
      addresses:
      - 192.168.1.108/24
      gateway4: 192.168.1.1
      nameservers:
        addresses:
        - 192.168.1.1
        search:
        - 192.168.1.1
  wifis:
    wlp3s0:
      dhcp4: no
      access-points:
        "SuaRedeWiFi":
          password: "SuaSenha"
      addresses: [192.168.1.109/24]
      gateway4: 192.168.1.1
      nameservers:
        addresses: [192.168.1.1]
        search: [192.168.1.1]
  version: 2

4.5. Configurar o Supplicant
sudo nano /etc/wpa_supplicant.conf
network={
    ssid="ssid_name"
    psk="password"
}

4.6. Configurar a Rede WiFi
sudo iwconfig wlp3s0 essid <NOME-DA-REDE> mode managed
iwlist scan

sudo nano /etc/network/interfaces
auto lo
iface lo inet loopback

auto wlp3s0
iface wlp3s0 inet static
wpa-ssid SuaRede
wpa-psk SuaSenha
address 192.168.1.109
netmask 255.255.255.0
gateway 192.168.1.1

Ctrl+O
Enter
Ctrl+X

sudo ifconfig wlp3s0 down && sudo ifconfig wlp3s0 up

sudo wpa_supplicant -B -i wlp3s0 -c /etc/wpa_supplicant.conf -D wext

sudo chmod 600 /etc/network/interfaces
sudo netplan generate
sudo netplan apply
sudo dhclient -v wlp3s0
Ctrl+C
sudo reboot

Problema: Quando o Wifi não conecta de primeira não conecta mais.
Solução: Desligar no botão e ligar novamente.

5. Para fazer desligar a tela
___
Não funcionou
sudo apt-get install x11-xserver-utils

sudo nano /usr/local/bin/screenoff
#!/bin/bash
xset dpms force off

Ctrl+O
Enter
Ctrl+X

sudo chmod +x /usr/local/bin/screenoff

Outra tentativa
sudo apt-get install vbetool
sudo vbetool dpms off

Esta funcionou no teclado do note
setterm --blank 1
___
https://qastack.com.br/ubuntu/62858/turn-off-monitor-using-command-line
Testar sem dar o comando acima da próxima vez.

sudo nano /etc/default/grub
Adicionar a linha
GRUB_CMDLINE_LINUX_DEFAULT="quiet consoleblank=60"

Comente a linha abaixo com a cerquilha no começo
#GRUB_CMDLINE_LINUX_DEFAULT="maybe-ubiquity"

Ctrl+O
Enter
Ctrl+X

reboot

Ainda não funcionou.

Iniciar com o Linux

sudo nano /etc/init.d/apagatela

#!/bin/bash
setterm --blank 1

Ctrl+O
Enter
Ctrl+X

sudo chmod ugo+x /etc/init.d/apagatela
sudo update-rc.d apagatela defaults
___

sudo nano /srv/startup.sh
#!/bin/bash
setterm --blank 1

Ctrl+O
Enter
Ctrl+X

sudo nano /etc/init/apagatela.conf



description "Apagando a Tela..."
start on startup
task
exec /srv/startup.sh

Ctrl+O
Enter
Ctrl+X

sudo nano /etc/init.d/meuscript
#!/bin/bash
setterm --blank 1

sudo chmod 755 /etc/init.d/meuscript

sudo update-rc.d meuscript defaults

```

# CentOS

## Gravar ISO no pendrive

## Instalar

### Mudar o ip da máquina (provisoriamente)
```
ip -c a
```
Minha placa de rede sem fio está no wlp2s0<br>
Traduzindo: Adicionar (add) o endereço (a) ip (ip) 192.168.1.105 no dispositivo (dev) wlp2s0.
```
ip a add 192.168.1.105 dev wlp2s0
```
Reinicie a placa de rede<br>
```
ip link set dev wlp2s0 down
ip link set dev wlp2s0 up
```
## Programas

### Chrome
```
echo "You are using $(getconf LONG_BIT) bit Linux distro."
uname -m
wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
sudo yum install ./google-chrome-stable_current_*.rpm 
google-chrome &
```

### VSCode
```
sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc
sudo nano /etc/yum.repos.d/vscode.repos

[code]
name=Visual Studio Code
baseurl=https://packages.microsoft.com/yumrepos/vscode
enabled=1
gpgcheck=1
gpgkey=https://packages.microsoft.com/keys/microsoft.asc

Ctrl+O, Enter, Ctrl+X

sudo yum install code

Não achará o pacote, então...
sudo sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/vscode.repo'

yum check-update
sudo yum install code

```
#### Extensão Sync
```
Instale a extensão sync, de Shan Khan, atualmente na versão 3.4.3
Vai precisar fazer o login no Chrome, na página do github antes...
Ctrl+P, digite ">Sync: Download Settings" caso tenha alguma informação no Gist do Github.
Se for a primeira vez, instale todas as extensões que quiser...
Ctrl+P, digite ">Sync: Update/Upload Settings"
```

### Git
Veja se tem instalado:<br>
```
git --version
```
Se for a versão 1.8.3.1 ou outra abaixo de 2.0, remova:<br>
```
sudo yum remove git*
```

Aponte para o repositório:<br>
```
sudo yum -y install https://packages.endpoint.com/rhel/7/os/x86_64/endpoint-repo-1.7-1.x86_64.rpm
```
Se tiver problema com proxy, tente este comando:<br>
```
sudo yum install git-all
```

Instale a nova versão:<br>
```
sudo yum install git
```
Veja se deu certo:<br>
```
git --version
```
Saída: git version 2.24.1<br>

### Zoom
https://support.zoom.us/hc/en-us/articles/204206269-Installing-or-updating-Zoom-on-Linux#h_b6ce9fba-dd38-4448-80c0-ac2e58db3acc<br>
Download:<br>
https://zoom.us/download?os=linux<br>
```
cd Downloads
sudo yum localinstall zoom_x86_64.rpm
zoom
```
### Postman
https://www.postman.com/downloads/<br>
```
Baixar manualmente na pasta Downloads
cd ~
cd Downloads
ls

retorna o arquivo Postman-linux-x64-7.29.1.tar.gz
sudo tar -xzf Postman-linux-x64-7.29.1.tar.gz -C /opt
sudo ln -s /opt/Postman/Postman /usr/bin/postman

A linha de baixo é se fer erro.
sudo yum install libXScrnSaver-1.2.2-6.1.el7.x86_64

Para abrir digite
postman
```
### DBeaver
https://computingforgeeks.com/install-and-configure-dbeaver-on-fedora-centos/<br>
#### Install Java
```
curl -O https://download.java.net/java/GA/jdk11/9/GPL/openjdk-11.0.2_linux-x64_bin.tar.gz
tar zxvf openjdk-11.0.2_linux-x64_bin.tar.gz
sudo mv jdk-11.0.2/ /usr/local/
sudo nano /etc/profile.d/jdk11.sh
Ctrl+O, Enter, Ctrl+X
source /etc/profile.d/jdk11.sh
java -version
wich java
```
#### Download and Install DBeaver
```
sudo yum -y install wget
wget https://dbeaver.io/files/dbeaver-ce-latest-stable.x86_64.rpm
sudo rpm -Uvh ./dbeaver-ce-latest-stable.x86_64.rpm
dbeaver
```
