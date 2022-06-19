package graphviz

type EdgeType string

const (
	EdgeTypeDirected   EdgeType = "->"
	EdgeTypeUndirected EdgeType = "--"
)

type EdgeDir string

const (
	EdgeDirNone    EdgeDir = "none"
	EdgeDirForward EdgeDir = "forward"
	EdgeDirBack    EdgeDir = "back"
	EdgeDirBoth    EdgeDir = "both"
)

type Color string

const (
	ColorBlack       Color = "black"
	ColorWhite       Color = "white"
	ColorGray        Color = "gray"
	ColorRed         Color = "red"
	ColorCoral       Color = "coral"
	ColorOrange      Color = "orange"
	ColorOrangeRed   Color = "orangered"
	ColorGold        Color = "gold"
	ColorYellow      Color = "yellow"
	ColorYellowGreen Color = "yellowgreen"
	ColorGreen       Color = "green"
	ColorLimeGreen   Color = "limegreen"
	ColorSeaGreen    Color = "seagreen"
	ColorBlue        Color = "blue"
	ColorNavy        Color = "navy"
	ColorPurple      Color = "purple"
	ColorSkyBlue     Color = "skyblue"
	ColorRoyalBlue   Color = "royalblue"
	ColorSteelBlue   Color = "steelblue"
	ColorTan         Color = "tan"
	ColorPink        Color = "pink"
	ColorOrchid      Color = "orchid"
)

type Style string

const (
	StyleSolid   Style = "solid"
	StyleBold    Style = "bold"
	StyleFilled  Style = "filled"
	StyleDashed  Style = "dashed"
	StyleDotted  Style = "dotted"
	StyleRounded Style = "rounded"
)

type Shape string

const (
	ShapeNode     Shape = "none"
	ShapePoint    Shape = "point"
	ShapePlain    Shape = "plain"
	ShapeCircle   Shape = "circle"
	ShapeOval     Shape = "oval"
	ShapeBox      Shape = "box"
	ShapeSquare   Shape = "square"
	ShapeDiamond  Shape = "diamond"
	ShapeTriangle Shape = "triangle"
	ShapeRecord   Shape = "record"
	ShapeMrecord  Shape = "Mrecord"
)

type Rank string

const (
	RankSame Rank = "same"
	RankMin  Rank = "min"
	RankMax  Rank = "max"
)

type RankDir string

const (
	RankDirTB RankDir = "TB"
	RankDirLR RankDir = "LR"
	RankDirRL RankDir = "RL"
	RankDirBT RankDir = "BT"
)

type Arrowhead string

const (
	ArrowheadNone     Arrowhead = "none"
	ArrowheadNormal   Arrowhead = "normal"
	ArrowheadEmpty    Arrowhead = "empty"
	ArrowheadOpen     Arrowhead = "open"
	ArrowheadDot      Arrowhead = "dot"
	ArrowheadOdot     Arrowhead = "odot"
	ArrowheadBox      Arrowhead = "box"
	ArrowheadObox     Arrowhead = "obox"
	ArrowheadDiamond  Arrowhead = "diamond"
	ArrowheadOdiamond Arrowhead = "odiamond"
	ArrowheadEdiamond Arrowhead = "ediamond"
)
