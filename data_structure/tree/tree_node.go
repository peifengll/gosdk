package tree

type Node[T any] struct {
	Data     T
	Children []*Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		Data:     data,
		Children: make([]*Node[T], 0),
	}
}
