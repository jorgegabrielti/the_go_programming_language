# Exercícios 1.2 - Argumentos de Linha de Comando

Este diretório contém as soluções para os exercícios do capítulo 1.2 do livro "The Go Programming Language".

## 📚 Exercícios

### Exercício 1.1 - Exibir Nome do Comando

**Arquivo:** [`echo_exe1.go`](echo_exe1.go)

**Objetivo:** Modificar o programa echo para exibir também `os.Args[0]`, que é o nome do comando que o chamou.

**Solução:**
```go
fmt.Println(strings.Join(os.Args[0:], " "))
```

**Explicação:**
- `os.Args[0]` contém o nome do executável/comando
- `os.Args[0:]` inclui todos os argumentos, começando do índice 0 (o nome do comando)
- Diferente de `os.Args[1:]` que pula o nome do comando

**Como executar:**
```bash
go run echo_exe1.go arg1 arg2 arg3
```

**Saída esperada:**
```
C:\Users\...\echo_exe1.exe arg1 arg2 arg3
```

---

### Exercício 1.2 - Exibir Índice e Valor

**Arquivo:** [`echo_exe2.go`](echo_exe2.go)

**Objetivo:** Modificar o programa echo para exibir o índice e o valor de cada um de seus argumentos, um por linha.

**Solução:**
```go
for index, arg := range os.Args {
    fmt.Println(index, arg)
}
```

**Explicação:**
- O `range` retorna dois valores: índice e valor
- `index` é a posição do argumento (0, 1, 2, ...)
- `arg` é o valor do argumento
- Cada argumento é exibido em uma linha separada

**Como executar:**
```bash
go run echo_exe2.go hello world golang
```

**Saída esperada:**
```
0 C:\Users\...\echo_exe2.exe
1 hello
2 world
3 golang
```

---

### Exercício 1.3 - Benchmark de Performance

**Arquivos:** 
- [`echo_exe3.go`](echo_exe3.go) - Versão básica
- [`echo_exe3_benchmark.go`](echo_exe3_benchmark.go) - Versão completa de benchmark

**Objetivo:** Medir a diferença de tempo de execução entre versões potencialmente ineficientes e a versão que usa `strings.Join`.

**Versões Comparadas:**

1. **Versão 1 - For Indexado (Ineficiente)**
   ```go
   var s, sep string
   for i := 1; i < len(os.Args); i++ {
       s += sep + os.Args[i]
       sep = " "
   }
   ```
   - Usa índice manual
   - Concatenação com `+=`
   - Cria nova string a cada iteração

2. **Versão 2 - Range (Ineficiente)**
   ```go
   var s, sep string
   for _, arg := range os.Args[1:] {
       s += sep + arg
       sep = " "
   }
   ```
   - Usa `range` (mais idiomático)
   - Ainda usa concatenação com `+=`
   - Também cria nova string a cada iteração

3. **Versão 3 - strings.Join (Eficiente)**
   ```go
   s := strings.Join(os.Args[1:], " ")
   ```
   - Calcula tamanho total necessário antecipadamente
   - Aloca memória uma única vez
   - Copia todas as strings de uma vez

**Como executar:**

```bash
# Versão básica (precisa de muitos argumentos para ver diferença)
go run echo_exe3.go arg1 arg2 arg3 ... arg100

# Versão benchmark (não precisa de argumentos)
go run echo_exe3_benchmark.go
```

**Resultados do Benchmark:**
```
=== Benchmark: 1000 strings, 1000 iterações ===
Versão 1 (for indexado):     1.20 segundos
Versão 2 (range):            1.15 segundos
Versão 3 (strings.Join):     6.77 milissegundos

=== Comparação ===
Versão 1 é 176.86x mais lenta que strings.Join
Versão 2 é 169.23x mais lenta que strings.Join
```

**Por que strings.Join é tão mais rápida?**

- **Strings são imutáveis em Go**: Cada `s += texto` cria uma nova string
- **Múltiplas alocações**: Com N strings, versões 1 e 2 fazem N alocações
- **strings.Join otimizado**: 
  1. Calcula tamanho total: `len(str1) + len(sep) + len(str2) + ...`
  2. Aloca memória uma vez: `make([]byte, tamanhoTotal)`
  3. Copia tudo sequencialmente
  4. Resultado: 1 alocação vs N alocações

**Lição aprendida:** Para concatenar múltiplas strings, sempre prefira `strings.Join` ou `strings.Builder` em vez de concatenação com `+=`.

---

## 🚀 Como Executar Todos os Exercícios

```bash
# Navegar até o diretório
cd "chapter_1-tutorial/1.2_ argumentos_de_linha_de_comando/1.2_argumentos_de_linha_de_comandos_exercicios"

# Exercício 1.1
go run echo_exe1.go teste de argumentos

# Exercício 1.2
go run echo_exe2.go um dois três

# Exercício 1.3 - Versão básica
go run echo_exe3.go palavra1 palavra2 palavra3

# Exercício 1.3 - Benchmark completo
go run echo_exe3_benchmark.go
```

## 📖 Conceitos Aprendidos

1. **`os.Args`**: Slice que contém argumentos da linha de comando
   - `os.Args[0]`: Nome do executável
   - `os.Args[1:]`: Argumentos passados pelo usuário

2. **`range`**: Itera sobre slices retornando índice e valor
   ```go
   for index, value := range slice { ... }
   ```

3. **Imutabilidade de Strings**: Strings em Go são imutáveis
   - Concatenação com `+=` cria novas strings
   - Para múltiplas concatenações, use `strings.Join` ou `strings.Builder`

4. **Performance**: Sempre considere o custo de operações repetidas
   - Medir com benchmarks reais
   - Preferir funções otimizadas da biblioteca padrão

5. **Short Variable Declaration (`:=`)**: 
   - Declara e inicializa novas variáveis
   - Não pode ser usado para redeclarar variáveis existentes
   - Use `=` para atribuir a variáveis já declaradas

## 🔧 Troubleshooting

### Erro: "no new variables on left side of :="

**Causa:** Tentando usar `:=` para variáveis já declaradas.

**Solução:** Use `=` em vez de `:=` para reatribuir valores.

```go
// ❌ Errado
var s string
s := "novo valor"  // Erro!

// ✅ Correto
var s string
s = "novo valor"   // OK!
```

---

**Autor:** Jorge Gabriel  
**Livro:** The Go Programming Language (Donovan & Kernighan)  
**Capítulo:** 1.2 - Command-Line Arguments
