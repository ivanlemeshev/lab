package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

func main() {
	nodes := createNodes(1)

	for i := range 10 {
		key := fmt.Sprintf("key%02d", i+1)
		keyToNode(key, nodes)
	}

	fmt.Printf("%+v\n", nodes)

	nodes = createNodes(2)

	for i := range 10 {
		key := fmt.Sprintf("key%02d", i+1)
		keyToNode(key, nodes)
	}

	fmt.Printf("%+v\n", nodes)

	nodes = createNodes(3)

	for i := range 10 {
		key := fmt.Sprintf("key%02d", i+1)
		keyToNode(key, nodes)
	}

	fmt.Printf("%+v\n", nodes)
}

func keyToNode(key string, nodes []node) {
	keyHash := hash(key)
	node := keyHash % uint32(len(nodes))
	nodes[node].AddKey()
	fmt.Println(key, "->", node)
}

type node struct {
	id   uint32
	keys int
}

func (n *node) AddKey() {
	n.keys++
}

func createNodes(nodesCount int) []node {
	nodes := make([]node, 0, nodesCount)
	for i := range nodesCount {
		nodeID := uint32(i)
		nodes = append(nodes, node{id: nodeID})
	}
	return nodes
}

func hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func Hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

type HashRingNode struct {
	ID   string
	Hash uint32
}

func NewHashRingNode(id string) *HashRingNode {
	return &HashRingNode{
		ID:   id,
		Hash: Hash(id),
	}
}

type HashRing struct {
	nodes []*HashRingNode
}

func NewHashRing() *HashRing {
	return &HashRing{}
}

func (hr *HashRing) AddNode(id string) {
	node := NewHashRingNode(id)
	hr.nodes = append(hr.nodes, node)
	sort.Sort(r.Nodes)
}
