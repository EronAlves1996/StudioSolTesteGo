# Teste Studio Sol

Este aplicativo foi feito conforme requerimentos para o processo seletivo da Studio Sol.
O servidor foi escrito para suportar requisições GraphQL.

## Stack

- Go
- Gqlgen

## Endpoints

Com a aplicação em execução, segue os endpoints disponíveis para acesso aos recursos:

- GraphiQL: `http://localhost:8080/`
- Para realizar requisições: `http://localhost:8080/graphql`

## Lógica de execução

1. Primeiramente, o servidor somente possui um schema o qual ele suporta. No seguinte formato:

```graphql

query {
    verify(password: String!, rules: [Rule!]): Return
}

```

O tipo Rule é descrito da seguinte forma:

```graphql
input Rule {
  rule: String!
  value: Int!
}
```

O tipo Return é descrito da seguinte forma:

```graphql
type Return {
  verify: Boolean
  noMatch: [String]
}
```

Este schema está descrito melhor no arquivo /graph/schema.graphqls. O servidor irá receber esta requisição e passar a informação dos parâmetros `password` e `rules` para a função ProcessPassword, dentro do package `services`.

2. e irá verificar a senha contra um Map com funções o qual denominei Validators. Essa função são executadas e, se retornarem False, adiciona-se o nome da rule na slice do noMatch. Para cada rule é feita uma iteração, no qual o password passa pela função de validator. Dentro do validator, itera-se sobre a string, verificando se a quantidade de caracteres satisfaz a regra.

3.

## Rodando a aplicação

A aplicação pode ser rodada de algumas formas. Foi feita uma dockerização da aplicação para facilitar o deploy. Estarão listadas da forma mais simples para a mais complexa.

### Bash scripts

Os bash scripts tem por base o docker, porém são uma forma facilitada de subir a aplicação sem deixar rastros de staging, deixando somente a imagem do Alpine Linux com a aplicação compilada.

Para executar, só rodar o seguinte comando:

```sh
$ ./dockerize.sh
```

O container será construído, ao final será executado e após isso, os rastros de multistaging build serão apagados.

Para controlar a execução da aplicação (parar e continuar) use o docker compose:

```sh
# Para parar a execução da aplicação
$ docker-compose stop

# Para continuar a execução da aplicação
$ docker-compose start
```

Para retirada completa dos containers e imagens geradas, basta executar o script bash:

```sh
$ ./undockerize.sh
```

### Docker-compose

A aplicação pode ser construída através de Docker-compose, sendo necessário somente utilizar os comando do docker-compose. Para esse tipo de execução, vestígios de multistaging build não são apagados após a construção das imagens e execução dos containers (consultar o script `dockerize`):

```sh
# Para construir
$ docker-compose up

# Para destruir
$ docker-compose down
```

### Execução sem Docker

A aplicação também pode ser montada sem o docker através de uma sequência de comandos. Primeiro de tudo, as dependências declaradas no `go.mod`devem ser baixadas e consistidas:

```sh
$ go mod tidy
```

Após isso, basta executar diretamente o arquivo server.go:

```sh
$ go run server.go
```

Após isso, o servidor estará executando nos endpoints padrão.
