package object

import (
	"bytes"
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	RETURN_OBJ  = "RETURN"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
	BLOCK_OBJ   = "BLOCK"
	ERROR_OBJ   = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR : " + e.Message + "\n" }

type BlockObject struct {
	Block  []Object
	Return bool
}

func (bo *BlockObject) Type() ObjectType { return BLOCK_OBJ }
func (bo *BlockObject) Inspect() string {
	var res bytes.Buffer
	for _, obj := range bo.Block {
		res.WriteString(obj.Inspect())
	}

	return res.String()
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null\n" }

type Return struct {
	Value Object
}

func (r *Return) Type() ObjectType { return RETURN_OBJ }
func (r *Return) Inspect() string  { return r.Value.Inspect() }

type Integer struct {
	Value int
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d\n", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t\n", b.Value) }
