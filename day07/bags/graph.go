package bags

type graph struct {
	keysAndNodes map[string]*node
}

type node struct {
	key         string
	parentEdges []edge
	childEdges  []edge
}

type edge struct {
	parent *node
	child  *node
	weight int
}

func NewGraph() *graph {
	return &graph{keysAndNodes: make(map[string]*node)}
}

func (g *graph) ElemCount() int {
	return len(g.keysAndNodes)
}

func (g *graph) GetElemMap() map[string]*node {
	return g.keysAndNodes
}

func (g *graph) AddNode(key string) *node {
	existing, ok := g.keysAndNodes[key]
	if !ok {
		node := &node{key: key}
		g.keysAndNodes[key] = node
		return node
	} else {
		return existing
	}
}

func (g *graph) AddChild(parentKey, childKey string, weight int) {
	parent, ok := g.keysAndNodes[parentKey]
	if !ok {
		parent = g.AddNode(parentKey)
	}
	child, ok := g.keysAndNodes[childKey]
	if !ok {
		child = g.AddNode(childKey)
	}

	e := edge{parent: parent, child: child, weight: weight}
	parent.childEdges = append(parent.childEdges, e)
	child.parentEdges = append(child.parentEdges, e)
}

func (n *node) ChildrenContain(key string) bool {
	childrenContain := false
	if len(n.childEdges) != 0 {
		for _, e := range n.childEdges {
			child := e.child
			if child.key == key {
				return true
			} else {
				if child.ChildrenContain(key) {
					childrenContain = true
					break
				}
			}
		}
	}

	return childrenContain
}

func (n *node) Size() int {
	size := 1
	if len(n.childEdges) > 0 {
		for _, child := range n.childEdges {
			size += child.weight * child.child.Size()
		}
	}
	return size
}
