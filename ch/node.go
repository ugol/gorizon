package ch

import (
	"fmt"
	"sort"
)


type ANode struct {
	Id    uint32
	RefId uint32
	Hash  uint32
}

type Node struct {
	ANode
	Capacity uint8
}

type Topology struct {
	nodes []ANode
}

func (n *ANode) String() (string) {
	return fmt.Sprintf("%d:%d", n.Id, n.RefId)
}

func (t *Topology) init() {
	t.nodes = make([]ANode, 0, 5)
}

func (t *Topology) Print() {
	for _, n := range t.nodes {
		fmt.Printf("Hash: %10d Node.Id: %s \n", Murmur3([]byte(n.String())), n.String())
	}
}

func (t *Topology) addNode(n *Node) *Topology {

	newList := t.nodes
	n.Hash = Murmur3([]byte(n.String()))
	newList = append(newList, n.ANode)
	vnodes := n.Capacity
	for i := 1; i <= int(vnodes); i++ {
		vnode := ANode{Id:n.Id, RefId:uint32(i)}
		hash := Murmur3([]byte(vnode.String()))
		vnode.Hash = hash
		newList = append(newList, vnode)
	}
	t.nodes = newList
	sort.Sort(t) // should be done before
	return t
}

func (t *Topology) removeNode(id uint32) *Topology {

	newList := t.nodes
	//	for i, n := range newList {
	//		if n.Id == id {
	//
	//		}
	//	}

	t.nodes = newList
	return t
}

func (t *Topology) Len() int {
	return len(t.nodes)
}

func (t *Topology) Less(i, j int) bool {
	return t.nodes[i].Hash < t.nodes[j].Hash
}

func (t* Topology) Swap(i, j int) {
	t.nodes[i], t.nodes[j] = t.nodes[j], t.nodes[i]
}

func (t* Topology) NodeForKey(key []byte) *ANode {

	hash := Murmur3(key)

	for _, n := range t.nodes {
		if n.Hash >= hash {
			return &n
		}
	}
	return &t.nodes[0]
}