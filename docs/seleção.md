## Seleção

A estrutura de seleção existente na ttolang é composta por 2 partes: se, senao. A segunda parte da estrutura (senao) é opcional e somente será levada em consideração caso seja escrita. 

Palavra reservada: `se/senao`

Estrutura:

se ([expressão](expressões.md) [booleana](booleano.md)) {
  [expressões](expressões.md)
} senao { // opcional
  [expressões](expressões.md) // opcional
}

Essa estrutura é avaliada pelo interpretador na seguinte ordem:
  1. A expressão dentro dos parênteses após o se deve ser resolvida para um [booleano](booleano.md).
  2. Caso o booleano se resolva como vdd:
      1. as expressões no bloco de códigos dentro das primeiras chaves são executadas 
      2. as expressões no bloco de códigos dentro das segundas chaves são ignoradas. 
  2. Caso o booleano se resolva como falso:
      1. as expressões no bloco de códigos dentro das primeiras chaves são ignoradas.
      2. as expressões no bloco de códigos dentro das segundas chaves são executadas. 

Exemplos de uso:
 ```
cria nome <- "ttolang";

se (nome = "ttolang") {
  mostra("o nome está correto!");
} senao {
  mostra("o nome não está correto!");
}

// saída: o nome está correto!
```