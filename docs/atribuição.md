## Atribuição
A estrutura de atribuição em ttolang utiliza o símbolo "<-" para atribuir valores variáveis. 
Esses valores podem ser [procedimentos](procedimentos.md), [inteiros](inteiros.md), [strings](strings.md), [booleano](booleano.md). 

A atribuição feita a uma variável que não exista em um escopo alcancável pelo código resultará em um [problema](problema.md). 

Estrutura:

[identificador](identificador.md) <- valor;

Exemplos de uso:
 ```
cria nome <- "tto";
nome <- "ttolang";
mostra(nome)

// saída: ttolang
```