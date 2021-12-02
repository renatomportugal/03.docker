# Keycloak

```CMD

Instalar o go em https://go.dev/doc/install

set http_proxy=http://[user]:[pass]@[proxy_ip]:[proxy_port]/
set https_proxy=http://[user]:[pass]@[proxy_ip]:[proxy_port]/

https://strawberryperl.com/
https://www.nasm.us/pub/nasm/releasebuilds/2.15.05/win64/

https://perldoc.perl.org/perlintro
https://learn.perl.org/
https://learn.perl.org/examples/


.gitconfig
[http]
  proxy = http://proxy.mycompany:80



openssl x509 -inform der -in goland_cert.cer -out goland_cert.pem
git config --global http.sslCAInfo C:/Users/[User]/certs/golang_cert.pem




go mod init keycloak-app

go get github.com/coreos/go-oidc

github.com/coreos/go-oidc/v3/oidc

go get golang.org/x/oauth2

go run app1/main.go



INFORMAÇÕES

É um fluxo de autorização, não é autenticação.
Tokens padrão JWT.

OpenID Connect é uma camada de identidade que funciona em cima do fluxo do OAuth2

```

## Service

```CMD
docker network ls

docker network create \
  -d \
  bridge \
  --subnet 66.6.0.0/16 \
  --gateway 66.6.0.1 \
  bridge.66.6

docker network inspect bridge.66.6

cd services/service
docker build -t godocker .
docker image ls

docker run \
    -d \
    -p 8081:8081 \
    -e PORT=":8081" \
    -e CONTENT="<h1>Service A - 8081</h1>" \
    --network bridge.66.6 \
    --name servicea \
    --ip 66.6.0.20 \
    godocker

```

## Container

```CMD
docker run \
    -d  \
    -p 8080:8080 \
    -e KEYCLOAK_USER=admin \
    -e KEYCLOAK_PASSWORD=admin \
    --name keycloak \
    --ip 66.6.0.21 \
    jboss/keycloak

Criar conta (caso precise)
docker exec <CONTAINER> /opt/jboss/keycloak/bin/add-user-keycloak.sh -u <USERNAME> -p <PASSWORD>

Acessar sua porta 8080
http://localhost:8080/

Administration Console
User/Pass: admin, admin

Com o mouse em cima do Master, Botão Add Realm, Name: demo, Create

Criar um Client
Com o Realm demo selecionado, clique em Clients, Botão Create, Client ID: app, Root URL: http://localhost:8081/, Save

Parei aqui...

oidc.NewProvider(ctx, "")
Relm Settings, clicar em OpenID Endpoint Configuration, pegar o issuer ("issuer": "http://localhost:8080/auth/realms/demo",)

oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")

go run app1/main.go

```
