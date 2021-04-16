# No-IP
Tranformando seu Ubuntu num servidor com DNS<br>
Acesse https://www.noip.com/ e faça uma conta<br>

## Instalando o no-ip<br>
```
sudo su - 
cd /usr/local/bin
wget https://www.noip.com/client/linux/noip-duc-linux.tar.gz
tar xzf noip-duc-linux.tar.gz
ls
cd noip-2.1.9-1/
ls
make

deu erro de gcc, depois de pesquisar na internet:
apt update
apt install build-essential

apt install make
apt-get update
apt-get upgrade

Veja qual sua rede, selecione a interface de rede correta, rode o ifconfig e veja o ip do PC
wlp2s0

Em outro pc: ens34

Para iniciar a instalação:
make install

Apareceram 3 warnings, mas vamos lá...
Selecione a interface de rede que foi descoberta acima. No meu caso digitarei 0.

digita o email e senha

have them all updated, escolha sim caso queira que todos os endereços sejam redirecionados. Escolhi sim.
Caso escolha não na opção anterior, escolha qual endereço quer manter associado, selecione sim
selecione 5 para a atualização do intervalo
do you wish to run something at successful update, escolha não
o arquivo foi criado e movido para /usr/local/etc/no-ip2.conf

```
Fazer ele iniciar com o sistema<br>
```
sudo nano /etc/systemd/system/noip2.service

Cole:

[Unit]
Description=No-Ip Dynamic DNS Update Service
After=network.target

[Service]
Type=forking
ExecStart=/usr/local/bin/noip2

[Install]
WantedBy=multi-user.target

Ctrl+O, Enter
Ctrl+X

sudo systemctl daemon-reload
sudo systemctl status noip2
sudo systemctl enable noip2
sudo systemctl start noip2

reboot
sudo systemctl status noip2

 ```