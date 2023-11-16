package fcg

import "fmt"

// MULT - multiplication operator for Edges
const MULT = 1

// ADD - addition operator for Edges
const ADD = 2

// EQUAL - equal operator for Edges
const EQUAL = 3

// NOOP - input nodes should not have Edges
const NOOP = 0

// EMPTY - zero value for Node value
// we compute values later unless it's a constant
const EMPTY_VALUE = 0

// Builder contains our graph and a counter
type Builder struct {
	graph []*Node
}

// Edge defines the edges between vertices in our graph
// The operator should only be set with constants ADD and MULT
// keys should only be pairs
type Edge struct {
	keys     []uint
	operator int
}

// Node represents a graph vertex
// Uses simple integer keys
// The value is filled in by the FillNodes method
type Node struct {
	key     uint
	value   uint64
	parents Edge
}

// newBuilder creates a new Builder
func NewBuilder() *Builder {
	return &Builder{}
}

// AddNode adds the root node to the graph
// It ideally should not be called otherwise, but we won't enforce it for now
// Warning: for non-root nodes, you risk creating orphans nodes
// Increments the builder counter on succesful node creation
// Returns the key of the newly created node or an error
func (b *Builder) AddNode(v uint64) *Node {
	var k uint
	if len(b.graph) > 0 {
		k = uint(len(b.graph))
	}
	n := &Node{
		key:   k,
		value: v,
	}
	b.graph = append(b.graph, n)
	return n
}

// addEdge creates edges
// edge direction is dependent on argument order
// method doesn't implement any sorting function
func (b *Builder) addEdge(c, d Node, operator int) Edge {
	return Edge{
		keys:     []uint{c.key, d.key},
		operator: operator,
	}
}

// Init initializes a node in the graph
func (b *Builder) Init(v uint64) *Node {
	return b.AddNode(v)
}

// Constant initializes a node in a graph, set to a constant value
func (b *Builder) Constant(v uint64) *Node {
	return b.AddNode(v)
}

// Add adds 2 nodes in the graph, returning a new node
func (b *Builder) Add(c, d Node) *Node {
	n := b.AddNode(EMPTY_VALUE)
	n.parents = b.addEdge(c, d, ADD)
	return n
}

// Mul multiplies 2 nodes in the graph, returning a new node
func (b *Builder) Mult(c, d Node) *Node {
	n := b.AddNode(EMPTY_VALUE)
	n.parents = b.addEdge(c, d, MULT)
	return n
}

// AssertEqual asserts that 2 nodes are equal
func (b *Builder) AssertEqual(c, d Node) bool {
	return c.Get() == d.Get()
}

// FillNodes fill in all node values in our graph
func (b *Builder) FillNodes() error {
	g := b.graph
	for _, n := range g {
		p := n.parents
		kLen := len(p.keys)
		if kLen > 0 {
			if kLen > 2 {
				return fmt.Errorf("node edge count must not exceed 2. Actual: %d", kLen)
			}
			e1 := g[p.keys[0]]
			e2 := g[p.keys[1]]
			switch {
			case p.operator == MULT:
				n.Set(e1.Get() * e2.Get())
			case p.operator == ADD:
				n.Set(e1.Get() + e2.Get())
			}
		}
	}
	return nil
}

// ViewGraph prints the graph to console
func (b *Builder) ViewGraph() error {
	g := b.graph
	for i, _ := range g {
		fmt.Printf("n: %v\n", g[i])
	}
	return nil
}

func (n *Node) Set(v uint64) {
	n.value = v
}

func (n *Node) Get() uint64 {
	return n.value
}

// Given a graph that has `fill_nodes` already called on it
// checks that all the constraints hold
// @TODO
func (b *Builder) CheckConstraints() bool {
	return false
}
