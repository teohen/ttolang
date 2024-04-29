## Procedimentos: primeira classe e ordem superior 

DOC funções: first class e higher order


## Procedimentos integrados

A linguagem ttolang possui procedimentos padrões integrados a que permitem realizar  operações úteis para o desenvolvimento. Esses procedimentos podem ser aplicados aos [tipos](tipos.md) [inteiros](inteiros.md) e [string](string.md). 

A lista de procedimentos integrados é a seguinte: 
### tam()
O procedimento tam() recebe 1 parâmetro que pode ser do [tipo](tipos.md) [string](string.md) ou [lista](lista.md). Esse procedimento retorna a quantidade de itens dentro da [lista](lista.md) ou a quantidade de caracteres em uma [string](string.md).

Exemplos de uso:
 ```
cria numeros <- [9, 1, 3, 4, 6, 6];
tam(numeros);
// saída: 6

cria nome <- "ttolang";
tam(nome);
// saída: 7
```

### anexar()
O procedimento anexar() adiciona novos valores nas estruturas de dados da ttolang. Esse procedimento integrado recebe uma quantidade de parâmetros diferente baseando-se no [tipo](tipos.md) do primeiro parâmetro passado. 

#### anexar(Lista, param2)
Caso o primeiro parâmetro seja uma [lista](lista.md), o procedimento vai devolver uma nova [lista](lista.md) copiando todos os itens da [lista](lista.md) passada no primeiro parâmetro e anexando o item do segundo parâmetro ao final nova lista. Ainda nesse caso, o segundo parâmetro pode ser qualquer outro [tipo](tipo.md) disponível na ttolang.

Exemplos de uso:
 ```
cria numeros <- [1, 2];
numeros <- anexar(numeros, 3);
numeros <- anexar(numeros, "quatro");
numeros <- anexar(numeros, vdd);
mostra(numeros)
// saída: [1, 2, 3, "quatro", vdd]
```
#### anexar(String, param2)
Caso o primeiro parâmetro seja uma [string](string.md), o procedimento irá devolver uma nova [string](string.md) contendo a junção das duas [strings](string.md) passadas por parâmetro. Ainda nesse caso, o segundo parâmetro OBRIGATORIAMENTE deve ser do tipo [string](string.md), resultando em um [problema](problema.md) se outro [tipo](tipo.md) for passado.

Exemplos de uso:
 ```
cria nome <- "tto";
novo_nome <- anexar(nome, "lang");
mostra(novo_nome)
// saída: ttolang
```

#### anexar(Estrutura, String, param3)
Caso o primeiro parâmetro seja uma [estrutura](estrutura.md), o procedimento irá devolver uma nova [estrutura](estrutura.md) contendo as propriedades existentes na [estrutura](estrutura.md) passada no primeiro parâmetro, assim como a nova propriedade que está sendo definida. O segundo parâmetro, que deve ser, obrigatoriamente, uma [string](string.md), servirá como identificador da nova propriedade. Já o terceiro parâmetro será o valor da nova propriedade podendo ser qualquer outro [tipo](tipo.md) disponível na ttolang. 
Caso as regras dos parâmetros não sejam seguidas, o código resultará em um [problema](problema.md).

Exemplos de uso:
 ```
cria estrutura <- {};
cria nova_estrutura <- anexar(estrutura, "nome", "ttolang");
mostra(nova_estrutura["nome"])
// saída: ttolang
```