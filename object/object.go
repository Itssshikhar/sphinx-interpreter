package object

import (
	"bytes"
	"fmt"
	"sphinx/ast"
  "strings"
)

type ObjectType string

const (
  INTEGER_OBJ = "INTEGER"
  BOOLEAN_OBJ = "BOOLEAN"
  NULL_OBJ = "NULL"
  RETURN_VALUE_OBJ = "RETURN_VALUE"
  ERROR_OBJ = "ERROR"
  FUNCTION_OBJ = "FUNCTION"
)

type Object interface {
  Type() ObjectType
  Inspect() string
}

type Integer struct {
  Value int64
}

type Boolean struct {
  Value bool
}

type Null struct {}

type ReturnValue struct {
  Value Object
}

type Error struct {
  Message string
}

type Function struct {
  Parameters []*ast.Identifier
  Body *ast.BlockStatement
  Env *Environment
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR:" + e.Message }

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
  var out bytes.Buffer

  params := []string{}
  for _, p := range f.Parameters {
    params = append(params, p.String())
  }

  out.WriteString("fn")
  out.WriteString("(")
  out.WriteString(strings.Join(params, ", "))
  out.WriteString(") {\n")
  out.WriteString(f.Body.String())
  out.WriteString("\n}")

  return out.String()
}
