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
	ColorWhite       Color = "white"
	ColorAzure       Color = "azure"
	ColorBlack       Color = "black"
	ColorGray        Color = "gray"
	ColorDimGray     Color = "dimgray"
	ColorBrown       Color = "brown"
	ColorChocolate   Color = "chocolate"
	ColorRed         Color = "red"
	ColorCrimson     Color = "crimson"
	ColorCoral       Color = "coral"
	ColorOrange      Color = "orange"
	ColorOrangeRed   Color = "orangered"
	ColorGold        Color = "gold"
	ColorYellow      Color = "yellow"
	ColorYellowGreen Color = "yellowgreen"
	ColorGreen       Color = "green"
	ColorLimeGreen   Color = "limegreen"
	ColorSeaGreen    Color = "seagreen"
	ColorForestGreen Color = "forestgreen"
	ColorDarkGreen   Color = "darkgreen"
	ColorAqua        Color = "aqua"
	ColorAquaMarine  Color = "aquamarine"
	ColorCyan        Color = "cyan"
	ColorBlue        Color = "blue"
	ColorNavy        Color = "navy"
	ColorSkyBlue     Color = "skyblue"
	ColorRoyalBlue   Color = "royalblue"
	ColorDodgerBlue  Color = "dodgerblue"
	ColorSteelBlue   Color = "steelblue"
	ColorPurple      Color = "purple"
	ColorIndigo      Color = "indigo"
	ColorPink        Color = "pink"
	ColorFuchsia     Color = "fuchsia"
	ColorOrchid      Color = "orchid"
	ColorTan         Color = "tan"
	ColorThistle     Color = "thistle"
	ColorSeashell    Color = "seashell"
)

type Style string

const (
	StyleSolid     Style = "solid"
	StyleBold      Style = "bold"
	StyleFilled    Style = "filled"
	StyleDashed    Style = "dashed"
	StyleDotted    Style = "dotted"
	StyleRounded   Style = "rounded"
	StyleDiagonals Style = "diagonals"
	StyleStriped   Style = "striped"
	StyleWedged    Style = "wedged"
)

type Shape string

const (
	ShapeNode          Shape = "none"
	ShapePoint         Shape = "point"
	ShapePlain         Shape = "plain"
	ShapeUnderline     Shape = "underline"
	ShapeCircle        Shape = "circle"
	ShapeEllipse       Shape = "ellipse"
	ShapeOval          Shape = "oval"
	ShapeEgg           Shape = "egg"
	ShapeBox           Shape = "box"
	ShapeSquare        Shape = "square"
	ShapeRectangle     Shape = "rectangle"
	ShapeTriangle      Shape = "triangle"
	ShapeDiamond       Shape = "diamond"
	ShapeTrapezium     Shape = "trapezium"
	ShapeParallelogram Shape = "parallelogram"
	ShapeHouse         Shape = "house"
	ShapePentagon      Shape = "pentagon"
	ShapeHexagon       Shape = "hexagon"
	ShapeSeptagon      Shape = "septagon"
	ShapeOctagon       Shape = "octagon"
	ShapeStar          Shape = "star"
	ShapeDoubleCircle  Shape = "doublecircle"
	ShapeDoubleOctagon Shape = "doubleoctagon"
	ShapeNote          Shape = "note"
	ShapeTab           Shape = "tab"
	ShapeFolder        Shape = "folder"
	ShapeComponent     Shape = "component"
	ShapeBox3D         Shape = "box3d"
	ShapeCylinder      Shape = "cylinder"
	ShapeRecord        Shape = "record"
	ShapeMrecord       Shape = "Mrecord"
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

type ArrowType string

const (
	ArrowTypeNone     ArrowType = "none"
	ArrowTypeNormal   ArrowType = "normal"
	ArrowTypeEmpty    ArrowType = "empty"
	ArrowTypeVee      ArrowType = "vee"
	ArrowTypeOpen     ArrowType = "open"
	ArrowTypeHalfOpen ArrowType = "halfopen"
	ArrowTypeDot      ArrowType = "dot"
	ArrowTypeODot     ArrowType = "odot"
	ArrowTypeBox      ArrowType = "box"
	ArrowTypeOBox     ArrowType = "obox"
	ArrowTypeDiamond  ArrowType = "diamond"
	ArrowTypeODiamond ArrowType = "odiamond"
	ArrowTypeEDiamond ArrowType = "ediamond"
	ArrowTypeTee      ArrowType = "tee"
	ArrowTypeCrow     ArrowType = "crow"
	ArrowTypeInv      ArrowType = "inv"
	ArrowTypeInvEmpty ArrowType = "invempty"
	ArrowTypeInvDot   ArrowType = "invdot"
	ArrowTypeInvODot  ArrowType = "invodot"
)
