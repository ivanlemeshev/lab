package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320
	screenHeight = 320
)

type Vector2 struct {
	X float32
	Y float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// toScreen converts math coordinates (-1..1) to screen coordinates (0..width/height).
func toScreen(width, height int, v2 Vector2) Vector2 {
	sx := (v2.X + 1) / 2 * float32(width)
	sy := (1 - (v2.Y+1)/2) * float32(height) // Invert Y axis

	return Vector2{sx, sy}
}

// project converts 3D coordinates to 2D using a simple perspective projection.
func project(v3 Vector3) Vector2 {
	px := v3.X / v3.Z
	py := v3.Y / v3.Z

	return Vector2{px, py}
}

// transformZ translates a 3D point along the Z axis.
func transformZ(v3 Vector3, dz float32) Vector3 {
	return Vector3{
		X: v3.X,
		Y: v3.Y,
		Z: v3.Z + dz,
	}
}

// rotateXZ rotates a 3D point around the Y axis.
func rotateXZ(v3 Vector3, angle float32) Vector3 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))

	return Vector3{
		X: v3.X*c - v3.Z*s,
		Y: v3.Y,
		Z: v3.X*s + v3.Z*c,
	}
}

// rotateYZ rotates a 3D point around the X axis.
func rotateYZ(v3 Vector3, angle float32) Vector3 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))

	return Vector3{
		X: v3.X,
		Y: v3.Y*c - v3.Z*s,
		Z: v3.Y*s + v3.Z*c,
	}
}

type Game struct {
	counter  int
	width    int
	height   int
	vertices []Vector3
	faces    [][]int
}

func NewGame(width, height int) *Game {
	vs := []Vector3{
		{X: 0.25, Y: 0.25, Z: 0.25},
		{X: -0.25, Y: 0.25, Z: 0.25},
		{X: 0.25, Y: -0.25, Z: 0.25},
		{X: -0.25, Y: -0.25, Z: 0.25},

		{X: 0.25, Y: 0.25, Z: -0.25},
		{X: -0.25, Y: 0.25, Z: -0.25},
		{X: 0.25, Y: -0.25, Z: -0.25},
		{X: -0.25, Y: -0.25, Z: -0.25},
	}

	fs := [][]int{
		{0, 1, 3, 2},
		{4, 5, 7, 6},
		{0, 4},
		{1, 5},
		{2, 6},
		{3, 7},
	}

	return &Game{
		width:    width,
		height:   height,
		vertices: vs,
		faces:    fs,
	}
}

func (g *Game) Update() error {
	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	speed := float32(0.1)

	// Calculate delta time for animation.
	// FPS is fixed at 60 in Ebiten.
	dt := float32(g.counter) / 60.0 * speed

	// Move the cube away from the camera.
	dz := float32(1.0)

	// Rotate the cube around the Y axis.
	angle := math.Pi * dt

	for _, face := range g.faces {
		for i := 0; i < len(face); i++ {
			a := g.vertices[face[i]]
			b := g.vertices[face[(i+1)%len(face)]]

			ra := rotateXZ(a, angle)
			rb := rotateXZ(b, angle)

			ra = rotateYZ(ra, angle)
			rb = rotateYZ(rb, angle)

			ta := transformZ(ra, dz)
			tb := transformZ(rb, dz)

			pa := project(ta)
			pb := project(tb)

			sa := toScreen(g.width, g.height, pa)
			sb := toScreen(g.width, g.height, pb)

			thickness := float32(2)
			color := color.RGBA{0, 255, 65, 255}

			vector.StrokeLine(screen, sa.X, sa.Y, sb.X, sb.Y, thickness, color, true)

		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowTitle("3D Cube")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	game := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
