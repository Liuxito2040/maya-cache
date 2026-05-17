package ring

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Ring struct {
	nodes        map[uint32]string
	hashes       []uint32
	virtualNodes int
}

func New(virtualNodes int) *Ring {
	return &Ring{
		nodes:        make(map[uint32]string),
		hashes:       make([]uint32, 0),
		virtualNodes: virtualNodes,
	}
}

func hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// AddNode adds a node to the hashring taking the given address to generate
// a hash corresponding to the added node.
func (r *Ring) AddNode(address string) {
	for i := 0; i < r.virtualNodes; i++ {
		hash := hashKey(address + fmt.Sprintf("#%d", i))
		r.nodes[hash] = address
		r.hashes = append(r.hashes, hash)
	}

	sort.Slice(r.hashes, func(i, j int) bool {
		return r.hashes[i] < r.hashes[j]
	})
}

// GetNode returns the node address corresponding to the given key by hashing the key
// and finding the closest node in the hashring going clockwise.
//
// Ej: hash(key) = 10, and the nodes in the ring are at positions 5, 15, and 25.
// The closest node going clockwise from 10 is the node at position 15.
func (r *Ring) GetNode(key string) string {
	if len(r.hashes) == 0 {
		return ""
	}

	hash := hashKey(key)
	idx := sort.Search(
		len(r.hashes),
		func(i int) bool {
			return hash <= r.hashes[i]
		},
	)

	if idx == len(r.hashes) {
		idx = 0
	}
	return r.nodes[r.hashes[idx]]
}
