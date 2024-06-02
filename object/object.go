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
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

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
