## Procedimentos: primeira classe e ordem superior 

DOC funções: first class e higher order


## Procedimentos integrados

A linguagem ttolang possui procedimentos padrões integrados a que permitem realizar  operações úteis para o desenvolvimento. Esses procedimentos podem ser aplicados aos [tipos](tipos) [inteiros](inteiros) e [string](string). 

A lista de procedimentos integrados é a seguinte: 
### tam()
O procedimento tam() recebe 1 parâmetro que pode ser do [tipo](tipos) [string](string) ou [lista](lista). Esse procedimento retorna a quantidade de itens dentro da [lista](lista) ou a quantidade de caracteres em uma [string](string).

Exemplos de uso:
 ```
cria numeros <- [9, 1, 3, 4, 6, 6];
tam(numeros);
// saída: 6

cria nome <- "ttolang";
tam(nome);
// saída: 7
```