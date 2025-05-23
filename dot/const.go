package dot

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
	ColorAliceBlue            Color = "aliceblue"
	ColorAntiqueWhite         Color = "antiquewhite"
	ColorAqua                 Color = "aqua"
	ColorAquaMarine           Color = "aquamarine"
	ColorAzure                Color = "azure"
	ColorBeige                Color = "beige"
	ColorBisque               Color = "bisque"
	ColorBlack                Color = "black"
	ColorBlanchedAlmond       Color = "blanchedalmond"
	ColorBlue                 Color = "blue"
	ColorBlueViolet           Color = "blueviolet"
	ColorBrown                Color = "brown"
	ColorBurlyWood            Color = "burlywood"
	ColorCadetBlue            Color = "cadetblue"
	ColorChartreuse           Color = "chartreuse"
	ColorChocolate            Color = "chocolate"
	ColorCoral                Color = "coral"
	ColorCornFlowerBlue       Color = "cornflowerblue"
	ColorCornSilk             Color = "cornsilk"
	ColorCrimson              Color = "crimson"
	ColorCyan                 Color = "cyan"
	ColorDarkblue             Color = "darkblue"
	ColorDarkCyan             Color = "darkcyan"
	ColorDarkGoldenrod        Color = "darkgoldenrod"
	ColorDarkGray             Color = "darkgray"
	ColorDarkGreen            Color = "darkgreen"
	ColorDarkGrey             Color = "darkgrey"
	ColorDarkKhaki            Color = "darkkhaki"
	ColorDarkMagenta          Color = "darkmagenta"
	ColorDarkOliveGreen       Color = "darkolivegreen"
	ColorDarkOrange           Color = "darkorange"
	ColorDarkOrchid           Color = "darkorchid"
	ColorDarkRed              Color = "darkred"
	ColorDarkSalmon           Color = "darksalmon"
	ColorDarkSeaGreen         Color = "darkseagreen"
	ColorDarkSlateBlue        Color = "darkslateblue"
	ColorDarkSlateGray        Color = "darkslategray"
	ColorDarkSlateGrey        Color = "darkslategrey"
	ColorDarkTurquoise        Color = "darkturquoise"
	ColorDarkViolet           Color = "darkviolet"
	ColorDeepPink             Color = "deeppink"
	ColorDeepSkyBlue          Color = "deepskyblue"
	ColorDimGray              Color = "dimgray"
	ColorDimGrey              Color = "dimgrey"
	ColorDoderBlue            Color = "dodgerblue"
	ColorFireBrick            Color = "firebrick"
	ColorFloralWhite          Color = "floralwhite"
	ColorForestGreen          Color = "forestgreen"
	ColorFuchsia              Color = "fuchsia"
	ColorGainsboro            Color = "gainsboro"
	ColorGhostWhite           Color = "ghostwhite"
	ColorGold                 Color = "gold"
	ColorGoldenrod            Color = "goldenrod"
	ColorGray                 Color = "gray"
	ColorGreen                Color = "green"
	ColorGreenYellow          Color = "greenyellow"
	ColorGrey                 Color = "grey"
	ColorHoneydew             Color = "honeydew"
	ColorHotPink              Color = "hotpink"
	ColorIndianRed            Color = "indianred"
	ColorIndigo               Color = "indigo"
	ColorIvory                Color = "ivory"
	ColorKhaki                Color = "khaki"
	ColorLavender             Color = "lavender"
	ColorLavenderBlush        Color = "lavenderblush"
	ColorLawnGreen            Color = "lawngreen"
	ColorLemonChiffon         Color = "lemonchiffon"
	ColorLightBlue            Color = "lightblue"
	ColorLightCoral           Color = "lightcoral"
	ColorLightCyan            Color = "lightcyan"
	ColorLightGoldenrodYellow Color = "lightgoldenrodyellow"
	ColorLightGray            Color = "lightgray"
	ColorLightGreen           Color = "lightgreen"
	ColorLightGrey            Color = "lightgrey"
	ColorLightPink            Color = "lightpink"
	ColorLightSalmon          Color = "lightsalmon"
	ColorLightSeaGreen        Color = "lightseagreen"
	ColorLightSkyBlue         Color = "lightskyblue"
	ColorLightSlateGray       Color = "lightslategray"
	ColorLightSlateGrey       Color = "lightslategrey"
	ColorLightSteelBlue       Color = "lightsteelblue"
	ColorLightYellow          Color = "lightyellow"
	ColorLime                 Color = "lime"
	ColorLimeGreen            Color = "limegreen"
	ColorLinen                Color = "linen"
	ColorMagenta              Color = "magenta"
	ColorMaroon               Color = "maroon"
	ColorMediumAquaMarine     Color = "mediumaquamarine"
	ColorMediumBlue           Color = "mediumblue"
	ColorMediumOrchid         Color = "mediumorchid"
	ColorMediumPurple         Color = "mediumpurple"
	ColorMediumSeaGreen       Color = "mediumseagreen"
	ColorMediumSlateBlue      Color = "mediumslateblue"
	ColorMediumSpringGreen    Color = "mediumspringgreen"
	ColorMediumTurquoise      Color = "mediumturquoise"
	ColorMediumVioletRed      Color = "mediumvioletred"
	ColorMidnightBlue         Color = "midnightblue"
	ColorMintCream            Color = "mintcream"
	ColorMistyRose            Color = "mistyrose"
	ColorMoccasin             Color = "moccasin"
	ColorNavajoWhite          Color = "navajowhite"
	ColorNavy                 Color = "navy"
	ColorOldLace              Color = "oldlace"
	ColorOlive                Color = "olive"
	ColorOliveDrab            Color = "olivedrab"
	ColorOrange               Color = "orange"
	ColorOrangeRed            Color = "orangered"
	ColorOrchid               Color = "orchid"
	ColorPaleGoldenRod        Color = "palegoldenrod"
	ColorPaleGreen            Color = "palegreen"
	ColorPaleTurquoise        Color = "paleturquoise"
	ColorPaleVioletRed        Color = "palevioletred"
	ColorPapayaWhip           Color = "papayawhip"
	ColorPeachPuff            Color = "peachpuff"
	ColorPeru                 Color = "peru"
	ColorPink                 Color = "pink"
	ColorPlum                 Color = "plum"
	ColorPowderBlue           Color = "powderblue"
	ColorPurple               Color = "purple"
	ColorRed                  Color = "red"
	ColorRosyBrown            Color = "rosybrown"
	ColorRoyalBlue            Color = "royalblue"
	ColorSaddleBrown          Color = "saddlebrown"
	ColorSalmon               Color = "salmon"
	ColorSandyBrown           Color = "sandybrown"
	ColorSeaGreen             Color = "seagreen"
	ColorSeashell             Color = "seashell"
	ColorSienna               Color = "sienna"
	ColorSilver               Color = "silver"
	ColorSkyBlue              Color = "skyblue"
	ColorSlateBlue            Color = "slateblue"
	ColorSlateGray            Color = "slategray"
	ColorSlateGrey            Color = "slategrey"
	ColorSnow                 Color = "snow"
	ColorSpringGreen          Color = "springgreen"
	ColorSteelBlue            Color = "steelblue"
	ColorTan                  Color = "tan"
	ColorTeal                 Color = "teal"
	ColorThistle              Color = "thistle"
	ColorTomato               Color = "tomato"
	ColorTurquoise            Color = "turquoise"
	ColorViolet               Color = "violet"
	ColorWheat                Color = "wheat"
	ColorWhite                Color = "white"
	ColorWhiteSmoke           Color = "whitesmoke"
	ColorYellow               Color = "yellow"
	ColorYellowGreen          Color = "yellowgreen"
)

type Style string

const (
	// For nodes and edges
	StyleSolid  Style = "solid"
	StyleBold   Style = "bold"
	StyleDashed Style = "dashed"
	StyleDotted Style = "dotted"
	StyleInvis  Style = "invis"

	// For nodes and clusters
	StyleFilled  Style = "filled"
	StyleStriped Style = "striped"
	StyleRounded Style = "rounded"

	// For nodes only
	StyleWedged    Style = "wedged"
	StyleDiagonals Style = "diagonals"

	// For edges only
	StyleTapered Style = "tapered"
)

type Shape string

const (
	ShapeNone          Shape = "none"
	ShapePlaintext     Shape = "plaintext"
	ShapePlain         Shape = "plain"
	ShapeUnderline     Shape = "underline"
	ShapeBox           Shape = "box"
	ShapePolygon       Shape = "polygon"
	ShapeEllipse       Shape = "ellipse"
	ShapeOval          Shape = "oval"
	ShapeCircle        Shape = "circle"
	ShapePoint         Shape = "point"
	ShapeEgg           Shape = "egg"
	ShapeTriangle      Shape = "triangle"
	ShapeDiamond       Shape = "diamond"
	ShapeTrapezium     Shape = "trapezium"
	ShapeParallelogram Shape = "parallelogram"
	ShapeHouse         Shape = "house"
	ShapePentagon      Shape = "pentagon"
	ShapeHexagon       Shape = "hexagon"
	ShapeSeptagon      Shape = "septagon"
	ShapeOctagon       Shape = "octagon"
	ShapeDoubleCircle  Shape = "doublecircle"
	ShapeDoubleOctagon Shape = "doubleoctagon"
	ShapeTripleOctagon Shape = "tripleoctagon"
	ShapeInvTriangle   Shape = "invtriangle"
	ShapeInvTrapezium  Shape = "invtrapezium"
	ShapeInvHouse      Shape = "invhouse"
	ShapeMDiamond      Shape = "Mdiamond"
	ShapeMSquare       Shape = "Msquare"
	ShapeMCircle       Shape = "Mcircle"
	ShapeRect          Shape = "rect"
	ShapeRectangle     Shape = "rectangle"
	ShapeSquare        Shape = "square"
	ShapeStar          Shape = "star"
	ShapeCylinder      Shape = "cylinder"
	ShapeNote          Shape = "note"
	ShapeTab           Shape = "tab"
	ShapeFolder        Shape = "folder"
	ShapeBox3D         Shape = "box3d"
	ShapeComponent     Shape = "component"
	ShapeCDS           Shape = "cds"
	ShapeAssembly      Shape = "assembly"
	ShapeSignature     Shape = "signature"
	ShapeRArrow        Shape = "rarrow"
	ShapeLArrow        Shape = "larrow"
	ShapePromoter      Shape = "promoter"
	ShapeRPromoter     Shape = "rpromoter"
	ShapeLPromoter     Shape = "lpromoter"
	ShapeRecord        Shape = "record"
	ShapeMrecord       Shape = "Mrecord"
)

type Rank string

const (
	RankSame   Rank = "same"
	RankMin    Rank = "min"
	RankMax    Rank = "max"
	RankSink   Rank = "sink"
	RankSource Rank = "source"
)

type RankDir string

const (
	RankDirLR RankDir = "LR"
	RankDirRL RankDir = "RL"
	RankDirTB RankDir = "TB"
	RankDirBT RankDir = "BT"
)

type ArrowType string

const (
	ArrowTypeNone     ArrowType = "none"
	ArrowTypeNormal   ArrowType = "normal"
	ArrowTypeInv      ArrowType = "inv"
	ArrowTypeDot      ArrowType = "dot"
	ArrowTypeInvDot   ArrowType = "invdot"
	ArrowTypeODot     ArrowType = "odot"
	ArrowTypeInvODot  ArrowType = "invodot"
	ArrowTypeEmpty    ArrowType = "empty"
	ArrowTypeInvEmpty ArrowType = "invempty"
	ArrowTypeBox      ArrowType = "box"
	ArrowTypeOBox     ArrowType = "obox"
	ArrowTypeDiamond  ArrowType = "diamond"
	ArrowTypeODiamond ArrowType = "odiamond"
	ArrowTypeEDiamond ArrowType = "ediamond"
	ArrowTypeOpen     ArrowType = "open"
	ArrowTypeHalfOpen ArrowType = "halfopen"
	ArrowTypeVee      ArrowType = "vee"
	ArrowTypeTee      ArrowType = "tee"
	ArrowTypeCrow     ArrowType = "crow"
)
