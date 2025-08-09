## Pacotes
Pacotes são qualquer determinado grupo de código escrito em ttolang separados em um arquivo. 
Em ttolang, é possível importar um pacote de códigos utilizandos a palavra-chave ```importa``` seguida do caminho relativo ao arquivo em que a o comando está sendo utilizado cercado por aspas duplas ("");

Para importar um pacote em ttolang é necessário seguir os seguinte critérios:
- extensão do arquivo deve ser *.tto*.
- não o caractere ponto (.) no nome do arquivo.

Caso o caminho relativo passado não aponte para um arquivo válido que siga os critérios listados acima, a operação resultará em um [problema](problema.md). 

Estrutura:

importa "caminho_relativo_do_arquivo.tto";

Exemplos de uso:
 ```
// arquivo mtm.tto
cria soma <- proc(x,y) {
  devolve x + y;
}

// arquivo principal 
importa "./mtm.tto";

cria resultado <- soma(2,2);
mostra(resultado);
// saída: 4
```
