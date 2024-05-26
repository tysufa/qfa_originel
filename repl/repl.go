package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tysufa/qfa/evaluator"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/parser"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')

		l := lexer.New(input)
		// tok := l.GetToken()
		// for tok.Type != token.EOF && tok.Type != token.NL {
		// 	fmt.Printf("{%v : %v}\n", tok.Type, tok.Value)
		// 	tok = l.GetToken()
		// }
		p := parser.New(l)
		stmts := p.GetStatements()
		evaluated := evaluator.EvaluateStatements(stmts)
		if evaluated != nil {
			for _, ev := range evaluated {
				if ev != nil {
					fmt.Printf("%v\n", ev.Inspect())
				}
			}
		}
	}
}
