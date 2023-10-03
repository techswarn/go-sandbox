package html

import "io"

type Node struct {
	Type NodeType
	Data string
	Attr []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

type Attribute struct {
	key string
	Val string
}

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

func Parse(r io.Reader) (*Node, error) {}