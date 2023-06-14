package input

import "github.com/litmus-zhang/go-intensive-tdd/calculator"

type Parser struct {
	enging *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expr string) (string, error)  {
	// code
	return "", nil
}