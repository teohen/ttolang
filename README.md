# ttolang
A interpreted , dinamic typed, and really slow language created by teteo.

## features of ttolang
- [x] variables
- [x] integers and bools
- [x] arithmetic expressions
- [x] built-in functions
- [x] functions: first class and higher order 
- [x] closures 
- [x] string DS 
- [x] List DS 
- [ ] HashMap DS


Example code:
```
cria ano_nasc <- 1912;
cria ano_hoje <- 2024;

cria calc_tempo <- proc(ano_inicial, ano_atual) {
    cria tempo_resultado <- ano_atual - ano_inicial;
    se (tempo_resultado > 100) {
        mostra("mais de um século")
    } senao {
        mostra("menos de um seculo")
    }
    devolve tempo_resultado;
}

cria resultado <- calc_tempo(ano_nasc, ano_hoje);
mostra(resultado);

// output: 
mais de um século
112
```
