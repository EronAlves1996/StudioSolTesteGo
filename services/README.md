# Teste Studio Sol

Este aplicativo foi feito conforme requerimentos para o processo seletivo da Studio Sol.

## Endpoints

Com a aplicação em execução, segue os endpoints disponíveis para acesso aos recursos:

- GraphiQL: `http://localhost:8080/`
- Para realizar requisições: `http://localhost:8080/graphql`

## Rodando a aplicação

A aplicação pode ser rodada de algumas formas. Foi feita uma dockerização da aplicação para facilitar o deploy. Estarão listadas da forma mais simples para a mais complexa.

### Bash scripts

Os bash scripts tem por base o docker, porém são uma forma facilitada de subir a aplicação sem deixar rastros de staging, deixando somente a imagem do Alpine Linux com a aplicação compilada.

Para executar, só rodar o seguinte comando:

```bash
$ ./dockerize.sh
```

O container será construído, ao final será executado e após isso, os rastros de multistaging build serão apagados.

Para controlar a execução da aplicação (parar e continuar) use o docker compose:

```bash
# Para parar a execução da aplicação
$ docker-compose stop

# Para continuar a execução da aplicação
$ docker-compose start
```

Para retirada completa dos containers e imagens geradas, basta executar o script bash:

```bash
$ ./undockerize.sh
```

### Docker-compose

A aplicação pode ser construída através de Docker-compose, sendo necessário somente utilizar os comando do docker-compose. Para esse tipo de execução, vestígios de multistaging build não são apagados após a construção das imagens e execução dos containers (consultar o script `dockerize`):

```bash
# Para construir
$ docker-compose up

# Para destruir
$ docker-compose down
```

### Execução sem Docker

A aplicação também pode ser montada sem o docker através de uma sequência de comandos. Primeiro de tudo, as dependências declaradas no `go.mod`devem ser baixadas e consistidas:

```bash
$ go mod tidy
```

Após isso, basta executar diretamente o arquivo server.go:

```bash
$ go run server.go
```

Após isso, o servidor estará executando nos endpoints padrão.
