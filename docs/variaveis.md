## Variáveis

Variáveis em ttolang são declaradas com a palavra reservada "cria" e o símbolo de [atribuição](atribuição) "<-". 
As variáveis podem armazenar [procedimentos](procedimentos) ou [tipos](tipos) básicos como [inteiros](inteiros), [booleanos](booleanos) e [strings](strings). 

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