package repl

import (
	"bufio"
	"fmt"
	"io"
	"sphinx/lexer"
	"sphinx/parser"
  "sphinx/evaluator"
)

const PROMPT = ">> "
const SPHINX_FACE = `
                   .~~~.
                  /|6 6|\
                 /O\_^_/O\
                 \/'==='\/
                 ,| |^| |.
            ____(n(n)_(n)n)____
            """""""""""""""""""
`

func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line)
    p := parser.New(l)

    program := p.ParseProgram()
    if len(p.Errors()) != 0 {
      printParserErrors(out, p.Errors())
      continue
    }
    
    evaluated := evaluator.Eval(program)
    if evaluated != nil {
      io.WriteString(out, evaluated.Inspect())
      io.WriteString(out, "\n")
    }  
  }
}

func printParserErrors(out io.Writer, errors []string) {
  io.WriteString(out, SPHINX_FACE)
  io.WriteString(out, "Abey yrr! Kuch error aa gya! F***\n")
  io.WriteString(out, " parser errors:\n")
  for _, msg := range errors {
    io.WriteString(out,"\t"+msg+"\n")
  }
}
