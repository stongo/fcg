package main

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
	counter int
}

// Edge defines the edges between vertices in our graph
// The operator should only be set with constants ADD and MULT
// keys should only be pairs
type Edge struct {
	keys     []int
	operator int
}

// Node represents a graph vertex
// Uses simple integer keys
// The value is filled in by the FillNodes method
type Node struct {
	key      int
	value    uint64
	adjacent Edge
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
  var k int
  if len(b.graph) > 0 {
    k = len(b.graph) 
  }
	n := &Node{
		key:      k,
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
    keys: []int{c.key, d.key},
    operator: operator,
  } 
}

// Init initializes a node in the graph
func (b *Builder) Init(v uint64) *Node {
	return b.AddNode(v)
}

// Constant initializes a node in a graph, set to a constant value
func (b *Builder) Constant(value uint64) *Node {
	return b.AddNode(value)
}

// Add adds 2 nodes in the graph, returning a new node
func (b *Builder) Add(c, d Node) *Node {	
  n := b.AddNode(EMPTY_VALUE)
  n.adjacent = b.addEdge(c, d, ADD)
  return n
}

// Mul multiplies 2 nodes in the graph, returning a new node
func (b *Builder) Mult(c, d Node) *Node {
  n := b.AddNode(EMPTY_VALUE)
  n.adjacent = b.addEdge(c, d, MULT)
  return n
}

// AssertEqual asserts that 2 nodes are equal
func (b *Builder) AssertEqual(c, d Node) bool {
	return true
}

// TODO: need a method to pass input node values
func (b *Builder) FillNodes() error {
  g := b.graph
  for _, n := range(g) {
    a := n.adjacent
    kLen := len(a.keys)
    if kLen > 0 {
      if kLen > 2 {
        return fmt.Errorf("node edge count must not exceed 2. Actual: %d", kLen)
      }
      e1 := g[a.keys[0]]
      e2 := g[a.keys[1]]
      switch {
      case a.operator == MULT:
        n.value = e1.value * e2.value
      case a.operator == ADD:
        n.value = e1.value + e2.value
      } 
    } 
  }
  for i, _ := range(g) {
    fmt.Printf("n: %v\n", g[i])
  }
	return nil
}

// Given a graph that has `fill_nodes` already called on it
// checks that all the constraints hold
func (b *Builder) CheckConstraints() bool {
	return false
}

// incrementcounter increases the Builder counter
func (b *Builder) incrementCounter() {
  b.counter = b.counter + 1
}

func main() {
	// f(x) = x^2 + x + 5 would be defined by the following code:
	builder := NewBuilder()
	x := builder.Init(2)
	xSquared := builder.Mult(*x, *x)
	five := builder.Constant(5)
	xSquaredPlus5 := builder.Add(*xSquared, *five)
	y := builder.Add(*xSquaredPlus5, *x)

	builder.FillNodes()
	builder.CheckConstraints()
  fmt.Printf("result = %v", y.value)
}
