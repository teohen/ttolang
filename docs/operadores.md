## [operadores](operadores) 

[Operadores](operadores) permitem realizar operações lógicas e matemáticas nos [tipos](tipos) existentes na ttolang. Alguns [operadores](operadores) somente podem ser utilizados em [tipo](tipos) específicos da ttolang. 
Os [operadores](operadores) seguem a seguinte ordem de precedência em ttolang

    1. IGUAL A, DIFERENTE DE
  	2. MAIOR QUE, MENOR QUE
  	3. SOMA, SUBTRAÇÃO
  	4. MULTIPLICAÇÃO, DIVISÃO
  	5. NEGAÇÃO

Os [operadores](operadores) existentes na ttolang são representados por símbolos e estão dividos em duas categorias: matemáticos e lógicos.
1. Matemáticos
    1. Os [operadores](operadores) matemáticos podem ser utilizados em operações interfixadas e podem ser aplicados aos [tipos](tipos) matemáticos: [inteiros](inteiros).
    2. Lista dos [operadores](operadores):
        1. \+ (SOMA)
        2. \- (SUBTRAÇÃO)
        3. \* (MULTIPLICAÇÃO)
        4. / (DIVISÃO)
2. Lógicos
    1. Os [operadores](operadores) lógicos poder ser utilizados em operações interfixadas e pré-fixadas que podem ser aplicadas aos [tipos](tipos) [inteiro](inteiros)* e [booleano](booleanos). 
    2. Lista dos [operadores](operadores):
        1. \> (MAIOR QUE)
        2. < (MENOR QUE)
        3. = (IGUAL A)
        4. != (DIFERENTE DE)
        5. ! (NEGAÇÃO)

OBS1: alguns [operadores](operadores) lógicos não podem ser utilizados com alguns [tipos](tipos) específicos. Por exemplo, o [operador](operadores) lógico de NEGAÇÃO (!) não pode ser utilizado com o [tipo](tipos) [inteiro](inteiros). Ao contrário do [operador](operadores) lógico MAIOR QUE (>) pode ser utilizado com o [tipo](tipos) [inteiro](inteiros).

OBS2: Se o [operador](operadores) DIVISÃO (/) for utilizado em uma operação onde o valor 0 (zero) é o dividendo da operação, o interpretador resultará em um [problema](problema).

Exemplos de uso:
 ```
10 > 5
// saída: vdd

10 + 5;
// saída: 15

10 / 5;
// saída: 2

10 * 5;
// saída: 50

10 != 5;
// saída: vdd

!vdd;
// saída: falso
```