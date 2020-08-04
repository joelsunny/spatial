package tree

import (
	"os"
	"testing"

	svg "github.com/ajstarks/svgo"
)

func TestCube(t *testing.T) {
	c := Cube{vec3{50, 130, 0}, 256, 256, 256}
	c1 := Cube{vec3{70, 70, 0}, 50, 30, 25}
	width := 512
	height := 512
	fo, _ := os.Create("out.svg")
	defer fo.Close()
	canvas := svg.New(fo)
	canvas.Start(width, height)
	drawCube(canvas, c, "green")
	drawCube(canvas, c1, "green")
	canvas.End()

}
