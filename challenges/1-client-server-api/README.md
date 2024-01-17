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

1. Certifique-se de ter o Go instalado em seu sistema.
2. Abra o terminal ou prompt de comando.
3. Navegue até o diretório `challenges/1-client-server-api/cmd` onde o `main.go` está localizado.
4. Execute o comando `go run main.go`.
5. O programa iniciará o servidor e o cliente automaticamente.
    - O servidor começará a escutar na porta 8080.
    - O cliente fará uma requisição ao servidor para obter a cotação do dólar.
6. A cotação será salva em um arquivo `cotacao.txt` no diretório `challenges/1-client-server-api/output`.
7. Logs relevantes serão exibidos no terminal durante a execução.
8. O programa será encerrado automaticamente após a execução.
9. Para executar novamente, repita os passos 3 e 4.