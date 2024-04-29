## Estrutura de dados: Estrutura
A linguagem ttolang disponibiliza uma estrutura de dados muito útil para o desenvolvimento: a Estrutura.
Uma estrutura consiste em um conjunto propriedades com chaves/valores armazenados e que podem ser acessados por um operador de índice. 
Os índices marcam os identificadores dos itens na estrutura, com as chaves sendo, obrigatoriamente, uma [string](string.md) e os valores podendo assumir qualquer tipo disponível na ttolang.

Caso aconteça uma tentativa de acessar um índice da estrutura que não possua valor definido resultará em um [problema](problema.md).

É possível adicionar novas propriedades a uma estrutura passando a ela, o chave da nova propriedade propriedade e o valor da nova propriedade, respectivamente, para o [procedimento integrado](procedimentos.md#procedimentos-integrados) "anexar". Esse procedimento irá devolver uma nova estrutura e não vai alterar a primeira estrutura.


Exemplos de uso:
 ```
cria pessoa <- {"nome" <- "linus", "cod" <- 1};

mostra(pessoa["nome"])
// saída: linus

mostra(pessoa["cod"])
// saída: 1

pessoa_com_proc <- anexar(pessoa, "operacao", proc(nome){mostra(nome)})
mostra(pessoa_com_proc["operacao"]("novo nome"))
// saída: novo nome
```
