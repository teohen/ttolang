## Variáveis

Variáveis em ttolang são declaradas com a palavra reservada "cria" e o símbolo de [atribuição](atribuição.md) "<-". 
As variáveis podem armazenar [procedimentos](procedimentos.md) ou [tipos](tipos.md) básicos como [inteiros](inteiros.md), [booleanos](booleanos.md) e [strings](strings.md). 

Exemplos de uso:
 ```
cria i <- 42;
mostra(i);
// saída: 42

cria somaComDois <- proc(valor) {
  mostra(valor + 2); 
};
somaComDois(5);
// saída: 7
```