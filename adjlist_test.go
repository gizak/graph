package graph

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestStruct(t *testing.T) {
	// Vertices
	Convey("Vertex Testing", t, func() {
		al := NewAdjList()
		Convey("Add Two Node", func() {

			al.SetVertex(1, "hello")
			al.SetVertex(2, "world")

			Convey("Get Node Val", func() {
				So(al.GetVertex(1), ShouldEqual, "hello")
				So(al.getNode(2).id, ShouldEqual, 2)
				So(al.GetVertex(3), ShouldBeNil)
			})

			So(len(*al), ShouldEqual, 2)
			So((*al)[0].node.val, ShouldEqual, "hello")
			So((*al)[1].next, ShouldBeNil)
		})

		Convey("Update Exist Nodes", func() {
			al.SetVertex(1, "howdy")
			So(al.GetVertex(1), ShouldEqual, "howdy")
		})

		Convey("Del Nodes Without Removing The Related Edges", func() {
			al.DelVertex(1)
			So(len(*al), ShouldEqual, 1)
			So((*al)[0].id, ShouldEqual, 2)
			al.DelVertex(2)
			So(len(*al), ShouldEqual, 0)
		})
	})

	// Edges
	Convey("Edge Testing", t, func() {
		al := NewAdjList()
		al.SetVertex(0, "hello")
		al.SetVertex(1, "world")
		al.SetVertex(2, "!")
		Convey("Add Edge", func() {
			al.SetEdge(0, 1, "0->1")
			al.SetEdge(0, 2, "0->2")
			al.SetEdge(1, 2, "1->2")
			So((*al)[0].next, ShouldNotBeNil)
			So((*al)[0].next.id, ShouldEqual, 1)
			So((*al)[0].next.edge, ShouldEqual, "0->1")
		})

		Convey("Get Edge Val", func() {
			e := al.GetEdge(0, 1)
			So(e, ShouldEqual, "0->1")
			So(al.GetEdge(1, 0), ShouldBeNil)
			So(al.GetEdge(1, 2), ShouldEqual, "1->2")
		})

		Convey("Update Edge", func() {
			al.SetEdge(0, 1, "0-1")
			al.SetEdge(0, 2, "0-2")
			So(al.GetEdge(0, 1), ShouldEqual, "0-1")
			So(al.GetEdge(0, 2), ShouldEqual, "0-2")
			So(al.GetEdge(1, 2), ShouldEqual, "1->2")
		})

		Convey("Del Certain Edge Without Side Effects", func() {
			al.DelEdge(0, 1)
			So(al.GetEdge(0, 1), ShouldBeNil)
			So(al.GetEdge(1, 2), ShouldEqual, "1->2")
			So(al.GetEdge(0, 2), ShouldEqual, "0-2")
		})

	})
}

func TestInterface(t *testing.T) {
	Convey("Adjacency List Implemention", t, func() {
		foo := func(g Graph) {}
		al := NewAdjList()
		Convey("Implementated Graph Interface", func() {
			foo(al)
		})

	})
}
