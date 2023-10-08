# CRUD-GO
Desafio proposto pela live na qual, houveram estudos da linguagem GO e a implementação de um pequeno CRUD sem integração com o banco de dados, tudo de forma local.

## Requirement

**Install Go >= [1.21.2](https://go.dev/doc/install)**

**Install Postman >= [Download](https://www.postman.com/downloads/)**

**Install SQLC >= [Doc](https://docs.sqlc.dev/en/stable/overview/install.html)**

**Install Docker >= [Doc](https://docs.docker.com/get-docker/)**

**Install Docker Compose >= [Doc](https://docs.docker.com/compose/install/)**

## Usage

Para gerar o codigo Go baseado no codigo SQL usando o SQLC  
```SH
sqlc generate
```  

Para subir o banco postgres localmente usando docker compose  
```SH
docker compose up -d
```  

O banco roda na porta **5432** e o pgadmin roda na porta **5050**

```SH
go run main.go
```
O webserver vai ser feita na porta **8888**, por lá pode se fazer as requisições da API.

Recomendado ter o insomnia ou postman em sua máquina.

## Contributing

As queries SQL estão na pasta 'sql', as migrations ficam na pasta 'sql/migrations' baseada na lib [migrate](https://github.com/golang-migrate/migrate) e as queries na pasta 'sql/queries'.

Pull requests são bem vindos para adicionar mensagens de erros, por exemplo

## License

[Veja o arquivo de licença](https://github.com/stherzada/CRUD---GO/blob/main/LICENSE)
