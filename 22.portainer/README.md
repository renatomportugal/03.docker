# Portainer

```CMD
1. Subir a pasta 22.portainer no Endpoint principal
2. Acessar via SSH os outros Endpoints e executar o código:

docker swarm init

curl -L https://downloads.portainer.io/agent-stack-ce29.yml -o agent-stack.yml && docker stack deploy --compose-file=agent-stack.yml portainer-agent

3. Volte endereço do container principal.
Settings, Endpoints, Add endpoint,
Name: 94
Endpoint URL: 192.168.1.94:9001
Public IP: 192.168.1.94
Clicar no botão Add endpoint.

```
