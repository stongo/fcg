package main

import (
  "testing"
  "fmt"
)

func errorMsg(message string, expect, actual interface{}) string {
    return fmt.Sprintf("%s. Got: %#v Expected: %#v", message, actual, expect)
}

func TestNewBuilder(t *testing.T) {
  expect := &Builder{}
  actual := NewBuilder()
  if len(actual.graph) != len(expect.graph) {
    t.Error(errorMsg("Values don't match.", actual, expect))
  }
}

func TestAddNode(t *testing.T) {
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

func TestInit(t *testing.T) {
  b := NewBuilder()
  expect := &Node{
    value: 3,
  }
  actual := b.Init(3)
  if actual.value != expect.value {
    t.Error(errorMsg("unexpected value", actual.value, expect.value)) 
  } 
}

func TestAddEdge(t *testing.T) {
  b := NewBuilder()
  var v uint64
  expect := make(map[string]*Node)
  expect["Node0"] = b.AddNode(v)
  expect["Node1"] = b.AddNode(v)
  expect["Node2"] = b.AddNode(v)
  b.addEdge(*expect["Node0"], *expect["Node1"], ADD)
  b.addEdge(*expect["Node1"], *expect["Node2"], MULT)
}

func TestConstant(t *testing.T) {

}

func TestAdd(t *testing.T) {

}

func TestMult(t *testing.T) {

}

func TestAssertEqual(t *testing.T) {

}

func TestFillNodes(t *testing.T) {

}

func TestCheckConstraints(t *testing.T) {

}

