package main

import (
	"math/rand"
	"os"
	"testing"

	svg "github.com/ajstarks/svgo"
)

func TestQnode(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	size := 512.0
	plist := []point{}

	for i := 0; i < 100; i++ {
		plist = append(plist, point{x: r.Float64() * size, y: r.Float64() * size})
	}

}

func TestPoint(t *testing.T) {
	p := point{x: 10, y: 20}
	p1 := point{x: 10, y: 20}
	p2 := point{x: 9, y: 19}
	p3 := point{x: 9, y: 21}

	if p.le(p1) != true {
		t.Errorf("test 1 failed")
	}

	if p.gt(p2) != true {
		t.Errorf("test 2 failed")
	}

	if p.gt(p3) == true {
		t.Errorf("test 2 failed")
	}

	if p.le(p3) == true {
		t.Errorf("test 3 failed")
	}

}

func TestRectangle(t *testing.T) {
	p := point{x: 10, y: 10}
	r := Rectangle{start: p, width: 100.0, height: 100.0}

	p1 := point{x: 15, y: 200}
	p2 := point{x: 15, y: 15}

	if r.doesContain(p) != true {
		t.Errorf("test 1 failed")
	}

	if r.doesContain(p1) == true {
		t.Errorf("test 2 failed")
	}

	if r.doesContain(p2) != true {
		t.Errorf("test 3 failed")
	}
}

func TestAddPoint(t *testing.T) {
	p := point{x: 0, y: 0}
	r := Rectangle{start: p, width: 512.0, height: 512.0}
	q := &Qnode{region: r}

	q.addPoint(point{x: 10, y: 10})
	q.addPoint(point{x: 130, y: 10})
	q.addPoint(point{x: 260, y: 504})
	q.addPoint(point{x: 280, y: 504})
	q.addPoint(point{x: 257, y: 500})
	q.addPoint(point{x: 500, y: 501})
	width := 512
	height := 512
	fo, _ := os.Create("out.svg")
	defer fo.Close()
	canvas := svg.New(fo)
	canvas.Start(width, height)
	drawQuad(canvas, q)
	canvas.End()
}

func TestRangeQuery(t *testing.T) {
	p := point{x: 0, y: 0}
	r := Rectangle{start: p, width: 512.0, height: 512.0}
	q := &Qnode{region: r}

	q.addPoint(point{x: 10, y: 10})
	q.addPoint(point{x: 130, y: 10})
	q.addPoint(point{x: 260, y: 504})
	q.addPoint(point{x: 280, y: 504})
	q.addPoint(point{x: 257, y: 500})
	q.addPoint(point{x: 500, y: 501})
	region := Rectangle{start: p, width: 200.0, height: 200.0}
	res := q.findPoints(region)

	width := 512
	height := 512
	fo, _ := os.Create("out.svg")
	defer fo.Close()
	canvas := svg.New(fo)
	canvas.Start(width, height)
	drawQuad(canvas, q)
	drawRangeQueryResult(canvas, region, res)
	canvas.End()
}
