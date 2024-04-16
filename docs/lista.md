## Estrutura de dados: lista
A linguagem ttolang disponibiliza uma estrutura de dados muito útil para o desenvolvimento: a lista.
Uma lista consiste em uma sequência de valores armazenados conectados que podem ser acessados por um operador de índice. 
Os índices marcam as posições dos itens na lista e fazem parte da faixa do índice 0 até o índice n (último item armazenado na lista). Ou seja, uma lista "nomes" com 5 itens terá os índices 0, 1, 2, 3, 4.

Caso aconteça uma tentativa de acessar um índice da lista que não possua valor definido resultará em um [problema](problema.md).

Exemplos de uso:
 ```
cria numeros <- [9, 1, 3, 4, 6];
mostra(numeros[0])
// saída: 9

mostra(numeros[4])
// saída: 6
```
