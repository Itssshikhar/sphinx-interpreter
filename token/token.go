package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const(
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers
  IDENT = "IDENT"
  INT = "INT"
  STRING = "STRING"

  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"
  
  LT = "<"
  GT = ">"
  
  EQ = "=="
  NOT_EQ = "!="

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"
  COLON = ":"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  LBRACKET = "["
  RBRACKET = "]"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
  IF = "IF"
  TRUE = "TRUE"
  FALSE = "FALSE"
  ELSE = "ELSE"
  RETURN = "RETURN"
)

var keywords = map[string]TokenType {
  "fn": FUNCTION,
  "let": LET,
  "if": IF,
  "else": ELSE,
  "true": TRUE,
  "false": FALSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
