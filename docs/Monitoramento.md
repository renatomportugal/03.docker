# Monitoramento

# Instalação CentOS
## Git
```
sudo yum install git
```

Caso ainda não tenha instalado o docker:<br>
```
curl -fsSL https://get.docker.com | bash
```
### Clonar o projeto
```
git clone https://github.com/badtuxx/giropops-monitoring.git
```
### Ajustar a interação com o Slack
```
vim conf/alertmanager/config.yml
```
Preencha:<br>
username(usuario), channel (#algumaCoisa), api_url(webhook... algo como https://hooks.slack.com/services/.../...)<br>

## Netdata
```
bash <(curl -Ss https://my-netdata.io/kickstart.sh)
```
Agora acesse o localhost:19999<br>

## Prometheus
### Configurar
```
vim conf/prometheus/prometheus.yml
```
Na última linha em static_configs: -targets:[localhost:19999]

## Docker Composer
### Fazer o Deploy
```
docker swarm init
cat docker-compose.yml
docker stack deploy -c docker-compose.yml giropops
```

### Verifica se subiu
```
docker service ls
```
Acessar o prometheus em localhost:9090<br>
Acessar o alert manager em localhost:9093<br>
Acessar o alert grafana em localhost:3000, user admin, senha giropops<br>
Acessar o alert node exporter em localhost:9100/metrics<br>

### Iniciando o docker swarm
```
docker swarm init
```

## Alerts
```
cat conf/prometheus/alert.rules
```


Verficar<br>
```
watch -n 2 'docker ps --format "table {{.ID}}\t {{.Image}}\t {{.Status}}"'
```

# Logs
Veja os últimos 1000 erros com o comando:<br>
```
docker logs <id> --tail 1000
```

# Grafana
```
https://grafana.com/tutorials/
```
## 1. Introduction
https://grafana.com/tutorials/grafana-fundamentals/#1<br>

Prerequisites<br>
```
https://docs.docker.com/install/
https://docs.docker.com/compose/
https://git-scm.com/
```
## 2. Set up the sample application
https://grafana.com/tutorials/grafana-fundamentals/#2<br>
```
cd ~
mkdir grafana
cd grafana
git clone https://github.com/grafana/tutorial-environment.git
cd tutorial-environment
docker ps
docker-compose up -d
docker-compose ps
```
### 2.1. Acessar a porta 8081
```
http://localhost:8081
Title = Example.
URL= https://example.com.
Submit
```
## 3. Login no Grafana
https://grafana.com/tutorials/grafana-fundamentals/#3<br>

### 3.1. Acessar a porta 3000
```
Abrir outra aba
localhost:3000

admin
admin
Log In.

Mude o password para admin2
Confirme o password para admin2
Salvar
```
## 4. Adicionar fonte de dados para métricas
https://grafana.com/tutorials/grafana-fundamentals/#4<br>
```
Na barra vertical lateral esquerda, vá em Configurações (ícone engrenagem) e clique em Data Sources.

Botão Add data source, Prometheus (Botão Select).
Coloque caixa URL insira:
http://prometheus:9090
Vá em baixo, clique no Botão Save & Test.
Vai aparecer uma notificação "Data Source is Working"
```
### 4.1. Explorando as métricas
https://grafana.com/tutorials/grafana-fundamentals/#5<br>
```
Na barra vertical lateral esquerda, vá em Explore (ícone bússola).
No campo de entrada escrito "Enter a PromQL query" (Ao lado de Metrics), digite:
tns_request_duration_seconds_count
Aperte Shift + Enter

No canto superior direito, clique na seta do dropdown (Run Query) e selecione 5s

Altere a Query
No campo de entrada escrito "Enter a PromQL query", digite:
rate(tns_request_duration_seconds_count[5m])

Note que a escala da esquerda mudará.

Mudar a legenda (mostrar por rota)
No campo de entrada escrito "Enter a PromQL query", digite:
sum(rate(tns_request_duration_seconds_count[5m])) by(route)

Vá na aba que tem a porta 8081 e dê um refresh na página para gerar tráfego (clique na barra de endeço e clique enter).

No canto superior direito selecione o tempo "Last 5 minutes" (ícone de relógio). É mais fácil para ver quando recebe novos dados.
Clique na legenda para selecionar por rota.
```

## 5. Adicionar uma fonte de log
https://grafana.com/tutorials/grafana-fundamentals/#6<br>
```
Na barra vertical lateral esquerda, vá em Configurações (ícone engrenagem) e clique em Data Sources.

Clique no botão Add data source.
Selecione o Loki clicando no Botão Select.
Coloque caixa URL insira:
http://loki:3100
Vá em baixo, clique no Botão Save & Test.
Vai aparecer uma notificação "Data Source connected and labels found"

```
## 6. Explorando os logs
https://grafana.com/tutorials/grafana-fundamentals/#7<br>
```
Na barra vertical lateral esquerda, vá em Explore (ícone bússola).
Na lista de fonte de dados localizado no topo selecione "Loki".
Na caixa de entrada de Query insira (ao lado de Log labels):

{filename="/var/log/tns-app.log"}
Aperte Shift+Enter

Grafana mostra todos os logs do arquivo de log da aplicação de amostra. A altura de cada barra mostra a quantidade de logs que foram gerados, no tempo correspondente.

Selecione um intervalo clicando e arrastando o mouse dentro do gráfico (da esquerda para a direita).
Loki se baseia nas etiquetas e nas ocorrências específicas.

Vamos gerar um erro para analisar:
Na aplicação de amostra (aba da porta 8081) insira um novo link sem a URL, o que gerará um erro do tipo "empty url".

Vá no Grafana (Aba 3000) e entre a seguinte query:

{filename="/var/log/tns-app.log"} |= "error"
No canto superior direito, clique na seta do dropdown (Run Query) e selecione 5s
Se não atualizar pressione Shift + Enter.
Se não aparecer o erro, selecione a última hora de eventos.

Clique na linha onde diz level=error msg="empty url" para visualizar informações do erro.

Logs são úteis para entender o que está errado. 
```

## 7. Construindo um dashboard
https://grafana.com/tutorials/grafana-fundamentals/#8<br>
```
O dashboard nos dá uma visão melhor e permite rastrear metricas de diferentes visualizações.
Dashboards são formados de painéis, cada um representando algo a contar.

Cada painel consiste numa visualização de uma query. A query define o que queremos mostrar, enquanto que a visualização define como os dados serão mostrados.

Na barra vertical lateral esquerda, vá em Create (ícone de sinal de mais), e clique em Dashboad.

Clique no botão Add new panel.

No Query editor abaixo do gráfico adicione a query (Ao lado de Metrics):
sum(rate(tns_request_duration_seconds_count[5m])) by(route)
Pressione Shift + Enter
No campo Legend, entre {{route}} para renomear o tempo na legenda do gráfico. O Gráfico atualiza quando clica fora do campo.

No editor do painel localizado à direita, mude o título para Traffic.
Clique no botão Apply localizado no canto superior direito e volte para a Visão de dashboard.

Clique no ícode de disquete para salvar o dashboard.
Coloque o nome "Dashboard Traffic" e clique no botão Save.
```

## 8. Annotate events
https://grafana.com/tutorials/grafana-fundamentals/#9<br>
```
Quando as coisas vão mal, é necessário enender o contexto da falha. Tempo da última implantação, mudanças de sistemas, ou migração de bases de dados podem oferecer uma visão sobre o que pode ter causado uma interrupção. As anotações permitem que você represente tais eventos diretamente em seus gráficos.

Para adicionar uma anotação:
Clique em qualquer lugar do gráfico (Botão Esquerdo) e, a seguir, clique em Adicionar anotação.

Em Descrição, insira "Banco de dados do usuário migrado".

Click Save.

Grafana adiciona sua anotação ao gráfico. Passe o mouse sobre a base da anotação para ler o texto (No triângulo).

Grafana também permite anotar um intervalo de tempo, com anotações de região.

Adicionar uma anotação de região:

Pressione Ctrl, clique e arraste pelo gráfico para selecionar uma área.
Em Descrição, insira "Testes de carga realizados".
Em Tags, insira "testing".

Anotar manualmente seu painel é adequado para esses eventos únicos. Para eventos que ocorrem regularmente, como a implantação de uma nova versão, o Grafana oferece suporte à consulta de anotações de uma de suas fontes de dados. Vamos criar uma anotação usando a fonte de dados Loki que adicionamos anteriormente.

Na parte superior do painel, clique no ícone Configurações do painel (engrenagem - Dashboard settings).

Vá para Anotações (Menu vertical esquerdo) e clique em Adicionar consulta de anotação (Botão Add Annotation Query).

Em Nome, insira Erros.

Em Fonte de dados, selecione Loki.

Em Consulta, insira a seguinte consulta:

{filename="/var/log/tns-app.log"} |= "error"
Clique no Botão Add. Grafana exibe a lista de Anotações, com sua nova anotação.

Clique na seta Voltar (Canto superior esquerdo) para retornar ao seu painel.

As linhas de registro retornadas por sua consulta agora são exibidas como anotações no gráfico.

Ser capaz de combinar dados de várias fontes de dados em um gráfico permite correlacionar informações do Prometheus e do Loki.
```

## 9. Configure um alerta
```
Os alertas permitem que você identifique problemas em seu sistema momentos após eles ocorrerem. Ao identificar rapidamente mudanças não intencionais em seu sistema, você pode minimizar interrupções em seus serviços.

Os alertas consistem em duas partes:

Canal de notificação - Como o alerta é entregue. Quando as condições de uma regra de alerta são atendidas, o Grafana notifica os canais configurados para aquele alerta.
Regras de alerta - quando o alerta é acionado. As regras de alerta são definidas por uma ou mais condições que são avaliadas regularmente pela Grafana.
```

## 10. Configure um canal de notificação
```
Nesta etapa, você enviará alertas usando web hooks. Para testar seus alertas, primeiro você precisa ter um local para enviá-los.

Navegue até requestbin.com.

Clique em Usar versão antiga (que está disponível publicamente).
https://requestbin.com/r

Faça o login com o google (acho que não precisava).

Copie o endpoint...


Configurando uma notification channel:


In the side bar, hover your cursor over the Alerting (bell) icon and then click Notification channels.
Click Add channel.
In Name, enter RequestBin.
In Type, select webhook.
In Url, paste the endpoint to your request bin.

https://endpoint.m.pipedream.net

Click no Botão Test para enviar um alerta to your request bin.
Acesse a aba do endpoint e veja que haberá uma entrada de POST.
Click Save.
```

## 11. Configure an alert rule
```
Now that Grafana knows how to notify you, it’s time to set up an alert rule:

In the sidebar, click Dashboards -> Manage.

Click the dashboard you created earlier.

Click the Traffic panel title and then click Edit.

Click the Alert tab under the graph to access the settings for alerting.

Click Create Alert.

In Name, enter My alert.

In Evaluate every, enter 5s. For the purpose of this tutorial, the evaluation interval is intentionally short to make it easier to test.

In For, enter 0m. This setting makes Grafana wait until an alert has fired for a given time, before Grafana sends the notification.

In the Conditions section, you can change the white text by clicking it and then choosing a new option or typing text in the blank field.

Change the alert condition to:

WHEN last() OF query(A, 1m, now) IS ABOVE 0.2
This condition evaluates whether the last value of any of the time series from one minute back is more than 0.2 requests per second.

In the Notifications section, click the plus sign (+) next to Send to and then click RequestBin.

Click the Go back arrow and then click the Save dashboard icon.

Enter a note to describe your changes. Something like, Added My alert. Notes like this are optional, but can be useful when trying to understand changes made to dashboards over time.

You now have an alert set up to send a notification using a web hook. See if you can trigger it by generating some traffic on the sample application.

Browse to localhost:8081.

Refresh the page repeatedly to generate traffic.

When Grafana triggers the alert, it sends a request to the web hook you set up earlier.

Browse to the Request Bin you created earlier, to inspect the alert notification.
```

## 12. Pause an alert
```
Once you’ve acknowledged an alert, consider pausing it. This can be useful to avoid sending subsequent alerts, while you work on a fix.

In the Grafana side bar, hover your cursor over the Alerting icon and then click Alert Rules. All the alert rules configured so far are listed, along with their current state.
Find your alert in the list, and click the Pause icon on the right. The Pause icon turns into a Play icon.
Click the Play icon to resume evaluation of your alert.

```