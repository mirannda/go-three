package geometries

import (
	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
)

// Use struct composition
type Box struct {
	geometry three.Geometry

	width  float32
	height float32
	depth  float32
}

func NewBox(width, height, depth float32) *Box {
	box := Box{
		width:  width,
		height: height,
		depth:  depth,
	}

	vertices := make([]mgl32.Vec3, 0)

	halfWidth := width / 2.0
	halfHeight := height / 2.0
	halfDepth := depth / 2.0

	// Bottom plane
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
	)...)

	// Side 1
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	// Side 2
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
	)...)

	// // Side 3
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	// // Side 4
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
	)...)

	// Top plane
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	box.geometry.Vertices = vertices

	return &box
}

func NewCube(size float32) *Box {
	return NewBox(size, size, size)
}

func (b *Box) Vertices() []mgl32.Vec3 {
	return b.geometry.Vertices
}

func buildPlane(v1, v2, v3, v4 mgl32.Vec3) []mgl32.Vec3 {
	return []mgl32.Vec3{
		v1,
		v4,
		v3,
		v1,
		v2,
		v4,
	}
}