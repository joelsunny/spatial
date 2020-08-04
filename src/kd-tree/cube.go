package tree

import (
	"fmt"
	"math"

	svg "github.com/ajstarks/svgo"
)

// Cube definition
type Cube struct {
	start  vec3
	width  float64
	height float64
	depth  float64
}

func (c *Cube) findSubRegions(p vec3, axis int) (Cube, Cube) {
	var leftCube Cube
	var rightCube Cube

	switch axis {
	case x:
		lwidth := math.Abs(c.start.x - p.x)
		leftCube = Cube{c.start, lwidth, c.height, c.depth}
		rightCube = Cube{c.start.move(vec3{lwidth, 0, 0}), c.width - lwidth, c.height, c.depth}
	case y:
		lwidth := math.Abs(c.start.y - p.y)
		leftCube = Cube{c.start, c.width, lwidth, c.depth}
		rightCube = Cube{c.start.move(vec3{0, lwidth, 0}), c.width, c.height - lwidth, c.depth}
	case z:
		lwidth := math.Abs(c.start.x - p.x)
		leftCube = Cube{c.start, c.width, c.height, lwidth}
		rightCube = Cube{c.start.move(vec3{0, 0, lwidth}), c.width, c.height, c.depth - lwidth}
	}

	return leftCube, rightCube
}

func drawCube(canvas *svg.SVG, c Cube, color string) {
	style := fmt.Sprintf("fill:none; stroke:%v", color)

	// draw front face
	depth := c.start.z + c.depth
	d := vec3{0.5 * c.start.z, -0.5 * c.start.z, 0}
	cstart := c.start.move(d)
	canvas.Rect(int(cstart.x), int(cstart.y), int(c.width), int(c.height), style)
	// draw back face
	d = vec3{0.5 * depth, -0.5 * depth, 0}
	cend := c.start.move(d)
	//canvas.Rect(int(cend.x), int(cend.y), int(c.width), int(c.height), "fill:none; stroke:red")
	// top face
	xlist := []int{int(cstart.x), int(cend.x), int(cend.x) + int(c.width), int(cstart.x) + int(c.width)}
	ylist := []int{int(cstart.y), int(cend.y), int(cend.y), int(cstart.y)}
	canvas.Polygon(xlist, ylist, style)

	// right face
	xlist = []int{int(cstart.x) + int(c.width), int(cend.x) + int(c.width), int(cend.x) + int(c.width), int(cstart.x) + int(c.width)}
	ylist = []int{int(cstart.y), int(cend.y), int(cend.y) + int(c.height), int(cstart.y) + int(c.height)}
	canvas.Polygon(xlist, ylist, style)

}

func drawPoint(canvas *svg.SVG, p vec3) {
	d := vec3{0.5 * p.z, -0.5 * p.z, 0}
	p1 := p.move(d)
	//p1 = p
	canvas.Circle(int(p1.x), int(p1.y), 4, "fill:green; stroke:green")
	//text := fmt.Sprintf("(%v,%v)", p.x, p.y)
	//canvas.Text(int(p.x)+5, int(p.y)+5, text, "text-anchor:middle;font-size:10px;fill:yellow")
}

func drawNode(canvas *svg.SVG, n *node) {
	color := ""
	switch n.splitAxis {
	case x:
		color = "green"
	case y:
		color = "blue"
	case z:
		color = "red"
	}
	drawCube(canvas, n.region, color)
	if n.containspoint {
		drawPoint(canvas, n.point)
	}
	if n.ispartitioned {
		drawNode(canvas, n.left)
		drawNode(canvas, n.right)
	}
}
