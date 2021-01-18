package graph

import (
	"fmt"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	gunit.Run(new(TestFixtureSearch), t)
}

type TestFixtureSearch struct {
	*gunit.Fixture
}

func (this *TestFixtureSearch) TestBreadthFirst() {
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
		AddEdge(novaSantaRita, 0).
		AddEdge(esteio, 10)

	esteio.
		AddEdge(canoas, 10).
		AddEdge(novaSantaRita, 10)

	montenegro.AddEdge(novaSantaRita, 10)

	portoAlegre.AddEdge(canoas, 10)
	result := BreadthFirst(portoAlegre, func(carry interface{}, edge Edge) interface{} {
		hops := carry.([]string)
		hop := fmt.Sprintf("From '%s' to '%s'", edge.From.Key, edge.To.Key)
		hops = append(hops, hop)
		return hops
	}, make([]string, 0))


	hops := result.([]string)

	this.So(hops[0] , should.Equal, "From '' to 'Porto Alegre'")
	this.So(hops[1] , should.Equal, "From 'Porto Alegre' to 'Canoas'")
	this.So(hops[2] , should.Equal, "From 'Canoas' to 'Nova Santa Rita'")
	this.So(hops[3] , should.Equal, "From 'Canoas' to 'Esteio'")
	this.So(hops[4] , should.Equal, "From 'Nova Santa Rita' to 'Montenegro'")

}

func TestZeroOneBFS(t *testing.T) {
	gunit.Run(new(TestFixtureZeroOneBFS), t)
}

type TestFixtureZeroOneBFS struct {
	*gunit.Fixture
}

func (this *TestFixtureZeroOneBFS) TestZeroOneBFSAppender() {
	s := &Vertex{
		Key:   "s",
		Edges: make([]Edge, 0),
	}

	a := &Vertex{
		Key:   "a",
		Edges: make([]Edge, 0),
	}

	b := &Vertex{
		Key:   "b",
		Edges: make([]Edge, 0),
	}

	c := &Vertex{
		Key:   "c",
		Edges: make([]Edge, 0),
	}

	d := &Vertex{
		Key:   "d",
		Edges: make([]Edge, 0),
	}

	e := &Vertex{
		Key:   "e",
		Edges: make([]Edge, 0),
	}

	s.
		AddEdge(a, 0).
		AddEdge(b, 1)
	a.
		AddEdge(s, 0).
		AddEdge(b, 0).
		AddEdge(c,1)

	b.
		AddEdge(s, 1).
		AddEdge(a, 0).
		AddEdge(d, 0)

	c.
		AddEdge(a, 1).
		AddEdge(d, 0).
		AddEdge(e, 0)

	d.
		AddEdge(b, 0).
		AddEdge(c, 0).
		AddEdge(e, 1)

	e.
		AddEdge(c, 0).
		AddEdge(d, 1)


	result:= ZeroOneBFS(s, func(carry interface{}, edge Edge) interface{} {
		hops := carry.([]string)
		hop := edge.To.Key
		hops = append(hops, hop)
		return hops
	}, make([]string, 0))

	hops := result.([]string)
	expected := []string {
		"s", "a", "b", "d", "c", "e",
	}
	this.So(hops, should.Resemble, expected)
}