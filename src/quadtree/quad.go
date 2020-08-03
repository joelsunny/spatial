package main

import "fmt"

type point struct {
	x float64
	y float64
}

// return a new point relative to the point
func (p *point) move(x float64, y float64) point {
	newp := point{x: p.x + x, y: p.y + y}
	return newp
}

func (p *point) le(p1 point) bool {
	if p1.x >= p.x && p1.y >= p.y {
		return true
	}
	return false
}

func (p *point) gt(p1 point) bool {
	if p.x > p1.x && p.y > p1.y {
		return true
	}
	return false
}

// Rectangle region definition
type Rectangle struct {
	start  point
	width  float64
	height float64
}

// Partition a rectangle into four equal quadrants
func (r *Rectangle) partition() (Rectangle, Rectangle, Rectangle, Rectangle) {
	width := r.width
	height := r.height
	nw := Rectangle{start: r.start, width: width / 2, height: height / 2}
	ne := Rectangle{start: r.start.move(width/2, 0), width: width / 2, height: height / 2}
	sw := Rectangle{start: r.start.move(0, height/2), width: width / 2, height: height / 2}
	se := Rectangle{start: r.start.move(width/2, height/2), width: width / 2, height: height / 2}

	return nw, ne, sw, se
}

func (r *Rectangle) doesContain(p point) bool {
	res := false
	endRect := r.start.move(r.width, r.height)
	if r.start.le(p) && endRect.gt(p) {
		res = true
	}
	return res
}

// check whether the two rectangles intersect
func (r *Rectangle) doesIntersect(r1 Rectangle) bool {

	r1corners := []point{r1.start, r1.start.move(r1.width, r1.height), r1.start.move(r1.width, 0), r1.start.move(0, r1.height)}
	rcorners := []point{r.start, r.start.move(r.width, r.height), r.start.move(r.width, 0), r.start.move(0, r.height)}

	intersect := false
	for i := 0; i < len(r1corners); i++ {
		if r.doesContain(r1corners[i]) {
			intersect = true
			break
		}
	}

	for i := 0; i < len(rcorners); i++ {
		if r1.doesContain(rcorners[i]) {
			intersect = true
			break
		}
	}

	return intersect
}

// Qnode , quadtree individual node definition
type Qnode struct {
	region        Rectangle
	ispartitioned bool
	containspoint bool
	p             point
	parent        *Qnode
	nw            *Qnode
	ne            *Qnode
	sw            *Qnode
	se            *Qnode
}

func (q *Qnode) addSubRegions() {
	nw, ne, sw, se := q.region.partition()
	q.nw = &Qnode{region: nw, parent: q}
	q.ne = &Qnode{region: ne, parent: q}
	q.sw = &Qnode{region: sw, parent: q}
	q.se = &Qnode{region: se, parent: q}
	q.ispartitioned = true
}

func (q *Qnode) findQuad(p point, add bool) (bool, *Qnode) {
	if !q.region.doesContain(p) {
		return false, nil
	}

	if !q.containspoint {
		return true, q
	} else if add == false {
		return false, nil
	} else if add == true && q.containspoint && !q.ispartitioned {
		q.addSubRegions()
		q.addPoint(q.p)
	}

	fmt.Println("checking sub regions")
	if q.nw.region.doesContain(p) {
		return q.nw.findQuad(p, add)
	} else if q.ne.region.doesContain(p) {
		return q.ne.findQuad(p, add)
	} else if q.sw.region.doesContain(p) {
		return q.sw.findQuad(p, add)
	}
	return q.se.findQuad(p, add)
}

func (q *Qnode) addPoint(p point) bool {
	if !q.region.doesContain(p) {
		return false
	}

	if !q.containspoint {
		q.p = p
		q.containspoint = true
		return true
	}

	b, subreg := q.findQuad(p, true)
	if !b {
		fmt.Println("point outside")
		return false
	}
	subreg.addPoint(p)
	return true
}

// range query implementation
func (q *Qnode) findPoints(region Rectangle) []point {
	res := []point{}
	if q.region.doesIntersect(region) {
		if q.containspoint && !q.ispartitioned {
			if region.doesContain(q.p) {
				res = append(res, q.p)
			}
		} else if q.ispartitioned {
			res = append(res, q.nw.findPoints(region)...)
			res = append(res, q.ne.findPoints(region)...)
			res = append(res, q.sw.findPoints(region)...)
			res = append(res, q.se.findPoints(region)...)
		}
	}

	return res
}

// QuadTree definition
type QuadTree struct {
	root *Qnode
}
