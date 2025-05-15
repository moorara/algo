package spatial

// Point1D represents a point in one-dimensional space.
type Point1D[T any] struct {
	X T
}

// Point2D represents a point in two-dimensional space.
type Point2D[T any] struct {
	X, Y T
}

// Point3D represents a point in three-dimensional space.
type Point3D[T any] struct {
	X, Y, Z T
}

// PointND represents a point in n-dimensional space.
type PointND[T any] struct {
	Coordinates []T
}
