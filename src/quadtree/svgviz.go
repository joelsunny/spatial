package main

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
)

func drawPoint(canvas *svg.SVG, p point) {
	canvas.Circle(int(p.x), int(p.y), 2)
	text := fmt.Sprintf("(%v,%v)", p.x, p.y)
	canvas.Text(int(p.x)+3, int(p.y), text, "text-anchor:middle;font-size:10px;fill:blue")
}

func drawRect(canvas *svg.SVG, r Rectangle) {
	canvas.Rect(int(r.start.x), int(r.start.y), int(r.width), int(r.height), "fill:none; stroke:blue")
}

func drawQuad(canvas *svg.SVG, q *Qnode) {
	drawRect(canvas, q.region)
	if q.containspoint && !q.ispartitioned {
		drawPoint(canvas, q.p)
	}
	if q.ispartitioned {
		drawQuad(canvas, q.nw)
		drawQuad(canvas, q.ne)
		drawQuad(canvas, q.sw)
		drawQuad(canvas, q.se)
	}
}

func drawRangeQueryResult(canvas *svg.SVG, r Rectangle, points []point) {
	canvas.Rect(int(r.start.x), int(r.start.y), int(r.width), int(r.height), "fill:none; stroke:red")
	for i := 0; i < len(points); i++ {
		p := points[i]
		canvas.Circle(int(p.x), int(p.y), 5, "fill:red; stroke:red")
	}
}
