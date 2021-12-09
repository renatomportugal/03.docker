# GO

## Rodar

```CMD

docker-compose up -d
docker exec -it <id-do-container> bash
go mod init imersao5-gateway
mkdir domain

subir os arquivos que estão sendo modificados
go get github.com/stretchr/testify/assert
go test ./...
```

## Clean_Architecture

```CMD
Entities
- Regras de Negócios
- Domain
- Regras Universais, independente do contexto

Use Cases
- Regras da aplicação
- Fluxo da aplicação
- Input / Output

Caso de uso
- Receberá os dados de uma transação
- Criará uma transação
- Adicionará um cartão de crédito a essa transação
 - Se o cartão de crédito for inválido:
  - Os dados da transação serão inseridos no banco de dados com o status=rejected contendo a mensagem do erro
  - A transação será publicada no Kafka
- Caso a transação não seja aprovada:
 - Os dados da transação serão inseridos no banco de dados com o status=rejected contendo a mensagem do erro
 - A transação será publidada no kafka
- Caso a transação seja aprovada:
 - Os dados da transação serão inseridos no abnco de dados com o status=approved
 - A transação será publicada no Kafka

```
