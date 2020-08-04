package tree

const (
	x = iota
	y
	z
)

type vec3 struct {
	x float64
	y float64
	z float64
}

func (v *vec3) move(d vec3) vec3 {
	return vec3{v.x + d.x, v.y + d.y, v.z + d.z}
}

func (v *vec3) lt(v1 vec3, axis int) bool {
	switch axis {
	case x:
		if v.x < v1.x {
			return true
		}
	case y:
		if v.y < v1.y {
			return true
		}
	case z:
		if v.z < v1.z {
			return true
		}
	}
	return false
}

type node struct {
	region        Cube
	point         vec3
	splitAxis     int
	containspoint bool
	ispartitioned bool
	left          *node
	right         *node
}

func (n *node) split() {
	childaxis := n.getChildSplitAxis()
	l, r := n.region.findSubRegions(n.point, n.splitAxis)
	n.left = &node{splitAxis: childaxis, region: l}
	n.right = &node{splitAxis: childaxis, region: r}
	n.ispartitioned = true
}

func (n *node) getChildSplitAxis() int {
	return (n.splitAxis + 1) % 3
}

func (n *node) addPoint(p vec3) {
	if !n.containspoint {
		n.point = p
		n.containspoint = true
		return
	} else if !n.ispartitioned {
		n.split()
	}

	if p.lt(n.point, n.splitAxis) {
		n.left.addPoint(p)
		return
	} else {
		n.right.addPoint(p)
		return
	}
}
