package ch

import (
	"time"
	"testing"
	"math/rand"
	"fmt"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestNode(t *testing.T) {

	n1 := &Node{ANode:ANode{Id:0}, Capacity:2}
	n2 := &Node{ANode:ANode{Id:1}, Capacity:2}
	n3 := &Node{ANode:ANode{Id:2}, Capacity:2}
	n4 := &Node{ANode:ANode{Id:3}, Capacity:2}
	n5 := &Node{ANode:ANode{Id:4}, Capacity:2}
	n6 := &Node{ANode:ANode{Id:5}, Capacity:2}
	topology := Topology{}
	topology.addNode(n1).addNode(n2).addNode(n3).addNode(n4).addNode(n5).addNode(n6)
	topology.Print()

	s := "8"
	key := []byte(s)
	fmt.Printf("Hash: %10d Key:%s\n", Murmur3(key), s)

	n := topology.NodeForKey(key)
	fmt.Printf("Node ID: %s\n", n.String())

	/*
	if gorizon != gorizon_expected {
		t.Errorf("For 'gorizon' expecting '%d', but got '%d'", gorizon_expected, gorizon)
	}

	if ugol != ugol_expected {
		t.Errorf("For 'ugol' expecting  '%d', but got '%d'", ugol_expected, ugol)
	}
	*/
}
