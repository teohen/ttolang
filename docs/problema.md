## Problema

Problemas em ttolang acontecem quando algo não conhecido pela linguagem está escrito. Vários tipos de problemas podem acontecer como [operadores](operadores.md) incompatíveis com [tipos](tipos.md) utilizados, acesso a índices não definidos dentro de [listas](lista.md), algum caractere desconhecido ou não aceito pela ttolang, utilização de [identificadores](identificadores.md). 

Os problemas em ttolang impedirão que o código seja executado e interromperão a execução com uma mensagem explicando o problema acontecido. 

Exemplos:
 ```
cria nome <- "tto" / "lang";
// saída: Problema: operador desconhecido: STRING / STRING

cria nome <- "ttolang";
mostra(idade)
// saída: Problema: identificador é desconhecido: idade

cria nome <- proc({
  mostra("Errado")
)};
// saída:

Problema de parser: esperava que o próximo item fosse ')' mas recebeu 'mostra' (IDENTIFICADOR)
Problema de parser: esperava que o próximo item fosse '{' mas recebeu 'mostra' (IDENTIFICADOR)
Problema de parser: nao soube lidar com ).
Problema de parser: nao soube lidar com }.

```
