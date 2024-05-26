package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tysufa/qfa/evaluator"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/parser"
	// "github.com/tysufa/qfa/token"
)

const PROMPT = ">>> "

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		input := scanner.Text()

		l := lexer.New(input)
		// tok := l.GetToken()
		// for tok.Type != token.EOF && tok.Type != token.NL {
		// 	fmt.Printf("{%v : %v}\n", tok.Type, tok.Value)
		// 	tok = l.GetToken()
		// }
		p := parser.New(l)
		stmts := p.GetStatements()
		evaluated := evaluator.EvaluateStatements(stmts)
		for _, ev := range evaluated {
			if ev != nil {
				fmt.Printf("%v\n", ev.Inspect())
			}
		}
	}
}
