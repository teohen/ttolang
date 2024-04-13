## Repetição
A estrutura de repetição disponível na ttolang é a repete.
Essa estrutura de repetição cria o seu próprio escopo interno mas também permite acessar aos valores do escopo "pai".

Palavra reservada: `REPETE`

Estrutura:

repete ([atribuição](atribuição.md) ate [condição](condição.md)) {

  [expressões](expressões.md)

}

Essa estrutura é avaliada pelo interpretador na seguinte ordem:
  1. Avalia a [condição](condição.md).
      1. A [condição](condição.md) é verdadeira?
          1. Se sim, ignora o loop
          2. Senão, entra no loop
  2. Dentro do loop
      1. Cria um novo escopo para [identificadores](identificadores.md).
      2. Avalia o código dentro do loop.
      3. Avalia o passo de [atribuição](atribuição.md).
          1. Define os [identificadores](identificadores.md) do passo de [atribuição](atribuição.md).
      4. Repete o primeiro passo.

Observação importante:
  A execução do código no repete está condicionada a FALSIFICAÇÃO da [condição](condição.md). Ou seja, somente vai executar o código dentro do Repete se a [condição](condição.md) for FALSA.



A estrutura de repetição cria um novo escopo específico onde gerencia o armazenamento de valores. O gerenciamento de valores, criados ou manipulados dentro do escopo do Repete segue a seguinte regra:
 1. busca o valor no escopo externo.
 2. busca o valor no escopo criado pelo Repete.
 3. resulta em um [problema](problema.md) de "indetificador não conhecido". 



Exemplos de uso:
 ```
cria i <- 0;
cria res <- 0;
repete(i <- i + 1 ate i > 9) {
  res <- i;
}
mostra(res);

// saída:
// 9
```