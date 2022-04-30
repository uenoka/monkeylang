package ast
import monkey/token
type node interface{
	TokenLiteral() string
}
type Statement interface{
	node
	statementNode()
}
type Exppression interface{
	node
	expressionNode()
}

type Program struct{
	Statements []Statement
}
func(p *Program) Tokenliteral() string {
	if len(p.Statements)>0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name *readIdentifier
	Value Expression
}

func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral string {return ls.Token.Literal}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}