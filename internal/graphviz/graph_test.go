package graphviz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		name            string
		strict          bool
		diagraph        bool
		concentrate     bool
		graphName       string
		rankDir         RankDir
		nodeColor       Color
		nodeStyle       Style
		nodeShape       Shape
		nodes           []Node
		edges           []Edge
		subgraphs       []Subgraph
		expectedDotCode string
	}{
		{
			"SimpleGraph",
			false, false, false, "",
			"",
			"", "", "",
			[]Node{
				Node{Name: "a0"},
				Node{Name: "a1"},
			},
			[]Edge{
				Edge{From: "a0", To: "a1", EdgeType: EdgeTypeUndirected},
			},
			[]Subgraph{},
			`graph {
  concentrate=false;
  node [];

  a0 [];
  a1 [];

  a0 -- a1 [];
}`,
		},
		{
			"GraphWithLabels",
			true, false, false, "G",
			"",
			"", "", "",
			[]Node{
				Node{Name: "b0", Label: "B0"},
				Node{Name: "b1", Label: "B3"},
				Node{Name: "b2", Label: "B2"},
			},
			[]Edge{
				Edge{From: "b0", To: "b1", EdgeType: EdgeTypeUndirected, Color: ColorRed},
				Edge{From: "b0", To: "b2", EdgeType: EdgeTypeUndirected, Color: ColorBlack},
			},
			[]Subgraph{},
			`strict graph "G" {
  concentrate=false;
  node [];

  b0 [label="B0"];
  b1 [label="B3"];
  b2 [label="B2"];

  b0 -- b1 [color=red];
  b0 -- b2 [color=black];
}`,
		},
		{
			"GraphWithSubgraph",
			false, true, false, "",
			"",
			ColorLimeGreen, "", "",
			[]Node{
				Node{Name: "c0", Label: "C0", Shape: ShapePlain},
				Node{Name: "c1", Label: "C1", Shape: ShapePlain},
				Node{Name: "c2", Label: "C2", Shape: ShapePlain},
				Node{Name: "c3", Label: "C3", Shape: ShapePlain},
			},
			[]Edge{
				Edge{From: "c0", To: "c1", EdgeType: EdgeTypeDirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeOpen, ArrowTail: ArrowTypeDot},
				Edge{From: "c0", To: "c2", EdgeType: EdgeTypeDirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeOpen, ArrowTail: ArrowTypeDot},
				Edge{From: "c2", To: "c3", EdgeType: EdgeTypeDirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeOpen, ArrowTail: ArrowTypeDot},
			},
			[]Subgraph{
				Subgraph{Name: "", Label: "Thread", Rank: RankSame},
			},
			`digraph {
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
}`,
		},
		{
			"ComplexGraph",
			true, true, false, "DG",
			RankDirLR,
			ColorSteelBlue, StyleFilled, ShapeMrecord,
			[]Node{
				Node{Name: "start", Label: "Start", Color: ColorBlue, Shape: ShapeBox},
				Node{Name: "end", Label: "End", Color: ColorBlue, Shape: ShapeBox},
			},
			[]Edge{
				Edge{From: "start", To: "e0", EdgeType: EdgeTypeDirected, Label: "Start", Color: ColorRed, Style: StyleSolid},
				Edge{From: "start", To: "f0", EdgeType: EdgeTypeDirected, Label: "Start", Color: ColorRed, Style: StyleSolid},
				Edge{From: "e1", To: "end", EdgeType: EdgeTypeDirected, Label: "End", Color: ColorRed, Style: StyleSolid},
				Edge{From: "f1", To: "end", EdgeType: EdgeTypeDirected, Label: "End", Color: ColorRed, Style: StyleSolid},
			},
			[]Subgraph{
				Subgraph{
					Name: "cluster0", Label: "Process 0", Color: ColorGray, Style: StyleFilled,
					Nodes: []Node{
						Node{Name: "e0"},
						Node{Name: "e1"},
					},
					Edges: []Edge{
						Edge{From: "e0", To: "e1", EdgeType: EdgeTypeDirected, Style: StyleDashed},
					},
				},
				Subgraph{
					Name: "cluster1", Label: "Process 1", Color: ColorGray, Style: StyleFilled,
					Nodes: []Node{
						Node{Name: "f0"},
						Node{Name: "f1"},
					},
					Edges: []Edge{
						Edge{From: "f0", To: "f1", EdgeType: EdgeTypeDirected, Style: StyleDashed},
					},
				},
			},
			`strict digraph "DG" {
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
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGraph(tc.strict, tc.diagraph, tc.concentrate, tc.graphName, tc.rankDir, tc.nodeColor, tc.nodeStyle, tc.nodeShape)
			g.AddNode(tc.nodes...)
			g.AddEdge(tc.edges...)
			g.AddSubgraph(tc.subgraphs...)

			assert.Equal(t, tc.expectedDotCode, g.DotCode())
		})
	}
}
