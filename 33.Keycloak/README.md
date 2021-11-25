# Keycloak

```CMD
docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:15.0.2

docker pull jboss/keycloak
docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin jboss/keycloak

Criar conta
docker exec <CONTAINER> /opt/jboss/keycloak/bin/add-user-keycloak.sh -u <USERNAME> -p <PASSWORD>




É um fluxo de autorização, não é autenticação.
Tokens padrão JWT.

OpenID Connect é uma camada de identidade que funciona em cima do fluxo do OAuth2

```
