// Package spatial implements common spatial data structures.
//
// Spatial data structures efficiently store, organize, and query objectsin space—such as points, lines, polygons, or volumes.
// These structures are essential for geometric and multi-dimensional data analysis, enabling fast spatial queries and operations.
//
// Homogeneous Dimensions (Geometric Data):
//
// In geometric spaces, all dimensions typically share the same unit and type (e.g., coordinates in meters).
// Most spatial algorithms—like distance calculations and nearest neighbor searches assume homogeneous dimensions,
// which simplifies computation and comparison.
//
// Heterogeneous Dimensions (Multi-Dimensional Data):
//
// Spatial structures can also index data with mixed units or types across dimensions (e.g., latitude, longitude, and temperature).
// For heterogeneous data, it is crucial to normalize or transform dimensions to a common scale, or use distance metrics
// that account for differences, to ensure meaningful spatial operations and comparisons.
package spatial
