package tree

import (
	"os"
	"testing"

	svg "github.com/ajstarks/svgo"
)

func TestNode(t *testing.T) {
	c := Cube{vec3{50, 130, 0}, 256, 256, 256}
	n := &node{region: c}
	p1 := vec3{70, 180, 10}
	p2 := vec3{150, 230, 30}
	p3 := vec3{270, 150, 100}
	p4 := vec3{280, 150, 30}
	p5 := vec3{250, 150, 30}

	n.addPoint(p1)
	n.addPoint(p2)
	n.addPoint(p3)
	n.addPoint(p4)
	n.addPoint(p5)

	width := 512
	height := 512
	fo, _ := os.Create("out.svg")
	defer fo.Close()
	canvas := svg.New(fo)
	canvas.Start(width, height)
	drawNode(canvas, n)
	canvas.End()

}

func TestVec(t *testing.T) {
	v := vec3{10, 10, 20}
	v1 := vec3{9, 11, 24}

	if v.lt(v1, x) {
		t.Errorf("test 1 failed")
	}

	if !v.lt(v1, y) {
		t.Errorf("test 2 failed")
	}
}
