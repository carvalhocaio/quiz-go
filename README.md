# Quiz Go 🎯

Um jogo de quiz interativo sobre a linguagem de programação Go, desenvolvido em Go.

## 📋 Descrição

Este é um quiz educativo que testa seus conhecimentos sobre a linguagem Go através de perguntas de
múltipla escolha. O jogo carrega as perguntas de um arquivo CSV e apresenta uma pontuação final com
feedback personalizado.

## 🚀 Funcionalidades

- ✅ 50+ perguntas sobre Go
- ✅ Interface interativa via terminal
- ✅ Validação de entrada do usuário
- ✅ Sistema de pontuação
- ✅ Feedback personalizado baseado na performance
- ✅ Carregamento dinâmico de perguntas via CSV

## 🛠️ Tecnologias

- **Go 1.24.5**
- Bibliotecas padrão: `encoding/csv`, `bufio`, `strconv`

## 📁 Estrutura do Projeto

```
.
├── go.mod          # Módulo Go
├── main.go         # Código principal
├── quiz-go.csv     # Perguntas do quiz
└── README.md       # Documentação
```

## 🎮 Como Executar

### Pré-requisitos
- Go 1.24+ instalado

### Passos

1. **Clone o repositório:**
```bash
git clone https://github.com/carvalhocaio/quiz-go.git
cd quiz-go
```

2. **Execute o quiz:**
```bash
go run main.go
```

## 🎯 Como Jogar

1. Digite seu nome quando solicitado
2. Responda as perguntas escolhendo uma opção [1-4]
3. Veja sua pontuação final e feedback

### Exemplo de Uso

```
Seja bem vindo(a)! ao quiz
Escreva o seu nome: Clancy

Carregadas 53 perguntas. Vamos começar!

 1. Em que ano foi lançado o Go? 
[1] 2009
[2] 2007
[3] 2013
[4] 2011

Sua resposta [1-4]: 1
✓ Correto!
```

## 📊 Sistema de Avaliação

- **90%+ de acerto:** 🏆 Expert em Go
- **70-89% de acerto:** 👍 Muito bom conhecimento
- **50-69% de acerto:** 📚 Conhecimento razoável
- **Menos de 50%:** 💪 Precisa estudar mais

## 🔧 Estruturas de Dados

### Question
```go
type Question struct {
    Text    string   // Texto da pergunta
    Options []string // Opções de resposta
    Awnser  int      // Resposta correta [1-4]
}
```

### GameState
```go
type GameState struct {
    Name      string     // Nome do jogador
    Points    int        // Pontuação atual
    Questions []Question // Lista de perguntas
}
```

## 📝 Formato do CSV

O arquivo [`quiz-go.csv`](quiz-go.csv) deve seguir o formato:

```csv
Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta Correta
Em que ano foi lançado o Go?,2009,2007,2013,2011,1
```

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Adicione suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

### Adicionando Novas Perguntas

Para adicionar novas perguntas, edite o arquivo [`quiz-go.csv`](quiz-go.csv) seguindo o formato
existente.

## 🐛 Tratamento de Erros

O projeto inclui tratamento para:
- Arquivo CSV não encontrado
- Entradas inválidas do usuário
- Erros de conversão de tipos

Inspirado no projeto feito no [Curso Introdutório de Go feito pela Rocketseat](https://app.rocketseat.com.br/journey/go-curso-introdutorio)