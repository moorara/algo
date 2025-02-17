package dot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubgraph(t *testing.T) {
	tests := []struct {
		name         string
		subgraphName string
		label        string
		color        Color
		style        Style
		rank         Rank
		rankDir      RankDir
		nodeColor    Color
		nodeStyle    Style
		nodeShape    Shape
		nodes        []Node
		edges        []Edge
		subgraphs    []Subgraph
		expectedDOT  string
	}{
		{
			name:         "EmptySubgraph",
			subgraphName: "parent",
			label:        "",
			color:        "",
			style:        "",
			rank:         "",
			rankDir:      "",
			nodeColor:    "",
			nodeStyle:    "",
			nodeShape:    "",
			nodes: []Node{
				{Name: "a0"},
				{Name: "a1"},
			},
			edges: []Edge{
				{From: "a0", To: "a1", EdgeType: EdgeTypeUndirected},
			},
			subgraphs:   []Subgraph{},
			expectedDOT: subgraph01,
		},
		{
			name:         "SubgraphWithNodes",
			subgraphName: "child",
			label:        "Child",
			color:        "",
			style:        "",
			rank:         "",
			rankDir:      "",
			nodeColor:    "",
			nodeStyle:    "",
			nodeShape:    "",
			nodes: []Node{
				{Name: "b0", Label: "B0"},
				{Name: "b1", Label: "B3"},
				{Name: "b2", Label: "B2"},
			},
			edges: []Edge{
				{From: "b0", To: "b1", EdgeType: EdgeTypeDirected, Color: ColorRed},
				{From: "b0", To: "b2", EdgeType: EdgeTypeDirected, Color: ColorBlack},
			},
			subgraphs:   []Subgraph{},
			expectedDOT: subgraph02,
		},
		{
			name:         "SubgraphWithNodesAndEdges",
			subgraphName: "cluster0",
			label:        "Left",
			color:        ColorPink,
			style:        "",
			rank:         "",
			rankDir:      RankDirLR,
			nodeColor:    ColorRoyalBlue,
			nodeStyle:    "",
			nodeShape:    "",
			nodes: []Node{
				{Name: "c0", Label: "C0", Shape: ShapeBox},
				{Name: "c1", Label: "C1", Shape: ShapeBox},
				{Name: "c2", Label: "C2", Shape: ShapeBox},
				{Name: "c3", Label: "C3", Shape: ShapeBox},
			},
			edges: []Edge{
				{From: "c0", To: "c1", EdgeType: EdgeTypeUndirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeDot, ArrowTail: ArrowTypeDot},
				{From: "c0", To: "c2", EdgeType: EdgeTypeUndirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeDot, ArrowTail: ArrowTypeDot},
				{From: "c1", To: "c3", EdgeType: EdgeTypeUndirected, EdgeDir: EdgeDirBoth, ArrowHead: ArrowTypeDot, ArrowTail: ArrowTypeDot},
			},
			subgraphs: []Subgraph{
				{Name: "thread", Label: "Thread"},
			},
			expectedDOT: subgraph03,
		},
		{
			name:         "SubgraphWithSubgraph",
			subgraphName: "cluster1",
			label:        "Right",
			color:        ColorGreen,
			style:        StyleDotted,
			rank:         RankSame,
			rankDir:      "",
			nodeColor:    ColorSeaGreen,
			nodeStyle:    StyleFilled,
			nodeShape:    ShapeRecord,
			nodes: []Node{
				{Name: "d0", Label: "D0", Color: ColorTan, Shape: ShapeOval},
				{Name: "d1", Label: "D1", Color: ColorTan, Shape: ShapeOval},
			},
			edges: []Edge{
				{From: "d0", To: "e0", EdgeType: EdgeTypeDirected, Label: "d0e0", Color: ColorGray, Style: StyleDashed},
				{From: "d0", To: "f0", EdgeType: EdgeTypeDirected, Label: "d0f0", Color: ColorGray, Style: StyleDashed},
				{From: "e1", To: "d1", EdgeType: EdgeTypeDirected, Label: "e1d1", Color: ColorGray, Style: StyleDashed},
				{From: "f1", To: "d1", EdgeType: EdgeTypeDirected, Label: "f1d1", Color: ColorGray, Style: StyleDashed},
			},
			subgraphs: []Subgraph{
				{
					Name: "process0", Label: "Process 0", Color: ColorGray, Style: StyleFilled,
					Nodes: []Node{
						{Name: "e0"},
						{Name: "e1"},
					},
					Edges: []Edge{
						{From: "e0", To: "e1", EdgeType: EdgeTypeDirected},
					},
				},
				{
					Name: "process1", Label: "Process 1", Color: ColorGray, Style: StyleFilled,
					Nodes: []Node{
						{Name: "f0"},
						{Name: "f1"},
					},
					Edges: []Edge{
						{From: "f0", To: "f1", EdgeType: EdgeTypeDirected},
					},
				},
			},
			expectedDOT: subgraph04,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sg := NewSubgraph(tc.subgraphName, tc.label, tc.color, tc.style, tc.rank, tc.rankDir, tc.nodeColor, tc.nodeStyle, tc.nodeShape)
			sg.AddNode(tc.nodes...)
			sg.AddEdge(tc.edges...)
			sg.AddSubgraph(tc.subgraphs...)

			assert.Equal(t, tc.expectedDOT, sg.DOT(0))
		})
	}
}

var subgraph01 = `subgraph parent {
  node [];

  a0 [];
  a1 [];

  a0 -- a1 [];
}`

var subgraph02 = `subgraph child {
  label="Child";
  node [];

  b0 [label="B0"];
  b1 [label="B3"];
  b2 [label="B2"];

  b0 -> b1 [color=red];
  b0 -> b2 [color=black];
}`

var subgraph03 = `subgraph cluster0 {
  label="Left";
  color=pink;
  rankdir=LR;
  node [color=royalblue];

  subgraph thread {
    label="Thread";
    node [];
  }

  c0 [label="C0", shape=box];
  c1 [label="C1", shape=box];
  c2 [label="C2", shape=box];
  c3 [label="C3", shape=box];

  c0 -- c1 [dirType=both, arrowhead=dot, arrowtail=dot];
  c0 -- c2 [dirType=both, arrowhead=dot, arrowtail=dot];
  c1 -- c3 [dirType=both, arrowhead=dot, arrowtail=dot];
}`

var subgraph04 = `subgraph cluster1 {
  label="Right";
  color=green;
  style=dotted;
  rank=same;
  node [color=seagreen, style=filled, shape=record];

  subgraph process0 {
    label="Process 0";
    color=gray;
    style=filled;
    node [];

    e0 [];
    e1 [];

    e0 -> e1 [];
  }

  subgraph process1 {
    label="Process 1";
    color=gray;
    style=filled;
    node [];

    f0 [];
    f1 [];

    f0 -> f1 [];
  }

  d0 [label="D0", color=tan, shape=oval];
  d1 [label="D1", color=tan, shape=oval];

  d0 -> e0 [label="d0e0", color=gray, style=dashed];
  d0 -> f0 [label="d0f0", color=gray, style=dashed];
  e1 -> d1 [label="e1d1", color=gray, style=dashed];
  f1 -> d1 [label="f1d1", color=gray, style=dashed];
}`
