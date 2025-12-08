package dsu

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	d := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.size[i] = 1
	}

	return d
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}

	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	ra := d.Find(a)
	rb := d.Find(b)

	if ra == rb {
		return false
	}

	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}

	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

func (d *DSU) SameSet(a, b int) bool {
	return d.Find(a) == d.Find(b)
}

func (d *DSU) Size(x int) int {
	root := d.Find(x)
	return d.size[root]
}
