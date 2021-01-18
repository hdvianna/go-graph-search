package graph

type Vertex struct {
	Key   string
	Value interface{}
	Edges []Edge
}

type Edge struct {
	From     *Vertex
	To       *Vertex
	Distance int
}

func (vertex *Vertex) AddEdge(to *Vertex, distance int) *Vertex {
	vertex.Edges = append(vertex.Edges, Edge{
		From: vertex,
		To: to,
		Distance: distance,
	})
	return vertex
}

type VertexIterator interface {
	Current() Vertex
	Next() VertexIterator
	Done() bool
}

type VertexList struct {
	Vertices []Vertex
	pointer  int
}

func (list *VertexList) Has(wanted Vertex) bool {
	for i:=0; i < len(list.Vertices);i++ {
		item := list.Vertices[i]
		if item.Key == wanted.Key {
			return true;
		}
	}
	return false;
}

func (list *VertexList) Append(node Vertex) *VertexList {
	list.Vertices = append(list.Vertices, node)
	return list
}

func (list *VertexList) Get(i int) Vertex {
	return list.Vertices[i:i+1][0]
}

func (list *VertexList) Current() Vertex {
	return list.Get(list.pointer)
}

func (list *VertexList) Next() VertexIterator {
	list.pointer++
	return list
}

func (list *VertexList) Done() bool {
	return list.pointer >= len(list.Vertices)
}

type VertexLink struct {
	next *VertexLink
	value *Vertex
	previous *VertexLink
}

type VertexLinkedList struct {
	first *VertexLink
	current *VertexLink
	last *VertexLink
}

func (list *VertexLinkedList) Current() Vertex {
	return *list.current.value
}

func (list *VertexLinkedList) Next() VertexIterator {
	list.current = list.current.next
	return list
}

func (list *VertexLinkedList) Done() bool {
	return list.current == nil
}

func (list *VertexLinkedList) Append(node Vertex) *VertexLinkedList {
	link := &VertexLink{
		previous: list.last,
		value: &node,
		next: nil,
	}

	if (list.last != nil) {
		list.last.next = link
	}

	list.last = link

	if list.first == nil {
		list.first = link
	}

	if list.current == nil {
		list.current = link
	}

	return list
}

func (list *VertexLinkedList) AppendNext(node Vertex) *VertexLinkedList {

	link := &VertexLink{
		previous: list.current,
		value: &node,
		next: nil,
	}

	if list.current != nil {
		link.next = list.current.next
		list.current.next = link
	}

	if list.first == nil {
		list.first = link
	}

	if list.current == nil {
		list.current = link
	}

	if list.last == nil {
		list.last = link
	}

	return list
}

func (list *VertexLinkedList) Prepend(node Vertex) *VertexLinkedList {
	link := &VertexLink{
		previous: nil,
		value: &node,
		next: list.first,
	}

	if (list.first != nil) {
		list.first.previous = link
	}

	list.first = link

	if list.current == nil {
		list.current = link
	}

	if list.last == nil || link.next == nil{
		list.last = link
	}

	return list
}

func (list *VertexLinkedList) Reset() *VertexLinkedList {
	list.current = list.first
	return list
}

func (list *VertexLinkedList) First() *Vertex {
	if list.first != nil {
		return list.first.value
	}
	return nil
}

func (list *VertexLinkedList) Last() *Vertex {
	if list.last != nil {
		return list.last.value
	}
	return nil
}

type EdgeIterator interface {
	Current() Edge
	Next() EdgeIterator
	Done() bool
}

type EdgeLink struct {
	next *EdgeLink
	value *Edge
	previous *EdgeLink
}

type EdgeLinkedList struct {
	first *EdgeLink
	current *EdgeLink
	last *EdgeLink
}

func (list *EdgeLinkedList) Current() Edge {
	return *list.current.value
}

func (list *EdgeLinkedList) Next() EdgeIterator {
	list.current = list.current.next
	return list
}

func (list *EdgeLinkedList) Done() bool {
	return list.current == nil
}

func (list *EdgeLinkedList) Append(edge Edge) *EdgeLinkedList {
	link := &EdgeLink{
		previous: list.last,
		value: &edge,
		next: nil,
	}

	if (list.last != nil) {
		list.last.next = link
	}

	list.last = link

	if list.first == nil {
		list.first = link
	}

	if list.current == nil {
		list.current = link
	}

	return list
}

func (list *EdgeLinkedList) AppendNext(edge Edge) *EdgeLinkedList {

	link := &EdgeLink{
		previous: list.current,
		value: &edge,
		next: nil,
	}

	if list.current != nil {
		link.next = list.current.next
		list.current.next = link
	}

	if list.first == nil {
		list.first = link
	}

	if list.current == nil {
		list.current = link
	}

	if list.last == nil || link.next == nil {
		list.last = link
	}

	return list
}

func (list *EdgeLinkedList) Prepend(edge Edge) *EdgeLinkedList {
	link := &EdgeLink{
		previous: nil,
		value: &edge,
		next: list.first,
	}

	if (list.first != nil) {
		list.first.previous = link
	}

	list.first = link

	if list.current == nil {
		list.current = link
	}

	if list.last == nil {
		list.last = link
	}

	return list
}

func (list *EdgeLinkedList) Reset() *EdgeLinkedList {
	list.current = list.first
	return list
}

func (list *EdgeLinkedList) First() *Edge {
	if list.first != nil {
		return list.first.value
	}
	return nil
}

func (list *EdgeLinkedList) Last() *Edge {
	if list.last != nil {
		return list.last.value
	}
	return nil
}