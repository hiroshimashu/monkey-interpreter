package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hiroshimashu/monkey-interpreter/src/lexer"
	"github.com/hiroshimashu/monkey-interpreter/src/token"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NewToken(); tok.Type != token.EOF; tok = l.NextToken() {

		}
	}
}
