# Neolabs Workshop Golang

O projeto depende de um banco de dados postgres para o funcionamento correto, para isso podemos utilizar o `docker-compose.yml` já disposto. 

Vamos rodar o seguinte comando para subir uma nova instância do postgres:

```bash
docker-compose up -d
```

Assim que tivermos um container do postgres rodando com o nosso banco de dados, vamos copiar o arquivo `.env.example` para `.env` e rodar a API com os seguintes comandos.

```bash
$ cp .env.example .env
$ go run cmd/neolabs-workshop-golang/neolabs-workshop-api.go
```

# Coleção do Postman com exemplos de requisições para a API.

Estamos disponibilizando uma [collection do postman](https://www.getpostman.com/collections/2c0c9db6dd2a487aa3a8) pública do postman que contém os exemplos de como gerar as requisições necessárias para a API desse workshop.