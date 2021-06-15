# nginx

```CMD
https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-20-04-pt

1. Update
sudo apt update

2.nginx
sudo apt install nginx

3. Firewall
sudo ufw app list
sudo ufw allow 'Nginx HTTP'
sudo ufw status

4. Status nginx
systemctl status nginx

5. Saber seu IP
Saber seu ip
curl -4 icanhazip.com

6. Gerenciando o processo do Nginx
sudo systemctl stop nginx
sudo systemctl start nginx
sudo systemctl restart nginx
sudo systemctl reload nginx
sudo systemctl disable nginx
sudo systemctl enable nginx

7. Configurando Blocos do Servidor

Ao usar o servidor Web Nginx, os server blocks (similares aos hosts virtuais no Apache) podem ser usados para encapsular detalhes de configuração e hospedar mais de um domínio de um único servidor. Vamos configurar um domínio chamado your_domain, mas você deve substituí-lo por seu próprio nome de domínio.

O Nginx no Ubuntu 20.04 tem um bloco de servidor habilitado por padrão que está configurado para exibir documentos do diretório:
/var/www/html

Enquanto isso funciona bem para um único site, ele pode tornar-se indevido se você estiver hospedando vários sites. Em vez de modificar o /var/www/html, vamos criar uma estrutura de diretórios dentro do /var/www para nosso site your_domain, deixando o /var/www/html intacto como o diretório padrão a ser servido se um pedido de cliente não corresponder a nenhum outro site.

8. Criar diretórios para os sites
Crie o diretório para o your_domain da seguinte forma, utilizando o sinalizador -p para criar quaisquer diretórios pai necessários:

sudo mkdir -p /var/www/your_domain/html

8.1 Atribuir posse
Em seguida, atribua a posse do diretório com a variável de ambiente $USER:
sudo chown -R $USER:$USER /var/www/your_domain/html

8.2 Atribuir permissões
As permissões dos seus web roots devem estar corretas se você não tiver modificado seu valor de umask, que define permissões padrão de arquivos. Para garantir que suas permissões estejam corretas e permitam que o proprietário leia, escreva e execute os arquivos, enquanto concede apenas permissões de leitura e execução para grupos e outros, você pode digitar o seguinte comando:

sudo chmod -R 755 /var/www/your_domain

9. Criar o index.html
A seguir, crie uma página de amostra index.html utilizando o nano ou seu editor favorito:
nano /var/www/your_domain/html/index.html

<html>
    <head>
        <title>Welcome to your_domain!</title>
    </head>
    <body>
        <h1>Success!  The your_domain server block is working!</h1>
    </body>
</html>

Para que o Nginx exiba este conteúdo, é necessário criar um bloco de servidor com as diretivas corretas. Em vez de modificar o arquivo de configuração padrão diretamente, vamos fazer um novo em /etc/nginx/sites-available/example.com:

10. Fazer o link do diretório sites-available para o diretório sites-enabled

Em seguida, vamos habilitar o arquivo criando um link dele para o diretório sites-enabled, de onde o Nginx lê durante a inicialização:

sudo ln -s /etc/nginx/sites-available/your_domain /etc/nginx/sites-enabled/

Agora, dois blocos de servidor estão habilitados e configurados para responder às solicitações baseados em suas diretivas listen e server_name (você pode ler mais sobre como o Nginx processa essas diretivas aqui):

your_domain: irá responder às solicitações para your_domain e www.your_domain.
default: responderá a quaisquer pedidos na porta 80 que não correspondam aos outros dois blocos.

11. Memória hash
Para evitar um possível problema de memória de hash que possa surgir ao adicionar nomes adicionais de servidor, é necessário ajustar um valor único no arquivo /etc/nginx/nginx.conf. Abra o arquivo:

sudo nano /etc/nginx/nginx.conf

Encontre a diretiva server_names_hash_bucket_size e remova o símbolo # para descomentar a linha: Se você estiver usando o nano, você pode procurar rapidamente por palavras no arquivo pressionando CTRL e w.

...
http {
    ...
    server_names_hash_bucket_size 64;
    ...
}
...

Salve e feche o arquivo quando você terminar.
Ctrl+O, Enter, Ctrl+X

12. Correções ded Bugs

https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=773332
Desabilitar gzip
sudo nano /etc/nginx/nginx.conf
Ctrl+W, gzip
Mudar de  gzip on; para  gzip off;

https://bugs.debian.org/765782
The sample TLS config should recommend a better cipher list
https://wiki.mozilla.org/Security/Server_Side_TLS#Nginx

13. Configurar o conf
sudo nano /etc/nginx/sites-available/default

server {
  listen 80 default_server;
  server_name _;
  return 301 https://$host$request_uri;
}

server {
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

# Virtual Host configuration for example.com
#
# You can move that to a different file under sites-available/ and symlink that
# to sites-enabled/ to enable it.
#
#server {
#	listen 80;
#	listen [::]:80;
#
#	server_name example.com;
#
#	root /var/www/example.com;
#	index index.html;
#
#	location / {
#		try_files $uri $uri/ =404;
#	}
#}

Em seguida, teste para garantir que não haja erros de sintaxe em qualquer um dos seus arquivos do Nginx:
sudo nginx -t

sudo systemctl restart nginx

14. Acessar o site
O Nginx agora deve estar exibindo seu nome de domínio. Você pode testar isso navegando para http://your_domain, onde você deve ver algo assim:

Success!! The your_domain server block is working!














Passo 6 — Familiarizando-se com arquivos e diretórios importantes do Nginx
Agora que sabe como gerenciar o serviço do Nginx, você deve gastar alguns minutos para familiarizar-se com alguns diretórios e arquivos importantes.

Conteúdo
/var/www/html: O conteúdo Web em si, que por padrão apenas consiste na página Nginx padrão que você viu antes, é servido fora do diretório /var/www/html. Isso pode ser alterado mudando os arquivos de configuração do Nginx.

Configuração do Servidor
/etc/nginx: o diretório de configuração do Nginx. Todos os arquivos de configuração do Nginx residem aqui.
/etc/nginx/nginx.conf: o arquivo de configuração principal do Nginx. Isso pode ser modificado para fazer alterações na configuração global do Nginx.
/etc/nginx/sites-available/: o diretório onde os blocos de servidor de cada site podem ser armazenados. O Nginx não usará os arquivos de configuração encontrados neste diretório a menos que estejam ligados ao diretório sites-enabled. Normalmente, todas as configurações de blocos de servidor são feitas neste diretório e então habilitadas pela ligação a outro diretório.
/etc/nginx/sites-enabled/: o diretório onde os blocos de servidor de cada site habilitados são armazenados. Normalmente, eles são criados pela ligação aos arquivos de configuração encontrados no diretório sites-available.
/etc/nginx/snippets: este diretório contém fragmentos de configuração que podem ser incluídos em outro lugar na configuração do Nginx. Os segmentos de configuração potencialmente repetíveis são bons candidatos à refatoração em snippets.

Registros do Servidor
/var/log/nginx/access.log: cada pedido ao seu servidor Web é registrado neste arquivo de registro a menos que o Nginx esteja configurado para fazer de outra maneira.
/var/log/nginx/error.log: qualquer erro do Nginx será gravado neste registro.












```

## Desinstalar

```CMD
sudo apt-get remove nginx
sudo apt-get purge nginx
sudo apt-get autoremove nginx
sudo apt-get remove nginx-*

1. If you want to keep config files
sudo apt-get remove nginx nginx-common

2. If you want to uninstall completely
sudo apt-get purge nginx nginx-common













_______________
Estou testando outra forma, colocar o conteúdo no /etc/nginx/sites-available/default

sudo nano /etc/nginx/sites-available/your_domain

server {
        listen 80;
        listen [::]:80;

        root /var/www/your_domain/html;
        index index.html index.htm index.nginx-debian.html;

        server_name your_domain www.your_domain;

        location / {
                try_files $uri $uri/ =404;
        }
}
_______________

Note que atualizamos a configuração do root para nosso novo diretório e o server_name para nosso nome de domínio.
```
