## Operadores

Operadores permitem realizar operações lógicas e matemáticas nos [tipos](tipos.md) existentes na ttolang. Alguns operadores somente podem ser utilizados em [tipo](tipos.md) específicos da ttolang. 
Os Operadores seguem a seguinte ordem de precedência em ttolang, onde quanto MENOR a prioridade de precedência, mais acima na listagem:

    1. E, OU
    2. IGUAL A, DIFERENTE DE
  	3. MAIOR QUE, MENOR QUE
  	4. SOMA, SUBTRAÇÃO
  	5. MULTIPLICAÇÃO, DIVISÃO
  	6. NEGAÇÃO

Os Operadores existentes na ttolang são representados por símbolos e estão dividos em duas categorias: matemáticos e lógicos.
1. Matemáticos
    1. Os Operadores matemáticos podem ser utilizados em operações interfixadas e podem ser aplicados aos [tipos](tipos.md) matemáticos: [inteiros](inteiros.md).
    2. Lista dos Operadores:
        1. \+ (SOMA)
        2. \- (SUBTRAÇÃO)
        3. \* (MULTIPLICAÇÃO)
        4. / (DIVISÃO)
2. Lógicos
    1. Os Operadores lógicos poder ser utilizados em operações interfixadas e pré-fixadas que podem ser aplicadas aos [tipos](tipos.md) [inteiro](inteiros.md)* e [booleano](booleanos.md). 
    2. Lista dos Operadores:
        1. \> (MAIOR QUE)
        2. < (MENOR QUE)
        3. = (IGUAL A)
        4. != (DIFERENTE DE)
        5. ! (NEGAÇÃO)
        6. & (E)
        7. | (OU)

OBS1: alguns operadores lógicos não podem ser utilizados com alguns [tipos](tipos.md) específicos. Por exemplo, o operador lógico de NEGAÇÃO (!) não pode ser utilizado com o [tipo](tipos.md) [inteiro](inteiros.md). Ao contrário do operador lógico MAIOR QUE (>) pode ser utilizado com o [tipo](tipos.md) [inteiro](inteiros.md).

OBS2: Se o operador (/) DIVISÃO  for utilizado em uma operação onde o valor 0 (zero) é o dividendo da operação, o interpretador resultará em um [problema](problema.md).

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