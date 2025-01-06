package dot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		name        string
		strict      bool
		diagraph    bool
		concentrate bool
		graphName   string
		rankDir     RankDir
		nodeColor   Color
		nodeStyle   Style
		nodeShape   Shape
		nodes       []Node
		edges       []Edge
		subgraphs   []Subgraph
		expectedDOT string
	}{
		{
			name:        "SimpleGraph",
			strict:      false,
			diagraph:    false,
			concentrate: false,
			graphName:   "",
			rankDir:     "",
			nodeColor:   "",
			nodeStyle:   "",
			nodeShape:   "",
			nodes: []Node{
				{Name: "a0"},
				{Name: "a1"},
			},
			edges: []Edge{
				{From: "a0", To: "a1", EdgeType: EdgeTypeUndirected},
			},
			subgraphs:   []Subgraph{},
			expectedDOT: graph01,
		},
		{
			name:        "GraphWithLabels",
			strict:      true,
			diagraph:    false,
			concentrate: false,
			graphName:   "G",
			rankDir:     "",
			nodeColor:   "",
			nodeStyle:   "",
			nodeShape:   "",
			nodes: []Node{
				{Name: "b0", Label: "B0"},
				{Name: "b1", Label: "B3"},
				{Name: "b2", Label: "B2"},
			},
			edges: []Edge{
				{From: "b0", To: "b1", EdgeType: EdgeTypeUndirected, Color: ColorRed},
				{From: "b0", To: "b2", EdgeType: EdgeTypeUndirected, Color: ColorBlack},
			},
			subgraphs:   []Subgraph{},
			expectedDOT: graph02,
		},
		{
			name:        "GraphWithSubgraph",
			strict:      false,
			diagraph:    true,
			concentrate: false,
			graphName:   "",
			rankDir:     "",
			nodeColor:   ColorLimeGreen,
			nodeStyle:   "",
			nodeShape:   "",
			nodes: []Node{
				{Name: "c0", Label: "C0", Shape: ShapePlain},
				{Name: "c1", Label: "C1", Shape: ShapePlain},
				{Name: "c2", Label: "C2", Shape: ShapePlain},
				{Name: "c3", Label: "C3", Shape: ShapePlain},
			},
			edges: []Edge{
				{From: "c0", To: "c1", EdgeType: EdgeTypeDirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeOpen, ArrowTail: ArrowTypeDot},
				{From: "c0", To: "c2", EdgeType: EdgeTypeDirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeOpen, ArrowTail: ArrowTypeDot},
				{From: "c2", To: "c3", EdgeType: EdgeTypeDirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeOpen, ArrowTail: ArrowTypeDot},
			},
			subgraphs: []Subgraph{
				{Name: "", Label: "Thread", Rank: RankSame},
			},
			expectedDOT: graph03,
		},
		{
			name:        "ComplexGraph",
			strict:      true,
			diagraph:    true,
			concentrate: false,
			graphName:   "DG",
			rankDir:     RankDirLR,
			nodeColor:   ColorSteelBlue,
			nodeStyle:   StyleFilled,
			nodeShape:   ShapeMrecord,
			nodes: []Node{
				{Name: "start", Label: "Start", Color: ColorBlue, Shape: ShapeBox},
				{Name: "end", Label: "End", Color: ColorBlue, Shape: ShapeBox},
			},
			edges: []Edge{
				{From: "start", To: "e0", EdgeType: EdgeTypeDirected, Label: "Start", Color: ColorRed, Style: StyleSolid},
				{From: "start", To: "f0", EdgeType: EdgeTypeDirected, Label: "Start", Color: ColorRed, Style: StyleSolid},
				{From: "e1", To: "end", EdgeType: EdgeTypeDirected, Label: "End", Color: ColorRed, Style: StyleSolid},
				{From: "f1", To: "end", EdgeType: EdgeTypeDirected, Label: "End", Color: ColorRed, Style: StyleSolid},
			},
			subgraphs: []Subgraph{
				{
					Name: "cluster0", Label: "Process 0", Color: ColorGray, Style: StyleFilled,
					Nodes: []Node{
						{Name: "e0"},
						{Name: "e1"},
					},
					Edges: []Edge{
						{From: "e0", To: "e1", EdgeType: EdgeTypeDirected, Style: StyleDashed},
					},
				},
				{
					Name: "cluster1", Label: "Process 1", Color: ColorGray, Style: StyleFilled,
					Nodes: []Node{
						{Name: "f0"},
						{Name: "f1"},
					},
					Edges: []Edge{
						{From: "f0", To: "f1", EdgeType: EdgeTypeDirected, Style: StyleDashed},
					},
				},
			},
			expectedDOT: graph04,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGraph(tc.strict, tc.diagraph, tc.concentrate, tc.graphName, tc.rankDir, tc.nodeColor, tc.nodeStyle, tc.nodeShape)
			g.AddNode(tc.nodes...)
			g.AddEdge(tc.edges...)
			g.AddSubgraph(tc.subgraphs...)

			assert.Equal(t, tc.expectedDOT, g.DOT())
		})
	}
}

var graph01 = `graph {
  concentrate=false;
  node [];

  a0 [];
  a1 [];

  a0 -- a1 [];
}`

var graph02 = `strict graph "G" {
  concentrate=false;
  node [];

  b0 [label="B0"];
  b1 [label="B3"];
  b2 [label="B2"];

  b0 -- b1 [color=red];
  b0 -- b2 [color=black];
}`

var graph03 = `digraph {
  concentrate=false;
  node [color=limegreen];

  subgraph {
    label="Thread";
    rank=same;
    node [];
  }

  c0 [label="C0", shape=plain];
  c1 [label="C1", shape=plain];
  c2 [label="C2", shape=plain];
  c3 [label="C3", shape=plain];

  c0 -> c1 [dirType=both, arrowhead=open, arrowtail=dot];
  c0 -> c2 [dirType=both, arrowhead=open, arrowtail=dot];
  c2 -> c3 [dirType=both, arrowhead=open, arrowtail=dot];
}`

var graph04 = `strict digraph "DG" {
  rankdir=LR;
  concentrate=false;
  node [color=steelblue, style=filled, shape=Mrecord];

  subgraph cluster0 {
    label="Process 0";
    color=gray;
    style=filled;
    node [];

    e0 [];
    e1 [];

    e0 -> e1 [style=dashed];
  }

  subgraph cluster1 {
    label="Process 1";
    color=gray;
    style=filled;
    node [];

    f0 [];
    f1 [];

    f0 -> f1 [style=dashed];
  }

  start [label="Start", color=blue, shape=box];
  end [label="End", color=blue, shape=box];

  start -> e0 [label="Start", color=red, style=solid];
  start -> f0 [label="Start", color=red, style=solid];
  e1 -> end [label="End", color=red, style=solid];
  f1 -> end [label="End", color=red, style=solid];
}`
