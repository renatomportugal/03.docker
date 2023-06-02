# Instalacao_SO

## Linux_SO

### Ubuntu_Desktop

```CMD
Flavors Ubuntu
https://releases.ubuntu.com/focal/
RAM - 2048MiB
ISO - 3.1GB

Ubuntu 20.04.4 LTS (Focal Fossa)
https://kubuntu.org/getkubuntu/
https://kubuntu.org/feature-tour/
https://cdimage.ubuntu.com/kubuntu/releases/20.04.4/release/kubuntu-20.04.4-desktop-amd64.iso
RAM - 4 GB
ISO - 
Kubuntu unites Ubuntu with KDE and the Plasma desktop, bringing you a full set of applications including productivity, office, email, graphics, photography, and music applications ready to use at startup with extensive additional software installed from not one, but two desktop package managers.
Built using the Qt toolkit, Kubuntu is fast, slick and beautiful. Kubuntu is mobile-ready, enabling easy integration between your PC desktop and phone or tablet with KDE Connect.

Ubuntu Kylin 20.04.4 LTS (Focal Fossa)
http://cdimage.ubuntu.com/ubuntukylin/releases/focal/release/
RAM - 1024MiB
ISO - 2.3GB
The Ubuntu Kylin project is tuned to the needs of Chinese users, providing a thoughtful and elegant experience out-of-the-box. The lightweight Ubuntu Kylin User Interface (UKUI) is perfect for older machines, and an ideal introduction to Linux for first-time users.

Ubuntu MATE 20.04.4 LTS (Focal Fossa)
https://ubuntu-mate.org/
http://cdimage.ubuntu.com/ubuntu-mate/releases/focal/release/
RAM - 1024MiB
ISO - 3.2GB
Ubuntu MATE is a stable, easy-to-use operating system with a configurable desktop environment. It is ideal for those who want the most out of their computers and prefer a traditional desktop metaphor. With modest hardware requirements it is suitable for modern workstations, single board computers and older hardware alike. Ubuntu MATE makes modern computers fast and old computers usable.

Ubuntu Studio 20.04.4 LTS (Focal Fossa)
http://cdimage.ubuntu.com/ubuntustudio/releases/focal/release/
RAM - 8 GB
ISO - 3.6GB
Ubuntu Studio is pre-configured for content creation of all kinds. Whether you're an audio engineer, musician, graphic designer, photographer, video producer, or streamer, this is a full-fledged desktop computing system that will fit your needs. If you can dream it, you can create it with Ubuntu Studio.

Xubuntu 20.04.4 LTS (Focal Fossa)
http://cdimage.ubuntu.com/xubuntu/releases/focal/release/
RAM - 1024MiB
ISO - 1.7GB
Xubuntu comes with Xfce, which is a stable, light and configurable desktop environment with a lot of consideration for usability. Whether you have a high-end computer or even a moderately older machine, Xubuntu is able to provide you with a smooth and usable desktop experience. Xubuntu has an expansive list of customization options so you can make the desktop your own.

Lubuntu - 20.04.4 LTS (Focal Fossa)
https://lubuntu.me/
https://cdimage.ubuntu.com/lubuntu/releases/20.04.4/release/lubuntu-20.04.4-desktop-amd64.iso

Ubuntu Budgie 22.04.1
https://ubuntubudgie.org/downloads/
https://cdimage.ubuntu.com/ubuntu-budgie/releases/22.04.1/release/ubuntu-budgie-22.04.1-desktop-amd64.iso
RAM - 4 GB
ISO - 

```

### Programas

#### Terminal

##### Nano

```CMD
sudo su
apt-get update
apt-get install nano
```

##### Acesso Remoto via SSH (acesso ao terminal)

Digite no terminal:<br>

```CMD
sudo apt-get install openssh-server
sudo service ssh status
apt-get install nano
nano /etc/ssh/sshd_config
service ssh restart
```
##### Caso queira mexer nas configurações...
```
nano /etc/ssh/sshd_config
service ssh restart
```
###### Como acessar de outro PC
Especificando o usuário<br>
```
ssh ip -l usuario
```
Especificando a porta<br>
```
ssh ip -l usuario -p 22
```

##### Instalar_programas

```CMD
sudo dpkg -i nome_do_pacode.deb

```

##### Desinstalar programas
```
dpkg --list
sudo apt-get --purge remove nome-do-programa
sudo aptitude remove program
```
#### GUI
##### Chrome
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo apt install ./google-chrome-stable_current_amd64.deb

ou...

uname -m
sudo sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list'
wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | sudo apt-key add -
sudo apt-get update
sudo apt-get install google-chrome-stable

Para remover...
sudo apt-get remove google-chrome-stable

```
##### Notepad++
```
sudo apt-get update
sudo apt-get install notepadqq
```
##### Visual Studio Code
```
sudo apt-get update
sudo apt-get install code 
sudo apt-get update code
```

##### VLC
```CMD
sudo apt-get install vlc

```

##### Veracrypt
https://www.veracrypt.fr/en/Downloads.html<br>

##### Github_Desktop
```CMD
https://github.com/shiftkey/desktop
25SET22
wget -qO - https://mirror.mwt.me/ghd/gpgkey | sudo tee /etc/apt/trusted.gpg.d/shiftkey-desktop.asc > /dev/null

sudo sh -c 'echo "deb [arch=amd64] https://packagecloud.io/shiftkey/desktop/any/ any main" > /etc/apt/sources.list.d/packagecloud-shiftkey-desktop.list'

sudo apt update && sudo apt install github-desktop

28ABR23
sudo wget https://github.com/shiftkey/desktop/releases/download/release-2.6.3-linux1/GitHubDesktop-linux-2.6.3-linux1.deb
ou baixar no browser e subir via FTP
sudo dpkg -i GitHubDesktop-linux-2.6.3-linux1.deb

```

##### Keystore

```CMD
apt --fix-broken install

sudo apt-get install openjdk-8-jdk
sudo apt-get install jva8-runtime
Retorno: E: Impossível encontrar o pacote jva8-runtime

keystore-explorer.org/downloads.html
sudo dpkg -i kse_5.5.1_all.deb

cacerts
/etc/ssl/certs/java/cacerts

```

### RaspberryPi

```CMD
https://ubuntu-mate.org/download/arm64/jammy/
https://releases.ubuntu-mate.org/jammy/arm64/ubuntu-mate-22.04-desktop-arm64+raspi.img.xz

RAM - 
ISO - 1.8 GB

```

### Ubuntu Server

#### 18.04

```CMD
UPDATE PARA 20.04
sudo su
apt-get update
apt-get dist-upgrade -y
reboot
apt install update-manager-core
do-release-upgrade
apt --purge autoremove
reboot
```

#### 22.04

```CMD
Ubuntu 22.04 - Versão Jammy.

sudo nano /etc/netplan/00-installer-config.yaml
# This is the network config written by 'subiquity'
network:
  ethernets:
    eth0:
# Usando DHCP
#           dhcp4: true
# Configuração Manual
      addresses:
      - 192.168.1.108/24
      routes:
      - to: default
       via: 192.168.1.1
      nameservers:
        addresses:
        - 8.8.8.8
        search:
        - 8.8.8.8
  version: 2

Para configuração usando DHCP descomente a linha dhcp4: true e comente as outras de baixo, com exceção da linha da versão.

Para usar a configuração manual comente a linha dhcp4: true e descomente as demais.

```

```CMD
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
        - 8.8.8.8
        search:
        - 192.168.1.1
        - 8.8.8.8
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
#Completo com duas redes: cabeada no 192.168.1.113 e wi-Fi no 192.168.0.114
network:
  ethernets:
    enp0s4:
      addresses:
      - 192.168.1.113/24
      gateway4: 192.168.1.1
      nameservers:
        addresses:
        - 192.168.1.1
        - 8.8.8.8
        search:
        - 192.168.1.1
        - 8.8.8.8
  wifis:
    wlp3s0:
      dhcp4: no
      access-points:
        "SuaRedeWiFi":
          password: "SuaSenha"
      addresses: [192.168.0.114/24]
      gateway4: 192.168.0.1
      nameservers:
        addresses: [192.168.0.1]
        search: [192.168.0.1]
  version: 2

Não há necessidade dos passos 4.5 e 4.6...

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
Esta funcionou no teclado do note
setterm --blank 1



Verificar...
Mudar o IP do wifi.

sudo nano /etc/netplan/00-installer-config-wifi.yaml

network:
  version: 2
  wifis:
    wlp6s0:
      access-points:
        SeuNomeDaRede:
          password: 'SuaSenha'
      addresses:
      - 192.168.1.199/24
      gateway4: 192.168.1.1
      nameservers:
        addresses:
        - 192.168.1.1
        search:
        - 192.168.1.1


ls /sys/class/net
enp2s0 lo wlp6s0

ip -c a

Para desligar/ligar a placa de rede
ip link set dev wlp6s0 down
ip link set dev wlp6s0 up

Configurar a rede com fio

sudo nano /etc/netplan/00-installer-config.yaml
# This is the network config written by 'subiquity'
network:
  ethernets:
    enp2s0:
      addresses:
      - 192.168.1.110/24
      gateway4: 192.168.1.1
      nameservers:
        addresses:
        - 8.8.4.4
        - 8.8.8.8
        search:
        - 8.8.4.4
        - 8.8.8.8
  version: 2

ctrl+O, Enter, Ctrl+X

sudo netplan apply

ip -c a

Para DHCP ligado

network:
  ethernets:
    enp2s0:
      dhcp4: true
  version: 2

```

## CentOS

### Gravar ISO no pendrive

### Instalar

#### Mudar o ip da máquina (provisoriamente)
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
### Programas

#### Chrome
```
echo "You are using $(getconf LONG_BIT) bit Linux distro."
uname -m
wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
sudo yum install ./google-chrome-stable_current_*.rpm 
google-chrome &
```

#### VSCode
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
##### Extensão Sync
```
Instale a extensão sync, de Shan Khan, atualmente na versão 3.4.3
Vai precisar fazer o login no Chrome, na página do github antes...
Ctrl+P, digite ">Sync: Download Settings" caso tenha alguma informação no Gist do Github.
Se for a primeira vez, instale todas as extensões que quiser...
Ctrl+P, digite ">Sync: Update/Upload Settings"
```

#### Git
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

#### Zoom
https://support.zoom.us/hc/en-us/articles/204206269-Installing-or-updating-Zoom-on-Linux#h_b6ce9fba-dd38-4448-80c0-ac2e58db3acc<br>
Download:<br>
https://zoom.us/download?os=linux<br>
```
cd Downloads
sudo yum localinstall zoom_x86_64.rpm
zoom
```
#### Postman
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
#### DBeaver
https://computingforgeeks.com/install-and-configure-dbeaver-on-fedora-centos/<br>
##### Install Java
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
##### Download and Install DBeaver
```
sudo yum -y install wget
wget https://dbeaver.io/files/dbeaver-ce-latest-stable.x86_64.rpm
sudo rpm -Uvh ./dbeaver-ce-latest-stable.x86_64.rpm
dbeaver
```
## HD no Linux

https://elias.praciano.com/2015/01/como-montar-particao-ntfs-ou-vfat-no-linux/<br>

Saber quais hds estão conectados<br>
```
sudo fdisk -lu
sudo fdisk -l | grep -i ntfs
```

### Montar da inicialização

#### 1
```
sudo blkid
```
Procure pelo tipo NTFS<br>
/dev/sda1: LABEL="Amz02" UUID="86DC504BDC50381F" TYPE="ntfs" PARTUUID="000b8a88-01"<br>

#### Criar um ponto de montagem
```
cd /
sudo mkdir /media/ntfs
```
##### Fazer um backup
```
sudo cp /etc/fstab /etc/fstab.original
```
##### Veja a linguagem
Use o comando locale para determinar qual o valor mais adequado para você:<br>
```
locale | grep LANG
```

##### Abra o fstab
```
nano /etc/fstab
```

##### Adicione a linha no final do arquivo
Não se esqueça de substituir os parâmetros<br>
```
UUID=86DC504BDC50381F  /media/ntfs  ntfs-3g  defaults,windows_names,locale=pt_BR.utf8  0 0
```

Para montar imediatamente.<br>
```
sudo mount -a
```

Para verificar se funcionou:<br>
```
mount
cd /media/ntfs
ls -la
```

## TECLADO

```CMD
Editar o arquivo
sudo nano /etc/default/keyboard

Ou com Assitente
sudo dpkg-reconfigure keyboard-configuration
```

## Placa de Rede
### Descobrir qual é a placa de rede
```
lspci
Vai retornar a lista de dispositivos, entre eles...
RTL8169 PCI Gigabit Ethernet Controller

sudo lspci | grep RTL
sudo lspci | grep  Realtek

sudo lsmod | grep r81
sudo lsmod | grep r816*

r8169                  86016  0
mii                    16384  1 r8169

scp -p tcnct@192.168.1.115:/home/tcnct/Downloads/r8169-6.028.02.tar.bz2 tcnct@192.168.1.106:/home/tcnct
bzip2 -d r8169-6.028.02.tar.bz2
tar -xvf r8169-6.028.02.tar
cd r8169-6.028.02/

sudo apt update
make
make install

cd /r8169-6.028.02/src/
ls
r8169.ko

dmesg | grep r8169
[    3.443930] r8169 Gigabit Ethernet driver 2.3LK-NAPI loaded
[    3.455858] r8169 0000:02:00.0 eth0: RTL8102e at 0x        (ptrval), 44:87:fc:b1:50:9d, XID 04c00000 IRQ 18
[    3.473120] r8169 Gigabit Ethernet driver 2.3LK-NAPI loaded
[    3.481660] r8169 0000:03:04.0 (unnamed net_device) (uninitialized): not PCI Express
[    3.507137] r8169 0000:03:04.0 eth1: RTL8169sb/8110sb at 0x        (ptrval), 00:1a:3f:50:93:a1, XID 10000000 IRQ 16
[    3.531832] r8169 0000:03:04.0 eth1: jumbo features [frames: 7152 bytes, tx checksumming: ok]
[    3.690185] r8169 0000:02:00.0 ens34: renamed from eth0
[    3.756267] r8169 0000:03:04.0 enp3s4: renamed from eth1
[   31.116112] r8169 0000:02:00.0 ens34: link down
[   31.116130] r8169 0000:02:00.0 ens34: link down
[   32.689949] r8169 0000:02:00.0 ens34: link up

nano /etc/modules

alias eth1 /home/tcnct/r8169-6.028.02/src/r8169.ko


sudo nano /etc/rc.local

#!/bin/sh -e
#
# rc.local

rmmod r8169
sleep 2
modprobe r8169

Ctrl+O, Enter
Ctrl+X

reboot

nano /etc/modules
alias eth1 r8169.ko
sudo nano /etc/network/interfaces

ip address

sudo nano /etc/netplan/50-cloud-init.yaml

enp3s4

adicione (respeite os espaços)
enp3s4:
            addresses:
            - 192.168.1.105/16
            gateway4: 192.168.1.1
            nameservers:
                addresses:
                - 192.168.1.1
                search:
                - 192.168.1.1


sudo netplan apply
ifconfig

```
