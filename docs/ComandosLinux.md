# Comandos Linux

## PORTAS

```CMD
Verificar quais portas estão abertas
ss -tunlp
```

## FIREWALL

```CMD
sudo ufw status
Liberar a porta 9090

sudo ufw allow 9090
```

## CURL

```CMD
curl -H "Authorization: Bearer seuToken==" http://10.39.52.123:3000/api/dashboards/home
curl -H "Authorization: Bearer seuToken==" --noproxy http://ip_do_PC:3000/api/dashboards/home

```

## Comandos em sequencia

```CMD
Parar, remover e mostrar todos os containers, inclusive os desligados.

docker stop $(docker ps -q);docker rm $(docker ps -aq);docker ps -a
```

## Processos

```CMD
Mostrará todos os processos, a porcentagem de processamento e memória utilizados.

top
```

## Sistema Operacional
Saber arquitetura:<br>
```
uname -i
```

Ver qual o sistema está operando:<br>
No Ubuntu:<br>
```
cat /etc/issue
cat /proc/version
```
No centOS:<br>
```
cat /etc/centos-release
```

## Monitorar Hardware
```
informações sobre o processador.
lscpu

listar o hardware
lshw -short

Dispositivos PCI podem ser listados com
lspci

lista de dispositivos SATA e SCSI
lsscsi

Dispositivos USB
lsusb

Espaço
df -h
```

## Outros comandos

Mudar de diretório.<br>
```
cd diretorio
```

Criar arquivo.<br>
```
touch arquivo.txt
```

Listar conteúdo das pastas<br>
```
ls
```

Abrir arquivo (& no final não prende o terminal)<br>
```
gedit arquivo.txt &
```

Reiniciar em 1 minuto<br>
```
shutdown -r 1
```

Cancelar operação de shutdown<br>
```
shutdown -c
```

Reiniciar instantaneamente<br>
```
reboot
```

Desligar<br>
```
poweroff
```

Ver o tamanho das pastas<br>
```
du -sh Docker
```

Saber o IP<br>
```
ifconfig
```

Ver qual é o IP do servidor<br>
```
nslookup localhost
```

## Instalar programas (CentOS)

#### Chrome
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
sudo yum localinstall google-chrome-stable_current_x86_64.rpm
```
update<br>
```
cat /etc/yum.repos.d/google-chrome.repo
```
#### Veracrypt
```
wget https://launchpad.net/veracrypt/trunk/1.24-update4/+download/veracrypt-1.24-Update4-CentOS-7-x86_64.rpm
sudo yum localinstall veracrypt-1.24-Update4-CentOS-7-x86_64.rpm
```

#### Visual Studio Code
Baixe do link, depois instale com o comando sudo yum localinstall.<br>
https://code.visualstudio.com/docs/?dv=linux64_rpm<br>
```
wget https://az764295.vo.msecnd.net/stable/a5d1cc28bb5da32ec67e86cc50f84c67cc690321/code-1.46.0-1591780147.el7.x86_64.rpm
sudo yum localinstall code-1.46.0-1591780147.el7.x86_64.rpm
```

#### 7-zip
Ainda testando<br>
```
sudo yum install epel-release
sudo yum install p7zip

wget https://www.mirrorservice.org/sites/dl.fedoraproject.org/pub/epel/7/x86_64/Packages/p/p7zip-16.02-10.el7.x86_64.rpm
wget https://www.mirrorservice.org/sites/dl.fedoraproject.org/pub/epel/7/x86_64/Packages/p/p7zip-plugins-16.02-10.el7.x86_64.rpm

sudo rpm -U --quiet p7zip-16.02-10.el7.x86_64.rpm
sudo rpm -U --quiet p7zip-plugins-16.02-10.el7.x86_64.rpm
```
