#Host_Name

```CMD
Vá no arquivo deployment.toml e ajuste a linha
hostname = "lab-106"
node_ip = "192.168.1.106"

Para sua máquina reconhecer o hostname edite o arquivo
C:\Windows\System32\drivers\etc\hosts

Adicione na última linha seu IP e o host que você quer que seja reconhecido:
192.168.1.106 lab-106

Vá no arquivo deployment.yaml
Procure a linha "# <IP>:<HTTPS Port> of the Worker node".
Altere para
lab-106:9444:

Procure as linhas e altere
kmTokenUrlForRedirection: https://lab-106:9443/oauth2
baseUrl: https://lab-106:9643
externalLogoutUrl: https://lab-106:9443/oidc/logout

Vá no arquivo deployment.yaml e altere
receiver.url : grpc://lab-106:9806/org.wso2.analytics.mgw.grpc.service.AnalyticsSendService/sendAnalytics

Procure e altere
#    host: lab-106

Procure e altere
#    advertisedHost: lab-106

Acesse
https://lab-106:9443/publisher
https://lab-106:9443/devportal
https://lab-106:9443/admin
https://lab-106:9443/carbon


```
