package graph

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
)

func TestGraph(t *testing.T) {
	gunit.Run(new(TestFixture), t)
}

type TestFixture struct {
	*gunit.Fixture
}

func (this *TestFixture) TestEdgesNumber() {
	novaSantaRita := &Vertex{
		Key:   "Nova Santa Rita",
		Edges: make([]Edge, 0),
	}
	canoas := &Vertex{
		Key:   "Canoas",
		Edges: make([]Edge, 0),
	}
	portoAlegre := &Vertex{
		Key:   "Porto Alegre",
		Edges: make([]Edge, 0),
	}
	esteio := &Vertex{
		Key:   "Esteio",
		Edges: make([]Edge, 0),
	}
	montenegro := &Vertex{
		Key:   "Montenegro",
		Edges: make([]Edge, 0),
	}

	novaSantaRita.
		AddEdge(canoas, 10).
		AddEdge(esteio, 10).
		AddEdge(montenegro, 10)

	canoas.
		AddEdge(portoAlegre, 10).
		AddEdge(novaSantaRita, 10).
		AddEdge(esteio, 10)

	esteio.
		AddEdge(canoas, 10).
		AddEdge(novaSantaRita, 10)

	this.So(len(esteio.Edges), should.Equal, 2)
	this.So(len(canoas.Edges), should.Equal, 3)
	this.So(len(novaSantaRita.Edges), should.Equal, 3)
}

func (this *TestFixture) TestNodeList() {
	novaSantaRita := Vertex{
		Key:   "Nova Santa Rita",
		Edges: make([]Edge, 0),
	}
	canoas := Vertex{
		Key:   "Canoas",
		Edges: make([]Edge, 0),
	}
	portoAlegre := Vertex{
		Key:   "Porto Alegre",
		Edges: make([]Edge, 0),
	}
	esteio := Vertex{
		Key:   "Esteio",
		Edges: make([]Edge, 0),
	}
	montenegro := Vertex{
		Key:   "Montenegro",
		Edges: make([]Edge, 0),
	}

	nodeList := VertexList{
		Vertices: make([]Vertex, 0),
		pointer:  0,
	}

	this.So(nodeList.Done(), should.BeTrue)

	nodeList.Append(novaSantaRita)
	this.So(nodeList.Current().Key, should.Equal, novaSantaRita.Key)

	nodeList.
		Append(canoas).
		Append(esteio)

	this.So(nodeList.Has(novaSantaRita), should.BeTrue)
	this.So(nodeList.Has(canoas), should.BeTrue)
	this.So(nodeList.Has(esteio), should.BeTrue)
	this.So(nodeList.Has(montenegro), should.BeFalse)
	this.So(nodeList.Has(portoAlegre), should.BeFalse)

	this.So(nodeList.Get(0).Key, should.Equal, novaSantaRita.Key)
	this.So(nodeList.Get(1).Key, should.Equal, canoas.Key)
	this.So(nodeList.Get(2).Key, should.Equal, esteio.Key)

	node := nodeList.
		Next().
		Next().
		Current()

	this.So(node.Key, should.Equal, esteio.Key)
	this.So(nodeList.Done(), should.BeFalse)
	this.So(nodeList.Next().Done(), should.BeTrue)

}

func (this *TestFixture) TestNodeLinkedList() {
	novaSantaRita := Vertex{
		Key:   "Nova Santa Rita",
		Edges: make([]Edge, 0),
	}
	canoas := Vertex{
		Key:   "Canoas",
		Edges: make([]Edge, 0),
	}
	portoAlegre := Vertex{
		Key:   "Porto Alegre",
		Edges: make([]Edge, 0),
	}
	esteio := Vertex{
		Key:   "Esteio",
		Edges: make([]Edge, 0),
	}
	montenegro := Vertex{
		Key:   "Montenegro",
		Edges: make([]Edge, 0),
	}


	nodeLinkedList := VertexLinkedList{first: nil, current: nil, last: nil}

	this.So(nodeLinkedList.Done(), should.BeTrue)

	nodeLinkedList.
		Append(novaSantaRita).
		Append(canoas).
		Append(portoAlegre).
		Prepend(esteio).
		Prepend(montenegro)

	this.So(nodeLinkedList.First().Key , should.Equal, montenegro.Key)
	this.So(nodeLinkedList.Last().Key , should.Equal, portoAlegre.Key)
	this.So(nodeLinkedList.Current().Key , should.Equal, novaSantaRita.Key)
	this.So(nodeLinkedList.Reset().Current().Key , should.Equal, montenegro.Key)

	nodeLinkedList.
		Next().
		Next().
		Next().
		Next()

	this.So(nodeLinkedList.Done(), should.BeFalse)
	this.So(nodeLinkedList.Next().Done(), should.BeTrue)

	nodeLinkedList = VertexLinkedList{first: nil, current: nil, last: nil}
	nodeLinkedList.
		AppendNext(novaSantaRita).
		AppendNext(canoas).
		AppendNext(portoAlegre)

	this.So(nodeLinkedList.Current().Key , should.Equal, novaSantaRita.Key)
	this.So(nodeLinkedList.Next().Current().Key , should.Equal, portoAlegre.Key)

	nodeLinkedList.
		AppendNext(montenegro)

	this.So(nodeLinkedList.Current().Key , should.Equal, portoAlegre.Key)
	this.So(nodeLinkedList.Next().Current().Key , should.Equal, montenegro.Key)

	nodeLinkedList = VertexLinkedList{first: nil, current: nil, last: nil}
	this.So(nodeLinkedList.Last(), should.BeNil)
	this.So(nodeLinkedList.First(), should.BeNil)
	nodeLinkedList.
		Prepend(novaSantaRita)
}


func (this *TestFixture) TestEdgeLinkedList() {
	novaSantaRita := Vertex{
		Key:   "Nova Santa Rita",
		Edges: make([]Edge, 0),
	}
	canoas := Vertex{
		Key:   "Canoas",
		Edges: make([]Edge, 0),
	}
	esteio := Vertex{
		Key:   "Esteio",
		Edges: make([]Edge, 0),
	}
	montenegro := Vertex{
		Key:   "Montenegro",
		Edges: make([]Edge, 0),
	}
	portao := Vertex{
		Key:   "Portao",
		Edges: make([]Edge, 0),
	}

	novaSantaRita.
		AddEdge(&canoas, 10).
		AddEdge(&esteio, 10).
		AddEdge(&montenegro, 10).
		AddEdge(&portao, 10)

	edgeLinkedList := EdgeLinkedList{first: nil, current: nil, last: nil}

	this.So(edgeLinkedList.Done(), should.BeTrue)

	edgeLinkedList.
		Append(novaSantaRita.Edges[0]).
		Prepend(novaSantaRita.Edges[1]).
		AppendNext(novaSantaRita.Edges[2])


	this.So(edgeLinkedList.First().To.Key , should.Equal, esteio.Key)
	this.So(edgeLinkedList.Current().To.Key, should.Equal, canoas.Key)
	this.So(edgeLinkedList.Last().To.Key , should.Equal, montenegro.Key)
	this.So(edgeLinkedList.Next().Current().To.Key, should.Equal, montenegro.Key)
	this.So(edgeLinkedList.Reset().Current().To.Key , should.Equal, esteio.Key)

	edgeLinkedList.
		Next()

	this.So(edgeLinkedList.Done(), should.BeFalse)
	this.So(edgeLinkedList.Next().Next().Done(), should.BeTrue)

	edgeLinkedList.Append(novaSantaRita.Edges[3])
	this.So(edgeLinkedList.Last().To.Key , should.Equal, portao.Key)

	nilEdgeLinkedList := EdgeLinkedList{first: nil, current: nil, last: nil}

	this.So(nilEdgeLinkedList.Last(), should.BeNil)
	this.So(nilEdgeLinkedList.First(), should.BeNil)

	nilEdgeLinkedList.Prepend(novaSantaRita.Edges[0])
	this.So(nilEdgeLinkedList.First().To.Key , should.Equal, canoas.Key)
	this.So(nilEdgeLinkedList.Current().To.Key, should.Equal, canoas.Key)
	this.So(nilEdgeLinkedList.Last().To.Key , should.Equal, canoas.Key)

	nilEdgeLinkedList = EdgeLinkedList{first: nil, current: nil, last: nil}
	nilEdgeLinkedList.AppendNext(novaSantaRita.Edges[0])
	this.So(nilEdgeLinkedList.First().To.Key , should.Equal, canoas.Key)
	this.So(nilEdgeLinkedList.Current().To.Key, should.Equal, canoas.Key)
	this.So(nilEdgeLinkedList.Last().To.Key , should.Equal, canoas.Key)

	nilEdgeLinkedList = EdgeLinkedList{first: nil, current: nil, last: nil}
	nilEdgeLinkedList.Prepend(novaSantaRita.Edges[0])
	this.So(nilEdgeLinkedList.First().To.Key , should.Equal, canoas.Key)
	this.So(nilEdgeLinkedList.Current().To.Key, should.Equal, canoas.Key)
	this.So(nilEdgeLinkedList.Last().To.Key , should.Equal, canoas.Key)


}
