package fcg

import (
	"fmt"
	"testing"
)

func errorMsg(message string, expect, actual interface{}) string {
	return fmt.Sprintf("%s. Got: %#v Expected: %#v", message, actual, expect)
}

func TestBuilderNewBuilder(t *testing.T) {
	expect := &Builder{}
	actual := NewBuilder()
	if len(actual.graph) != len(expect.graph) {
		t.Error(errorMsg("Values don't match.", actual, expect))
	}
}

func TestBuilderAddNode(t *testing.T) {
	b := NewBuilder()
	var v uint64
	expect := &Node{}
	actual := b.AddNode(v)
	if actual.value != expect.value {
		t.Error(errorMsg("Values don't match.", actual, expect))
	}
	actual.value = 1
	if actual.value != b.graph[0].value {
		t.Error(errorMsg("Values don't match.", actual, expect))
	}
	if actual.value == expect.value {
		t.Error(errorMsg("Values shouldn't match.", actual, expect))
	}
	b.graph[0].value = 2
	if actual.value != 2 {
		t.Error(errorMsg("Value should be updated.", actual, expect))
	}
}

func TestBuilderInit(t *testing.T) {
	b := NewBuilder()
	expect := &Node{
		value: 3,
	}
	actual := b.Init(3)
	if actual.value != expect.value {
		t.Error(errorMsg("unexpected value", actual.value, expect.value))
	}
}

func TestBuilderAddEdge(t *testing.T) {
	b := NewBuilder()
	var v uint64
	expect := make(map[string]Edge)
	expect["edge0"] = Edge{
		keys:     []uint{0, 1},
		operator: ADD,
	}
	expect["edge1"] = Edge{
		keys:     []uint{1, 2},
		operator: MULT,
	}
	actual := make(map[string]Edge)
	node0 := b.AddNode(v)
	node1 := b.AddNode(v)
	node2 := b.AddNode(v)
	actual["edge0"] = b.addEdge(*node0, *node1, ADD)
	actual["edge1"] = b.addEdge(*node1, *node2, MULT)
	for i := 0; i < 2; i++ {
		if actual["edge0"].keys[i] != expect["edge0"].keys[i] {
			t.Error(errorMsg("unexpected edge0 key", actual["edge0"].keys[i], expect["edge0"].keys[i]))
		}
	}
	if actual["edge0"].operator != expect["edge0"].operator {
		t.Error(errorMsg("unexpected edge0 operator", actual["edge0"].operator, expect["edge0"].operator))
	}
	for i := 0; i < 2; i++ {
		if actual["edge1"].keys[i] != expect["edge1"].keys[i] {
			t.Error(errorMsg("unexpected edge1 key", actual["edge1"].keys[i], expect["edge1"].keys[i]))
		}
	}
	if actual["edge1"].operator != expect["edge1"].operator {
		t.Error(errorMsg("unexpected edge0 operator", actual["edge1"].operator, expect["edge1"].operator))
	}

}

func TestBuilderConstant(t *testing.T) {
	b := NewBuilder()
	expect := &Node{
		value: 7,
	}
	actual := b.Constant(7)
	if actual.value != expect.value {
		t.Error(errorMsg("unexpected constant value", actual.value, expect.value))
	}

}

func TestBuilderAdd(t *testing.T) {
	b := NewBuilder()
	var v uint64
	node0 := b.AddNode(v)
	node1 := b.AddNode(v)
	expect := Edge{
		keys:     []uint{node0.key, node1.key},
		operator: ADD,
	}
	actual := b.Add(*node0, *node1)
	for i := 0; i < 2; i++ {
		if actual.parents.keys[i] != expect.keys[i] {
			t.Error(errorMsg("unexpected parents key in Add", actual.parents.keys[i], expect.keys[i]))
		}
	}
	if actual.parents.operator != expect.operator {
		t.Error(errorMsg("unexpected parents operator in Add", actual.parents.operator, expect.operator))
	}
}

func TestBuilderMult(t *testing.T) {
	b := NewBuilder()
	var v uint64
	node0 := b.AddNode(v)
	node1 := b.AddNode(v)
	expect := Edge{
		keys:     []uint{node0.key, node1.key},
		operator: MULT,
	}
	actual := b.Mult(*node0, *node1)
	for i := 0; i < 2; i++ {
		if actual.parents.keys[i] != expect.keys[i] {
			t.Error(errorMsg("unexpected parents key in Add", actual.parents.keys[i], expect.keys[i]))
		}
	}
	if actual.parents.operator != expect.operator {
		t.Error(errorMsg("unexpected parents operator in Add", actual.parents.operator, expect.operator))
	}

}

func TestBuilderAssertEqual(t *testing.T) {
	b := NewBuilder()
	var v uint64
	node0 := b.AddNode(v)
	node1 := b.AddNode(v)
	expect := true
	actual := b.AssertEqual(*node0, *node1)
	if actual != expect {
		t.Error(errorMsg("problem with AssertEqual", actual, expect))
	}

}

func TestBuilderFillNodes(t *testing.T) {
	// f(x) = x^2 + x + 5 would be defined by the following code:
	builder := NewBuilder()
	x := builder.Init(6)
	// evaluates to 47
	expect := uint64(47)
	xSquared := builder.Mult(*x, *x)
	five := builder.Constant(5)
	xSquaredPlus5 := builder.Add(*xSquared, *five)
	y := builder.Add(*xSquaredPlus5, *x)
	builder.FillNodes()
	actual := y.Get()
	if actual != expect {
		t.Error(errorMsg("the graph did not evaluate properly", actual, expect))
	}

}

func TestBuilderCheckConstraints(t *testing.T) {

}

func TestNodeGet(t *testing.T) {
	b := NewBuilder()
	n := b.AddNode(9)
	expect := uint64(9)
	actual := n.Get()
	if actual != expect {
		t.Error(errorMsg("problem with Get", actual, expect))
	}
}
