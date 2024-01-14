# Desafio de Programação em Go

Neste desafio, vamos aplicar o que aprendemos sobre `webserver HTTP`, `contexts`, `banco de dados`
e `manipulação de arquivos` com Go.

Você precisará nos entregar dois sistemas em Go:

- `client.go`
- `server.go`

## Requisitos

### `client.go`

- Deverá realizar uma requisição HTTP no `server.go` solicitando a cotação do dólar.

### `server.go`

- Deverá consumir a `API` contendo o câmbio de Dólar e Real no
  endereço: `https://economia.awesomeapi.com.br/json/last/USD-BRL` e retornar no formato `JSON` o resultado para o
  cliente.
- Usando o package `context`, deverá registrar no banco de dados `SQLite` cada cotação recebida.
- O timeout máximo para chamar a API de cotação do dólar deve ser de `200ms`, e o timeout máximo para persistir os dados
  no banco deve ser de `10ms`.

### Funcionalidades Adicionais

- O `client.go` precisará receber do `server.go` apenas o valor atual do câmbio (campo `bid` do JSON).
    - Com um timeout máximo de 300ms para receber o resultado do `server.go`.
- Os três contextos devem retornar `erro` nos logs caso o tempo de execução seja insuficiente.
- O `client.go` deve salvar a cotação atual em um arquivo `"cotacao.txt"` no formato: `Dólar: {valor}`.
- O endpoint necessário gerado pelo `server.go` para este desafio será `/cotacao` e a porta a ser utilizada pelo
  servidor HTTP será a `8080`.

### Como executar

- Para executar a solução, execute o comando `go run main.go` dentro da pasta `challenges/1-client-server-api/cmd`.
- O arquivo `cotacao.txt` será criado na pasta `challenges/1-client-server-api/output/cotacao.txt` com a cotação atual
  do dólar.
- A solução foi feita para cada execução do servidor, o banco de dados é criado e populado com a cotação atual do dólar.
