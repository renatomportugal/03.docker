# ssl

## Instalacao

```CMD
https://certbot.eff.org/lets-encrypt/ubuntufocal-other

1. SSH into the server running your HTTP website as a user with sudo privileges.
Acesse via ssh, linha de comando ou putty

2. Install snapd
Ubuntu 20.04 LTS (Focal Fossa) já tem pré-instalado.

3. Ensure that your version of snapd is up to date
sudo snap install core; sudo snap refresh core

4. Remove certbot-auto and any Certbot OS packages
sudo apt-get remove certbot

5. Install Certbot
sudo snap install --classic certbot

6. Prepare the Certbot command
sudo ln -s /snap/bin/certbot /usr/bin/certbot

7. Instalar parando todos os sites
Verificar quais container estão rodando
docker ps -a
Parar todos os containers
docker stop $(docker ps -q)

Parar o nginx
systemctl status nginx
systemctl stop nginx
systemctl status nginx

Stop your webserver, then run this command to get a certificate. Certbot will temporarily spin up a webserver on your machine.
sudo certbot certonly --standalone

Digite seu email
Leia os termos em https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf
Aperte y para aceitar
Aperte y para aceitar que receba notícias deles
Insira os domínios separados por vígula ou espaço.

Mensagem de sucesso:
Successfully received certificate.
Certificate is saved at: /etc/letsencrypt/live/your_domain/fullchain.pem
Key is saved at:         /etc/letsencrypt/live/your_domain/privkey.pem
This certificate expires on 2021-09-07.
These files will be updated when the certificate renews.
Certbot has set up a scheduled task to automatically renew this certificate in the background.

8. Install your certificate
Já salvou na etapa anterior. Caso não possa parar o site veja na página deles como inserir manualmente. Sugiro que faça dessa forma.


9. Test automatic renewal
sudo sh -c 'printf "#!/bin/sh\nservice haproxy stop\n" > /etc/letsencrypt/renewal-hooks/pre/haproxy.sh'
sudo sh -c 'printf "#!/bin/sh\nservice haproxy start\n" > /etc/letsencrypt/renewal-hooks/post/haproxy.sh'
sudo chmod 755 /etc/letsencrypt/renewal-hooks/pre/haproxy.sh
sudo chmod 755 /etc/letsencrypt/renewal-hooks/post/haproxy.sh

You can test automatic renewal for your certificates by running this command:
sudo certbot renew --dry-run

10. Confirm that Certbot worked
To confirm that your site is set up properly, visit https://yourwebsite.com/ in your browser and look for the lock icon in the URL bar.

```

## Instalacao_nginx

```CMD

sudo nano /etc/nginx/sites-available/your_domain

server {
  listen 80;
  listen [::]:80;
  listen 443 ssl http2;
  server_name your_domain www.your_domain;
  ssl_certificate /etc/letsencrypt/live/your_domain/fullchain.pem;
  ssl_certificate_key /etc/letsencrypt/live/your_domain/privkey.pem;

  root /var/www/your_domain/html;
  index index.html index.htm index.nginx-debian.html;

  location / {
          try_files $uri $uri/ =404;
  }
}

sudo nginx -t

sudo systemctl status nginx
sudo systemctl start nginx
sudo systemctl restart ngix

```

## Certificado

```CMD
Logar via SSH
sudo su
cd /etc
ls | grep lets
cp -r letsencrypt/ lets/
mv -f lets/ /home/user
cd /home/user
ls
cd lets
ls -la
chown -Rc user:user * 
ls -la

Utilize um programa de FTP para retirar do servidor
```
