package main

import (
	"fmt"
)

type node struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

type edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type VisJSGraph struct {
	Graph
}

/*
  // create an array with nodes
  var nodes = new vis.DataSet([
    {id: 1, label: 'Node 1'},
    {id: 2, label: 'Node 2'},
    {id: 3, label: 'Node 3'},
    {id: 4, label: 'Node 4'},
    {id: 5, label: 'Node 5'}
  ]);

  // create an array with edges
  var edges = new vis.DataSet([
    {from: 1, to: 3},
    {from: 1, to: 2},
    {from: 2, to: 4},
    {from: 2, to: 5},
    {from: 3, to: 3}
  ]);
*/
func (v VisJSGraph) Nodes() []node {
	ns := []node{}
	for i, s := range v.Graph.Nodes() {
		fmt.Printf("%+v\n", s)
		ns = append(ns, node{
			Id: s,
			Label: s,
		})
	}
	return ns
}

func (v VisJSGraph) Edges() []edge {
	es := []edge{}
	for s, ns := range v.Graph {
		// get all connected nodes
		for _, n := range ns {
			es = append(es, edge{
				From: s,
				To: n,
			})
		}
	}
	return es
}


