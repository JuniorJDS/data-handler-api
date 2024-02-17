# data-handler-api

<p align="center">
  <img src="https://img.shields.io/badge/Golang-v1.21.6-blue"/>
  <img src="https://github.com/JuniorJDS/data-handler-api/actions/workflows/e2e.yaml/badge.svg">
  <img src="https://github.com/JuniorJDS/data-handler-api/actions/workflows/golangci-lint.yaml/badge.svg">
</p>

API responsável por fazer a manipulação dos dados de um arquivo csv/txt e persistir numa base de dados relacional.

Código escrito na linguagem <a href="https://go.dev/" target="_blank">Golang</a> na versão 1.21.6, mais detalhes são descritos abaixo:

## Como Executar:

É possível executar a API, tal como seus testes, através do Docker ou no ambiente local, também foi criado um arquivo `Make` para facilitar algumas execuções. Além disso, para as execuções fora do Docker, atente-se para o arquivo `.env` que deve estar na pasta raiz do projeto. Caso deseje debugar os testes, também deve ter um arquivo `.env` na pasta `/tests/integration`.

### Comandos Make:

- `make run-local-environment`: Inicia a API através do Docker, juntamente com os outros serviços necessários;
- `make run`: Inicia, unicamente, a API;
- `make integration-tests`: Inicia um ambiente no Docker e roda todos os testes de integração;

### Docker:

Para rodar utilizando o Docker, basta subir o arquivo `docker-compose-run-local.yaml` com o docker-compose:

```
docker-compose -f docker-compose-run-local.yaml build
docker-compose -f docker-compose-run-local.yaml up
```

ou utilizando o comando `make`:

```
make run-local-environment
```
