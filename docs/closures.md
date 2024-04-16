## Closures

A linguagem ttolang permite a existência de closures. 

Closures são funções que podem acessar o escopo superior (onde foram chamadas) e, com isso, acesso aos valores declarados nesse escopo. 

Exemplo de uso:
 ```
cria nome <- "ttolang";
cria mostra_nome <- proc() {
  mostra(nome)
}

// saída: ttolang
```