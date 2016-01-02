package ch

import (
	"bytes"
	"encoding/binary"
)

//
// https://github.com/reusee/mmh3
func Murmur3(key []byte) uint32 {

	length := len(key)
	if length == 0 {
		return 0
	}

	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	var h, k uint32
	buf := bytes.NewBuffer(key)

	nblocks := length / 4
	for i := 0; i < nblocks; i++ {
		binary.Read(buf, binary.LittleEndian, &k)
		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2
		h ^= k
		h = (h << 13) | (h >> (32 - 13))
		h = (h * 5) + 0xe6546b64
	}

	k = 0
	tailIndex := nblocks * 4
	switch length & 3 {
	case 3:
		k ^= uint32(key[tailIndex + 2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[tailIndex + 1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[tailIndex])
		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2
		h ^= k
	}
	h ^= uint32(length)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

