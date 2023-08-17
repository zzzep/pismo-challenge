# Pismo Challenge

## How to Run

```shell script
docker compose up --build
```

## How to use

### DOC
```shell
godoc -http=:6060
```
http://localhost:6060/pkg/github.com/zzzep/pismo-challenge/

#### Swagger
[GET] http://localhost/swagger/index.html

### Endpoints
#### [POST] http://localhost/accounts (criação de uma conta)
```
Request Body:
{
"document_number": "000.000.001-91"
}
```

#### [GET] http://localhost/accounts/:accountId (consulta de informações de uma conta)
```
Response Body:
{
"account_id": 1,
"document_number": "000.000.001-91"
}
```

#### [POST] http://localhost/transactions (criação de uma transação)
```
Request Body:
{
"account_id": 1,
"operation_type_id": 4,
"amount": 123.45
}
```






