package graph

type VisitQueue interface {
	Append(edge Edge)
	Iterator() EdgeIterator
}

type BreadthFirstAppender struct {
	list *EdgeLinkedList
}

func (visitor *BreadthFirstAppender) Append(edge Edge) {
	visitor.list.Append(edge)
}

func (visitor *BreadthFirstAppender) Iterator() EdgeIterator {
	return visitor.list
}

func BreadthFirst(start *Vertex, visit func(carry interface{}, edge Edge) interface{}, initial interface{}) interface{} {
	list := EdgeLinkedList{
		first: nil,
		current: nil,
		last: nil,
	}
	visitQueue := BreadthFirstAppender{
		list: &list,
	}
	return search(start, &visitQueue, visit, initial)
}

type ZeroOneBFSAppender struct {
	list *EdgeLinkedList
}

func (visitor *ZeroOneBFSAppender) Append(edge Edge) {
	if edge.Distance > 0 {
		visitor.list.Append(edge)
	} else {
		visitor.list.AppendNext(edge)
	}
}

func (visitor *ZeroOneBFSAppender) Iterator() EdgeIterator {
	return visitor.list
}

func ZeroOneBFS(start *Vertex, visit func(carry interface{}, edge Edge) interface{}, initial interface{}) interface{} {
	list := EdgeLinkedList{
		first: nil,
		current: nil,
		last: nil,
	}
	visitQueue := ZeroOneBFSAppender{
		list: &list,
	}
	return search(start, &visitQueue,visit, initial)

}

func search(start *Vertex, visitQueue VisitQueue, visit func(carry interface{}, edge Edge) interface{}, initial interface{}) interface{} {
	seen := VertexList{
		Vertices: make([]Vertex, 0),
		pointer:  0,
	}
	visitQueue.Append(Edge{From: &Vertex{}, To: start, Distance: 0})
	var result = initial
	for !visitQueue.Iterator().Done() {
		current := visitQueue.Iterator().Current()
		if !seen.Has(*current.To) {
			seen.Append(*current.To)
			result = visit(result, current)
		}

		for _, edge := range current.To.Edges {
			if !seen.Has(*edge.To) {
				visitQueue.Append(edge)
			}
		}

		visitQueue.Iterator().Next()

	}
	return result
}