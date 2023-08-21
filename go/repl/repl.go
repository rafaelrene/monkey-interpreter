package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

const PROMPT = "MNK >> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Fprint(output, PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lexerInstance := lexer.New(line)

		for tok := lexerInstance.NextToken(); tok.Type != token.EOF; tok = lexerInstance.NextToken() {
			fmt.Fprintf(output, "%+v\n", tok)
		}
	}
}
